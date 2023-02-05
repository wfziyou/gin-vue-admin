package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HkUserBrowsingHistoryApi struct {
}

var hkUserBrowsingHistoryService = service.ServiceGroupApp.GeneralServiceGroup.HkUserBrowsingHistoryService

// CreateHkUserBrowsingHistory 创建HkUserBrowsingHistory
// @Tags HkUserBrowsingHistory
// @Summary 创建HkUserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkUserBrowsingHistory true "创建HkUserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserBrowsingHistory/createHkUserBrowsingHistory [post]
func (hkUserBrowsingHistoryApi *HkUserBrowsingHistoryApi) CreateHkUserBrowsingHistory(c *gin.Context) {
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

// DeleteHkUserBrowsingHistory 删除HkUserBrowsingHistory
// @Tags HkUserBrowsingHistory
// @Summary 删除HkUserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkUserBrowsingHistory true "删除HkUserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserBrowsingHistory/deleteHkUserBrowsingHistory [delete]
func (hkUserBrowsingHistoryApi *HkUserBrowsingHistoryApi) DeleteHkUserBrowsingHistory(c *gin.Context) {
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

// DeleteHkUserBrowsingHistoryByIds 批量删除HkUserBrowsingHistory
// @Tags HkUserBrowsingHistory
// @Summary 批量删除HkUserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkUserBrowsingHistory/deleteHkUserBrowsingHistoryByIds [delete]
func (hkUserBrowsingHistoryApi *HkUserBrowsingHistoryApi) DeleteHkUserBrowsingHistoryByIds(c *gin.Context) {
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

// UpdateHkUserBrowsingHistory 更新HkUserBrowsingHistory
// @Tags HkUserBrowsingHistory
// @Summary 更新HkUserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkUserBrowsingHistory true "更新HkUserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserBrowsingHistory/updateHkUserBrowsingHistory [put]
func (hkUserBrowsingHistoryApi *HkUserBrowsingHistoryApi) UpdateHkUserBrowsingHistory(c *gin.Context) {
	var hkUserBrowsingHistory general.HkUserBrowsingHistory
	err := c.ShouldBindJSON(&hkUserBrowsingHistory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserBrowsingHistoryService.UpdateHkUserBrowsingHistory(hkUserBrowsingHistory); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkUserBrowsingHistory 用id查询HkUserBrowsingHistory
// @Tags HkUserBrowsingHistory
// @Summary 用id查询HkUserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query general.HkUserBrowsingHistory true "用id查询HkUserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserBrowsingHistory/findHkUserBrowsingHistory [get]
func (hkUserBrowsingHistoryApi *HkUserBrowsingHistoryApi) FindHkUserBrowsingHistory(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkUserBrowsingHistory, err := hkUserBrowsingHistoryService.GetHkUserBrowsingHistory(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkUserBrowsingHistory": rehkUserBrowsingHistory}, c)
	}
}

// GetHkUserBrowsingHistoryList 分页获取HkUserBrowsingHistory列表
// @Tags HkUserBrowsingHistory
// @Summary 分页获取HkUserBrowsingHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.HkUserBrowsingHistorySearch true "分页获取HkUserBrowsingHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserBrowsingHistory/getHkUserBrowsingHistoryList [get]
func (hkUserBrowsingHistoryApi *HkUserBrowsingHistoryApi) GetHkUserBrowsingHistoryList(c *gin.Context) {
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
