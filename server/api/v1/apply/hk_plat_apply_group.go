package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HkPlatApplyGroupApi struct {
}

var hkPlatApplyGroupService = service.ServiceGroupApp.ApplyServiceGroup.HkPlatApplyGroupService

// CreateHkPlatApplyGroup 创建HkPlatApplyGroup
// @Tags HkPlatApplyGroup
// @Summary 创建HkPlatApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkPlatApplyGroup true "创建HkPlatApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkPlatApplyGroup/createHkPlatApplyGroup [post]
func (hkPlatApplyGroupApi *HkPlatApplyGroupApi) CreateHkPlatApplyGroup(c *gin.Context) {
	var hkPlatApplyGroup apply.HkPlatApplyGroup
	err := c.ShouldBindJSON(&hkPlatApplyGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkPlatApplyGroupService.CreateHkPlatApplyGroup(hkPlatApplyGroup); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkPlatApplyGroup 删除HkPlatApplyGroup
// @Tags HkPlatApplyGroup
// @Summary 删除HkPlatApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkPlatApplyGroup true "删除HkPlatApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkPlatApplyGroup/deleteHkPlatApplyGroup [delete]
func (hkPlatApplyGroupApi *HkPlatApplyGroupApi) DeleteHkPlatApplyGroup(c *gin.Context) {
	var hkPlatApplyGroup apply.HkPlatApplyGroup
	err := c.ShouldBindJSON(&hkPlatApplyGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkPlatApplyGroupService.DeleteHkPlatApplyGroup(hkPlatApplyGroup); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkPlatApplyGroupByIds 批量删除HkPlatApplyGroup
// @Tags HkPlatApplyGroup
// @Summary 批量删除HkPlatApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkPlatApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkPlatApplyGroup/deleteHkPlatApplyGroupByIds [delete]
func (hkPlatApplyGroupApi *HkPlatApplyGroupApi) DeleteHkPlatApplyGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkPlatApplyGroupService.DeleteHkPlatApplyGroupByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkPlatApplyGroup 更新HkPlatApplyGroup
// @Tags HkPlatApplyGroup
// @Summary 更新HkPlatApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkPlatApplyGroup true "更新HkPlatApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkPlatApplyGroup/updateHkPlatApplyGroup [put]
func (hkPlatApplyGroupApi *HkPlatApplyGroupApi) UpdateHkPlatApplyGroup(c *gin.Context) {
	var hkPlatApplyGroup apply.HkPlatApplyGroup
	err := c.ShouldBindJSON(&hkPlatApplyGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkPlatApplyGroupService.UpdateHkPlatApplyGroup(hkPlatApplyGroup); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkPlatApplyGroup 用id查询HkPlatApplyGroup
// @Tags HkPlatApplyGroup
// @Summary 用id查询HkPlatApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkPlatApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkPlatApplyGroup/findHkPlatApplyGroup [get]
func (hkPlatApplyGroupApi *HkPlatApplyGroupApi) FindHkPlatApplyGroup(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkPlatApplyGroup, err := hkPlatApplyGroupService.GetHkPlatApplyGroup(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkPlatApplyGroup": rehkPlatApplyGroup}, c)
	}
}

// GetHkPlatApplyGroupList 分页获取HkPlatApplyGroup列表
// @Tags HkPlatApplyGroup
// @Summary 分页获取HkPlatApplyGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.HkPlatApplyGroupSearch true "分页获取HkPlatApplyGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkPlatApplyGroup/getHkPlatApplyGroupList [get]
func (hkPlatApplyGroupApi *HkPlatApplyGroupApi) GetHkPlatApplyGroupList(c *gin.Context) {
	var pageInfo applyReq.HkPlatApplyGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkPlatApplyGroupService.GetHkPlatApplyGroupInfoList(pageInfo); err != nil {
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
