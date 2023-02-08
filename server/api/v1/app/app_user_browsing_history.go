package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserBrowsingHistoryApi struct {
}

// DeleteUserBrowsingHistory 删除UserBrowsingHistory
// @Tags APP_UserBrowsingHistory
// @Summary 删除UserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.UserBrowsingHistory true "删除UserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/userBrowsingHistory/deleteUserBrowsingHistory [delete]
func (userBrowsingHistoryApi *UserBrowsingHistoryApi) DeleteUserBrowsingHistory(c *gin.Context) {
	var hkUserBrowsingHistory general.UserBrowsingHistory
	err := c.ShouldBindJSON(&hkUserBrowsingHistory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUserBrowsingHistoryService.DeleteUserBrowsingHistory(hkUserBrowsingHistory); err != nil {
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
	if err := appUserBrowsingHistoryService.DeleteUserBrowsingHistoryByIds(IDS); err != nil {
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
// @Param data query generalReq.UserBrowsingHistorySearch true "分页获取UserBrowsingHistory列表"
// @Success 200 {object}  response.PageResult{List=[]general.UserBrowsingHistory,msg=string} "返回general.UserBrowsingHistory"
// @Router /app/userBrowsingHistory/getUserBrowsingHistoryList [get]
func (userBrowsingHistoryApi *UserBrowsingHistoryApi) GetUserBrowsingHistoryList(c *gin.Context) {
	var pageInfo generalReq.UserBrowsingHistorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appUserBrowsingHistoryService.GetUserBrowsingHistoryInfoList(pageInfo); err != nil {
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
