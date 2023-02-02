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

// CreateHkReport 创建HkReport
// @Tags App_Report
// @Summary 创建HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReport true "创建HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/report/createHkReport [post]
func (reportApi *ReportApi) CreateHkReport(c *gin.Context) {
	var hkReport community.HkReport
	err := c.ShouldBindJSON(&hkReport)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkReportService.CreateHkReport(hkReport); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkReport 删除HkReport
// @Tags App_Report
// @Summary 删除HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReport true "删除HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/report/deleteHkReport [delete]
func (reportApi *ReportApi) DeleteHkReport(c *gin.Context) {
	var hkReport community.HkReport
	err := c.ShouldBindJSON(&hkReport)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkReportService.DeleteHkReport(hkReport); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkReportByIds 批量删除HkReport
// @Tags App_Report
// @Summary 批量删除HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/report/deleteHkReportByIds [delete]
func (reportApi *ReportApi) DeleteHkReportByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkReportService.DeleteHkReportByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkReport 更新HkReport
// @Tags App_Report
// @Summary 更新HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReport true "更新HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/report/updateHkReport [put]
func (reportApi *ReportApi) UpdateHkReport(c *gin.Context) {
	var hkReport community.HkReport
	err := c.ShouldBindJSON(&hkReport)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkReportService.UpdateHkReport(hkReport); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkReport 用id查询HkReport
// @Tags App_Report
// @Summary 用id查询HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/report/findHkReport [get]
func (reportApi *ReportApi) FindHkReport(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkReport, err := hkReportService.GetHkReport(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkReport": rehkReport}, c)
	}
}

// GetHkReportList 分页获取HkReport列表
// @Tags App_Report
// @Summary 分页获取HkReport列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkReportSearch true "分页获取HkReport列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/report/getHkReportList [get]
func (reportApi *ReportApi) GetHkReportList(c *gin.Context) {
	var pageInfo communityReq.HkReportSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkReportService.GetHkReportInfoList(pageInfo); err != nil {
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

// CreateHkReportReason 创建HkReportReason
// @Tags App_Report
// @Summary 创建HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReportReason true "创建HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/report/createHkReportReason [post]
func (reportApi *ReportApi) CreateHkReportReason(c *gin.Context) {
	var hkReportReason community.HkReportReason
	err := c.ShouldBindJSON(&hkReportReason)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkReportReasonService.CreateHkReportReason(hkReportReason); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkReportReason 删除HkReportReason
// @Tags App_Report
// @Summary 删除HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReportReason true "删除HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/report/deleteHkReportReason [delete]
func (reportApi *ReportApi) DeleteHkReportReason(c *gin.Context) {
	var hkReportReason community.HkReportReason
	err := c.ShouldBindJSON(&hkReportReason)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkReportReasonService.DeleteHkReportReason(hkReportReason); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkReportReasonByIds 批量删除HkReportReason
// @Tags App_Report
// @Summary 批量删除HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/report/deleteHkReportReasonByIds [delete]
func (reportApi *ReportApi) DeleteHkReportReasonByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkReportReasonService.DeleteHkReportReasonByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkReportReason 更新HkReportReason
// @Tags App_Report
// @Summary 更新HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReportReason true "更新HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/report/updateHkReportReason [put]
func (reportApi *ReportApi) UpdateHkReportReason(c *gin.Context) {
	var hkReportReason community.HkReportReason
	err := c.ShouldBindJSON(&hkReportReason)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkReportReasonService.UpdateHkReportReason(hkReportReason); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkReportReason 用id查询HkReportReason
// @Tags App_Report
// @Summary 用id查询HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/report/findHkReportReason [get]
func (reportApi *ReportApi) FindHkReportReason(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkReportReason, err := hkReportReasonService.GetHkReportReason(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkReportReason": rehkReportReason}, c)
	}
}

// GetHkReportReasonList 分页获取HkReportReason列表
// @Tags App_Report
// @Summary 分页获取HkReportReason列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkReportReasonSearch true "分页获取HkReportReason列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/report/getHkReportReasonList [get]
func (reportApi *ReportApi) GetHkReportReasonList(c *gin.Context) {
	var pageInfo communityReq.HkReportReasonSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkReportReasonService.GetHkReportReasonInfoList(pageInfo); err != nil {
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
