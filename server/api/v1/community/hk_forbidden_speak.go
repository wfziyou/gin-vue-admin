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

type HkForbiddenSpeakApi struct {
}

var hkForbiddenSpeakService = service.ServiceGroupApp.CommunityServiceGroup.HkForbiddenSpeakService

// CreateHkForbiddenSpeak 创建HkForbiddenSpeak
// @Tags HkForbiddenSpeak
// @Summary 创建HkForbiddenSpeak
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForbiddenSpeak true "创建HkForbiddenSpeak"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForbiddenSpeak/createHkForbiddenSpeak [post]
func (hkForbiddenSpeakApi *HkForbiddenSpeakApi) CreateHkForbiddenSpeak(c *gin.Context) {
	var hkForbiddenSpeak community.HkForbiddenSpeak
	err := c.ShouldBindJSON(&hkForbiddenSpeak)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForbiddenSpeakService.CreateHkForbiddenSpeak(hkForbiddenSpeak); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkForbiddenSpeak 删除HkForbiddenSpeak
// @Tags HkForbiddenSpeak
// @Summary 删除HkForbiddenSpeak
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForbiddenSpeak true "删除HkForbiddenSpeak"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForbiddenSpeak/deleteHkForbiddenSpeak [delete]
func (hkForbiddenSpeakApi *HkForbiddenSpeakApi) DeleteHkForbiddenSpeak(c *gin.Context) {
	var hkForbiddenSpeak community.HkForbiddenSpeak
	err := c.ShouldBindJSON(&hkForbiddenSpeak)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForbiddenSpeakService.DeleteHkForbiddenSpeak(hkForbiddenSpeak); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkForbiddenSpeakByIds 批量删除HkForbiddenSpeak
// @Tags HkForbiddenSpeak
// @Summary 批量删除HkForbiddenSpeak
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForbiddenSpeak"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkForbiddenSpeak/deleteHkForbiddenSpeakByIds [delete]
func (hkForbiddenSpeakApi *HkForbiddenSpeakApi) DeleteHkForbiddenSpeakByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForbiddenSpeakService.DeleteHkForbiddenSpeakByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkForbiddenSpeak 更新HkForbiddenSpeak
// @Tags HkForbiddenSpeak
// @Summary 更新HkForbiddenSpeak
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForbiddenSpeak true "更新HkForbiddenSpeak"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForbiddenSpeak/updateHkForbiddenSpeak [put]
func (hkForbiddenSpeakApi *HkForbiddenSpeakApi) UpdateHkForbiddenSpeak(c *gin.Context) {
	var hkForbiddenSpeak community.HkForbiddenSpeak
	err := c.ShouldBindJSON(&hkForbiddenSpeak)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForbiddenSpeakService.UpdateHkForbiddenSpeak(hkForbiddenSpeak); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkForbiddenSpeak 用id查询HkForbiddenSpeak
// @Tags HkForbiddenSpeak
// @Summary 用id查询HkForbiddenSpeak
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkForbiddenSpeak true "用id查询HkForbiddenSpeak"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForbiddenSpeak/findHkForbiddenSpeak [get]
func (hkForbiddenSpeakApi *HkForbiddenSpeakApi) FindHkForbiddenSpeak(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForbiddenSpeak, err := hkForbiddenSpeakService.GetHkForbiddenSpeak(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForbiddenSpeak": rehkForbiddenSpeak}, c)
	}
}

// GetHkForbiddenSpeakList 分页获取HkForbiddenSpeak列表
// @Tags HkForbiddenSpeak
// @Summary 分页获取HkForbiddenSpeak列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkForbiddenSpeakSearch true "分页获取HkForbiddenSpeak列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForbiddenSpeak/getHkForbiddenSpeakList [get]
func (hkForbiddenSpeakApi *HkForbiddenSpeakApi) GetHkForbiddenSpeakList(c *gin.Context) {
	var pageInfo communityReq.HkForbiddenSpeakSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForbiddenSpeakService.GetHkForbiddenSpeakInfoList(pageInfo); err != nil {
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
