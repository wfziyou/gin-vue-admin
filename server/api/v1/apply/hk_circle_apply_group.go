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

type HkCircleApplyGroupApi struct {
}

var hkCircleApplyGroupService = service.ServiceGroupApp.ApplyServiceGroup.HkCircleApplyGroupService

// CreateHkCircleApplyGroup 创建HkCircleApplyGroup
// @Tags HkCircleApplyGroup
// @Summary 创建HkCircleApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkCircleApplyGroup true "创建HkCircleApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleApplyGroup/createHkCircleApplyGroup [post]
func (hkCircleApplyGroupApi *HkCircleApplyGroupApi) CreateHkCircleApplyGroup(c *gin.Context) {
	var hkCircleApplyGroup apply.HkCircleApplyGroup
	err := c.ShouldBindJSON(&hkCircleApplyGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleApplyGroupService.CreateHkCircleApplyGroup(hkCircleApplyGroup); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkCircleApplyGroup 删除HkCircleApplyGroup
// @Tags HkCircleApplyGroup
// @Summary 删除HkCircleApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkCircleApplyGroup true "删除HkCircleApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleApplyGroup/deleteHkCircleApplyGroup [delete]
func (hkCircleApplyGroupApi *HkCircleApplyGroupApi) DeleteHkCircleApplyGroup(c *gin.Context) {
	var hkCircleApplyGroup apply.HkCircleApplyGroup
	err := c.ShouldBindJSON(&hkCircleApplyGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleApplyGroupService.DeleteHkCircleApplyGroup(hkCircleApplyGroup); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkCircleApplyGroupByIds 批量删除HkCircleApplyGroup
// @Tags HkCircleApplyGroup
// @Summary 批量删除HkCircleApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkCircleApplyGroup/deleteHkCircleApplyGroupByIds [delete]
func (hkCircleApplyGroupApi *HkCircleApplyGroupApi) DeleteHkCircleApplyGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleApplyGroupService.DeleteHkCircleApplyGroupByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkCircleApplyGroup 更新HkCircleApplyGroup
// @Tags HkCircleApplyGroup
// @Summary 更新HkCircleApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkCircleApplyGroup true "更新HkCircleApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleApplyGroup/updateHkCircleApplyGroup [put]
func (hkCircleApplyGroupApi *HkCircleApplyGroupApi) UpdateHkCircleApplyGroup(c *gin.Context) {
	var hkCircleApplyGroup apply.HkCircleApplyGroup
	err := c.ShouldBindJSON(&hkCircleApplyGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleApplyGroupService.UpdateHkCircleApplyGroup(hkCircleApplyGroup); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkCircleApplyGroup 用id查询HkCircleApplyGroup
// @Tags HkCircleApplyGroup
// @Summary 用id查询HkCircleApplyGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkCircleApplyGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleApplyGroup/findHkCircleApplyGroup [get]
func (hkCircleApplyGroupApi *HkCircleApplyGroupApi) FindHkCircleApplyGroup(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleApplyGroup, err := hkCircleApplyGroupService.GetHkCircleApplyGroup(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircleApplyGroup": rehkCircleApplyGroup}, c)
	}
}

// GetHkCircleApplyGroupList 分页获取HkCircleApplyGroup列表
// @Tags HkCircleApplyGroup
// @Summary 分页获取HkCircleApplyGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.HkCircleApplyGroupSearch true "分页获取HkCircleApplyGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleApplyGroup/getHkCircleApplyGroupList [get]
func (hkCircleApplyGroupApi *HkCircleApplyGroupApi) GetHkCircleApplyGroupList(c *gin.Context) {
	var pageInfo applyReq.HkCircleApplyGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkCircleApplyGroupService.GetHkCircleApplyGroupInfoList(pageInfo); err != nil {
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
