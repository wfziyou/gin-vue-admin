package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/general/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserBrowsingHistoryApi struct {
}

// CreateUserBrowsingHistory 创建UserBrowsingHistory
// @Tags APP_UserBrowsingHistory
// @Summary 创建UserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkUserBrowsingHistory true "创建UserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/userBrowsingHistory/createUserBrowsingHistory [post]
func (userBrowsingHistoryApi *UserBrowsingHistoryApi) CreateUserBrowsingHistory(c *gin.Context) {
	var hkUserBrowsingHistory general.HkUserBrowsingHistory
	err := c.ShouldBindJSON(&hkUserBrowsingHistory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserBrowsingHistoryService.CreateHkUserBrowsingHistory(hkUserBrowsingHistory); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserBrowsingHistory 删除UserBrowsingHistory
// @Tags APP_UserBrowsingHistory
// @Summary 删除UserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkUserBrowsingHistory true "删除UserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/userBrowsingHistory/deleteUserBrowsingHistory [delete]
func (userBrowsingHistoryApi *UserBrowsingHistoryApi) DeleteUserBrowsingHistory(c *gin.Context) {
	var hkUserBrowsingHistory general.HkUserBrowsingHistory
	err := c.ShouldBindJSON(&hkUserBrowsingHistory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserBrowsingHistoryService.DeleteHkUserBrowsingHistory(hkUserBrowsingHistory); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserBrowsingHistoryByIds 批量删除UserBrowsingHistory
// @Tags APP_UserBrowsingHistory
// @Summary 批量删除UserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/userBrowsingHistory/deleteUserBrowsingHistoryByIds [delete]
func (userBrowsingHistoryApi *UserBrowsingHistoryApi) DeleteUserBrowsingHistoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserBrowsingHistoryService.DeleteHkUserBrowsingHistoryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// GetUserBrowsingHistoryList 分页获取UserBrowsingHistory列表
// @Tags APP_UserBrowsingHistory
// @Summary 分页获取UserBrowsingHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.HkUserBrowsingHistorySearch true "分页获取UserBrowsingHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/userBrowsingHistory/getUserBrowsingHistoryList [get]
func (userBrowsingHistoryApi *UserBrowsingHistoryApi) GetUserBrowsingHistoryList(c *gin.Context) {
	var pageInfo generalReq.HkUserBrowsingHistorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkUserBrowsingHistoryService.GetHkUserBrowsingHistoryInfoList(pageInfo); err != nil {
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
