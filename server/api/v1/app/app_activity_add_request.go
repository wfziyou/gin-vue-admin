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

type ActivityAddRequestApi struct {
}

var hkActivityAddRequestService = service.ServiceGroupApp.AppServiceGroup.Community.ActivityAddRequestService

// CreateActivityAddRequest 创建ActivityAddRequest
// @Tags ActivityAddRequest
// @Summary 创建ActivityAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ActivityAddRequest true "创建ActivityAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /activityAddRequest/createActivityAddRequest [post]
func (hkActivityAddRequestApi *ActivityAddRequestApi) CreateActivityAddRequest(c *gin.Context) {
	var activityAddRequest community.ActivityAddRequest
	err := c.ShouldBindJSON(&activityAddRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityAddRequestService.CreateActivityAddRequest(&activityAddRequest); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteActivityAddRequest 删除ActivityAddRequest
// @Tags ActivityAddRequest
// @Summary 删除ActivityAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ActivityAddRequest true "删除ActivityAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /activityAddRequest/deleteActivityAddRequest [delete]
func (hkActivityAddRequestApi *ActivityAddRequestApi) DeleteActivityAddRequest(c *gin.Context) {
	var activityAddRequest community.ActivityAddRequest
	err := c.ShouldBindJSON(&activityAddRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityAddRequestService.DeleteActivityAddRequest(activityAddRequest); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteActivityAddRequestByIds 批量删除ActivityAddRequest
// @Tags ActivityAddRequest
// @Summary 批量删除ActivityAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ActivityAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /activityAddRequest/deleteActivityAddRequestByIds [delete]
func (hkActivityAddRequestApi *ActivityAddRequestApi) DeleteActivityAddRequestByIds(c *gin.Context) {
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

// UpdateActivityAddRequest 更新ActivityAddRequest
// @Tags ActivityAddRequest
// @Summary 更新ActivityAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.ActivityAddRequest true "更新ActivityAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /activityAddRequest/updateActivityAddRequest [put]
func (hkActivityAddRequestApi *ActivityAddRequestApi) UpdateActivityAddRequest(c *gin.Context) {
	var activityAddRequest community.ActivityAddRequest
	err := c.ShouldBindJSON(&activityAddRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityAddRequestService.UpdateActivityAddRequest(activityAddRequest); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindActivityAddRequest 用id查询ActivityAddRequest
// @Tags ActivityAddRequest
// @Summary 用id查询ActivityAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.ActivityAddRequest true "用id查询ActivityAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /activityAddRequest/findActivityAddRequest [get]
func (hkActivityAddRequestApi *ActivityAddRequestApi) FindActivityAddRequest(c *gin.Context) {
	var activityAddRequest community.ActivityAddRequest
	err := c.ShouldBindQuery(&activityAddRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkActivityAddRequest, err := hkActivityAddRequestService.GetActivityAddRequest(activityAddRequest.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkActivityAddRequest": rehkActivityAddRequest}, c)
	}
}

// GetActivityAddRequestList 分页获取ActivityAddRequest列表
// @Tags ActivityAddRequest
// @Summary 分页获取ActivityAddRequest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ActivityAddRequestSearch true "分页获取ActivityAddRequest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /activityAddRequest/getActivityAddRequestList [get]
func (hkActivityAddRequestApi *ActivityAddRequestApi) GetActivityAddRequestList(c *gin.Context) {
	var pageInfo communityReq.ActivityAddRequestSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkActivityAddRequestService.GetActivityAddRequestInfoList(pageInfo); err != nil {
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
