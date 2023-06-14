package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ForumPostsApi struct {
}

/*************************************
帖子
**************************************/

// GetRecommendPostsList 分页获取推荐帖子列表
// @Tags 帖子
// @Summary 分页获取推荐帖子列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.GetRecommendPostsSearch true "分页获取推荐帖子列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/forumPosts/getRecommendPostsList [get]
func (forumPostsApi *ForumPostsApi) GetRecommendPostsList(c *gin.Context) {
	var req communityReq.GetRecommendPostsSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var userId = utils.GetUserID(c)
	if list, total, err := appForumPostsService.GetRecommendPostsList(userId, req.ChannelId, req.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		appForumThumbsUpService.GetUserForumThumbsUp(userId, list)
		appUserCollectService.GetUserIsCollect(userId, list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// GetGlobalTopInfoList 获全局置顶资讯列表
// @Tags 帖子
// @Summary 获全局置顶资讯列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]community.ForumPostsBaseInfo,msg=string} "返回[]community.ForumPostsBaseInfo"
// @Router /app/forumPosts/getGlobalTopInfoList [get]
func (forumPostsApi *ForumPostsApi) GetGlobalTopInfoList(c *gin.Context) {
	userId := utils.GetUserID(c)
	if list, _, err := appForumPostsService.GetGlobalTopInfoList(userId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		appForumThumbsUpService.GetUserForumThumbsUp(userId, list)
		appUserCollectService.GetUserIsCollect(userId, list)
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// GetGlobalRecommendInfoList 分页获全局推荐资讯列表
// @Tags 帖子
// @Summary 分页获全局推荐资讯列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.GlobalRecommendInfoSearch true "分页获全局推荐资讯列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/forumPosts/getGlobalRecommendInfoList [get]
func (forumPostsApi *ForumPostsApi) GetGlobalRecommendInfoList(c *gin.Context) {
	var req communityReq.GlobalRecommendInfoSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := appForumPostsService.GetGlobalRecommendInfoList(userId, req.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		appForumThumbsUpService.GetUserForumThumbsUp(userId, list)
		appUserCollectService.GetUserIsCollect(userId, list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// GetNearbyRecommendPostsList 分页获附近推荐帖子列表
// @Tags 帖子
// @Summary 分页获附近推荐帖子列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.NearbyRecommendPostsSearch true "分页获附近推荐帖子列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/forumPosts/getNearbyRecommendPostsList [get]
func (forumPostsApi *ForumPostsApi) GetNearbyRecommendPostsList(c *gin.Context) {
	var req communityReq.NearbyRecommendPostsSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := appForumPostsService.GetNearbyRecommendPostsList(userId, req.CurPos, req.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		appForumThumbsUpService.GetUserForumThumbsUp(userId, list)
		appUserCollectService.GetUserIsCollect(userId, list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// CreateForumPosts 创建帖子
// @Tags 帖子
// @Summary 创建帖子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.CreateForumPostsReq true "创建帖子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/forumPosts/createForumPosts [post]
func (forumPostsApi *ForumPostsApi) CreateForumPosts(c *gin.Context) {
	var req communityReq.CreateForumPostsReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if req.CircleId != 0 {
		if _, err := appCircleService.GetCircleInfo(req.CircleId); err != nil {
			response.FailWithMessage("圈子不存在", c)
			return
		}

		if _, err := appCircleUserService.GetCircleUser(req.CircleId, userId); errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithMessage("用户不在圈子中", c)
			return
		} else if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
	req.UserId = userId
	if err := appForumPostsService.CreateForumPosts(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	} else {
		appUserService.UpdatePostsTime(req.CircleId, userId)
		response.OkWithMessage("创建成功", c)
	}
}

// CreateNews 创建资讯
// @Tags 帖子
// @Summary 创建资讯
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamCreateNews true "创建资讯"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/forumPosts/createNews [post]
func (forumPostsApi *ForumPostsApi) CreateNews(c *gin.Context) {
	var req communityReq.ParamCreateNews
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if _, err := appCircleService.GetCircleInfo(req.CircleId); err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}
	circleUser, err := appCircleUserService.GetCircleUser(req.CircleId, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("用户不在圈子中", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if circleUser.Power == community.CircleUserPowerGeneral {
		response.FailWithMessage("没有权限", c)
		return
	}

	if err := appForumPostsService.CreateNews(userId, req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	} else {
		appUserService.UpdatePostsTime(req.CircleId, userId)
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteForumPosts 删除帖子
// @Tags 帖子
// @Summary 删除帖子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除帖子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/forumPosts/deleteForumPosts [delete]
func (forumPostsApi *ForumPostsApi) DeleteForumPosts(c *gin.Context) {
	var req request.IdDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	posts, err := appForumPostsService.GetForumPosts(req.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("帖子不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appForumPostsService.DeleteForumPosts(posts); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSelfForumPosts 删除自己的帖子
// @Tags 帖子
// @Summary 删除自己的帖子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除自己的帖子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/forumPosts/deleteSelfForumPosts [delete]
func (forumPostsApi *ForumPostsApi) DeleteSelfForumPosts(c *gin.Context) {
	var req request.IdDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if data, err := appForumPostsService.GetForumPosts(req.ID); err != nil {
		response.FailWithMessage("帖子不存在", c)
		return
	} else if data.UserId != utils.GetUserID(c) {
		response.FailWithMessage("不是自己的帖子", c)
		return
	} else {
		if err := appForumPostsService.DeleteForumPosts(data); err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			response.FailWithMessage("删除失败", c)
		} else {
			response.OkWithMessage("删除成功", c)
		}
	}
}

// DeleteForumPostsByIds 批量删除帖子
// @Tags 帖子
// @Summary 批量删除帖子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除帖子"
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
// @Tags 帖子
// @Summary 更新ForumPosts
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ForumPosts true "更新ForumPosts"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/forumPosts/updateForumPosts [put]
func (forumPostsApi *ForumPostsApi) UpdateForumPosts(c *gin.Context) {
	var req community.ForumPosts
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if data, err := appForumPostsService.GetForumPosts(req.ID); err != nil {
		response.FailWithMessage("帖子不存在", c)
	} else if data.UserId != utils.GetUserID(c) {
		response.FailWithMessage("不是自己的帖子", c)
	}
	if err := appForumPostsService.UpdateForumPosts(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindForumPosts 用id查询帖子
// @Tags 帖子
// @Summary 用id查询帖子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询帖子"
// @Success 200 {object}  response.Response{data=community.ForumPosts,msg=string}  "返回community.ForumPosts"
// @Router /app/forumPosts/findForumPosts [get]
func (forumPostsApi *ForumPostsApi) FindForumPosts(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var userId = utils.GetUserID(c)
	if rehkForumPosts, err := appForumPostsService.FindForumPosts(userId, idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		if _, num, err := appForumThumbsUpService.GetForumThumbsUpEx(userId, []uint64{idSearch.ID}); err == nil && num > 0 {
			rehkForumPosts.ThumbsUp = 1
		}
		if _, num, err := appUserCollectService.GetUserCollectEx(userId, []uint64{idSearch.ID}); err == nil && num > 0 {
			rehkForumPosts.Collect = 1
		}
		if rehkForumPosts.CircleId > 0 {
			if _, err := appCircleUserService.GetCircleUser(rehkForumPosts.CircleId, userId); err == nil {
				rehkForumPosts.CircleInfo.HaveCircle = 1
			}
		}

		hkRecordBrowsingUserHomepageService.BrowsingUser(userId, rehkForumPosts.UserId)
		appUserBrowsingHistoryService.CreateUserBrowsingHistory(userId, rehkForumPosts.ID, rehkForumPosts.Category)
		response.OkWithData(rehkForumPosts, c)
	}
}

// GetForumPostsList 分页获取ForumPosts列表
// @Tags 帖子
// @Summary 分页获取ForumPosts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ForumPostsSearch true "分页获取ForumPosts列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/forumPosts/getForumPostsList [get]
func (forumPostsApi *ForumPostsApi) GetForumPostsList(c *gin.Context) {
	var pageInfo communityReq.ForumPostsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := appForumPostsService.GetForumPostsInfoList(userId, pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		var userId = utils.GetUserID(c)
		appForumThumbsUpService.GetUserForumThumbsUp(userId, list)
		appUserCollectService.GetUserIsCollect(userId, list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// CreateForumComment 创建评论
// @Tags 帖子
// @Summary 创建评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.CreateForumComment true "创建评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /app/forumPosts/createForumComment [post]
func (forumPostsApi *ForumPostsApi) CreateForumComment(c *gin.Context) {
	var req communityReq.CreateForumComment
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	posts, err := appForumPostsService.GetForumPosts(req.PostsId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("帖子不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if posts.PowerComment == community.ForumPostsPowerCommentClose {
		response.FailWithMessage("此帖子不能被评论", c)
		return
	}

	req.UserId = utils.GetUserID(c)
	if err := appForumCommentService.CreateForumComment(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteForumComment 删除评论
// @Tags 帖子
// @Summary 删除评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/forumPosts/deleteForumComment [delete]
func (forumPostsApi *ForumPostsApi) DeleteForumComment(c *gin.Context) {
	var req request.IdDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if data, err := appForumCommentService.GetForumComment(uint64(req.ID)); err != nil {
		response.OkWithMessage("评论不存在", c)
	} else {
		if data.UserId != utils.GetUserID(c) {
			response.OkWithMessage("不能删除别人的评论", c)
			return
		}
		if err := appForumCommentService.DeleteForumComment(data); err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			response.FailWithMessage("删除失败", c)
		} else {
			response.OkWithMessage("删除成功", c)
		}
	}
}

// DeleteForumCommentByIds 批量删除评论
// @Tags 帖子
// @Summary 批量删除评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论"
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
// @Tags 帖子
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
	var userId = utils.GetUserID(c)
	if rehkForumComment, err := appForumCommentService.GetForumComment(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		if _, num, err := appCommentThumbsUpService.GetCommentThumbsUpEx(userId, []uint64{idSearch.ID}); err == nil && num > 0 {
			rehkForumComment.ThumbsUp = 1
		}
		response.OkWithData(rehkForumComment, c)
	}
}

// GetForumCommentList 分页获取评论列表
// @Tags 帖子
// @Summary 分页获取评论列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ForumCommentSearch true "分页获取评论列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumComment,msg=string} "返回community.ForumComment"
// @Router /app/forumPosts/getForumCommentList [get]
func (forumPostsApi *ForumPostsApi) GetForumCommentList(c *gin.Context) {
	var pageInfo communityReq.ForumCommentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := appForumCommentService.GetForumCommentInfoList(userId, pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		appCommentThumbsUpService.GetUserCommentThumbsUp(userId, list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// CreateForumThumbsUp 帖子点赞
// @Tags 帖子
// @Summary 帖子点赞
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ForumThumbsUpReq true "帖子点赞"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/forumPosts/createForumThumbsUp [post]
func (forumPostsApi *ForumPostsApi) CreateForumThumbsUp(c *gin.Context) {
	var req communityReq.ForumThumbsUpReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.UserId = utils.GetUserID(c)
	if err := appForumThumbsUpService.CreateForumThumbsUp(community.ForumThumbsUp{
		UserId:  req.UserId,
		PostsId: req.PostsId,
	}); err != nil {
		global.GVA_LOG.Error("点赞失败!", zap.Error(err))
		response.FailWithMessage("点赞失败", c)
	} else {
		response.OkWithMessage("点赞成功", c)
	}
}

// DeleteForumThumbsUp 取消帖子点赞
// @Tags 帖子
// @Summary 取消帖子点赞
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.DeleteForumThumbsUp true "取消帖子点赞"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/forumPosts/deleteForumThumbsUp [delete]
func (forumPostsApi *ForumPostsApi) DeleteForumThumbsUp(c *gin.Context) {
	var req communityReq.DeleteForumThumbsUp
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.UserId = utils.GetUserID(c)
	if err := appForumThumbsUpService.DeleteForumThumbsUp(req); err != nil {
		global.GVA_LOG.Error("取消点赞失败!", zap.Error(err))
		response.FailWithMessage("取消点赞失败", c)
	} else {
		response.OkWithMessage("取消点赞成功", c)
	}
}

// CreateCommentThumbsUp 评论点赞
// @Tags 帖子
// @Summary 评论点赞
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.CommentThumbsUpReq true "评论点赞"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/forumPosts/createCommentThumbsUp [post]
func (forumPostsApi *ForumPostsApi) CreateCommentThumbsUp(c *gin.Context) {
	var req communityReq.CommentThumbsUpReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if _, err := appForumCommentService.GetForumComment(req.CommentId); err != nil {
		response.FailWithMessage("评论不存在", c)
		return
	}

	if err := appCommentThumbsUpService.CreateCommentThumbsUp(userId, req.CommentId); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCommentThumbsUp 取消评论点赞
// @Tags 帖子
// @Summary 取消评论点赞
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.DeleteCommentThumbsUp true "取消评论点赞"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/forumPosts/deleteCommentThumbsUp [delete]
func (forumPostsApi *ForumPostsApi) DeleteCommentThumbsUp(c *gin.Context) {
	var req communityReq.DeleteCommentThumbsUp
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.UserId = utils.GetUserID(c)
	if err := appCommentThumbsUpService.DeleteCommentThumbsUp(req); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetUserForumPostsList 分页获取用户帖子列表
// @Tags 帖子
// @Summary 分页获取用户帖子列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.UserForumPostsSearch true "分页获取用户帖子列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/forumPosts/getUserForumPostsList [get]
func (forumPostsApi *ForumPostsApi) GetUserForumPostsList(c *gin.Context) {
	var req communityReq.UserForumPostsSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	selectUserId := utils.GetUserID(c)
	userId := selectUserId
	if req.UserId > 0 && userId != req.UserId {
		userId = req.UserId
	}
	if list, total, err := appForumPostsService.GetUserForumPostsList(selectUserId, userId, req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		selectUserId := utils.GetUserID(c)
		appForumThumbsUpService.GetUserForumThumbsUp(selectUserId, list)
		appUserCollectService.GetUserIsCollect(selectUserId, list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}
