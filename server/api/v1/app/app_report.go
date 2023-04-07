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

type ReportApi struct {
}

/*************************************
举报
**************************************/

// CreateReport 举报
// @Tags 举报
// @Summary 举报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.Report true "举报"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/report/createReport [post]
func (reportApi *ReportApi) CreateReport(c *gin.Context) {
	var hkReport community.Report
	err := c.ShouldBindJSON(&hkReport)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	hkReport.UserId = utils.GetUserID(c)
	if err := appReportService.CreateReport(hkReport); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// FindReport 用id查询举报
// @Tags 举报
// @Summary 用id查询举报
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询举报"
// @Success 200 {object}  response.Response{data=community.Report,msg=string}  "返回community.Report"
// @Router /app/report/findReport [get]
func (reportApi *ReportApi) FindReport(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkReport, err := appReportService.GetReport(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rehkReport, c)
	}
}

// GetReportList 分页获取举报列表
// @Tags 举报
// @Summary 分页获取举报列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ReportSearch true "分页获取举报列表"
// @Success 200 {object}  response.PageResult{List=[]community.Report,msg=string} "返回community.Report"
// @Router /app/report/getReportList [get]
func (reportApi *ReportApi) GetReportList(c *gin.Context) {
	var pageInfo communityReq.ReportSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.UserId = utils.GetUserID(c)
	if list, total, err := appReportService.GetReportInfoList(pageInfo); err != nil {
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

// FindReportReason 用id查询举报原因
// @Tags 举报
// @Summary 用id查询举报原因
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询举报原因"
// @Success 200 {object}  response.Response{data=community.ReportReason,msg=string}  "返回community.ReportReason"
// @Router /app/report/findReportReason [get]
func (reportApi *ReportApi) FindReportReason(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkReportReason, err := appReportReasonService.GetReportReason(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rehkReportReason, c)
	}
}

// GetReportReasonList 分页获取举报原因列表
// @Tags 举报
// @Summary 分页获取举报原因列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ReportReasonSearch true "分页获取举报原因列表"
// @Success 200 {object}  response.PageResult{List=[]community.ReportReason,msg=string} "返回community.ReportReason"
// @Router /app/report/getReportReasonList [get]
func (reportApi *ReportApi) GetReportReasonList(c *gin.Context) {
	var pageInfo communityReq.ReportReasonSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appReportReasonService.GetReportReasonInfoList(pageInfo); err != nil {
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
