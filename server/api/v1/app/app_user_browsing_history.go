package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserBrowsingHistoryApi struct {
}

/*************************************
浏览历史
**************************************/

// DeleteUserBrowsingHistory 删除浏览历史
// @Tags 浏览历史
// @Summary 删除浏览历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除浏览历史"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/userBrowsingHistory/deleteUserBrowsingHistory [delete]
func (userBrowsingHistoryApi *UserBrowsingHistoryApi) DeleteUserBrowsingHistory(c *gin.Context) {
	var req request.IdDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := appUserBrowsingHistoryService.DeleteUserBrowsingHistoryById(userId, req.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserBrowsingHistoryByIds 批量删除浏览历史
// @Tags 浏览历史
// @Summary 批量删除浏览历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除浏览历史"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/userBrowsingHistory/deleteUserBrowsingHistoryByIds [delete]
func (userBrowsingHistoryApi *UserBrowsingHistoryApi) DeleteUserBrowsingHistoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := appUserBrowsingHistoryService.DeleteUserBrowsingHistoryByIds(userId, IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// DeleteAllUserBrowsingHistory 删除所有浏览历史
// @Tags 浏览历史
// @Summary 删除所有浏览历史
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body generalReq.DeleteAllUserBrowsingHistoryReq true "删除所有浏览历史"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/userBrowsingHistory/deleteAllUserBrowsingHistory [delete]
func (userBrowsingHistoryApi *UserBrowsingHistoryApi) DeleteAllUserBrowsingHistory(c *gin.Context) {
	var req generalReq.DeleteAllUserBrowsingHistoryReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := appUserBrowsingHistoryService.DeletAlleUserBrowsingHistory(userId, req.Category); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetUserBrowsingHistoryList 分页获取浏览历史列表
// @Tags 浏览历史
// @Summary 分页获取浏览历史列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.UserBrowsingHistorySearch true "分页获取浏览历史列表"
// @Success 200 {object}  response.PageResult{List=[]general.UserBrowsingHistoryInfo,msg=string} "返回general.UserBrowsingHistoryInfo"
// @Router /app/userBrowsingHistory/getUserBrowsingHistoryList [get]
func (userBrowsingHistoryApi *UserBrowsingHistoryApi) GetUserBrowsingHistoryList(c *gin.Context) {
	var pageInfo generalReq.UserBrowsingHistorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := appUserBrowsingHistoryService.GetUserBrowsingHistoryInfoList(userId, pageInfo); err != nil {
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
