package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api"
	"github.com/gin-gonic/gin"
)

type YunXinImRouter struct {
}

func (s *YunXinImRouter) InitYunXinImRouter(Router *gin.RouterGroup) {
	emailRouter := Router
	aliSmsApi := api.ApiGroupApp.AliSmsApi
	{
		emailRouter.POST("sendSms", aliSmsApi.SendAliSms) // 发送测试邮件
	}
}
