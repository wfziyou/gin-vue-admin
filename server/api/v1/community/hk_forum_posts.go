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

type HkForumPostsApi struct {
}

var hkForumPostsService = service.ServiceGroupApp.CommunityServiceGroup.HkForumPostsService

// CreateHkForumPosts 创建HkForumPosts
// @Tags HkForumPosts
// @Summary 创建HkForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumPosts true "创建HkForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumPosts/createHkForumPosts [post]
func (hkForumPostsApi *HkForumPostsApi) CreateHkForumPosts(c *gin.Context) {
	var hkForumPosts community.HkForumPosts
	err := c.ShouldBindJSON(&hkForumPosts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumPostsService.CreateHkForumPosts(hkForumPosts); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkForumPosts 删除HkForumPosts
// @Tags HkForumPosts
// @Summary 删除HkForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumPosts true "删除HkForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumPosts/deleteHkForumPosts [delete]
func (hkForumPostsApi *HkForumPostsApi) DeleteHkForumPosts(c *gin.Context) {
	var hkForumPosts community.HkForumPosts
	err := c.ShouldBindJSON(&hkForumPosts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumPostsService.DeleteHkForumPosts(hkForumPosts); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkForumPostsByIds 批量删除HkForumPosts
// @Tags HkForumPosts
// @Summary 批量删除HkForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkForumPosts/deleteHkForumPostsByIds [delete]
func (hkForumPostsApi *HkForumPostsApi) DeleteHkForumPostsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumPostsService.DeleteHkForumPostsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkForumPosts 更新HkForumPosts
// @Tags HkForumPosts
// @Summary 更新HkForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumPosts true "更新HkForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumPosts/updateHkForumPosts [put]
func (hkForumPostsApi *HkForumPostsApi) UpdateHkForumPosts(c *gin.Context) {
	var hkForumPosts community.HkForumPosts
	err := c.ShouldBindJSON(&hkForumPosts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumPostsService.UpdateHkForumPosts(hkForumPosts); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkForumPosts 用id查询HkForumPosts
// @Tags HkForumPosts
// @Summary 用id查询HkForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumPosts/findHkForumPosts [get]
func (hkForumPostsApi *HkForumPostsApi) FindHkForumPosts(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumPosts, err := hkForumPostsService.GetHkForumPosts(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForumPosts": rehkForumPosts}, c)
	}
}

// GetHkForumPostsList 分页获取HkForumPosts列表
// @Tags HkForumPosts
// @Summary 分页获取HkForumPosts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkForumPostsSearch true "分页获取HkForumPosts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumPosts/getHkForumPostsList [get]
func (hkForumPostsApi *HkForumPostsApi) GetHkForumPostsList(c *gin.Context) {
	var pageInfo communityReq.HkForumPostsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForumPostsService.GetHkForumPostsInfoList(pageInfo); err != nil {
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
