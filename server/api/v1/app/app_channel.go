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

type HkChannelApi struct {
}

var hkChannelService = service.ServiceGroupApp.AppServiceGroup.Community.HkChannelService

// CreateHkChannel 创建HkChannel
// @Tags HkChannel
// @Summary 创建HkChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkChannel true "创建HkChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkChannel/createHkChannel [post]
func (hkChannelApi *HkChannelApi) CreateHkChannel(c *gin.Context) {
	var hkChannel community.HkChannel
	err := c.ShouldBindJSON(&hkChannel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkChannelService.CreateHkChannel(hkChannel); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkChannel 删除HkChannel
// @Tags HkChannel
// @Summary 删除HkChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkChannel true "删除HkChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkChannel/deleteHkChannel [delete]
func (hkChannelApi *HkChannelApi) DeleteHkChannel(c *gin.Context) {
	var hkChannel community.HkChannel
	err := c.ShouldBindJSON(&hkChannel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkChannelService.DeleteHkChannel(hkChannel); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkChannelByIds 批量删除HkChannel
// @Tags HkChannel
// @Summary 批量删除HkChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkChannel/deleteHkChannelByIds [delete]
func (hkChannelApi *HkChannelApi) DeleteHkChannelByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkChannelService.DeleteHkChannelByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkChannel 更新HkChannel
// @Tags HkChannel
// @Summary 更新HkChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkChannel true "更新HkChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkChannel/updateHkChannel [put]
func (hkChannelApi *HkChannelApi) UpdateHkChannel(c *gin.Context) {
	var hkChannel community.HkChannel
	err := c.ShouldBindJSON(&hkChannel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkChannelService.UpdateHkChannel(hkChannel); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkChannel 用id查询HkChannel
// @Tags HkChannel
// @Summary 用id查询HkChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkChannel true "用id查询HkChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkChannel/findHkChannel [get]
func (hkChannelApi *HkChannelApi) FindHkChannel(c *gin.Context) {
	var hkChannel community.HkChannel
	err := c.ShouldBindQuery(&hkChannel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkChannel, err := hkChannelService.GetHkChannel(hkChannel.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkChannel": rehkChannel}, c)
	}
}

// GetHkChannelList 分页获取HkChannel列表
// @Tags HkChannel
// @Summary 分页获取HkChannel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkChannelSearch true "分页获取HkChannel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkChannel/getHkChannelList [get]
func (hkChannelApi *HkChannelApi) GetHkChannelList(c *gin.Context) {
	var pageInfo communityReq.HkChannelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkChannelService.GetHkChannelInfoList(pageInfo); err != nil {
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
