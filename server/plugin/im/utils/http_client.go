package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im/config"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	// Version current go sdk version
	Version               = "0.7.19"
	userAgent             = "cos-go-sdk-v5/" + Version
	contentTypeJson       = "application/x-www-form-urlencoded"
	defaultServiceBaseURL = "http://service.cos.myqcloud.com"
)

type HttpClient struct {
	client *http.Client

	Host      string
	UserAgent string
	AppKey    string
	AppSecret string
	BaseUrl   string
}

// Response API 响应
type Response struct {
	*http.Response
}

func newResponse(resp *http.Response) *Response {
	return &Response{
		Response: resp,
	}
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

func NewClient(config *config.YunXinIm) (*HttpClient, error) {
	client := &HttpClient{}
	err := client.Init(config)
	return client, err
}
func (httpClient *HttpClient) Init(config *config.YunXinIm) (_err error) {
	if config == nil {
		return &ErrorResponse{Message: "请配置 YunXinIm"}
	}
	var baseUrl string
	if len(config.Url) > 0 {
		baseUrl = config.Url
	} else {
		baseUrl = defaultServiceBaseURL
	}
	httpClient.client = &http.Client{}
	httpClient.BaseUrl = baseUrl
	httpClient.AppKey = config.AppKey
	httpClient.AppSecret = config.AppSecret
	return nil
}

func (httpClient *HttpClient) Post(name string, postData url.Values, result interface{}) (err error) {
	urlStr := httpClient.BaseUrl + name
	req, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(postData.Encode()))
	if err != nil {
		return err
	}

	var Nonce, CurTime, CheckSum string
	Nonce = strconv.Itoa(rand.Intn(10000))
	CurTime = strconv.FormatInt(time.Now().UTC().Unix(), 10)
	CheckSum = httpClient.getCheckSum(httpClient.AppSecret, Nonce, CurTime)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("AppKey", httpClient.AppKey)
	req.Header.Set("Nonce", Nonce)
	req.Header.Set("CurTime", CurTime)
	req.Header.Set("CheckSum", CheckSum)

	resp, err := httpClient.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if c := resp.StatusCode; 200 != c {
		return errors.New("调用失败")
	}
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(string(body))
	if result != nil {
		err = json.NewDecoder(resp.Body).Decode(result)
		if err == io.EOF {
			err = nil // ignore EOF errors caused by empty response body
		}
	}
	return err
}

func (httpClient *HttpClient) getCheckSum(appSecret string, nonce string, curTime string) string {
	h := sha1.New()
	h.Write([]byte(appSecret + nonce + curTime))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
