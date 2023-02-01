package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkForbiddenSpeakDurationRouter struct {
}

// InitHkForbiddenSpeakDurationRouter 初始化 HkForbiddenSpeakDuration 路由信息
func (s *HkForbiddenSpeakDurationRouter) InitHkForbiddenSpeakDurationRouter(Router *gin.RouterGroup) {
	hkForbiddenSpeakDurationRouter := Router.Group("hkForbiddenSpeakDuration").Use(middleware.OperationRecord())
	hkForbiddenSpeakDurationRouterWithoutRecord := Router.Group("hkForbiddenSpeakDuration")
	var hkForbiddenSpeakDurationApi = v1.ApiGroupApp.CommunityApiGroup.HkForbiddenSpeakDurationApi
	{
		hkForbiddenSpeakDurationRouter.POST("createHkForbiddenSpeakDuration", hkForbiddenSpeakDurationApi.CreateHkForbiddenSpeakDuration)   // 新建HkForbiddenSpeakDuration
		hkForbiddenSpeakDurationRouter.DELETE("deleteHkForbiddenSpeakDuration", hkForbiddenSpeakDurationApi.DeleteHkForbiddenSpeakDuration) // 删除HkForbiddenSpeakDuration
		hkForbiddenSpeakDurationRouter.DELETE("deleteHkForbiddenSpeakDurationByIds", hkForbiddenSpeakDurationApi.DeleteHkForbiddenSpeakDurationByIds) // 批量删除HkForbiddenSpeakDuration
		hkForbiddenSpeakDurationRouter.PUT("updateHkForbiddenSpeakDuration", hkForbiddenSpeakDurationApi.UpdateHkForbiddenSpeakDuration)    // 更新HkForbiddenSpeakDuration
	}
	{
		hkForbiddenSpeakDurationRouterWithoutRecord.GET("findHkForbiddenSpeakDuration", hkForbiddenSpeakDurationApi.FindHkForbiddenSpeakDuration)        // 根据ID获取HkForbiddenSpeakDuration
		hkForbiddenSpeakDurationRouterWithoutRecord.GET("getHkForbiddenSpeakDurationList", hkForbiddenSpeakDurationApi.GetHkForbiddenSpeakDurationList)  // 获取HkForbiddenSpeakDuration列表
	}
}
