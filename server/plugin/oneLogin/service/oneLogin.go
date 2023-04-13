package service

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/oneLogin/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/oneLogin/model"
	utils "github.com/flipped-aurora/gin-vue-admin/server/plugin/oneLogin/utils"
	"golang.org/x/text/encoding/unicode"
	"math/rand"
	"time"
)

type OneLoginService struct {
}

func (e *OneLoginService) GetErrorDes(code string) (des string) {
	switch code {
	case "103000":
		return "返回成功"
	case "103101":
		return "签名错误"
	case "103113":
		return "token 格式错误"
	case "103119":
		return "appid 不存在"
	case "103133":
		return "sourceid 不合法（服务端需要使用调用 SDK 时使用的 appid 去换取号码）"
	case "103211":
		return "其他错误"
	case "103412":
		return "无效的请求"
	case "103414":
		return "参数校验异常"
	case "103511":
		return "请求 ip 不在社区配置的服务器白名单内"
	case "103811":
		return "token 为空"
	case "104201":
		return "token 失效或不存在"
	case "105018":
		return "用户权限不足（使用了本机号码校验的 token 去调用本接口）"
	case "105019":
		return "应用未授权（开发者社区未勾选能力）获取更多资讯返回码 返回码描述"
	case "105312":
		return "套餐已用完"
	case "105313":
		return "非法请"
	default:
		return "不知道错误"
	}
}

func (e *OneLoginService) LoginTokenValidate(token string) (result *model.LoginTokenValidateRespone, err error) {
	client, err := utils.NewClient()
	if err != nil {
		return nil, errors.New("创建httpClient失败")
	}
	if len(token) == 0 {
		return nil, errors.New("请输入token")
	}
	msgId := fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000000))

	sign, err := utils.SHA256withRSA(global.GlobalConfig.Appid+token, global.GlobalConfig.RsaPrivateKey)
	if err != nil {
		return nil, errors.New("签名错误")
	}

	TimeLocation, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().Unix()
	curTime := time.Unix(now, 0).In(TimeLocation).Format("20060102150405000")

	req := model.LoginTokenValidateRequest{}
	req.Version = utils.Version
	req.Msgid = msgId
	req.Systemtime = curTime
	req.Strictcheck = global.GlobalConfig.StrictCheck
	req.Appid = global.GlobalConfig.Appid
	req.Token = token
	req.Sign = sign
	req.Encryptionalgorithm = utils.Encryptionalgorithm

	result = &model.LoginTokenValidateRespone{}
	err = client.Post("/unisdk/rsapi/loginTokenValidate", req, result)
	if err != nil {
		return nil, errors.New("调用httpClient失败")
	}
	if result.ResultCode == "103000" {
		if utils.Encryptionalgorithm == "RSA" {
			data, err := hex.DecodeString(result.Msisdn)
			if err != nil {
				return nil, errors.New("返回结果解密错误")
			}
			output, err := utils.RsaDecrypt(data, "")
			if err != nil {
				return nil, errors.New("返回结果解密错误")
			}
			var decoder = unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
			telephoneData, _ := decoder.Bytes(output)
			telephone := string(telephoneData)
			result.Telephone = telephone
		}
	} else {
		return nil, errors.New(fmt.Sprintf("ErrorCode(%s):%s", result.ResultCode, e.GetErrorDes(result.ResultCode)))
	}
	return
}
