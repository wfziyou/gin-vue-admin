package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HelpCenterApi struct {
}

// GetHelpTypeList (管理员)获取帮助类型列表
// @Tags 反馈中心
// @Summary (管理员)获取帮助类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]community.HelpType,msg=string} "返回[]community.HelpType"
// @Router /app/helpCenter/getHelpTypeList [get]
func (api *HelpCenterApi) GetHelpTypeList(c *gin.Context) {
	if list, err := hkHelpTypeService.GetHelpTypeInfoList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// GetHelpList 获取帮助列表
// @Tags 反馈中心
// @Summary 获取帮助列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HelpSearch true "获取帮助列表"
// @Success 200 {object}  response.PageResult{List=[]community.HelpBaseInfo,msg=string} "返回[]community.HelpBaseInfo"
// @Router /app/helpCenter/getHelpList [get]
func (api *HelpCenterApi) GetHelpList(c *gin.Context) {
	var pageInfo communityReq.HelpSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkHelpService.GetHelpList(pageInfo); err != nil {
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

// GetMainHelpList 获取主页帮助列表
// @Tags 反馈中心
// @Summary 获取主页帮助列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success  200   {object}  response.Response{data=community.MainHelp,msg=string}  "返回community.MainHelp"
// @Router /app/helpCenter/getMainHelpList [get]
func (api *HelpCenterApi) GetMainHelpList(c *gin.Context) {
	if obj, err := hkHelpService.GetMainHelpList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(obj, "获取成功", c)
	}
}

// FindHelp 用id查询帮助
// @Tags 反馈中心
// @Summary 用id查询帮助
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询帮助"
// @Success 200 {object} response.Response{data=community.HelpInfo,msg=string}  "返回community.HelpInfo"
// @Router /app/helpCenter/findHelp [get]
func (api *HelpCenterApi) FindHelp(c *gin.Context) {
	var req request.IdSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if help, err := hkHelpService.GetHelpInfo(req.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		userId := utils.GetUserID(c)
		if list, err := hkHelpService.GetHelpThumbsUpList(userId, req.ID); err == nil {
			if len(list) > 0 {
				help.ThumbsUp = list[0].ThumbsUp
			}
		}
		response.OkWithDetailed(help, "成功", c)
	}
}

// HelpThumbsUp 帮助点赞
// @Tags 反馈中心
// @Summary 帮助点赞
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.HelpThumbsUpReq true "帮助点赞"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/helpCenter/helpThumbsUp [post]
func (api *HelpCenterApi) HelpThumbsUp(c *gin.Context) {
	var req communityReq.HelpThumbsUpReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := hkHelpService.SetHelpThumbsUp(userId, req); err != nil {
		global.GVA_LOG.Error("点赞失败!", zap.Error(err))
		response.FailWithMessage("点赞失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

// CreateFeedback 创建反馈
// @Tags 反馈中心
// @Summary 创建反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.CreateFeedbackReq true "创建反馈"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/helpCenter/createFeedback [post]
func (api *HelpCenterApi) CreateFeedback(c *gin.Context) {
	var req communityReq.CreateFeedbackReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := hkFeedbackService.CreateFeedback(userId, req.Des, req.Attachment, req.Phone); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

// DeleteFeedback (管理员)删除反馈
// @Tags 反馈中心
// @Summary (管理员)删除反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.Feedback true "(管理员)删除反馈"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/helpCenter/deleteFeedback [delete]
func (api *HelpCenterApi) DeleteFeedback(c *gin.Context) {
	var req request.IdDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	obj, err := hkFeedbackService.GetFeedback(req.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("反馈不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hkFeedbackService.DeleteFeedback(obj); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFeedbackByIds (管理员)批量删除反馈
// @Tags 反馈中心
// @Summary (管理员)批量删除反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "(管理员)批量删除反馈"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/helpCenter/deleteFeedbackByIds [delete]
func (api *HelpCenterApi) DeleteFeedbackByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkFeedbackService.DeleteFeedbackByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFeedback (管理员)更新反馈
// @Tags 反馈中心
// @Summary (管理员)更新反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.Feedback true "(管理员)更新反馈"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/helpCenter/updateFeedback [put]
func (api *HelpCenterApi) UpdateFeedback(c *gin.Context) {
	var req communityReq.UpdateFeedbackReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hkFeedbackService.UpdateFeedback(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFeedback 用id查询反馈
// @Tags 反馈中心
// @Summary 用id查询反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询反馈"
// @Success 200 {object} response.Response{data=community.Feedback,msg=string}  "返回community.Feedback"
// @Router /app/helpCenter/findFeedback [get]
func (api *HelpCenterApi) FindFeedback(c *gin.Context) {
	var req request.IdSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkFeedback, err := hkFeedbackService.GetFeedback(req.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(rehkFeedback, "成功", c)
	}
}

// GetFeedbackList 分页获取反馈列表
// @Tags 反馈中心
// @Summary 分页获取反馈列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.FeedbackSearch true "分页获取反馈列表"
// @Success 200 {object}  response.PageResult{List=[]community.Feedback,msg=string} "返回[]community.Feedback"
// @Router /app/helpCenter/getFeedbackList [get]
func (api *HelpCenterApi) GetFeedbackList(c *gin.Context) {
	var pageInfo communityReq.FeedbackSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := hkFeedbackService.GetFeedbackInfoList(userId, pageInfo); err != nil {
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
