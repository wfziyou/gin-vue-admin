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

type HkBugReportApi struct {
}

var hkBugReportService = service.ServiceGroupApp.GeneralServiceGroup.HkBugReportService

// CreateHkBugReport 创建HkBugReport
// @Tags HkBugReport
// @Summary 创建HkBugReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkBugReport true "创建HkBugReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkBugReport/createHkBugReport [post]
func (hkBugReportApi *HkBugReportApi) CreateHkBugReport(c *gin.Context) {
	var hkBugReport general.HkBugReport
	err := c.ShouldBindJSON(&hkBugReport)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkBugReportService.CreateHkBugReport(hkBugReport); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkBugReport 删除HkBugReport
// @Tags HkBugReport
// @Summary 删除HkBugReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkBugReport true "删除HkBugReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkBugReport/deleteHkBugReport [delete]
func (hkBugReportApi *HkBugReportApi) DeleteHkBugReport(c *gin.Context) {
	var hkBugReport general.HkBugReport
	err := c.ShouldBindJSON(&hkBugReport)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkBugReportService.DeleteHkBugReport(hkBugReport); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkBugReportByIds 批量删除HkBugReport
// @Tags HkBugReport
// @Summary 批量删除HkBugReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkBugReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkBugReport/deleteHkBugReportByIds [delete]
func (hkBugReportApi *HkBugReportApi) DeleteHkBugReportByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkBugReportService.DeleteHkBugReportByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkBugReport 更新HkBugReport
// @Tags HkBugReport
// @Summary 更新HkBugReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkBugReport true "更新HkBugReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkBugReport/updateHkBugReport [put]
func (hkBugReportApi *HkBugReportApi) UpdateHkBugReport(c *gin.Context) {
	var hkBugReport general.HkBugReport
	err := c.ShouldBindJSON(&hkBugReport)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkBugReportService.UpdateHkBugReport(hkBugReport); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkBugReport 用id查询HkBugReport
// @Tags HkBugReport
// @Summary 用id查询HkBugReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkBugReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkBugReport/findHkBugReport [get]
func (hkBugReportApi *HkBugReportApi) FindHkBugReport(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkBugReport, err := hkBugReportService.GetHkBugReport(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkBugReport": rehkBugReport}, c)
	}
}

// GetHkBugReportList 分页获取HkBugReport列表
// @Tags HkBugReport
// @Summary 分页获取HkBugReport列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.HkBugReportSearch true "分页获取HkBugReport列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkBugReport/getHkBugReportList [get]
func (hkBugReportApi *HkBugReportApi) GetHkBugReportList(c *gin.Context) {
	var pageInfo generalReq.HkBugReportSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkBugReportService.GetHkBugReportInfoList(pageInfo); err != nil {
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
