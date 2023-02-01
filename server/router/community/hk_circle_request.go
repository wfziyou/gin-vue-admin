package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkCircleRequestRouter struct {
}

// InitHkCircleRequestRouter 初始化 HkCircleRequest 路由信息
func (s *HkCircleRequestRouter) InitHkCircleRequestRouter(Router *gin.RouterGroup) {
	hkCircleRequestRouter := Router.Group("hkCircleRequest").Use(middleware.OperationRecord())
	hkCircleRequestRouterWithoutRecord := Router.Group("hkCircleRequest")
	var hkCircleRequestApi = v1.ApiGroupApp.CommunityApiGroup.HkCircleRequestApi
	{
		hkCircleRequestRouter.POST("createHkCircleRequest", hkCircleRequestApi.CreateHkCircleRequest)   // 新建HkCircleRequest
		hkCircleRequestRouter.DELETE("deleteHkCircleRequest", hkCircleRequestApi.DeleteHkCircleRequest) // 删除HkCircleRequest
		hkCircleRequestRouter.DELETE("deleteHkCircleRequestByIds", hkCircleRequestApi.DeleteHkCircleRequestByIds) // 批量删除HkCircleRequest
		hkCircleRequestRouter.PUT("updateHkCircleRequest", hkCircleRequestApi.UpdateHkCircleRequest)    // 更新HkCircleRequest
	}
	{
		hkCircleRequestRouterWithoutRecord.GET("findHkCircleRequest", hkCircleRequestApi.FindHkCircleRequest)        // 根据ID获取HkCircleRequest
		hkCircleRequestRouterWithoutRecord.GET("getHkCircleRequestList", hkCircleRequestApi.GetHkCircleRequestList)  // 获取HkCircleRequest列表
	}
}
