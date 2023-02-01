package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkUserRouter struct {
}

// InitHkUserRouter 初始化 HkUser 路由信息
func (s *HkUserRouter) InitHkUserRouter(Router *gin.RouterGroup) {
	hkUserRouter := Router.Group("hkUser").Use(middleware.OperationRecord())
	hkUserRouterWithoutRecord := Router.Group("hkUser")
	var hkUserApi = v1.ApiGroupApp.CommunityApiGroup.HkUserApi
	{
		hkUserRouter.POST("createHkUser", hkUserApi.CreateHkUser)   // 新建HkUser
		hkUserRouter.DELETE("deleteHkUser", hkUserApi.DeleteHkUser) // 删除HkUser
		hkUserRouter.DELETE("deleteHkUserByIds", hkUserApi.DeleteHkUserByIds) // 批量删除HkUser
		hkUserRouter.PUT("updateHkUser", hkUserApi.UpdateHkUser)    // 更新HkUser
	}
	{
		hkUserRouterWithoutRecord.GET("findHkUser", hkUserApi.FindHkUser)        // 根据ID获取HkUser
		hkUserRouterWithoutRecord.GET("getHkUserList", hkUserApi.GetHkUserList)  // 获取HkUser列表
	}
}
