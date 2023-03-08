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

type HkUserSignApi struct {
}

var hkUserSignService = service.ServiceGroupApp.AppServiceGroup.Community.HkUserSignService

// CreateHkUserSign 创建HkUserSign
// @Tags HkUserSign
// @Summary 创建HkUserSign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUserSign true "创建HkUserSign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserSign/createHkUserSign [post]
func (hkUserSignApi *HkUserSignApi) CreateHkUserSign(c *gin.Context) {
	var hkUserSign community.HkUserSign
	err := c.ShouldBindJSON(&hkUserSign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserSignService.CreateHkUserSign(hkUserSign); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkUserSign 删除HkUserSign
// @Tags HkUserSign
// @Summary 删除HkUserSign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUserSign true "删除HkUserSign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserSign/deleteHkUserSign [delete]
func (hkUserSignApi *HkUserSignApi) DeleteHkUserSign(c *gin.Context) {
	var hkUserSign community.HkUserSign
	err := c.ShouldBindJSON(&hkUserSign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserSignService.DeleteHkUserSign(hkUserSign); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkUserSignByIds 批量删除HkUserSign
// @Tags HkUserSign
// @Summary 批量删除HkUserSign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserSign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkUserSign/deleteHkUserSignByIds [delete]
func (hkUserSignApi *HkUserSignApi) DeleteHkUserSignByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserSignService.DeleteHkUserSignByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkUserSign 更新HkUserSign
// @Tags HkUserSign
// @Summary 更新HkUserSign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUserSign true "更新HkUserSign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserSign/updateHkUserSign [put]
func (hkUserSignApi *HkUserSignApi) UpdateHkUserSign(c *gin.Context) {
	var hkUserSign community.HkUserSign
	err := c.ShouldBindJSON(&hkUserSign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserSignService.UpdateHkUserSign(hkUserSign); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkUserSign 用id查询HkUserSign
// @Tags HkUserSign
// @Summary 用id查询HkUserSign
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkUserSign true "用id查询HkUserSign"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserSign/findHkUserSign [get]
func (hkUserSignApi *HkUserSignApi) FindHkUserSign(c *gin.Context) {
	var hkUserSign community.HkUserSign
	err := c.ShouldBindQuery(&hkUserSign)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkUserSign, err := hkUserSignService.GetHkUserSign(hkUserSign.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkUserSign": rehkUserSign}, c)
	}
}

// GetHkUserSignList 分页获取HkUserSign列表
// @Tags HkUserSign
// @Summary 分页获取HkUserSign列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkUserSignSearch true "分页获取HkUserSign列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserSign/getHkUserSignList [get]
func (hkUserSignApi *HkUserSignApi) GetHkUserSignList(c *gin.Context) {
	var pageInfo communityReq.HkUserSignSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkUserSignService.GetHkUserSignInfoList(pageInfo); err != nil {
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
