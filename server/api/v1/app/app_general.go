package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
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

// CreateBugReport Bug反馈
// @Tags 常规方法
// @Summary Bug反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.BugReport true "Bug反馈"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/general/createBugReport [post]
func (generalApi *GeneralApi) CreateBugReport(c *gin.Context) {
	var req generalReq.BugReportCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appBugReportService.CreateBugReport(general.BugReport{
		UserId:                 utils.GetUserID(c),
		Title:                  req.Title,
		Content:                req.Content,
		ContentAttachment:      req.ContentAttachment,
		ExpectedResult:         req.ExpectedResult,
		ActualResult:           req.ActualResult,
		ActualResultAttachment: req.ActualResultAttachment,
		OtherInfo:              req.OtherInfo,
	}); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// FindBugReport 用id查询Bug反馈
// @Tags 常规方法
// @Summary 用id查询Bug反馈
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询Bug反馈"
// @Success 200 {object}  response.Response{data=general.BugReport,msg=string}  "返回general.BugReport"
// @Router /app/general/findBugReport [get]
func (generalApi *GeneralApi) FindBugReport(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkBugReport, err := appBugReportService.GetBugReport(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rehkBugReport, c)
	}
}

// GetBugReportList 分页获取Bug反馈列表
// @Tags 常规方法
// @Summary 分页获取Bug反馈列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.BugReportSearch true "分页获取Bug反馈列表"
// @Success 200 {object}  response.PageResult{List=[]general.BugReport,msg=string} "返回general.BugReport"
// @Router /app/general/getBugReportList [get]
func (generalApi *GeneralApi) GetBugReportList(c *gin.Context) {
	var pageInfo generalReq.BugReportSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.UserId = utils.GetUserID(c)
	if list, total, err := appBugReportService.AppGetBugReportInfoList(pageInfo); err != nil {
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

// FindMiniProgram 用id查询小程序
// @Tags 常规方法
// @Summary 用id查询小程序
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询小程序"
// @Success 200 {object}  response.Response{data=general.MiniProgramBaseInfo,msg=string}  "返回general.BugReport"
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
// @Tags FeedbackType
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
// @Tags Feedback
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
// @Tags Feedback
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
// @Tags Feedback
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
// @Tags Feedback
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
// @Tags Feedback
// @Summary 用id查询Feedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query general.Feedback true "用id查询Feedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
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
		response.OkWithData(gin.H{"rehkFeedback": rehkFeedback}, c)
	}
}

// GetFeedbackList 分页获取反馈列表
// @Tags Feedback
// @Summary 分页获取反馈列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.FeedbackSearch true "分页获取反馈列表"
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
