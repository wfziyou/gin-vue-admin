package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WalletRouter struct {
}

// InitWalletRouter 初始化 Wallet 路由信息
func (s *WalletRouter) InitWalletRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	router := appRouter.Group("wallet").Use(middleware.OperationRecord())
	routerWithoutRecord := appRouter.Group("wallet")
	var api = v1.ApiGroupApp.AppApiGroup.WalletApi
	{
		router.POST("createHkUserRecharge", api.CreateHkUserRecharge)               // 新建HkUserRecharge
		routerWithoutRecord.GET("findHkUserRecharge", api.FindHkUserRecharge)       // 根据ID获取HkUserRecharge
		routerWithoutRecord.GET("getHkUserRechargeList", api.GetHkUserRechargeList) // 获取HkUserRecharge列表
	}

	{
		routerWithoutRecord.GET("findHkGoldBill", api.FindHkGoldBill)       // 根据ID获取HkGoldBill
		routerWithoutRecord.GET("getHkGoldBillList", api.GetHkGoldBillList) // 获取HkGoldBill列表
	}

	{
		routerWithoutRecord.GET("findHkOrder", api.FindHkOrder)       // 根据ID获取HkOrder
		routerWithoutRecord.GET("getHkOrderList", api.GetHkOrderList) // 获取HkOrder列表
	}

	{
		routerWithoutRecord.GET("findHkUserBill", api.FindHkUserBill)       // 根据ID获取HkUserBill
		routerWithoutRecord.GET("getHkUserBillList", api.GetHkUserBillList) // 获取HkUserBill列表
	}
}
