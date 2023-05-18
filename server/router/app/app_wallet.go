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
		routerWithoutRecord.GET("getProductGoldList", api.GetProductGoldList) // 获取金币列表
		routerWithoutRecord.GET("getPayTypeList", api.GetPayTypeList)         // 获取支付类型列表
		router.POST("createOrder", api.CreateOrder)                           // 创建订单
		routerWithoutRecord.GET("getGoldBillList", api.GetGoldBillList)       // 获取金币账单列表
		routerWithoutRecord.GET("findGoldBill", api.FindGoldBill)             // 根据ID获取金币账单

		routerWithoutRecord.GET("getExtractTypeList", api.GetExtractTypeList) // 获取提现类型列表
		router.POST("createUserExtract", api.CreateUserExtract)               // 创建用户提现

		routerWithoutRecord.GET("getUserBillList", api.GetUserBillList) // 获取用户账单列表
		routerWithoutRecord.GET("findUserBill", api.FindUserBill)       // 根据ID获取用户账单
	}
}
