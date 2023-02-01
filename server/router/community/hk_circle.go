package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkCircleRouter struct {
}

// InitHkCircleRouter 初始化 HkCircle 路由信息
func (s *HkCircleRouter) InitHkCircleRouter(Router *gin.RouterGroup) {
	hkCircleRouter := Router.Group("hkCircle").Use(middleware.OperationRecord())
	hkCircleRouterWithoutRecord := Router.Group("hkCircle")
	var hkCircleApi = v1.ApiGroupApp.CommunityApiGroup.HkCircleApi
	{
		hkCircleRouter.POST("createHkCircle", hkCircleApi.CreateHkCircle)   // 新建HkCircle
		hkCircleRouter.DELETE("deleteHkCircle", hkCircleApi.DeleteHkCircle) // 删除HkCircle
		hkCircleRouter.DELETE("deleteHkCircleByIds", hkCircleApi.DeleteHkCircleByIds) // 批量删除HkCircle
		hkCircleRouter.PUT("updateHkCircle", hkCircleApi.UpdateHkCircle)    // 更新HkCircle
	}
	{
		hkCircleRouterWithoutRecord.GET("findHkCircle", hkCircleApi.FindHkCircle)        // 根据ID获取HkCircle
		hkCircleRouterWithoutRecord.GET("getHkCircleList", hkCircleApi.GetHkCircleList)  // 获取HkCircle列表
	}
}
