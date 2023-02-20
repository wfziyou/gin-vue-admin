package service

import (
	"errors"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im/model/response"
	utils "github.com/flipped-aurora/gin-vue-admin/server/plugin/im/utils"
)

type YunXinImService struct{}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(url *string, appKey *string) (_result *utils.HttpClient, _err error) {
	config := &config.YunXinIm{}
	if url != nil {
		config.Url = *url
	}
	if appKey != nil {
		config.AppKey = *appKey
	}
	_result, _err = utils.NewClient(config)
	return _result, _err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SendAliSms
//@description: 发送（阿里）短信
//@return: err error

// 模板使用json字符串 {"code":"xxx"} 对应你模板里面的变量key和value
func (e *YunXinImService) Register(req request.RegisterReq) (result *response.RegisterRsp, err error) {
	client, err := CreateClient(tea.String(global.GlobalConfig.Url), tea.String(global.GlobalConfig.AppKey))
	if err != nil {
		return
	}
	if len(req.Accid) == 0 {
		return nil, errors.New("请输入im账号")
	}

	_, err = client.Register(&req)
	if err != nil {
		return
	}
	return
}
