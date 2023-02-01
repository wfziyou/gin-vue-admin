package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkCircleAddRequestRouter struct {
}

// InitHkCircleAddRequestRouter 初始化 HkCircleAddRequest 路由信息
func (s *HkCircleAddRequestRouter) InitHkCircleAddRequestRouter(Router *gin.RouterGroup) {
	hkCircleAddRequestRouter := Router.Group("hkCircleAddRequest").Use(middleware.OperationRecord())
	hkCircleAddRequestRouterWithoutRecord := Router.Group("hkCircleAddRequest")
	var hkCircleAddRequestApi = v1.ApiGroupApp.CommunityApiGroup.HkCircleAddRequestApi
	{
		hkCircleAddRequestRouter.POST("createHkCircleAddRequest", hkCircleAddRequestApi.CreateHkCircleAddRequest)   // 新建HkCircleAddRequest
		hkCircleAddRequestRouter.DELETE("deleteHkCircleAddRequest", hkCircleAddRequestApi.DeleteHkCircleAddRequest) // 删除HkCircleAddRequest
		hkCircleAddRequestRouter.DELETE("deleteHkCircleAddRequestByIds", hkCircleAddRequestApi.DeleteHkCircleAddRequestByIds) // 批量删除HkCircleAddRequest
		hkCircleAddRequestRouter.PUT("updateHkCircleAddRequest", hkCircleAddRequestApi.UpdateHkCircleAddRequest)    // 更新HkCircleAddRequest
	}
	{
		hkCircleAddRequestRouterWithoutRecord.GET("findHkCircleAddRequest", hkCircleAddRequestApi.FindHkCircleAddRequest)        // 根据ID获取HkCircleAddRequest
		hkCircleAddRequestRouterWithoutRecord.GET("getHkCircleAddRequestList", hkCircleAddRequestApi.GetHkCircleAddRequestList)  // 获取HkCircleAddRequest列表
	}
}
