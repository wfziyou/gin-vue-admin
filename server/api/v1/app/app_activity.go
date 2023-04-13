package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ActivityApi struct {
}

var hkActivityService = service.ServiceGroupApp.AppServiceGroup.Community.ActivityService
var hkActivityClassifyService = service.ServiceGroupApp.AppServiceGroup.Community.ActivityClassifyService
var hkActivityUserService = service.ServiceGroupApp.AppServiceGroup.Community.ActivityUserService

// CreateActivity 创建活动
// @Tags 活动
// @Summary 创建活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.Activity true "创建活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/activity/createActivity [post]
func (activityApi *ActivityApi) CreateActivity(c *gin.Context) {
	var hkActivity community.Activity
	err := c.ShouldBindJSON(&hkActivity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityService.CreateActivity(hkActivity); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteActivity 删除活动
// @Tags 活动
// @Summary 删除活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.Activity true "删除活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/activity/deleteActivity [delete]
func (activityApi *ActivityApi) DeleteActivity(c *gin.Context) {
	var hkActivity community.Activity
	err := c.ShouldBindJSON(&hkActivity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityService.DeleteActivity(hkActivity); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteActivityByIds 批量删除活动
// @Tags 活动
// @Summary 批量删除活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/activity/deleteActivityByIds [delete]
func (activityApi *ActivityApi) DeleteActivityByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityService.DeleteActivityByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateActivity 更新活动
// @Tags 活动
// @Summary 更新活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.Activity true "更新活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/activity/updateActivity [put]
func (activityApi *ActivityApi) UpdateActivity(c *gin.Context) {
	var hkActivity community.Activity
	err := c.ShouldBindJSON(&hkActivity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityService.UpdateActivity(hkActivity); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindActivity 用id查询活动详情
// @Tags 活动
// @Summary 用id查询活动详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.Activity true "用id查询活动详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/activity/findActivity [get]
func (activityApi *ActivityApi) FindActivity(c *gin.Context) {
	var hkActivity community.Activity
	err := c.ShouldBindQuery(&hkActivity)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkActivity, err := hkActivityService.GetActivity(hkActivity.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkActivity": rehkActivity}, c)
	}
}

// GetActivityList 分页获取推荐活动列表
// @Tags 活动
// @Summary 分页获取推荐活动列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ActivitySearch true "分页获取推荐活动列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/activity/getActivityList [get]
func (activityApi *ActivityApi) GetActivityList(c *gin.Context) {
	var pageInfo communityReq.ActivitySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkActivityService.GetActivityInfoList(pageInfo); err != nil {
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

// CreateActivityClassify 创建活动分类
// @Tags 活动
// @Summary 创建活动分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ActivityClassify true "创建活动分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/activity/createActivityClassify [post]
func (activityApi *ActivityApi) CreateActivityClassify(c *gin.Context) {
	var hkActivityClassify community.ActivityClassify
	err := c.ShouldBindJSON(&hkActivityClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityClassifyService.CreateActivityClassify(hkActivityClassify); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteActivityClassify 删除活动分类
// @Tags 活动
// @Summary 删除活动分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ActivityClassify true "删除活动分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/activity/deleteActivityClassify [delete]
func (activityApi *ActivityApi) DeleteActivityClassify(c *gin.Context) {
	var hkActivityClassify community.ActivityClassify
	err := c.ShouldBindJSON(&hkActivityClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityClassifyService.DeleteActivityClassify(hkActivityClassify); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteActivityClassifyByIds 批量删除活动分类
// @Tags 活动
// @Summary 批量删除活动分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除活动分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/activity/deleteActivityClassifyByIds [delete]
func (activityApi *ActivityApi) DeleteActivityClassifyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityClassifyService.DeleteActivityClassifyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateActivityClassify 更新活动分类
// @Tags 活动
// @Summary 更新活动分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ActivityClassify true "更新活动分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/activity/updateActivityClassify [put]
func (activityApi *ActivityApi) UpdateActivityClassify(c *gin.Context) {
	var hkActivityClassify community.ActivityClassify
	err := c.ShouldBindJSON(&hkActivityClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityClassifyService.UpdateActivityClassify(hkActivityClassify); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindActivityClassify 用id查询活动分类
// @Tags 活动
// @Summary 用id查询活动分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.ActivityClassify true "用id查询活动分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/activity/findActivityClassify [get]
func (activityApi *ActivityApi) FindActivityClassify(c *gin.Context) {
	var hkActivityClassify community.ActivityClassify
	err := c.ShouldBindQuery(&hkActivityClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkActivityClassify, err := hkActivityClassifyService.GetActivityClassify(hkActivityClassify.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkActivityClassify": rehkActivityClassify}, c)
	}
}

// GetActivityClassifyList 分页获取活动分类列表
// @Tags 活动
// @Summary 分页获取活动分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ActivityClassifySearch true "分页获取活动分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/activity/getActivityClassifyList [get]
func (activityApi *ActivityApi) GetActivityClassifyList(c *gin.Context) {
	var pageInfo communityReq.ActivityClassifySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkActivityClassifyService.GetActivityClassifyInfoList(pageInfo); err != nil {
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

// JoinActivity 加入活动
// @Tags 活动
// @Summary 加入活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ActivityUser true "加入活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/activity/joinActivity [post]
func (activityApi *ActivityApi) JoinActivity(c *gin.Context) {
	var hkActivityUser community.ActivityUser
	err := c.ShouldBindJSON(&hkActivityUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityUserService.CreateActivityUser(hkActivityUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// ExitActivity 退出活动
// @Tags 活动
// @Summary 退出活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ActivityUser true "退出活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/activity/exitActivity [delete]
func (activityApi *ActivityApi) ExitActivity(c *gin.Context) {
	var hkActivityUser community.ActivityUser
	err := c.ShouldBindJSON(&hkActivityUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityUserService.DeleteActivityUser(hkActivityUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// KickOutActivityUsers 踢出活动的用户
// @Tags 活动
// @Summary 踢出活动的用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "踢出活动的用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/activity/kickOutActivityUsers [delete]
func (activityApi *ActivityApi) KickOutActivityUsers(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityUserService.DeleteActivityUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// FindActivityUser 用id查询活动的用户
// @Tags 活动
// @Summary 用id查询活动的用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.ActivityUser true "用id查询活动的用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/activity/findActivityUser [get]
func (activityApi *ActivityApi) FindActivityUser(c *gin.Context) {
	var hkActivityUser community.ActivityUser
	err := c.ShouldBindQuery(&hkActivityUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkActivityUser, err := hkActivityUserService.GetActivityUser(hkActivityUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkActivityUser": rehkActivityUser}, c)
	}
}

// GetActivityUserList 分页获取活动的用户列表
// @Tags 活动
// @Summary 分页获取活动的用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ActivityUserSearch true "分页获取活动的用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/activity/getActivityUserList [get]
func (activityApi *ActivityApi) GetActivityUserList(c *gin.Context) {
	var pageInfo communityReq.ActivityUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkActivityUserService.GetActivityUserInfoList(pageInfo); err != nil {
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

// GetGlobalRecommendActivityList 分页获全局推荐活动列表
// @Tags 活动
// @Summary 分页获全局推荐活动列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.GlobalRecommendActivitySearch true "分页获全局推荐活动列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/activity/getGlobalRecommendActivityList [get]
func (activityApi *ActivityApi) GetGlobalRecommendActivityList(c *gin.Context) {
	var pageInfo communityReq.GlobalRecommendActivitySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := hkActivityService.GetActivityInfoList(pageInfo); err != nil {
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("获取失败", c)
	//} else {
	//	response.OkWithDetailed(response.PageResult{
	//		List:     list,
	//		Total:    total,
	//		Page:     pageInfo.Page,
	//		PageSize: pageInfo.PageSize,
	//	}, "获取成功", c)
	//}
}
