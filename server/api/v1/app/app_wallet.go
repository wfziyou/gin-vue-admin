package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WalletApi struct {
}

// GetProductGoldList 获取金币列表
// @Tags 钱包
// @Summary 获取金币列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object}  response.Response{data=[]community.ProductGoldInfo,msg=string} "返回[]community.ProductGoldInfo"
// @Router /app/wallet/getProductGoldList [get]
func (api *WalletApi) GetProductGoldList(c *gin.Context) {
	if list, err := hkProductGoldService.GetProductGoldInfoList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// GetPayTypeList 获取支付类型列表
// @Tags 钱包
// @Summary 获取支付类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object}  response.Response{data=[]community.PayTypeInfo,msg=string} "返回[]community.PayTypeInfo"
// @Router /app/wallet/getPayTypeList [get]
func (api *WalletApi) GetPayTypeList(c *gin.Context) {
	if list, err := payTypeService.GetPayTypeInfoList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// CreateOrder 创建订单
// @Tags 钱包
// @Summary 创建订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamCreateOrder true "创建订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/wallet/createOrder [post]
func (api *WalletApi) CreateOrder(c *gin.Context) {
	var req communityReq.ParamCreateOrder
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := hkOrderService.CreateOrder(userId, req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// GetGoldBillList 分页获取金币账单列表
// @Tags 钱包
// @Summary 分页获取金币账单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.GoldBillSearch true "分页获取金币账单列表"
// @Success 200 {object}  response.Response{data=[]community.GoldBill,msg=string} "返回[]community.GoldBill"
// @Router /app/wallet/getGoldBillList [get]
func (api *WalletApi) GetGoldBillList(c *gin.Context) {
	var pageInfo communityReq.GoldBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := hkGoldBillService.GetGoldBillInfoList(userId, pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// FindGoldBill 用id查询金币账单
// @Tags 钱包
// @Summary 用id查询金币账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询金币账单"
// @Success 200 {object} response.Response{data=community.GoldBill,msg=string}  "返回community.GoldBill"
// @Router /app/wallet/findGoldBill [get]
func (api *WalletApi) FindGoldBill(c *gin.Context) {
	var req request.IdSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if rehkGoldBill, err := hkGoldBillService.GetGoldBill(req.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(rehkGoldBill, "成功", c)
	}
}

// GetExtractTypeList 获取提现类型列表
// @Tags 钱包
// @Summary 获取提现类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object}  response.Response{data=[]community.ExtractType,msg=string} "返回[]community.PayTypeInfo"
// @Router /app/wallet/getExtractTypeList [get]
func (api *WalletApi) GetExtractTypeList(c *gin.Context) {
	if list, err := extractTypeService.GetExtractTypeInfoList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// CreateUserExtract 创建用户提现
// @Tags 钱包
// @Summary 创建用户提现
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamUserExtract true "创建用户提现"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/wallet/createUserExtract [post]
func (api *WalletApi) CreateUserExtract(c *gin.Context) {
	var req communityReq.ParamUserExtract
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := hkUserRechargeService.CreateUserRecharge(userId, req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// GetUserBillList 分页获取用户账单列表
// @Tags 钱包
// @Summary 分页获取用户账单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.UserBillSearch true "分页获取用户账单列表"
// @Success 200 {object}  response.Response{data=[]community.UserBill,msg=string} "返回[]community.UserBill"
// @Router /app/wallet/getUserBillList [get]
func (api *WalletApi) GetUserBillList(c *gin.Context) {
	var req communityReq.UserBillSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := hkUserBillService.GetUserBillInfoList(userId, req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// FindUserBill 用id查询用户账单
// @Tags 钱包
// @Summary 用id查询用户账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.UserBill true "用id查询用户账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Success 200 {object} response.Response{data=community.UserBill,msg=string}  "返回community.UserBill"
// @Router /app/wallet/findUserBill [get]
func (api *WalletApi) FindUserBill(c *gin.Context) {
	var hkUserBill community.UserBill
	err := c.ShouldBindQuery(&hkUserBill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkUserBill, err := hkUserBillService.GetUserBill(hkUserBill.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(rehkUserBill, "成功", c)
	}
}
