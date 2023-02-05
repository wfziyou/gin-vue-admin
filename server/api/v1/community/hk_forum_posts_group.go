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

type HkForumPostsGroupApi struct {
}

var hkForumPostsGroupService = service.ServiceGroupApp.CommunityServiceGroup.HkForumPostsGroupService

// CreateHkForumPostsGroup 创建HkForumPostsGroup
// @Tags HkForumPostsGroup
// @Summary 创建HkForumPostsGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumPostsGroup true "创建HkForumPostsGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumPostsGroup/createHkForumPostsGroup [post]
func (hkForumPostsGroupApi *HkForumPostsGroupApi) CreateHkForumPostsGroup(c *gin.Context) {
	var hkForumPostsGroup community.HkForumPostsGroup
	err := c.ShouldBindJSON(&hkForumPostsGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumPostsGroupService.CreateHkForumPostsGroup(hkForumPostsGroup); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkForumPostsGroup 删除HkForumPostsGroup
// @Tags HkForumPostsGroup
// @Summary 删除HkForumPostsGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumPostsGroup true "删除HkForumPostsGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumPostsGroup/deleteHkForumPostsGroup [delete]
func (hkForumPostsGroupApi *HkForumPostsGroupApi) DeleteHkForumPostsGroup(c *gin.Context) {
	var hkForumPostsGroup community.HkForumPostsGroup
	err := c.ShouldBindJSON(&hkForumPostsGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumPostsGroupService.DeleteHkForumPostsGroup(hkForumPostsGroup); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkForumPostsGroupByIds 批量删除HkForumPostsGroup
// @Tags HkForumPostsGroup
// @Summary 批量删除HkForumPostsGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumPostsGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkForumPostsGroup/deleteHkForumPostsGroupByIds [delete]
func (hkForumPostsGroupApi *HkForumPostsGroupApi) DeleteHkForumPostsGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumPostsGroupService.DeleteHkForumPostsGroupByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkForumPostsGroup 更新HkForumPostsGroup
// @Tags HkForumPostsGroup
// @Summary 更新HkForumPostsGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumPostsGroup true "更新HkForumPostsGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumPostsGroup/updateHkForumPostsGroup [put]
func (hkForumPostsGroupApi *HkForumPostsGroupApi) UpdateHkForumPostsGroup(c *gin.Context) {
	var hkForumPostsGroup community.HkForumPostsGroup
	err := c.ShouldBindJSON(&hkForumPostsGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumPostsGroupService.UpdateHkForumPostsGroup(hkForumPostsGroup); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkForumPostsGroup 用id查询HkForumPostsGroup
// @Tags HkForumPostsGroup
// @Summary 用id查询HkForumPostsGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkForumPostsGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumPostsGroup/findHkForumPostsGroup [get]
func (hkForumPostsGroupApi *HkForumPostsGroupApi) FindHkForumPostsGroup(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumPostsGroup, err := hkForumPostsGroupService.GetHkForumPostsGroup(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForumPostsGroup": rehkForumPostsGroup}, c)
	}
}

// GetHkForumPostsGroupList 分页获取HkForumPostsGroup列表
// @Tags HkForumPostsGroup
// @Summary 分页获取HkForumPostsGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkForumPostsGroupSearch true "分页获取HkForumPostsGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumPostsGroup/getHkForumPostsGroupList [get]
func (hkForumPostsGroupApi *HkForumPostsGroupApi) GetHkForumPostsGroupList(c *gin.Context) {
	var pageInfo communityReq.HkForumPostsGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForumPostsGroupService.GetHkForumPostsGroupInfoList(pageInfo); err != nil {
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
