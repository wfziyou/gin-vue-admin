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

type HkCircleRequestApi struct {
}

var hkCircleRequestService = service.ServiceGroupApp.CommunityServiceGroup.HkCircleRequestService

// CreateHkCircleRequest 创建HkCircleRequest
// @Tags HkCircleRequest
// @Summary 创建HkCircleRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleRequest true "创建HkCircleRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleRequest/createHkCircleRequest [post]
func (hkCircleRequestApi *HkCircleRequestApi) CreateHkCircleRequest(c *gin.Context) {
	var hkCircleRequest community.HkCircleRequest
	err := c.ShouldBindJSON(&hkCircleRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleRequestService.CreateHkCircleRequest(hkCircleRequest); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkCircleRequest 删除HkCircleRequest
// @Tags HkCircleRequest
// @Summary 删除HkCircleRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleRequest true "删除HkCircleRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleRequest/deleteHkCircleRequest [delete]
func (hkCircleRequestApi *HkCircleRequestApi) DeleteHkCircleRequest(c *gin.Context) {
	var hkCircleRequest community.HkCircleRequest
	err := c.ShouldBindJSON(&hkCircleRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleRequestService.DeleteHkCircleRequest(hkCircleRequest); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkCircleRequestByIds 批量删除HkCircleRequest
// @Tags HkCircleRequest
// @Summary 批量删除HkCircleRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkCircleRequest/deleteHkCircleRequestByIds [delete]
func (hkCircleRequestApi *HkCircleRequestApi) DeleteHkCircleRequestByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleRequestService.DeleteHkCircleRequestByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkCircleRequest 更新HkCircleRequest
// @Tags HkCircleRequest
// @Summary 更新HkCircleRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleRequest true "更新HkCircleRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleRequest/updateHkCircleRequest [put]
func (hkCircleRequestApi *HkCircleRequestApi) UpdateHkCircleRequest(c *gin.Context) {
	var hkCircleRequest community.HkCircleRequest
	err := c.ShouldBindJSON(&hkCircleRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleRequestService.UpdateHkCircleRequest(hkCircleRequest); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkCircleRequest 用id查询HkCircleRequest
// @Tags HkCircleRequest
// @Summary 用id查询HkCircleRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkCircleRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleRequest/findHkCircleRequest [get]
func (hkCircleRequestApi *HkCircleRequestApi) FindHkCircleRequest(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleRequest, err := hkCircleRequestService.GetHkCircleRequest(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircleRequest": rehkCircleRequest}, c)
	}
}

// GetHkCircleRequestList 分页获取HkCircleRequest列表
// @Tags HkCircleRequest
// @Summary 分页获取HkCircleRequest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkCircleRequestSearch true "分页获取HkCircleRequest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleRequest/getHkCircleRequestList [get]
func (hkCircleRequestApi *HkCircleRequestApi) GetHkCircleRequestList(c *gin.Context) {
	var pageInfo communityReq.HkCircleRequestSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkCircleRequestService.GetHkCircleRequestInfoList(pageInfo); err != nil {
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
