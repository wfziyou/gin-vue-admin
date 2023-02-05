package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserBrowsingHistoryRouter struct {
}

func (router *UserBrowsingHistoryRouter) InitUserBrowsingHistoryRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	userBrowsingHistoryRouter := appRouter.Group("userBrowsingHistory").Use(middleware.OperationRecord())
	userBrowsingHistoryRouterWithoutRecord := appRouter.Group("userBrowsingHistory")
	var userBrowsingHistoryApi = v1.ApiGroupApp.AppApiGroup.UserBrowsingHistoryApi
	{
		userBrowsingHistoryRouter.DELETE("deleteUserBrowsingHistory", userBrowsingHistoryApi.DeleteUserBrowsingHistory)           //删除UserBrowsingHistory
		userBrowsingHistoryRouter.DELETE("deleteUserBrowsingHistoryByIds", userBrowsingHistoryApi.DeleteUserBrowsingHistoryByIds) //批量删除UserBrowsingHistory
	}
	{
		userBrowsingHistoryRouterWithoutRecord.GET("getUserBrowsingHistoryList", userBrowsingHistoryApi.GetUserBrowsingHistoryList) //分页获取UserBrowsingHistory列表
	}
	return userBrowsingHistoryRouter
}