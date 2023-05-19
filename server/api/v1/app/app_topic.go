package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TopicApi struct {
}

/*************************************
话题
**************************************/

// FindForumTopicGroup 用id查询话题分组
// @Tags 话题
// @Summary 用id查询话题分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询话题分组"
// @Success 200 {object}  response.Response{data=community.ForumTopicGroup,msg=string}  "返回community.ForumTopicGroup"
// @Router /app/topic/findForumTopicGroup [get]
func (topicApi *TopicApi) FindForumTopicGroup(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if forumTopicGroup, err := appForumTopicGroupService.GetForumTopicGroup(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(forumTopicGroup, c)
	}
}

// GetForumTopicGroupList 分页获取话题分组列表
// @Tags 话题
// @Summary 分页获取话题分组列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ForumTopicGroupSearch true "分页获取话题分组列表"
// @Success 200 {object}  response.PageResult{List=[]community.ForumTopicGroup,msg=string} "返回community.ForumTopicGroup"
// @Router /app/topic/getForumTopicGroupList [get]
func (topicApi *TopicApi) GetForumTopicGroupList(c *gin.Context) {
	var req communityReq.ForumTopicGroupSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appForumTopicGroupService.GetForumTopicGroupInfoList(req.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// GetForumTopicGroupListAll 获取话题分组列表
// @Tags 话题
// @Summary 获取话题分组列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ForumTopicGroupSearch true "获取话题分组列表"
// @Success 200 {object}  response.Response{data=[]community.ForumTopicGroup,msg=string} "返回community.ForumTopicGroup"
// @Router /app/topic/getForumTopicGroupListAll [get]
func (topicApi *TopicApi) GetForumTopicGroupListAll(c *gin.Context) {
	var req communityReq.ForumTopicGroupListAllSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if data, _, err := appForumTopicGroupService.GetForumTopicGroupInfoListAll(req.Keyword); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(data, "获取成功", c)
	}
}

// CreateForumTopic 创建话题
// @Tags 话题
// @Summary 创建话题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.CreateForumTopicReq true "创建话题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Success 200 {object}  response.Response{data=community.ForumTopic,msg=string} "返回community.ForumTopic"
// @Router /app/topic/createForumTopic [post]
func (topicApi *TopicApi) CreateForumTopic(c *gin.Context) {
	var req communityReq.CreateForumTopicReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if topic, err := appForumTopicService.GetForumTopicByName(req.Name); err == nil {
		response.OkWithDetailed(topic, "成功", c)
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		if newTopic, err := appForumTopicService.CreateForumTopic(community.ForumTopic{
			Name: req.Name,
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			response.OkWithDetailed(newTopic, "成功", c)
			return
		}
	} else {
		response.FailWithMessage(err.Error(), c)
		return
	}
}

// DeleteForumTopic 删除话题
// @Tags 话题
// @Summary 删除话题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除话题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/topic/deleteForumTopic [delete]
func (topicApi *TopicApi) DeleteForumTopic(c *gin.Context) {
	var req request.IdDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var data community.ForumTopic
	data.ID = req.ID
	if err := appForumTopicService.DeleteForumTopic(data); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteForumTopicByIds 批量删除话题
// @Tags 话题
// @Summary 批量删除话题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除话题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/topic/deleteForumTopicByIds [delete]
func (topicApi *TopicApi) DeleteForumTopicByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appForumTopicService.DeleteForumTopicByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateForumTopic 更新话题
// @Tags 话题
// @Summary 更新话题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ForumTopicUpdate true "更新话题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/topic/updateForumTopic [put]
func (topicApi *TopicApi) UpdateForumTopic(c *gin.Context) {
	var hkForumTopic communityReq.ForumTopicUpdate
	err := c.ShouldBindJSON(&hkForumTopic)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appForumTopicService.UpdateForumTopic(hkForumTopic); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindForumTopic 用id查询话题
// @Tags 话题
// @Summary 用id查询话题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询话题"
// @Success 200 {object}  response.PageResult{List=[]community.ForumTopic,msg=string} "返回community.ForumTopic"
// @Router /app/topic/findForumTopic [get]
func (topicApi *TopicApi) FindForumTopic(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumTopic, err := appForumTopicService.GetForumTopic(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rehkForumTopic, c)
	}
}

// GetForumTopicList 分页获取话题列表
// @Tags 话题
// @Summary 分页获取话题列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ForumTopicSearch true "分页获取话题列表"
// @Success 200 {object}  response.PageResult{List=[]community.ForumTopic,msg=string} "返回community.ForumTopic"
// @Router /app/topic/getForumTopicList [get]
func (topicApi *TopicApi) GetForumTopicList(c *gin.Context) {
	var pageInfo communityReq.ForumTopicSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appForumTopicService.GetForumTopicInfoList(pageInfo); err != nil {
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

// GetTopicForumPostsList 分页获取话题帖子列表
// @Tags 话题
// @Summary 分页获取话题帖子列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.TopicForumPostsSearch true "分页获取话题帖子列表"
// @Success 200 {object}  response.PageResult{List=[]community.ForumTopic,msg=string} "返回community.ForumTopic"
// @Router /app/topic/getTopicForumPostsList [get]
func (topicApi *TopicApi) GetTopicForumPostsList(c *gin.Context) {
	var req communityReq.TopicForumPostsSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appForumPostsService.GetTopicForumPostsList(req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// GetNearbyHotTopicList 获取附近热门话题列表
// @Tags 话题
// @Summary 获取附近热门话题列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.NearbyHotTopicSearch true "获取附近热门话题列表"
// @Success 200 {object}  response.PageResult{List=[]community.ForumTopic,msg=string} "返回community.ForumTopic"
// @Router /app/topic/getNearbyHotTopicList [get]
func (topicApi *TopicApi) GetNearbyHotTopicList(c *gin.Context) {
	var req communityReq.NearbyHotTopicSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := appForumTopicService.GetNearbyHotTopicList(req.CurPos); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
