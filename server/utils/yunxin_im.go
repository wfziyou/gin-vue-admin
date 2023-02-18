package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/google/go-querystring/query"
	"github.com/mozillazg/go-httpheader"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

type YunxinIm struct {
}

const (
	// Version current go sdk version
	Version               = "0.7.19"
	userAgent             = "cos-go-sdk-v5/" + Version
	contentTypeJson       = "application/json"
	defaultServiceBaseURL = "http://service.cos.myqcloud.com"
)

type sendOptions struct {
	// 基础 URL
	baseURL *url.URL
	// URL 中除基础 URL 外的剩余部分
	uri string
	// 请求方法
	method string

	body interface{}
	// url 查询参数
	optQuery interface{}
	// http header 参数
	optHeader interface{}
	// 用 result 反序列化 resp.Body
	result interface{}
	// 是否禁用自动调用 resp.Body.Close()
	// 自动调用 Close() 是为了能够重用连接
	disableCloseBody bool
}
type Client struct {
	client *http.Client

	Host      string
	UserAgent string
	BaseUrl   *url.URL

	common service
}

// Response API 响应
type Response struct {
	*http.Response
}

type ErrorResponse struct {
	XMLName   xml.Name       `xml:"Error"`
	Response  *http.Response `xml:"-"`
	Code      string
	Message   string
	Resource  string
	RequestID string `header:"x-cos-request-id,omitempty" url:"-" xml:"RequestId,omitempty"`
	TraceID   string `xml:"TraceId,omitempty"`
}

// Error returns the error msg
func (r *ErrorResponse) Error() string {
	RequestID := r.RequestID
	if RequestID == "" {
		RequestID = r.Response.Header.Get("X-Cos-Request-Id")
	}
	TraceID := r.TraceID
	if TraceID == "" {
		TraceID = r.Response.Header.Get("X-Cos-Trace-Id")
	}
	return fmt.Sprintf("%v %v: %d %v(Message: %v, RequestId: %v, TraceId: %v)",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Code, r.Message, RequestID, TraceID)
}

func newResponse(resp *http.Response) *Response {
	return &Response{
		Response: resp,
	}
}

// addURLOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func addURLOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	// 保留原有的参数，并且放在前面。因为 cos 的 url 路由是以第一个参数作为路由的
	// e.g. /?uploads
	q := u.RawQuery
	rq := qs.Encode()
	if q != "" {
		if rq != "" {
			u.RawQuery = fmt.Sprintf("%s&%s", q, qs.Encode())
		}
	} else {
		u.RawQuery = rq
	}
	return u.String(), nil
}

// addHeaderOptions adds the parameters in opt as Header fields to req. opt
// must be a struct whose fields may contain "header" tags.
func addHeaderOptions(header http.Header, opt interface{}) (http.Header, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return header, nil
	}

	h, err := httpheader.Header(opt)
	if err != nil {
		return nil, err
	}

	for key, values := range h {
		for _, value := range values {
			header.Add(key, value)
		}
	}
	return header, nil
}

func (c *Client) newRequest(ctx context.Context, baseURL *url.URL, uri, method string, body interface{}, optQuery interface{}, optHeader interface{}) (req *http.Request, err error) {
	uri, err = addURLOptions(uri, optQuery)
	if err != nil {
		return
	}
	u, _ := url.Parse(uri)
	urlStr := baseURL.ResolveReference(u).String()

	var reader io.Reader
	contentType := ""
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		contentType = contentTypeJson
		reader = bytes.NewReader(b)
	}

	req, err = http.NewRequest(method, urlStr, reader)
	if err != nil {
		return
	}

	req.Header, err = addHeaderOptions(req.Header, optHeader)
	if err != nil {
		return
	}
	if v := req.Header.Get("Content-Length"); req.ContentLength == 0 && v != "" && v != "0" {
		req.ContentLength, _ = strconv.ParseInt(v, 10, 64)
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if req.Header.Get("Content-Type") == "" && contentType != "" {
		req.Header.Set("Content-Type", contentTypeJson)
	}
	if c.Host != "" {
		req.Host = c.Host
	}
	return
}

func (c *Client) doAPI(ctx context.Context, req *http.Request, result interface{}, closeBody bool) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	defer func() {
		if closeBody {
			// Close the body to let the Transport reuse the connection
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
		}
	}()

	response := newResponse(resp)

	err = checkResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if result != nil {
		if w, ok := result.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = xml.NewDecoder(resp.Body).Decode(result)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err
}

func (c *Client) send(ctx context.Context, opt *sendOptions) (resp *Response, err error) {
	req, err := c.newRequest(ctx, opt.baseURL, opt.uri, opt.method, opt.body, opt.optQuery, opt.optHeader)
	if err != nil {
		return
	}

	resp, err = c.doAPI(ctx, req, opt.result, !opt.disableCloseBody)
	return
}

type service struct {
	client *Client
}

func (s *service) Post(ctx context.Context, name string, body interface{}) (*Response, error) {
	u := fmt.Sprintf("/%s", encodeURIComponent(name))
	sendOpt := sendOptions{
		baseURL: s.client.BaseUrl,
		uri:     u,
		method:  http.MethodPost,
		body:    body,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return resp, err
}

// 检查 response 是否是出错时的返回的 response
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

// NewClient init COS client
func NewClient() *Client {
	urlStr, _ := url.Parse("https://" + global.GVA_CONFIG.TencentCOS.Bucket + ".cos." + global.GVA_CONFIG.TencentCOS.Region + ".myqcloud.com")
	var httpClient *http.Client = &http.Client{}
	c := &Client{
		client:    httpClient,
		UserAgent: userAgent,
		BaseUrl:   urlStr,
	}
	c.common.client = c
	return c
}
