package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	// Version current go sdk version
	Encryptionalgorithm   = "RSA" //RSA,SM,Other
	Version               = "2.0"
	contentTypeJson       = "application/json"
	defaultServiceBaseURL = "https://onekey2.cmpassport.com"
	backServiceBaseURL    = "https://www.cmpassport.com"
)

type HttpClient struct {
	client  *http.Client
	BaseUrl string
}

func NewClient() (*HttpClient, error) {
	client := &HttpClient{}
	err := client.Init()
	return client, err
}
func (httpClient *HttpClient) Init() (_err error) {
	httpClient.client = &http.Client{}
	httpClient.BaseUrl = defaultServiceBaseURL
	return nil
}

type AppOneLoginService struct{}

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

func (httpClient *HttpClient) Post(name string, postData interface{}, result interface{}) (err error) {
	urlStr := httpClient.BaseUrl + name
	configData, _ := json.Marshal(postData)
	param := bytes.NewBuffer([]byte(configData))
	req, err := http.NewRequest(http.MethodPost, urlStr, param)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", contentTypeJson)

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
