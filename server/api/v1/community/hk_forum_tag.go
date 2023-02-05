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

type HkForumTagApi struct {
}

var hkForumTagService = service.ServiceGroupApp.CommunityServiceGroup.HkForumTagService

// CreateHkForumTag 创建HkForumTag
// @Tags HkForumTag
// @Summary 创建HkForumTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumTag true "创建HkForumTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTag/createHkForumTag [post]
func (hkForumTagApi *HkForumTagApi) CreateHkForumTag(c *gin.Context) {
	var hkForumTag community.HkForumTag
	err := c.ShouldBindJSON(&hkForumTag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTagService.CreateHkForumTag(hkForumTag); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkForumTag 删除HkForumTag
// @Tags HkForumTag
// @Summary 删除HkForumTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumTag true "删除HkForumTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTag/deleteHkForumTag [delete]
func (hkForumTagApi *HkForumTagApi) DeleteHkForumTag(c *gin.Context) {
	var hkForumTag community.HkForumTag
	err := c.ShouldBindJSON(&hkForumTag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTagService.DeleteHkForumTag(hkForumTag); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkForumTagByIds 批量删除HkForumTag
// @Tags HkForumTag
// @Summary 批量删除HkForumTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkForumTag/deleteHkForumTagByIds [delete]
func (hkForumTagApi *HkForumTagApi) DeleteHkForumTagByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTagService.DeleteHkForumTagByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkForumTag 更新HkForumTag
// @Tags HkForumTag
// @Summary 更新HkForumTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumTag true "更新HkForumTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumTag/updateHkForumTag [put]
func (hkForumTagApi *HkForumTagApi) UpdateHkForumTag(c *gin.Context) {
	var hkForumTag community.HkForumTag
	err := c.ShouldBindJSON(&hkForumTag)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTagService.UpdateHkForumTag(hkForumTag); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkForumTag 用id查询HkForumTag
// @Tags HkForumTag
// @Summary 用id查询HkForumTag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkForumTag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumTag/findHkForumTag [get]
func (hkForumTagApi *HkForumTagApi) FindHkForumTag(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumTag, err := hkForumTagService.GetHkForumTag(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForumTag": rehkForumTag}, c)
	}
}

// GetHkForumTagList 分页获取HkForumTag列表
// @Tags HkForumTag
// @Summary 分页获取HkForumTag列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkForumTagSearch true "分页获取HkForumTag列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTag/getHkForumTagList [get]
func (hkForumTagApi *HkForumTagApi) GetHkForumTagList(c *gin.Context) {
	var pageInfo communityReq.HkForumTagSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForumTagService.GetHkForumTagInfoList(pageInfo); err != nil {
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
