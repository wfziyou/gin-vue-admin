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

type HkActivityClassifyApi struct {
}

var hkActivityClassifyService = service.ServiceGroupApp.AppServiceGroup.Community.HkActivityClassifyService

// CreateHkActivityClassify 创建HkActivityClassify
// @Tags HkActivityClassify
// @Summary 创建HkActivityClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivityClassify true "创建HkActivityClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityClassify/createHkActivityClassify [post]
func (hkActivityClassifyApi *HkActivityClassifyApi) CreateHkActivityClassify(c *gin.Context) {
	var hkActivityClassify community.HkActivityClassify
	err := c.ShouldBindJSON(&hkActivityClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityClassifyService.CreateHkActivityClassify(hkActivityClassify); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkActivityClassify 删除HkActivityClassify
// @Tags HkActivityClassify
// @Summary 删除HkActivityClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivityClassify true "删除HkActivityClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivityClassify/deleteHkActivityClassify [delete]
func (hkActivityClassifyApi *HkActivityClassifyApi) DeleteHkActivityClassify(c *gin.Context) {
	var hkActivityClassify community.HkActivityClassify
	err := c.ShouldBindJSON(&hkActivityClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityClassifyService.DeleteHkActivityClassify(hkActivityClassify); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkActivityClassifyByIds 批量删除HkActivityClassify
// @Tags HkActivityClassify
// @Summary 批量删除HkActivityClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkActivityClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkActivityClassify/deleteHkActivityClassifyByIds [delete]
func (hkActivityClassifyApi *HkActivityClassifyApi) DeleteHkActivityClassifyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityClassifyService.DeleteHkActivityClassifyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkActivityClassify 更新HkActivityClassify
// @Tags HkActivityClassify
// @Summary 更新HkActivityClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivityClassify true "更新HkActivityClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkActivityClassify/updateHkActivityClassify [put]
func (hkActivityClassifyApi *HkActivityClassifyApi) UpdateHkActivityClassify(c *gin.Context) {
	var hkActivityClassify community.HkActivityClassify
	err := c.ShouldBindJSON(&hkActivityClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityClassifyService.UpdateHkActivityClassify(hkActivityClassify); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkActivityClassify 用id查询HkActivityClassify
// @Tags HkActivityClassify
// @Summary 用id查询HkActivityClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkActivityClassify true "用id查询HkActivityClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkActivityClassify/findHkActivityClassify [get]
func (hkActivityClassifyApi *HkActivityClassifyApi) FindHkActivityClassify(c *gin.Context) {
	var hkActivityClassify community.HkActivityClassify
	err := c.ShouldBindQuery(&hkActivityClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkActivityClassify, err := hkActivityClassifyService.GetHkActivityClassify(hkActivityClassify.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkActivityClassify": rehkActivityClassify}, c)
	}
}

// GetHkActivityClassifyList 分页获取HkActivityClassify列表
// @Tags HkActivityClassify
// @Summary 分页获取HkActivityClassify列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkActivityClassifySearch true "分页获取HkActivityClassify列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityClassify/getHkActivityClassifyList [get]
func (hkActivityClassifyApi *HkActivityClassifyApi) GetHkActivityClassifyList(c *gin.Context) {
	var pageInfo communityReq.HkActivityClassifySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkActivityClassifyService.GetHkActivityClassifyInfoList(pageInfo); err != nil {
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
