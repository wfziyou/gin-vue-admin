package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HkReportApi struct {
}

var hkReportService = service.ServiceGroupApp.CommunityServiceGroup.HkReportService

// CreateHkReport 创建HkReport
// @Tags HkReport
// @Summary 创建HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReport true "创建HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkReport/createHkReport [post]
func (hkReportApi *HkReportApi) CreateHkReport(c *gin.Context) {
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
// @Tags HkReport
// @Summary 删除HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReport true "删除HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkReport/deleteHkReport [delete]
func (hkReportApi *HkReportApi) DeleteHkReport(c *gin.Context) {
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
// @Tags HkReport
// @Summary 批量删除HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkReport/deleteHkReportByIds [delete]
func (hkReportApi *HkReportApi) DeleteHkReportByIds(c *gin.Context) {
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
// @Tags HkReport
// @Summary 更新HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReport true "更新HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkReport/updateHkReport [put]
func (hkReportApi *HkReportApi) UpdateHkReport(c *gin.Context) {
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
// @Tags HkReport
// @Summary 用id查询HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkReport/findHkReport [get]
func (hkReportApi *HkReportApi) FindHkReport(c *gin.Context) {
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
// @Tags HkReport
// @Summary 分页获取HkReport列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkReportSearch true "分页获取HkReport列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkReport/getHkReportList [get]
func (hkReportApi *HkReportApi) GetHkReportList(c *gin.Context) {
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
