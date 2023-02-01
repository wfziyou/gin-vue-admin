package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkForbiddenSpeakReasonRouter struct {
}

// InitHkForbiddenSpeakReasonRouter 初始化 HkForbiddenSpeakReason 路由信息
func (s *HkForbiddenSpeakReasonRouter) InitHkForbiddenSpeakReasonRouter(Router *gin.RouterGroup) {
	hkForbiddenSpeakReasonRouter := Router.Group("hkForbiddenSpeakReason").Use(middleware.OperationRecord())
	hkForbiddenSpeakReasonRouterWithoutRecord := Router.Group("hkForbiddenSpeakReason")
	var hkForbiddenSpeakReasonApi = v1.ApiGroupApp.CommunityApiGroup.HkForbiddenSpeakReasonApi
	{
		hkForbiddenSpeakReasonRouter.POST("createHkForbiddenSpeakReason", hkForbiddenSpeakReasonApi.CreateHkForbiddenSpeakReason)   // 新建HkForbiddenSpeakReason
		hkForbiddenSpeakReasonRouter.DELETE("deleteHkForbiddenSpeakReason", hkForbiddenSpeakReasonApi.DeleteHkForbiddenSpeakReason) // 删除HkForbiddenSpeakReason
		hkForbiddenSpeakReasonRouter.DELETE("deleteHkForbiddenSpeakReasonByIds", hkForbiddenSpeakReasonApi.DeleteHkForbiddenSpeakReasonByIds) // 批量删除HkForbiddenSpeakReason
		hkForbiddenSpeakReasonRouter.PUT("updateHkForbiddenSpeakReason", hkForbiddenSpeakReasonApi.UpdateHkForbiddenSpeakReason)    // 更新HkForbiddenSpeakReason
	}
	{
		hkForbiddenSpeakReasonRouterWithoutRecord.GET("findHkForbiddenSpeakReason", hkForbiddenSpeakReasonApi.FindHkForbiddenSpeakReason)        // 根据ID获取HkForbiddenSpeakReason
		hkForbiddenSpeakReasonRouterWithoutRecord.GET("getHkForbiddenSpeakReasonList", hkForbiddenSpeakReasonApi.GetHkForbiddenSpeakReasonList)  // 获取HkForbiddenSpeakReason列表
	}
}
