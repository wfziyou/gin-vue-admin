package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FocusApi struct {
}

var hkFocusUserService = service.ServiceGroupApp.AppServiceGroup.Community.FocusUserService

// GetFrequentBrowsingUserList 分页获取经常浏览用户列表
// @Tags 关注/粉丝
// @Summary 分页获取经常浏览用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.FrequentBrowsingUserSearch true "分页获取经常浏览用户列表"
// @Success 200 {object}  response.Response{data=[]community.UserBaseInfo,msg=string}  "返回community.ForumPosts"
// @Router /app/focus/getFrequentBrowsingUserList [get]
func (focusApi *FocusApi) GetFrequentBrowsingUserList(c *gin.Context) {
	var req communityReq.FrequentBrowsingUserSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
}

// GetFocusUserPostsList 分页获取关注用户列子列表
// @Tags 关注/粉丝
// @Summary 分页获取关注用户列子列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.FocusUserDynamicSearch true "分页获取关注用户列子列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/focus/getFocusUserPostsList [get]
func (focusApi *FocusApi) GetFocusUserPostsList(c *gin.Context) {
	var req communityReq.FocusUserDynamicSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
}

// FocusUser 关注用户
// @Tags 关注/粉丝
// @Summary 关注用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.FocusUser true "关注用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /app/focus/focusUser [post]
func (focusApi *FocusApi) FocusUser(c *gin.Context) {
	var req communityReq.FocusUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//if err := hkFocusUserService.CreateFocusUser(hkFocusUser); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
}

// FocusUserCancel 取消用户关注
// @Tags 关注/粉丝
// @Summary 取消用户关注
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.FocusUserCancel true "取消用户关注"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /app/focus/focusUserCancel [post]
func (focusApi *FocusApi) FocusUserCancel(c *gin.Context) {
	var req communityReq.FocusUserCancel
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var userId = utils.GetUserID(c)
	if err := hkFocusUserService.DeleteFocusUser(userId, req.UserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// GetFocusUserList 分页获取关注用户列表
// @Tags 关注/粉丝
// @Summary 分页获取关注用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.FocusUserSearch true "分页获取关注用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/focus/getFocusUserList [get]
func (focusApi *FocusApi) GetFocusUserList(c *gin.Context) {
	var pageInfo communityReq.FocusUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkFocusUserService.GetFocusUserInfoList(pageInfo); err != nil {
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

// GetFansList 分页获取粉丝列表
// @Tags 关注/粉丝
// @Summary 分页获取粉丝列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.FansSearch true "分页获取粉丝列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /app/focus/getFansList [get]
func (focusApi *FocusApi) GetFansList(c *gin.Context) {
	var pageInfo communityReq.FansSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
}

// UpdateFocusUser 更新关注用户
// @Tags 关注/粉丝
// @Summary 更新关注用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.FocusUser true "更新关注用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/focus/updateFocusUser [put]
func (focusApi *FocusApi) UpdateFocusUser(c *gin.Context) {
	var hkFocusUser community.FocusUser
	err := c.ShouldBindJSON(&hkFocusUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkFocusUserService.UpdateFocusUser(hkFocusUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFocusUser 查询关注用户信息
// @Tags 关注/粉丝
// @Summary 查询关注用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.FindFocusUser true "查询关注用户信息"
// @Success 200  {object}  response.Response{data=community.FocusUserInfo,msg=string}  "返回community.FocusUserInfo"
// @Router /app/focus/findFocusUser [get]
func (focusApi *FocusApi) FindFocusUser(c *gin.Context) {
	var req communityReq.FindFocusUser
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var userId = utils.GetUserID(c)
	if focusUserInfo, err := hkFocusUserService.GetFocusUser(userId, req.UserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(focusUserInfo, c)
	}
}
