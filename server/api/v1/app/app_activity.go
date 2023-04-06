package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HkActivityApi struct {
}

var hkActivityService = service.ServiceGroupApp.AppServiceGroup.Community.HkActivityService

// CreateHkActivity 创建活动
// @Tags HkActivity
// @Summary 创建活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivity true "创建活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivity/createHkActivity [post]
func (hkActivityApi *HkActivityApi) CreateHkActivity(c *gin.Context) {
	var hkActivity community.HkActivity
	err := c.ShouldBindJSON(&hkActivity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityService.CreateHkActivity(hkActivity); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkActivity 删除活动
// @Tags HkActivity
// @Summary 删除活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivity true "删除活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivity/deleteHkActivity [delete]
func (hkActivityApi *HkActivityApi) DeleteHkActivity(c *gin.Context) {
	var hkActivity community.HkActivity
	err := c.ShouldBindJSON(&hkActivity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityService.DeleteHkActivity(hkActivity); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkActivityByIds 批量删除活动
// @Tags HkActivity
// @Summary 批量删除活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkActivity/deleteHkActivityByIds [delete]
func (hkActivityApi *HkActivityApi) DeleteHkActivityByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityService.DeleteHkActivityByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkActivity 更新活动
// @Tags HkActivity
// @Summary 更新活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivity true "更新活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkActivity/updateHkActivity [put]
func (hkActivityApi *HkActivityApi) UpdateHkActivity(c *gin.Context) {
	var hkActivity community.HkActivity
	err := c.ShouldBindJSON(&hkActivity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityService.UpdateHkActivity(hkActivity); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkActivity 用id查询活动详情
// @Tags HkActivity
// @Summary 用id查询活动详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkActivity true "用id查询活动详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkActivity/findHkActivity [get]
func (hkActivityApi *HkActivityApi) FindHkActivity(c *gin.Context) {
	var hkActivity community.HkActivity
	err := c.ShouldBindQuery(&hkActivity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkActivity, err := hkActivityService.GetHkActivity(hkActivity.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkActivity": rehkActivity}, c)
	}
}

// GetHkActivityList 分页获取推荐活动列表
// @Tags HkActivity
// @Summary 分页获取推荐活动列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkActivitySearch true "分页获取推荐活动列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivity/getHkActivityList [get]
func (hkActivityApi *HkActivityApi) GetHkActivityList(c *gin.Context) {
	var pageInfo communityReq.HkActivitySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkActivityService.GetHkActivityInfoList(pageInfo); err != nil {
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
