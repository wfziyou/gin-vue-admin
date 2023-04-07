package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	generalRouterWithoutRecord := appRouter.Group("general")
	var generalApi = v1.ApiGroupApp.AppApiGroup.GeneralApi
	{
		generalRouterWithoutRecord.GET("findProtocol", generalApi.FindProtocol)             //用id查询协议
		generalRouterWithoutRecord.GET("findProtocolByName", generalApi.FindProtocolByName) //用名字查询协议
	}
}
