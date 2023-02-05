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

type HkForumCommentApi struct {
}

var hkForumCommentService = service.ServiceGroupApp.CommunityServiceGroup.HkForumCommentService

// CreateHkForumComment 创建HkForumComment
// @Tags HkForumComment
// @Summary 创建HkForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumComment true "创建HkForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumComment/createHkForumComment [post]
func (hkForumCommentApi *HkForumCommentApi) CreateHkForumComment(c *gin.Context) {
	var hkForumComment community.HkForumComment
	err := c.ShouldBindJSON(&hkForumComment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumCommentService.CreateHkForumComment(hkForumComment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkForumComment 删除HkForumComment
// @Tags HkForumComment
// @Summary 删除HkForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumComment true "删除HkForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumComment/deleteHkForumComment [delete]
func (hkForumCommentApi *HkForumCommentApi) DeleteHkForumComment(c *gin.Context) {
	var hkForumComment community.HkForumComment
	err := c.ShouldBindJSON(&hkForumComment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumCommentService.DeleteHkForumComment(hkForumComment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkForumCommentByIds 批量删除HkForumComment
// @Tags HkForumComment
// @Summary 批量删除HkForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkForumComment/deleteHkForumCommentByIds [delete]
func (hkForumCommentApi *HkForumCommentApi) DeleteHkForumCommentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumCommentService.DeleteHkForumCommentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkForumComment 更新HkForumComment
// @Tags HkForumComment
// @Summary 更新HkForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumComment true "更新HkForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumComment/updateHkForumComment [put]
func (hkForumCommentApi *HkForumCommentApi) UpdateHkForumComment(c *gin.Context) {
	var hkForumComment community.HkForumComment
	err := c.ShouldBindJSON(&hkForumComment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumCommentService.UpdateHkForumComment(hkForumComment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkForumComment 用id查询HkForumComment
// @Tags HkForumComment
// @Summary 用id查询HkForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkForumComment true "用id查询HkForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumComment/findHkForumComment [get]
func (hkForumCommentApi *HkForumCommentApi) FindHkForumComment(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumComment, err := hkForumCommentService.GetHkForumComment(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForumComment": rehkForumComment}, c)
	}
}

// GetHkForumCommentList 分页获取HkForumComment列表
// @Tags HkForumComment
// @Summary 分页获取HkForumComment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkForumCommentSearch true "分页获取HkForumComment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumComment/getHkForumCommentList [get]
func (hkForumCommentApi *HkForumCommentApi) GetHkForumCommentList(c *gin.Context) {
	var pageInfo communityReq.HkForumCommentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForumCommentService.GetHkForumCommentInfoList(pageInfo); err != nil {
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
