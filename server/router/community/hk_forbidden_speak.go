package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkForbiddenSpeakRouter struct {
}

// InitHkForbiddenSpeakRouter 初始化 HkForbiddenSpeak 路由信息
func (s *HkForbiddenSpeakRouter) InitHkForbiddenSpeakRouter(Router *gin.RouterGroup) {
	hkForbiddenSpeakRouter := Router.Group("hkForbiddenSpeak").Use(middleware.OperationRecord())
	hkForbiddenSpeakRouterWithoutRecord := Router.Group("hkForbiddenSpeak")
	var hkForbiddenSpeakApi = v1.ApiGroupApp.CommunityApiGroup.HkForbiddenSpeakApi
	{
		hkForbiddenSpeakRouter.POST("createHkForbiddenSpeak", hkForbiddenSpeakApi.CreateHkForbiddenSpeak)   // 新建HkForbiddenSpeak
		hkForbiddenSpeakRouter.DELETE("deleteHkForbiddenSpeak", hkForbiddenSpeakApi.DeleteHkForbiddenSpeak) // 删除HkForbiddenSpeak
		hkForbiddenSpeakRouter.DELETE("deleteHkForbiddenSpeakByIds", hkForbiddenSpeakApi.DeleteHkForbiddenSpeakByIds) // 批量删除HkForbiddenSpeak
		hkForbiddenSpeakRouter.PUT("updateHkForbiddenSpeak", hkForbiddenSpeakApi.UpdateHkForbiddenSpeak)    // 更新HkForbiddenSpeak
	}
	{
		hkForbiddenSpeakRouterWithoutRecord.GET("findHkForbiddenSpeak", hkForbiddenSpeakApi.FindHkForbiddenSpeak)        // 根据ID获取HkForbiddenSpeak
		hkForbiddenSpeakRouterWithoutRecord.GET("getHkForbiddenSpeakList", hkForbiddenSpeakApi.GetHkForbiddenSpeakList)  // 获取HkForbiddenSpeak列表
	}
}
