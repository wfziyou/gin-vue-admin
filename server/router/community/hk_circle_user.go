package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkCircleUserRouter struct {
}

// InitHkCircleUserRouter 初始化 HkCircleUser 路由信息
func (s *HkCircleUserRouter) InitHkCircleUserRouter(Router *gin.RouterGroup) {
	hkCircleUserRouter := Router.Group("hkCircleUser").Use(middleware.OperationRecord())
	hkCircleUserRouterWithoutRecord := Router.Group("hkCircleUser")
	var hkCircleUserApi = v1.ApiGroupApp.CommunityApiGroup.HkCircleUserApi
	{
		hkCircleUserRouter.POST("createHkCircleUser", hkCircleUserApi.CreateHkCircleUser)   // 新建HkCircleUser
		hkCircleUserRouter.DELETE("deleteHkCircleUser", hkCircleUserApi.DeleteHkCircleUser) // 删除HkCircleUser
		hkCircleUserRouter.DELETE("deleteHkCircleUserByIds", hkCircleUserApi.DeleteHkCircleUserByIds) // 批量删除HkCircleUser
		hkCircleUserRouter.PUT("updateHkCircleUser", hkCircleUserApi.UpdateHkCircleUser)    // 更新HkCircleUser
	}
	{
		hkCircleUserRouterWithoutRecord.GET("findHkCircleUser", hkCircleUserApi.FindHkCircleUser)        // 根据ID获取HkCircleUser
		hkCircleUserRouterWithoutRecord.GET("getHkCircleUserList", hkCircleUserApi.GetHkCircleUserList)  // 获取HkCircleUser列表
	}
}
