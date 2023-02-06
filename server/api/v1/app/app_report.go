package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ReportApi struct {
}

// CreateReport 举报
// @Tags App_Report
// @Summary 举报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReport true "举报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/report/createReport [post]
func (reportApi *ReportApi) CreateReport(c *gin.Context) {
	var hkReport community.HkReport
	err := c.ShouldBindJSON(&hkReport)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appReportService.CreateHkReport(hkReport); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// FindReport 用id查询Report
// @Tags App_Report
// @Summary 用id查询Report
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询Report"
// @Success 200 {object}  response.Response{data=community.HkReport,msg=string}  "返回community.HkReport"
// @Router /app/report/findReport [get]
func (reportApi *ReportApi) FindReport(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkReport, err := appReportService.GetHkReport(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkReport": rehkReport}, c)
	}
}

// GetReportList 分页获取Report列表
// @Tags App_Report
// @Summary 分页获取Report列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkReportSearch true "分页获取Report列表"
// @Success 200 {object}  response.PageResult{List=[]community.HkReport,msg=string} "返回community.HkReport"
// @Router /app/report/getReportList [get]
func (reportApi *ReportApi) GetReportList(c *gin.Context) {
	var pageInfo communityReq.HkReportSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appReportService.GetHkReportInfoList(pageInfo); err != nil {
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

// FindReportReason 用id查询ReportReason
// @Tags App_Report
// @Summary 用id查询ReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询ReportReason"
// @Success 200 {object}  response.Response{data=community.HkReportReason,msg=string}  "返回community.HkReportReason"
// @Router /app/report/findReportReason [get]
func (reportApi *ReportApi) FindReportReason(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkReportReason, err := appReportReasonService.GetHkReportReason(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkReportReason": rehkReportReason}, c)
	}
}

// GetReportReasonList 分页获取ReportReason列表
// @Tags App_Report
// @Summary 分页获取ReportReason列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkReportReasonSearch true "分页获取ReportReason列表"
// @Success 200 {object}  response.PageResult{List=[]community.HkReportReason,msg=string} "返回community.HkReportReason"
// @Router /app/report/getReportReasonList [get]
func (reportApi *ReportApi) GetReportReasonList(c *gin.Context) {
	var pageInfo communityReq.HkReportReasonSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appReportReasonService.GetHkReportReasonInfoList(pageInfo); err != nil {
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
