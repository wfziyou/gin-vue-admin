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

type HkReportReasonApi struct {
}

var hkReportReasonService = service.ServiceGroupApp.CommunityServiceGroup.HkReportReasonService

// CreateHkReportReason 创建HkReportReason
// @Tags HkReportReason
// @Summary 创建HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReportReason true "创建HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkReportReason/createHkReportReason [post]
func (hkReportReasonApi *HkReportReasonApi) CreateHkReportReason(c *gin.Context) {
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
// @Tags HkReportReason
// @Summary 删除HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReportReason true "删除HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkReportReason/deleteHkReportReason [delete]
func (hkReportReasonApi *HkReportReasonApi) DeleteHkReportReason(c *gin.Context) {
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
// @Tags HkReportReason
// @Summary 批量删除HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkReportReason/deleteHkReportReasonByIds [delete]
func (hkReportReasonApi *HkReportReasonApi) DeleteHkReportReasonByIds(c *gin.Context) {
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
// @Tags HkReportReason
// @Summary 更新HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkReportReason true "更新HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkReportReason/updateHkReportReason [put]
func (hkReportReasonApi *HkReportReasonApi) UpdateHkReportReason(c *gin.Context) {
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
// @Tags HkReportReason
// @Summary 用id查询HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkReportReason/findHkReportReason [get]
func (hkReportReasonApi *HkReportReasonApi) FindHkReportReason(c *gin.Context) {
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
// @Tags HkReportReason
// @Summary 分页获取HkReportReason列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkReportReasonSearch true "分页获取HkReportReason列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkReportReason/getHkReportReasonList [get]
func (hkReportReasonApi *HkReportReasonApi) GetHkReportReasonList(c *gin.Context) {
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
