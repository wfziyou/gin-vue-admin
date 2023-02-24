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

type GeneralApi struct{}

/*************************************
常规
**************************************/

// FindProtocol 用id查询协议
// @Tags APP_General
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
// @Tags APP_General
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

// GetProtocolList 分页获取Protocol列表
// @Tags APP_General
// @Summary 分页获取Protocol列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.ProtocolSearch true "分页获取Protocol列表"
// @Success 200 {object}  response.PageResult{List=[]general.Protocol,msg=string} "返回general.Protocol"
// @Router /app/general/getProtocolList [get]
func (generalApi *GeneralApi) GetProtocolList(c *gin.Context) {
	var pageInfo generalReq.ProtocolSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appProtocolService.GetProtocolInfoList(pageInfo); err != nil {
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

// CreateBugReport Bug反馈
// @Tags APP_General
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

// FindBugReport 用id查询BugReport
// @Tags APP_General
// @Summary 用id查询BugReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询BugReport"
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

// GetBugReportList 分页获取BugReport列表
// @Tags APP_General
// @Summary 分页获取BugReport列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.BugReportSearch true "分页获取BugReport列表"
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

// FindMiniProgram 用id查询MiniProgram
// @Tags APP_General
// @Summary 用id查询MiniProgram
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询MiniProgram"
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
