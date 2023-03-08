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

type HkActivityCollectApi struct {
}

var hkActivityCollectService = service.ServiceGroupApp.AppServiceGroup.Community.HkActivityCollectService

// CreateHkActivityCollect 创建HkActivityCollect
// @Tags HkActivityCollect
// @Summary 创建HkActivityCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivityCollect true "创建HkActivityCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityCollect/createHkActivityCollect [post]
func (hkActivityCollectApi *HkActivityCollectApi) CreateHkActivityCollect(c *gin.Context) {
	var hkActivityCollect community.HkActivityCollect
	err := c.ShouldBindJSON(&hkActivityCollect)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityCollectService.CreateHkActivityCollect(hkActivityCollect); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkActivityCollect 删除HkActivityCollect
// @Tags HkActivityCollect
// @Summary 删除HkActivityCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivityCollect true "删除HkActivityCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivityCollect/deleteHkActivityCollect [delete]
func (hkActivityCollectApi *HkActivityCollectApi) DeleteHkActivityCollect(c *gin.Context) {
	var hkActivityCollect community.HkActivityCollect
	err := c.ShouldBindJSON(&hkActivityCollect)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityCollectService.DeleteHkActivityCollect(hkActivityCollect); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkActivityCollectByIds 批量删除HkActivityCollect
// @Tags HkActivityCollect
// @Summary 批量删除HkActivityCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkActivityCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkActivityCollect/deleteHkActivityCollectByIds [delete]
func (hkActivityCollectApi *HkActivityCollectApi) DeleteHkActivityCollectByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityCollectService.DeleteHkActivityCollectByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkActivityCollect 更新HkActivityCollect
// @Tags HkActivityCollect
// @Summary 更新HkActivityCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivityCollect true "更新HkActivityCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkActivityCollect/updateHkActivityCollect [put]
func (hkActivityCollectApi *HkActivityCollectApi) UpdateHkActivityCollect(c *gin.Context) {
	var hkActivityCollect community.HkActivityCollect
	err := c.ShouldBindJSON(&hkActivityCollect)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityCollectService.UpdateHkActivityCollect(hkActivityCollect); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkActivityCollect 用id查询HkActivityCollect
// @Tags HkActivityCollect
// @Summary 用id查询HkActivityCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkActivityCollect true "用id查询HkActivityCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkActivityCollect/findHkActivityCollect [get]
func (hkActivityCollectApi *HkActivityCollectApi) FindHkActivityCollect(c *gin.Context) {
	var hkActivityCollect community.HkActivityCollect
	err := c.ShouldBindQuery(&hkActivityCollect)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkActivityCollect, err := hkActivityCollectService.GetHkActivityCollect(hkActivityCollect.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkActivityCollect": rehkActivityCollect}, c)
	}
}

// GetHkActivityCollectList 分页获取HkActivityCollect列表
// @Tags HkActivityCollect
// @Summary 分页获取HkActivityCollect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkActivityCollectSearch true "分页获取HkActivityCollect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityCollect/getHkActivityCollectList [get]
func (hkActivityCollectApi *HkActivityCollectApi) GetHkActivityCollectList(c *gin.Context) {
	var pageInfo communityReq.HkActivityCollectSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkActivityCollectService.GetHkActivityCollectInfoList(pageInfo); err != nil {
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
