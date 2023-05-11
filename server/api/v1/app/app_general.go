package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GeneralApi struct{}

var hkFeedbackService = service.ServiceGroupApp.AppServiceGroup.General.FeedbackService
var hkFeedbackTypeService = service.ServiceGroupApp.AppServiceGroup.General.FeedbackTypeService

/*************************************
常规
**************************************/

// FindProtocol 用id查询协议
// @Tags 常规方法
// @Summary 用id查询协议
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询协议"
// @Success 200 {object}  response.Response{data=general.Protocol,msg=string}  "返回general.Protocol"
// @Router /app/general/findProtocol [get]
func (generalApi *GeneralApi) FindProtocol(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkProtocol, err := appProtocolService.GetProtocol(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rehkProtocol, c)
	}
}

// FindProtocolByName 用名字查询协议
// @Tags 常规方法
// @Summary 用名字查询协议
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.FindProtocolByName true "用名字查询协议"
// @Success 200 {object}  response.Response{data=general.Protocol,msg=string}  "返回general.Protocol"
// @Router /app/general/findProtocolByName [get]
func (generalApi *GeneralApi) FindProtocolByName(c *gin.Context) {
	var req generalReq.FindProtocolByName
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkProtocol, err := appProtocolService.GetProtocolByName(req.Name); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rehkProtocol, c)
	}
}

// FindMiniProgram 用id查询小程序
// @Tags 常规方法
// @Summary 用id查询小程序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询小程序"
// @Success 200 {object}  response.Response{data=general.MiniProgramBaseInfo,msg=string}  "返回general.MiniProgramBaseInfo"
// @Router /app/general/findMiniProgram [get]
func (generalApi *GeneralApi) FindMiniProgram(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if data, err := appMiniProgramService.GetMiniProgram(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(data, c)
	}
}

// GetFeedbackTypeList 获取反馈类型列表
// @Tags 常规方法
// @Summary 获取反馈类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]general.FeedbackType,msg=string} "返回[]general.FeedbackType"
// @Router /app/general/getFeedbackTypeList [get]
func (generalApi *GeneralApi) GetFeedbackTypeList(c *gin.Context) {
	if list, err := hkFeedbackTypeService.GetFeedbackTypeInfoList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// CreateFeedback 创建反馈
// @Tags 常规方法
// @Summary 创建反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.Feedback true "创建反馈"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/general/createFeedback [post]
func (generalApi *GeneralApi) CreateFeedback(c *gin.Context) {
	var req generalReq.CreateFeedbackReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	obj, err := hkFeedbackTypeService.GetFeedbackType(req.TypeId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("类型不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := hkFeedbackService.CreateFeedback(userId, obj.ID, obj.Name, req.Des, req.Attachment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

// DeleteFeedback 删除反馈
// @Tags 常规方法
// @Summary 删除反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.Feedback true "删除反馈"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/general/deleteFeedback [delete]
func (generalApi *GeneralApi) DeleteFeedback(c *gin.Context) {
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

// DeleteFeedbackByIds 批量删除反馈
// @Tags 常规方法
// @Summary 批量删除反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除反馈"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/general/deleteFeedbackByIds [delete]
func (generalApi *GeneralApi) DeleteFeedbackByIds(c *gin.Context) {
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

// UpdateFeedback 更新反馈
// @Tags 常规方法
// @Summary 更新反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.Feedback true "更新反馈"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/general/updateFeedback [put]
func (generalApi *GeneralApi) UpdateFeedback(c *gin.Context) {
	var req generalReq.UpdateFeedbackReq
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
// @Tags 常规方法
// @Summary 用id查询Feedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query general.Feedback true "用id查询Feedback"
// @Success 200 {object} response.Response{data=general.Feedback,msg=string}  "返回general.Feedback"
// @Router /app/general/findFeedback [get]
func (generalApi *GeneralApi) FindFeedback(c *gin.Context) {
	var hkFeedback general.Feedback
	err := c.ShouldBindQuery(&hkFeedback)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkFeedback, err := hkFeedbackService.GetFeedback(hkFeedback.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(rehkFeedback, "成功", c)
	}
}

// GetFeedbackList 分页获取反馈列表
// @Tags 常规方法
// @Summary 分页获取反馈列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.FeedbackSearch true "分页获取反馈列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Success 200 {object}  response.PageResult{List=[]general.Feedback,msg=string} "返回[]general.Feedback"
// @Router /app/general/getFeedbackList [get]
func (generalApi *GeneralApi) GetFeedbackList(c *gin.Context) {
	var pageInfo generalReq.FeedbackSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkFeedbackService.GetFeedbackInfoList(pageInfo); err != nil {
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

// FindDraft 用id查询草稿
// @Tags 帖子
// @Summary 用id查询草稿
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询草稿"
// @Success 200 {object}  response.Response{data=community.ForumPosts,msg=string}  "返回community.ForumPosts"
// @Router /app/general/findDraft [get]
func (generalApi *GeneralApi) FindDraft(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumPosts, err := appForumPostsService.GetDraft(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rehkForumPosts, c)
	}
}

// GetDraftList 分页获取草稿列表
// @Tags 常规方法
// @Summary 分页获取草稿列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.DraftSearch true "分页获取草稿列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Success 200 {object}  response.PageResult{List=[]general.Feedback,msg=string} "返回[]general.Feedback"
// @Router /app/general/getDraftList [get]
func (generalApi *GeneralApi) GetDraftList(c *gin.Context) {
	var req generalReq.DraftSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := appForumPostsService.GetDraftList(userId, req); err != nil {
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

// DeleteAllDraft 删除所有草稿
// @Tags 常规方法
// @Summary 删除所有草稿
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body generalReq.DeleteAllDraftReq true "删除所有草稿"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/general/deleteAllDraft [delete]
func (generalApi *GeneralApi) DeleteAllDraft(c *gin.Context) {
	var req generalReq.DeleteAllDraftReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := appForumPostsService.DeleteAllDraft(userId, req.Category); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDraft 删除草稿
// @Tags 常规方法
// @Summary 删除草稿
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除草稿"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/general/deleteDraft [delete]
func (generalApi *GeneralApi) DeleteDraft(c *gin.Context) {
	var req request.IdDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := appForumPostsService.DeleteDraft(userId, req.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDraftByIds 批量删除草稿
// @Tags 常规方法
// @Summary 批量删除草稿
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除草稿"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/general/deleteDraftByIds [delete]
func (generalApi *GeneralApi) DeleteDraftByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := appForumPostsService.DeleteDraftByIds(userId, IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDraft 更新草稿
// @Tags 常规方法
// @Summary 更新草稿
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body generalReq.UpdateDraftReq true "更新草稿"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/general/updateDraft [post]
func (generalApi *GeneralApi) UpdateDraft(c *gin.Context) {
	var req generalReq.UpdateDraftReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := appForumPostsService.UpdateDraft(userId, req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetUserCoverImageList 获取用户主页封面列表
// @Tags 常规方法
// @Summary 获取用户主页封面列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.UserCoverImageSearch true "获取用户主页封面列表"
// @Success 200 {object}  response.PageResult{List=[]community.UserCoverImage,msg=string} "返回common.User"
// @Router /app/general/getUserCoverImageList [get]
func (generalApi *GeneralApi) GetUserCoverImageList(c *gin.Context) {
	var req communityReq.UserCoverImageSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkUserCoverImageService.GetUserCoverImageList(req); err != nil {
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
