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

type HkUserCircleApplyApi struct {
}

var hkUserCircleApplyService = service.ServiceGroupApp.CommunityServiceGroup.HkUserCircleApplyService

// CreateHkUserCircleApply 创建HkUserCircleApply
// @Tags HkUserCircleApply
// @Summary 创建HkUserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUserCircleApply true "创建HkUserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserCircleApply/createHkUserCircleApply [post]
func (hkUserCircleApplyApi *HkUserCircleApplyApi) CreateHkUserCircleApply(c *gin.Context) {
	var hkUserCircleApply community.HkUserCircleApply
	err := c.ShouldBindJSON(&hkUserCircleApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserCircleApplyService.CreateHkUserCircleApply(hkUserCircleApply); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkUserCircleApply 删除HkUserCircleApply
// @Tags HkUserCircleApply
// @Summary 删除HkUserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUserCircleApply true "删除HkUserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserCircleApply/deleteHkUserCircleApply [delete]
func (hkUserCircleApplyApi *HkUserCircleApplyApi) DeleteHkUserCircleApply(c *gin.Context) {
	var hkUserCircleApply community.HkUserCircleApply
	err := c.ShouldBindJSON(&hkUserCircleApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserCircleApplyService.DeleteHkUserCircleApply(hkUserCircleApply); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkUserCircleApplyByIds 批量删除HkUserCircleApply
// @Tags HkUserCircleApply
// @Summary 批量删除HkUserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkUserCircleApply/deleteHkUserCircleApplyByIds [delete]
func (hkUserCircleApplyApi *HkUserCircleApplyApi) DeleteHkUserCircleApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserCircleApplyService.DeleteHkUserCircleApplyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkUserCircleApply 更新HkUserCircleApply
// @Tags HkUserCircleApply
// @Summary 更新HkUserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUserCircleApply true "更新HkUserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserCircleApply/updateHkUserCircleApply [put]
func (hkUserCircleApplyApi *HkUserCircleApplyApi) UpdateHkUserCircleApply(c *gin.Context) {
	var hkUserCircleApply community.HkUserCircleApply
	err := c.ShouldBindJSON(&hkUserCircleApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserCircleApplyService.UpdateHkUserCircleApply(hkUserCircleApply); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkUserCircleApply 用id查询HkUserCircleApply
// @Tags HkUserCircleApply
// @Summary 用id查询HkUserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkUserCircleApply true "用id查询HkUserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserCircleApply/findHkUserCircleApply [get]
func (hkUserCircleApplyApi *HkUserCircleApplyApi) FindHkUserCircleApply(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkUserCircleApply, err := hkUserCircleApplyService.GetHkUserCircleApply(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkUserCircleApply": rehkUserCircleApply}, c)
	}
}

// GetHkUserCircleApplyList 分页获取HkUserCircleApply列表
// @Tags HkUserCircleApply
// @Summary 分页获取HkUserCircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkUserCircleApplySearch true "分页获取HkUserCircleApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserCircleApply/getHkUserCircleApplyList [get]
func (hkUserCircleApplyApi *HkUserCircleApplyApi) GetHkUserCircleApplyList(c *gin.Context) {
	var pageInfo communityReq.HkUserCircleApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkUserCircleApplyService.GetHkUserCircleApplyInfoList(pageInfo); err != nil {
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
