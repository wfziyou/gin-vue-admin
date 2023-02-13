package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ForumPostsApi struct {
}

/*************************************
帖子
**************************************/

// CreateForumPosts 创建ForumPosts
// @Tags App_ForumPosts
// @Summary 创建ForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ForumPosts true "创建ForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/forumPosts/createForumPosts [post]
func (forumPostsApi *ForumPostsApi) CreateForumPosts(c *gin.Context) {
	var req communityReq.CreateForumPostsReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.TopicId != 0 {
		if _, err := appForumTopicService.GetForumTopic(uint64(req.TopicId)); err != nil {
			response.FailWithMessage("话题不存在", c)
		}
	}
	if req.CircleId != 0 {
		if _, err := appCircleService.GetCircle(uint64(req.CircleId)); err != nil {
			response.FailWithMessage("圈子不存在", c)
		}
	}
	req.UserId = utils.GetUserID(c)
	if err := appForumPostsService.CreateForumPosts(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	} else {

		response.OkWithMessage("创建成功", c)
	}
}

// DeleteForumPosts 删除ForumPosts
// @Tags App_ForumPosts
// @Summary 删除ForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除ForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/forumPosts/deleteForumPosts [delete]
func (forumPostsApi *ForumPostsApi) DeleteForumPosts(c *gin.Context) {
	var hkForumPosts request.IdDelete
	err := c.ShouldBindJSON(&hkForumPosts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := appForumPostsService.DeleteForumPosts(hkForumPosts); err != nil {
	//	global.GVA_LOG.Error("删除失败!", zap.Error(err))
	//	response.FailWithMessage("删除失败", c)
	//} else {
	//	response.OkWithMessage("删除成功", c)
	//}
}

// DeleteForumPostsByIds 批量删除ForumPosts
// @Tags App_ForumPosts
// @Summary 批量删除ForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/forumPosts/deleteForumPostsByIds [delete]
func (forumPostsApi *ForumPostsApi) DeleteForumPostsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appForumPostsService.DeleteForumPostsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateForumPosts 更新ForumPosts
// @Tags App_ForumPosts
// @Summary 更新ForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ForumPosts true "更新ForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/forumPosts/updateForumPosts [put]
func (forumPostsApi *ForumPostsApi) UpdateForumPosts(c *gin.Context) {
	var hkForumPosts community.ForumPosts
	err := c.ShouldBindJSON(&hkForumPosts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := appForumPostsService.UpdateForumPosts(hkForumPosts); err != nil {
	//	global.GVA_LOG.Error("更新失败!", zap.Error(err))
	//	response.FailWithMessage("更新失败", c)
	//} else {
	//	response.OkWithMessage("更新成功", c)
	//}
}

// FindForumPosts 用id查询ForumPosts
// @Tags App_ForumPosts
// @Summary 用id查询ForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询ForumPosts"
// @Success 200 {object}  response.Response{data=community.ForumPosts,msg=string}  "返回community.ForumPosts"
// @Router /app/forumPosts/findForumPosts [get]
func (forumPostsApi *ForumPostsApi) FindForumPosts(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumPosts, err := appForumPostsService.GetForumPosts(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForumPosts": rehkForumPosts}, c)
	}
}

// GetForumPostsList 分页获取ForumPosts列表
// @Tags App_ForumPosts
// @Summary 分页获取ForumPosts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ForumPostsSearch true "分页获取ForumPosts列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPosts,msg=string} "返回community.ForumPosts"
// @Router /app/forumPosts/getForumPostsList [get]
func (forumPostsApi *ForumPostsApi) GetForumPostsList(c *gin.Context) {
	var pageInfo communityReq.ForumPostsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appForumPostsService.GetForumPostsInfoList(pageInfo); err != nil {
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

// CreateForumComment 创建ForumComment
// @Tags App_ForumPosts
// @Summary 创建ForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ForumComment true "创建ForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/forumPosts/createForumComment [post]
func (forumPostsApi *ForumPostsApi) CreateForumComment(c *gin.Context) {
	var hkForumComment community.ForumComment
	err := c.ShouldBindJSON(&hkForumComment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appForumCommentService.CreateForumComment(hkForumComment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteForumComment 删除ForumComment
// @Tags App_ForumPosts
// @Summary 删除ForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除ForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/forumPosts/deleteForumComment [delete]
func (forumPostsApi *ForumPostsApi) DeleteForumComment(c *gin.Context) {
	var hkForumComment request.IdDelete
	err := c.ShouldBindJSON(&hkForumComment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := appForumCommentService.DeleteForumComment(hkForumComment); err != nil {
	//	global.GVA_LOG.Error("删除失败!", zap.Error(err))
	//	response.FailWithMessage("删除失败", c)
	//} else {
	//	response.OkWithMessage("删除成功", c)
	//}
}

// DeleteForumCommentByIds 批量删除ForumComment
// @Tags App_ForumPosts
// @Summary 批量删除ForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ForumComment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/forumPosts/deleteForumCommentByIds [delete]
func (forumPostsApi *ForumPostsApi) DeleteForumCommentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appForumCommentService.DeleteForumCommentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// FindForumComment 用id查询ForumComment
// @Tags App_ForumPosts
// @Summary 用id查询ForumComment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询ForumComment"
// @Success 200 {object}  response.Response{data=community.ForumComment,msg=string}  "返回community.ForumComment"
// @Router /app/forumPosts/findForumComment [get]
func (forumPostsApi *ForumPostsApi) FindForumComment(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumComment, err := appForumCommentService.GetForumComment(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForumComment": rehkForumComment}, c)
	}
}

// GetForumCommentList 分页获取ForumComment列表
// @Tags App_ForumPosts
// @Summary 分页获取ForumComment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ForumCommentSearch true "分页获取ForumComment列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumComment,msg=string} "返回community.ForumComment"
// @Router /app/forumPosts/getForumCommentList [get]
func (forumPostsApi *ForumPostsApi) GetForumCommentList(c *gin.Context) {
	var pageInfo communityReq.ForumCommentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appForumCommentService.GetForumCommentInfoList(pageInfo); err != nil {
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

// CreateForumThumbsUp 帖子点赞
// @Tags App_ForumPosts
// @Summary 帖子点赞
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ForumThumbsUp true "帖子点赞"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/forumPosts/createForumThumbsUp [post]
func (forumPostsApi *ForumPostsApi) CreateForumThumbsUp(c *gin.Context) {
	var hkForumThumbsUp community.ForumThumbsUp
	err := c.ShouldBindJSON(&hkForumThumbsUp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appForumThumbsUpService.CreateForumThumbsUp(hkForumThumbsUp); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteForumThumbsUp 删除帖子点赞
// @Tags App_ForumPosts
// @Summary 删除帖子点赞
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除帖子点赞"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/forumPosts/deleteForumThumbsUp [delete]
func (forumPostsApi *ForumPostsApi) DeleteForumThumbsUp(c *gin.Context) {
	var hkForumThumbsUp request.IdDelete
	err := c.ShouldBindJSON(&hkForumThumbsUp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := appForumThumbsUpService.DeleteForumThumbsUp(hkForumThumbsUp); err != nil {
	//	global.GVA_LOG.Error("删除失败!", zap.Error(err))
	//	response.FailWithMessage("删除失败", c)
	//} else {
	//	response.OkWithMessage("删除成功", c)
	//}
}

// CreateCommentThumbsUp 评论点赞
// @Tags App_ForumPosts
// @Summary 评论点赞
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.CommentThumbsUp true "评论点赞"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/forumPosts/createCommentThumbsUp [post]
func (forumPostsApi *ForumPostsApi) CreateCommentThumbsUp(c *gin.Context) {
	var hkCommentThumbsUp community.CommentThumbsUp
	err := c.ShouldBindJSON(&hkCommentThumbsUp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appCommentThumbsUpService.CreateCommentThumbsUp(hkCommentThumbsUp); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCommentThumbsUp 删除评论点赞
// @Tags App_ForumPosts
// @Summary 删除评论点赞
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除评论点赞"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/forumPosts/deleteCommentThumbsUp [delete]
func (forumPostsApi *ForumPostsApi) DeleteCommentThumbsUp(c *gin.Context) {
	var hkCommentThumbsUp request.IdDelete
	err := c.ShouldBindJSON(&hkCommentThumbsUp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := appCommentThumbsUpService.DeleteCommentThumbsUp(hkCommentThumbsUp); err != nil {
	//	global.GVA_LOG.Error("删除失败!", zap.Error(err))
	//	response.FailWithMessage("删除失败", c)
	//} else {
	//	response.OkWithMessage("删除成功", c)
	//}
}
