package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/general"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GeneralApi struct{}

// FindProtocol 用id查询协议
// @Tags APP_General
// @Summary 用id查询协议
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询协议"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/general/findProtocol [get]
func (generalApi *GeneralApi) FindProtocol(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkProtocol, err := hkProtocolService.GetHkProtocol(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkProtocol": rehkProtocol}, c)
	}
}

// GetProtocolList 分页获取Protocol列表
// @Tags APP_General
// @Summary 分页获取Protocol列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.ProtocolSearch true "分页获取Protocol列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/general/getProtocolList [get]
func (generalApi *GeneralApi) GetProtocolList(c *gin.Context) {
	var pageInfo appReq.ProtocolSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkProtocolService.AppGetHkProtocolInfoList(pageInfo); err != nil {
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

// CreateHkBugReport Bug反馈
// @Tags APP_General
// @Summary Bug反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkBugReport true "Bug反馈"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/general/createBugReport [post]
func (generalApi *GeneralApi) CreateHkBugReport(c *gin.Context) {
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

// FindBugReport 用id查询BugReport
// @Tags APP_General
// @Summary 用id查询BugReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询BugReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/general/findBugReport [get]
func (generalApi *GeneralApi) FindBugReport(c *gin.Context) {
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

// GetBugReportList 分页获取BugReport列表
// @Tags APP_General
// @Summary 分页获取BugReport列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.BugReportSearch true "分页获取BugReport列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/general/getBugReportList [get]
func (generalApi *GeneralApi) GetBugReportList(c *gin.Context) {
	var pageInfo appReq.BugReportSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkBugReportService.AppGetHkBugReportInfoList(pageInfo); err != nil {
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
