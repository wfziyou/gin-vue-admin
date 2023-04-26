package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ActivityApi struct {
}

var hkActivityService = service.ServiceGroupApp.AppServiceGroup.Community.ActivityService
var hkActivityUserService = service.ServiceGroupApp.AppServiceGroup.Community.ActivityUserService

// CreateActivity 创建活动
// @Tags 活动
// @Summary 创建活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.CreateActivityReq true "创建活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/activity/createActivity [post]
func (activityApi *ActivityApi) CreateActivity(c *gin.Context) {
	var req communityReq.CreateActivityReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)

	if err := hkActivityService.CreateActivity(userId, req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		appUserService.UpdatePostsTime(userId)
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteActivity 删除活动
// @Tags 活动
// @Summary 删除活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/activity/deleteActivity [delete]
func (activityApi *ActivityApi) DeleteActivity(c *gin.Context) {
	var req request.IdDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityService.DeleteActivityById(req.ID); err != nil {
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
// @Param data body communityReq.UpdateActivityReq true "更新活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/activity/updateActivity [put]
func (activityApi *ActivityApi) UpdateActivity(c *gin.Context) {
	var req communityReq.UpdateActivityReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	_, err = hkActivityService.GetActivity(req.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("活动不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hkActivityService.UpdateActivity(req); err != nil {
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
// @Param data query request.IdSearch true "用id查询活动详情"
// @Success 200 {object} response.Response{data=community.ForumPosts,msg=string}  "返回community.ForumPosts"
// @Router /app/activity/findActivity [get]
func (activityApi *ActivityApi) FindActivity(c *gin.Context) {
	var req request.IdSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkActivity, err := hkActivityService.GetActivity(req.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(rehkActivity, "成功", c)
	}
}

// JoinActivity 加入活动
// @Tags 活动
// @Summary 加入活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.JoinActivityReq true "加入活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/activity/joinActivity [post]
func (activityApi *ActivityApi) JoinActivity(c *gin.Context) {
	var req communityReq.JoinActivityReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	activity, err := hkActivityService.GetActivity(req.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("活动不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)
	if activity.ActivityAddApprove == 0 {
		if activity.PayNum > 0 {
			user, err := appUserService.GetUser(userId)
			if err != nil {
				global.GVA_LOG.Error("查询失败!", zap.Error(err))
				response.FailWithMessage("查询失败", c)
				return
			}

			if user.UserExtend.CurrencyGold < activity.PayNum {
				response.FailWithMessage("金币不够", c)
				return
			}
			err = appUserService.DecreaseGold(userId, activity.PayNum, "加入活动", "", "")
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}
		}
		hkActivityUserService.AddActivityUser(community.ActivityUser{
			ActivityId: activity.ActivityId,
			UserId:     userId,
		})
	} else {
		if err := hkActivityAddRequestService.CreateActivityAddRequest(userId, activity.ID, req.Reason); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
		} else {
			response.OkWithMessage("创建成功", c)
		}
	}
}

// ExitActivity 退出活动
// @Tags 活动
// @Summary 退出活动
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ExitActivityReq true "退出活动"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/activity/exitActivity [post]
func (activityApi *ActivityApi) ExitActivity(c *gin.Context) {
	var req communityReq.ExitActivityReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	activity, err := hkActivityService.GetActivity(req.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("活动不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)
	activityUser, err := hkActivityUserService.GetActivityUser(activity.ID, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.OkWithMessage("成功", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hkActivityUserService.DeleteActivityUser(activityUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

// KickOutActivityUsers 踢出活动的用户
// @Tags 活动
// @Summary 踢出活动的用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.KickOutActivityUsersReq true "踢出活动的用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/activity/kickOutActivityUsers [delete]
func (activityApi *ActivityApi) KickOutActivityUsers(c *gin.Context) {
	var req communityReq.KickOutActivityUsersReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityUserService.DeleteActivityUserByUserIds(req.Id, req.Ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// FindActivityUser 查询活动的用户
// @Tags 活动
// @Summary 查询活动的用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.FindActivityUserReq true "查询活动的用户"
// @Success 200 {object} response.Response{data=community.ActivityUser,msg=string}  "返回community.ActivityUser"
// @Router /app/activity/findActivityUser [get]
func (activityApi *ActivityApi) FindActivityUser(c *gin.Context) {
	var req communityReq.FindActivityUserReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkActivityUser, err := hkActivityUserService.GetActivityUser(req.Id, req.UserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(rehkActivityUser, "成功", c)
	}
}

// GetActivityUserList 分页获取活动的用户列表
// @Tags 活动
// @Summary 分页获取活动的用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ActivityUserSearch true "分页获取活动的用户列表"
// @Success 200 {object}  response.PageResult{List=[]community.ActivityUser,msg=string} "返回[]community.ActivityUser"
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

// GetCircleRecommendActivityList 分页获圈子推荐活动列表
// @Tags 活动
// @Summary 分页获圈子推荐活动列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.CircleRecommendActivitySearch true "分页获圈子推荐活动列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/activity/getCircleRecommendActivityList [get]
func (activityApi *ActivityApi) GetCircleRecommendActivityList(c *gin.Context) {
	var req communityReq.CircleRecommendActivitySearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appForumPostsService.GetGlobalRecommendActivityList(req.PageInfo); err != nil {
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

// GetGlobalRecommendActivityList 分页获取全局推荐活动列表
// @Tags 活动
// @Summary 分页获取全局推荐活动列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.GlobalRecommendActivitySearch true "分页获取全局推荐活动列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/activity/getGlobalRecommendActivityList [get]
func (activityApi *ActivityApi) GetGlobalRecommendActivityList(c *gin.Context) {
	var req communityReq.GlobalRecommendActivitySearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appForumPostsService.GetGlobalRecommendActivityList(req.PageInfo); err != nil {
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

// CreateActivityDynamic 创建活动动态
// @Tags 活动
// @Summary 创建活动动态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.CreateActivityDynamicReq true "创建活动动态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/activity/createActivityDynamic [post]
func (activityApi *ActivityApi) CreateActivityDynamic(c *gin.Context) {
	var req communityReq.CreateActivityDynamicReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	activity, err := hkActivityService.GetActivity(req.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("活动不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	err = hkActivityService.CreateActivityDynamic(userId, activity, req.Content, req.Attachment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("成功", c)
}

// DeleteActivityDynamic 删除活动动态
// @Tags 活动
// @Summary 删除活动动态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除活动动态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/activity/deleteActivityDynamic [delete]
func (activityApi *ActivityApi) DeleteActivityDynamic(c *gin.Context) {
	var req request.IdDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	obj, err := hkActivityService.GetActivityDynamic(req.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)
	if obj.UserId != userId {
		response.FailWithMessage("不能删除他人的活动动态", c)
		return
	}

	err = hkActivityService.DeleteActivityDynamic(obj.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("成功", c)
}

// DeleteActivityDynamicByIds 批量删除活动动态
// @Tags 活动
// @Summary 批量删除活动动态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除活动动态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/activity/deleteActivityDynamicByIds [delete]
func (activityApi *ActivityApi) DeleteActivityDynamicByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
}

// GetActivityDynamicList 分页获取活动的动态列表
// @Tags 活动
// @Summary 分页获取活动的动态列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ActivityDynamicSearch true "分页获取活动的动态列表"
// @Success 200 {object}  response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回[]community.ForumPostsBaseInfo"
// @Router /app/activity/getActivityDynamicList [get]
func (activityApi *ActivityApi) GetActivityDynamicList(c *gin.Context) {
	var req communityReq.ActivityDynamicSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkActivityService.GetActivityDynamicList(req.Id, req.PageInfo); err != nil {
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

// DeleteActivityAddRequest 删除活动报名申请
// @Tags 活动
// @Summary 删除活动报名申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除活动报名申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/activity/deleteActivityAddRequest [delete]
func (activityApi *ActivityApi) DeleteActivityAddRequest(c *gin.Context) {
	var req request.IdDelete

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityAddRequestService.DeleteActivityAddRequestById(req.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteActivityAddRequestByIds 批量删除活动报名申请
// @Tags 活动
// @Summary 批量删除活动报名申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除活动报名申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/activity/deleteActivityAddRequestByIds [delete]
func (activityApi *ActivityApi) DeleteActivityAddRequestByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityAddRequestService.DeleteActivityAddRequestByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateActivityAddRequest 活动报名申请审批
// @Tags 活动
// @Summary 活动报名申请审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.UpdateActivityAddRequestReq true "活动报名申请审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/activity/updateActivityAddRequest [put]
func (activityApi *ActivityApi) UpdateActivityAddRequest(c *gin.Context) {
	var req communityReq.UpdateActivityAddRequestReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	request, err := hkActivityAddRequestService.GetActivityAddRequest(req.Id)
	if err != nil {
		response.FailWithMessage("没有找到活动请求", c)
		return
	}
	userId := utils.GetUserID(c)
	if err := hkActivityAddRequestService.UpdateActivityAddRequestStatus(request.ID, userId, req.CheckStatus); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

// FindActivityAddRequest 用id查询活动报名申请
// @Tags 活动
// @Summary 用id查询活动报名申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询活动报名申请"
// @Success 200 {object} response.Response{data=community.ActivityAddRequest,msg=string}  "返回community.ActivityAddRequest"
// @Router /app/activity/findActivityAddRequest [get]
func (activityApi *ActivityApi) FindActivityAddRequest(c *gin.Context) {
	var req request.IdSearch

	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkActivityAddRequest, err := hkActivityAddRequestService.GetActivityAddRequest(req.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(rehkActivityAddRequest, "成功", c)
	}
}

// GetActivityAddRequestList 分页获取活动报名申请列表
// @Tags 活动
// @Summary 分页获取活动报名申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ActivityAddRequestSearch true "分页获取活动报名申请列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/activity/getActivityAddRequestList [get]
func (activityApi *ActivityApi) GetActivityAddRequestList(c *gin.Context) {
	var req communityReq.ActivityAddRequestSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkActivityAddRequestService.GetActivityAddRequestInfoList(req); err != nil {
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
