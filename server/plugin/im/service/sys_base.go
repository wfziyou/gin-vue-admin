package service

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im/global"
	imReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/im/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im/model/response"
	utils "github.com/flipped-aurora/gin-vue-admin/server/plugin/im/utils"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type YunXinImService struct{}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(url *string, appKey *string, appSecret *string) (_result *utils.HttpClient, _err error) {
	config := &config.YunXinIm{}
	if url != nil {
		config.Url = *url
	}
	if appKey != nil {
		config.AppKey = *appKey
	}
	if appSecret != nil {
		config.AppSecret = *appSecret
	}
	_result, _err = utils.NewClient(config)
	return _result, _err
}

func (e *YunXinImService) Register(req imReq.RegisterReq) (result *response.RegisterRsp, err error) {
	client, err := CreateClient(tea.String(global.GlobalConfig.Url), tea.String(global.GlobalConfig.AppKey), tea.String(global.GlobalConfig.AppSecret))
	if err != nil {
		return
	}
	if len(req.Accid) == 0 {
		return nil, errors.New("请输入im账号")
	}
	postData := url.Values{}
	if len(req.Accid) > 0 {
		postData.Add("accid", req.Accid)
	}
	if len(req.Token) > 0 {
		postData.Add("token", req.Token)
	}
	if len(req.Name) > 0 {
		postData.Add("name", req.Name)
	}
	if len(req.Icon) > 0 {
		postData.Add("icon", req.Icon)
	}
	if len(req.Sign) > 0 {
		postData.Add("sign", req.Sign)
	}
	if len(req.Email) > 0 {
		postData.Add("email", req.Email)
	}
	if len(req.Birth) > 0 {
		postData.Add("birth", req.Birth)
	}
	if len(req.Mobile) > 0 {
		postData.Add("mobile", req.Mobile)
	}
	postData.Add("gender", strconv.Itoa(req.Gender))
	if len(req.Ex) > 0 {
		postData.Add("ex", req.Ex)
	}
	result = &response.RegisterRsp{}
	err = client.Post("/user/create.action", postData, result)
	return
}
func (e *YunXinImService) UpdateAction(req imReq.UpdateActionReq) (result *response.ImResponse, err error) {
	client, err := CreateClient(tea.String(global.GlobalConfig.Url), tea.String(global.GlobalConfig.AppKey), tea.String(global.GlobalConfig.AppSecret))
	if err != nil {
		return
	}
	if len(req.Accid) == 0 {
		return nil, errors.New("请输入im账号")
	}
	postData := url.Values{}
	if len(req.Accid) > 0 {
		postData.Add("accid", req.Accid)
	}
	if len(req.Token) > 0 {
		postData.Add("token", req.Token)
	}
	result = &response.ImResponse{}
	err = client.Post("/user/update.action", postData, result)
	return
}

func getCheckSum(appSecret string, nonce string, curTime string) string {
	h := sha1.New()
	h.Write([]byte(appSecret + nonce + curTime))
	t := sha1.New()
	io.WriteString(t, appSecret+nonce+curTime)
	sha := hex.EncodeToString(h.Sum(nil))
	aa := fmt.Sprintf("%x", t.Sum(nil))
	fmt.Println(aa)
	return sha
}
func testHttp() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.netease.im/nimserver/user/create.action", strings.NewReader("accid=cjb"))
	if err != nil {
		// handle error
	}

	var AppKey, AppSecret, Nonce, CurTime, CheckSum string
	AppKey = "fa7661a7cb7b9db8f47fe0e10ac8bae5"
	AppSecret = "54699e9dd7a0"
	Nonce = strconv.Itoa(rand.Intn(10000))
	CurTime = strconv.FormatInt(time.Now().UTC().Unix(), 10)

	CheckSum = getCheckSum(AppSecret, Nonce, CurTime)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("AppKey", AppKey)
	req.Header.Set("Nonce", Nonce)
	req.Header.Set("CurTime", CurTime)
	req.Header.Set("CheckSum", CheckSum)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
