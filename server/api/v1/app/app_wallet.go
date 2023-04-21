package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WalletApi struct {
}

// FindHkGoldBill 用id查询HkGoldBill
// @Tags 钱包
// @Summary 用id查询HkGoldBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkGoldBill true "用id查询HkGoldBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkGoldBill/findHkGoldBill [get]
func (api *WalletApi) FindHkGoldBill(c *gin.Context) {
	var hkGoldBill community.HkGoldBill
	err := c.ShouldBindQuery(&hkGoldBill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkGoldBill, err := hkGoldBillService.GetHkGoldBill(hkGoldBill.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkGoldBill": rehkGoldBill}, c)
	}
}

// GetHkGoldBillList 分页获取HkGoldBill列表
// @Tags 钱包
// @Summary 分页获取HkGoldBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkGoldBillSearch true "分页获取HkGoldBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkGoldBill/getHkGoldBillList [get]
func (api *WalletApi) GetHkGoldBillList(c *gin.Context) {
	var pageInfo communityReq.HkGoldBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkGoldBillService.GetHkGoldBillInfoList(pageInfo); err != nil {
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

// FindHkOrder 用id查询HkOrder
// @Tags 钱包
// @Summary 用id查询HkOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkOrder true "用id查询HkOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkOrder/findHkOrder [get]
func (api *WalletApi) FindHkOrder(c *gin.Context) {
	var hkOrder community.HkOrder
	err := c.ShouldBindQuery(&hkOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkOrder, err := hkOrderService.GetHkOrder(hkOrder.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkOrder": rehkOrder}, c)
	}
}

// GetHkOrderList 分页获取HkOrder列表
// @Tags 钱包
// @Summary 分页获取HkOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkOrderSearch true "分页获取HkOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkOrder/getHkOrderList [get]
func (api *WalletApi) GetHkOrderList(c *gin.Context) {
	var pageInfo communityReq.HkOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkOrderService.GetHkOrderInfoList(pageInfo); err != nil {
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

// FindHkUserBill 用id查询HkUserBill
// @Tags 钱包
// @Summary 用id查询HkUserBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkUserBill true "用id查询HkUserBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserBill/findHkUserBill [get]
func (api *WalletApi) FindHkUserBill(c *gin.Context) {
	var hkUserBill community.HkUserBill
	err := c.ShouldBindQuery(&hkUserBill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkUserBill, err := hkUserBillService.GetHkUserBill(hkUserBill.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkUserBill": rehkUserBill}, c)
	}
}

// GetHkUserBillList 分页获取HkUserBill列表
// @Tags 钱包
// @Summary 分页获取HkUserBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkUserBillSearch true "分页获取HkUserBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserBill/getHkUserBillList [get]
func (api *WalletApi) GetHkUserBillList(c *gin.Context) {
	var pageInfo communityReq.HkUserBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkUserBillService.GetHkUserBillInfoList(pageInfo); err != nil {
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

// CreateHkUserRecharge 创建HkUserRecharge
// @Tags 钱包
// @Summary 创建HkUserRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUserRecharge true "创建HkUserRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserRecharge/createHkUserRecharge [post]
func (api *WalletApi) CreateHkUserRecharge(c *gin.Context) {
	var hkUserRecharge community.HkUserRecharge
	err := c.ShouldBindJSON(&hkUserRecharge)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserRechargeService.CreateHkUserRecharge(&hkUserRecharge); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// FindHkUserRecharge 用id查询HkUserRecharge
// @Tags 钱包
// @Summary 用id查询HkUserRecharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkUserRecharge true "用id查询HkUserRecharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserRecharge/findHkUserRecharge [get]
func (api *WalletApi) FindHkUserRecharge(c *gin.Context) {
	var hkUserRecharge community.HkUserRecharge
	err := c.ShouldBindQuery(&hkUserRecharge)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkUserRecharge, err := hkUserRechargeService.GetHkUserRecharge(hkUserRecharge.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkUserRecharge": rehkUserRecharge}, c)
	}
}

// GetHkUserRechargeList 分页获取HkUserRecharge列表
// @Tags 钱包
// @Summary 分页获取HkUserRecharge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkUserRechargeSearch true "分页获取HkUserRecharge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserRecharge/getHkUserRechargeList [get]
func (api *WalletApi) GetHkUserRechargeList(c *gin.Context) {
	var pageInfo communityReq.HkUserRechargeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkUserRechargeService.GetHkUserRechargeInfoList(pageInfo); err != nil {
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
