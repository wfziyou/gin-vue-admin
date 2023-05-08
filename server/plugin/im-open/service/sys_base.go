package service

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open/global"
	imReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open/model/response"
	utils "github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open/utils"
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
	config := &config.OpenIm{}
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

func (e *YunXinImService) UserCreateAction(req imReq.RegisterReq) (result *response.RegisterRsp, err error) {
	client, err := CreateClient(tea.String(global.GlobalConfig.Url), tea.String(global.GlobalConfig.AppKey), tea.String(global.GlobalConfig.AppSecret))
	if err != nil {
		return
	}
	req.Secret = global.GlobalConfig.AppSecret
	result = &response.RegisterRsp{}
	err = client.Post("/auth/user_register", req, result)
	return
}

func (e *YunXinImService) UserGetUinfosAction(req imReq.UserGetUinfosActionReq) (result *response.UserGetUinfosActionRsp, err error) {
	client, err := CreateClient(tea.String(global.GlobalConfig.Url), tea.String(global.GlobalConfig.AppKey), tea.String(global.GlobalConfig.AppSecret))
	if err != nil {
		return
	}
	if len(req.Accids) == 0 {
		return nil, errors.New("请输入im账号")
	}

	postData := url.Values{}
	if data, err := json.Marshal(req.Accids); err == nil {
		postData.Add("accids", string(data))
	}

	result = &response.UserGetUinfosActionRsp{}
	err = client.Post("/user/getUinfos.action", postData, result)
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
