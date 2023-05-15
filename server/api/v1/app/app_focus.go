package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
// @Success 200 {object}  response.Response{data=[]community.UserBaseInfo,msg=string}  "返回[]community.UserBaseInfo"
// @Router /app/focus/getFrequentBrowsingUserList [get]
func (focusApi *FocusApi) GetFrequentBrowsingUserList(c *gin.Context) {
	var req communityReq.FrequentBrowsingUserSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := hkRecordBrowsingUserHomepageService.GetFrequentBrowsingUserList(userId, req.PageInfo); err != nil {
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

// GetFocusUserPostsList 分页获取关注用户动态列表
// @Tags 关注/粉丝
// @Summary 分页获取关注用户动态列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.FocusUserDynamicSearch true "分页获取关注用户动态列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/focus/getFocusUserPostsList [get]
func (focusApi *FocusApi) GetFocusUserPostsList(c *gin.Context) {
	var req communityReq.FocusUserDynamicSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)
	if list, total, err := appForumPostsService.GetFocusUserPostsList(userId, req.PageInfo); err != nil {
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
	userInfo := utils.GetUserInfo(c)
	if userInfo.GetUserId() == req.UserId {
		response.FailWithMessage("不能关注自己", c)
		return
	}
	focusUser, err := appUserService.GetUser(req.UserId)
	if err != nil {
		global.GVA_LOG.Error("关注用户不存在!", zap.Error(err))
		response.FailWithMessage("关注用户不存在", c)
		return
	}
	otherFocus, err := hkFocusUserService.GetFocusUserObj(req.UserId, userInfo.GetUserId())
	isOtherFocus := false
	if errors.Is(err, gorm.ErrRecordNotFound) {
		isOtherFocus = false
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		isOtherFocus = true
	}

	_, err = hkFocusUserService.GetFocusUserObj(userInfo.GetUserId(), req.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		mutual := 0
		if isOtherFocus == true {
			mutual = 1
		}
		if err := hkFocusUserService.CreateFocusUser(userInfo.GetUserId(), userInfo.NickName, focusUser, mutual); err != nil {
			global.GVA_LOG.Error("关注用户失败", zap.Error(err))
			response.FailWithMessage("关注用户失败", c)
			return
		} else {
			if isOtherFocus == true {
				otherFocus.Mutual = mutual
				hkFocusUserService.UpdateFocusUser(otherFocus)
			}
			appUserService.UpdateUserNumFocus(userInfo.GetUserId())
			appUserService.UpdateUserNumFan(req.UserId)
			response.OkWithMessage("成功", c)
			return
		}
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithMessage("成功", c)
		return
	}
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
	_, err = appUserService.GetUser(req.UserId)
	if err != nil {
		global.GVA_LOG.Error("关注用户不存在!", zap.Error(err))
		response.FailWithMessage("关注用户不存在", c)
		return
	}
	otherFocus, err := hkFocusUserService.GetFocusUserObj(req.UserId, userId)
	isOtherFocus := false
	if errors.Is(err, gorm.ErrRecordNotFound) {
		isOtherFocus = false
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		isOtherFocus = true
	}

	focus, err := hkFocusUserService.GetFocusUserObj(userId, req.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithMessage("成功", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		if err := hkFocusUserService.DeleteFocusUserObj(focus); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			if isOtherFocus == true {
				otherFocus.Mutual = 0
				hkFocusUserService.UpdateFocusUser(otherFocus)
			}
			appUserService.UpdateUserNumFocus(userId)
			appUserService.UpdateUserNumFan(req.UserId)
		}
	}
	response.OkWithMessage("成功", c)
}

// GetFocusUserList 分页获取关注用户列表
// @Tags 关注/粉丝
// @Summary 分页获取关注用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.FocusUserSearch true "分页获取关注用户列表"
// @Success 200 {object} response.PageResult{List=[]community.FocusUserInfo,msg=string} "返回[]community.FocusUserInfo"
// @Router /app/focus/getFocusUserList [get]
func (focusApi *FocusApi) GetFocusUserList(c *gin.Context) {
	var pageInfo communityReq.FocusUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := hkFocusUserService.GetFocusUserInfoList(userId, pageInfo); err != nil {
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
// @Success 200 {object} response.PageResult{List=[]community.FocusUserInfo,msg=string} "返回[]community.FocusUserInfo"
// @Router /app/focus/getFansList [get]
func (focusApi *FocusApi) GetFansList(c *gin.Context) {
	var req communityReq.FansSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var userId = utils.GetUserID(c)
	if list, total, err := hkFocusUserService.GetFansList(userId, req.PageInfo); err != nil {
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

// UpdateFocusUser 更新关注用户
// @Tags 关注/粉丝
// @Summary 更新关注用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.UpdateFocusUserReq true "更新关注用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/focus/updateFocusUser [put]
func (focusApi *FocusApi) UpdateFocusUser(c *gin.Context) {
	var req communityReq.UpdateFocusUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := hkFocusUserService.UpdateFocusUserEx(userId, req); err != nil {
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
