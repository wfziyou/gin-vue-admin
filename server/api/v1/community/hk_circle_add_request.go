package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HkCircleAddRequestApi struct {
}

var hkCircleAddRequestService = service.ServiceGroupApp.CommunityServiceGroup.HkCircleAddRequestService

// CreateHkCircleAddRequest 创建HkCircleAddRequest
// @Tags HkCircleAddRequest
// @Summary 创建HkCircleAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleAddRequest true "创建HkCircleAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleAddRequest/createHkCircleAddRequest [post]
func (hkCircleAddRequestApi *HkCircleAddRequestApi) CreateHkCircleAddRequest(c *gin.Context) {
	var hkCircleAddRequest community.HkCircleAddRequest
	err := c.ShouldBindJSON(&hkCircleAddRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleAddRequestService.CreateHkCircleAddRequest(hkCircleAddRequest); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkCircleAddRequest 删除HkCircleAddRequest
// @Tags HkCircleAddRequest
// @Summary 删除HkCircleAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleAddRequest true "删除HkCircleAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleAddRequest/deleteHkCircleAddRequest [delete]
func (hkCircleAddRequestApi *HkCircleAddRequestApi) DeleteHkCircleAddRequest(c *gin.Context) {
	var hkCircleAddRequest community.HkCircleAddRequest
	err := c.ShouldBindJSON(&hkCircleAddRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleAddRequestService.DeleteHkCircleAddRequest(hkCircleAddRequest); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkCircleAddRequestByIds 批量删除HkCircleAddRequest
// @Tags HkCircleAddRequest
// @Summary 批量删除HkCircleAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkCircleAddRequest/deleteHkCircleAddRequestByIds [delete]
func (hkCircleAddRequestApi *HkCircleAddRequestApi) DeleteHkCircleAddRequestByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleAddRequestService.DeleteHkCircleAddRequestByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkCircleAddRequest 更新HkCircleAddRequest
// @Tags HkCircleAddRequest
// @Summary 更新HkCircleAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleAddRequest true "更新HkCircleAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleAddRequest/updateHkCircleAddRequest [put]
func (hkCircleAddRequestApi *HkCircleAddRequestApi) UpdateHkCircleAddRequest(c *gin.Context) {
	var hkCircleAddRequest community.HkCircleAddRequest
	err := c.ShouldBindJSON(&hkCircleAddRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleAddRequestService.UpdateHkCircleAddRequest(hkCircleAddRequest); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkCircleAddRequest 用id查询HkCircleAddRequest
// @Tags HkCircleAddRequest
// @Summary 用id查询HkCircleAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkCircleAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleAddRequest/findHkCircleAddRequest [get]
func (hkCircleAddRequestApi *HkCircleAddRequestApi) FindHkCircleAddRequest(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleAddRequest, err := hkCircleAddRequestService.GetHkCircleAddRequest(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircleAddRequest": rehkCircleAddRequest}, c)
	}
}

// GetHkCircleAddRequestList 分页获取HkCircleAddRequest列表
// @Tags HkCircleAddRequest
// @Summary 分页获取HkCircleAddRequest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkCircleAddRequestSearch true "分页获取HkCircleAddRequest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleAddRequest/getHkCircleAddRequestList [get]
func (hkCircleAddRequestApi *HkCircleAddRequestApi) GetHkCircleAddRequestList(c *gin.Context) {
	var pageInfo communityReq.HkCircleAddRequestSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkCircleAddRequestService.GetHkCircleAddRequestInfoList(pageInfo); err != nil {
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
