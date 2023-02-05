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

type HkForumTagGroupApi struct {
}

var hkForumTagGroupService = service.ServiceGroupApp.CommunityServiceGroup.HkForumTagGroupService

// CreateHkForumTagGroup 创建HkForumTagGroup
// @Tags HkForumTagGroup
// @Summary 创建HkForumTagGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumTagGroup true "创建HkForumTagGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTagGroup/createHkForumTagGroup [post]
func (hkForumTagGroupApi *HkForumTagGroupApi) CreateHkForumTagGroup(c *gin.Context) {
	var hkForumTagGroup community.HkForumTagGroup
	err := c.ShouldBindJSON(&hkForumTagGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTagGroupService.CreateHkForumTagGroup(hkForumTagGroup); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkForumTagGroup 删除HkForumTagGroup
// @Tags HkForumTagGroup
// @Summary 删除HkForumTagGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumTagGroup true "删除HkForumTagGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTagGroup/deleteHkForumTagGroup [delete]
func (hkForumTagGroupApi *HkForumTagGroupApi) DeleteHkForumTagGroup(c *gin.Context) {
	var hkForumTagGroup community.HkForumTagGroup
	err := c.ShouldBindJSON(&hkForumTagGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTagGroupService.DeleteHkForumTagGroup(hkForumTagGroup); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkForumTagGroupByIds 批量删除HkForumTagGroup
// @Tags HkForumTagGroup
// @Summary 批量删除HkForumTagGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumTagGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkForumTagGroup/deleteHkForumTagGroupByIds [delete]
func (hkForumTagGroupApi *HkForumTagGroupApi) DeleteHkForumTagGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTagGroupService.DeleteHkForumTagGroupByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkForumTagGroup 更新HkForumTagGroup
// @Tags HkForumTagGroup
// @Summary 更新HkForumTagGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumTagGroup true "更新HkForumTagGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumTagGroup/updateHkForumTagGroup [put]
func (hkForumTagGroupApi *HkForumTagGroupApi) UpdateHkForumTagGroup(c *gin.Context) {
	var hkForumTagGroup community.HkForumTagGroup
	err := c.ShouldBindJSON(&hkForumTagGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTagGroupService.UpdateHkForumTagGroup(hkForumTagGroup); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkForumTagGroup 用id查询HkForumTagGroup
// @Tags HkForumTagGroup
// @Summary 用id查询HkForumTagGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkForumTagGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumTagGroup/findHkForumTagGroup [get]
func (hkForumTagGroupApi *HkForumTagGroupApi) FindHkForumTagGroup(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumTagGroup, err := hkForumTagGroupService.GetHkForumTagGroup(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForumTagGroup": rehkForumTagGroup}, c)
	}
}

// GetHkForumTagGroupList 分页获取HkForumTagGroup列表
// @Tags HkForumTagGroup
// @Summary 分页获取HkForumTagGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkForumTagGroupSearch true "分页获取HkForumTagGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTagGroup/getHkForumTagGroupList [get]
func (hkForumTagGroupApi *HkForumTagGroupApi) GetHkForumTagGroupList(c *gin.Context) {
	var pageInfo communityReq.HkForumTagGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForumTagGroupService.GetHkForumTagGroupInfoList(pageInfo); err != nil {
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
