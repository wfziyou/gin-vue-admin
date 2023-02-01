package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkUserBrowsingHistoryRouter struct {
}

// InitHkUserBrowsingHistoryRouter 初始化 HkUserBrowsingHistory 路由信息
func (s *HkUserBrowsingHistoryRouter) InitHkUserBrowsingHistoryRouter(Router *gin.RouterGroup) {
	hkUserBrowsingHistoryRouter := Router.Group("hkUserBrowsingHistory").Use(middleware.OperationRecord())
	hkUserBrowsingHistoryRouterWithoutRecord := Router.Group("hkUserBrowsingHistory")
	var hkUserBrowsingHistoryApi = v1.ApiGroupApp.GeneralApiGroup.HkUserBrowsingHistoryApi
	{
		hkUserBrowsingHistoryRouter.POST("createHkUserBrowsingHistory", hkUserBrowsingHistoryApi.CreateHkUserBrowsingHistory)   // 新建HkUserBrowsingHistory
		hkUserBrowsingHistoryRouter.DELETE("deleteHkUserBrowsingHistory", hkUserBrowsingHistoryApi.DeleteHkUserBrowsingHistory) // 删除HkUserBrowsingHistory
		hkUserBrowsingHistoryRouter.DELETE("deleteHkUserBrowsingHistoryByIds", hkUserBrowsingHistoryApi.DeleteHkUserBrowsingHistoryByIds) // 批量删除HkUserBrowsingHistory
		hkUserBrowsingHistoryRouter.PUT("updateHkUserBrowsingHistory", hkUserBrowsingHistoryApi.UpdateHkUserBrowsingHistory)    // 更新HkUserBrowsingHistory
	}
	{
		hkUserBrowsingHistoryRouterWithoutRecord.GET("findHkUserBrowsingHistory", hkUserBrowsingHistoryApi.FindHkUserBrowsingHistory)        // 根据ID获取HkUserBrowsingHistory
		hkUserBrowsingHistoryRouterWithoutRecord.GET("getHkUserBrowsingHistoryList", hkUserBrowsingHistoryApi.GetHkUserBrowsingHistoryList)  // 获取HkUserBrowsingHistory列表
	}
}
