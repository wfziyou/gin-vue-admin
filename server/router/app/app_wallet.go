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
		router.POST("createUserRecharge", api.CreateUserRecharge)               // 新建用户充值
		routerWithoutRecord.GET("findUserRecharge", api.FindUserRecharge)       // 根据ID获取用户充值
		routerWithoutRecord.GET("getUserRechargeList", api.GetUserRechargeList) // 获取用户充值列表
		routerWithoutRecord.GET("findGoldBill", api.FindGoldBill)               // 根据ID获取金币账单
		routerWithoutRecord.GET("getGoldBillList", api.GetGoldBillList)         // 获取金币账单列表
		routerWithoutRecord.GET("findOrder", api.FindOrder)                     // 根据ID获取订单
		routerWithoutRecord.GET("getOrderList", api.GetOrderList)               // 获取订单列表
		routerWithoutRecord.GET("findUserBill", api.FindUserBill)               // 根据ID获取用户账单
		routerWithoutRecord.GET("getUserBillList", api.GetUserBillList)         // 获取用户账单列表
	}
}
