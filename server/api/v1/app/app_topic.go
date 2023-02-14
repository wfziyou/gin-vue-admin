package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TopicApi struct {
}

/*************************************
话题
**************************************/

// FindForumTopicGroup 用id查询ForumTopicGroup
// @Tags App_Topic
// @Summary 用id查询ForumTopicGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询ForumTopicGroup"
// @Success 200 {object}  response.Response{data=community.ForumTopicGroup,msg=string}  "返回community.ForumTopicGroup"
// @Router /app/topic/findForumTopicGroup [get]
func (topicApi *TopicApi) FindForumTopicGroup(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumTopicGroup, err := appForumTopicGroupService.GetForumTopicGroup(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForumTopicGroup": rehkForumTopicGroup}, c)
	}
}

// GetForumTopicGroupList 分页获取ForumTopicGroup列表
// @Tags App_Topic
// @Summary 分页获取ForumTopicGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ForumTopicGroupSearch true "分页获取ForumTopicGroup列表"
// @Success 200 {object}  response.PageResult{List=[]community.ForumTopicGroup,msg=string} "返回community.ForumTopicGroup"
// @Router /app/topic/getForumTopicGroupList [get]
func (topicApi *TopicApi) GetForumTopicGroupList(c *gin.Context) {
	var pageInfo communityReq.ForumTopicGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appForumTopicGroupService.GetForumTopicGroupInfoList(pageInfo); err != nil {
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

// GetForumTopicGroupListAll 获取ForumTopicGroup列表
// @Tags App_Topic
// @Summary 获取ForumTopicGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ForumTopicGroupSearch true "获取ForumTopicGroup列表"
// @Success 200 {object}  response.DataResult{data=[]community.ForumTopicGroup,msg=string} "返回community.ForumTopicGroup"
// @Router /app/topic/getForumTopicGroupListAll [get]
func (topicApi *TopicApi) GetForumTopicGroupListAll(c *gin.Context) {
	var pageInfo communityReq.ForumTopicGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if data, _, err := appForumTopicGroupService.GetForumTopicGroupInfoListAll(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(data, "获取成功", c)
	}
}

// CreateForumTopic 创建ForumTopic
// @Tags App_Topic
// @Summary 创建ForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.CreateForumTopicReq true "创建ForumTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/topic/createForumTopic [post]
func (topicApi *TopicApi) CreateForumTopic(c *gin.Context) {
	var req communityReq.CreateForumTopicReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appForumTopicService.CreateForumTopic(community.ForumTopic{
		Name:         req.Name,
		CoverImage:   req.CoverImage,
		TopicGroupId: req.TopicGroupId,
		Type:         req.Type,
		CircleId:     req.CircleId,
	}); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteForumTopic 删除ForumTopic
// @Tags App_Topic
// @Summary 删除ForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除ForumTopic"
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

// DeleteForumTopicByIds 批量删除ForumTopic
// @Tags App_Topic
// @Summary 批量删除ForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ForumTopic"
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

// UpdateForumTopic 更新ForumTopic
// @Tags App_Topic
// @Summary 更新ForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ForumTopic true "更新ForumTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/topic/updateForumTopic [put]
func (topicApi *TopicApi) UpdateForumTopic(c *gin.Context) {
	var hkForumTopic community.ForumTopic
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

// FindForumTopic 用id查询ForumTopic
// @Tags App_Topic
// @Summary 用id查询ForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询ForumTopic"
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
		response.OkWithData(gin.H{"rehkForumTopic": rehkForumTopic}, c)
	}
}

// GetForumTopicList 分页获取ForumTopic列表
// @Tags App_Topic
// @Summary 分页获取ForumTopic列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ForumTopicSearch true "分页获取ForumTopic列表"
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
