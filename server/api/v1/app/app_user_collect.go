package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserCollectApi struct {
}

/*************************************
收藏
**************************************/

// CreateUserCollect 收藏帖子
// @Tags 收藏
// @Summary 收藏帖子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body generalReq.UserCollectReq true "收藏帖子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /app/userCollect/createUserCollect [post]
func (userCollectApi *UserCollectApi) CreateUserCollect(c *gin.Context) {
	var req generalReq.UserCollectReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if posts, err := appForumPostsService.GetForumPosts(req.PostsId); err != nil {
		response.FailWithMessage("帖子不存在", c)
		return
	} else if err := appUserCollectService.CreateUserCollect(userId, posts); err != nil {
		global.GVA_LOG.Error("收藏失败!", zap.Error(err))
		response.FailWithMessage("收藏失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

// DeleteUserCollect 删除收藏
// @Tags 收藏
// @Summary 删除收藏
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除收藏"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/userCollect/deleteUserCollect [delete]
func (userCollectApi *UserCollectApi) DeleteUserCollect(c *gin.Context) {
	var req request.IdDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var data = general.UserCollect{}
	data.ID = req.ID
	if err := appUserCollectService.DeleteUserCollect(data); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserCollectByIds 批量删除收藏
// @Tags 收藏
// @Summary 批量删除收藏
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除收藏"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/userCollect/deleteUserCollectByIds [delete]
func (userCollectApi *UserCollectApi) DeleteUserCollectByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUserCollectService.DeleteUserCollectByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// GetUserCollectList 分页获取收藏列表
// @Tags 收藏
// @Summary 分页获取收藏列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.UserCollectSearch true "分页获取收藏列表"
// @Success 200 {object}  response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回general.UserCollect"
// @Router /app/userCollect/getUserCollectList [get]
func (userCollectApi *UserCollectApi) GetUserCollectList(c *gin.Context) {
	var pageInfo generalReq.UserCollectSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.UserId = utils.GetUserID(c)
	if list, total, err := appUserCollectService.GetUserCollectInfoList(pageInfo); err != nil {
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
