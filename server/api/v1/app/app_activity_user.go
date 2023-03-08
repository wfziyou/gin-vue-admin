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

type HkActivityUserApi struct {
}

var hkActivityUserService = service.ServiceGroupApp.AppServiceGroup.Community.HkActivityUserService

// CreateHkActivityUser 创建HkActivityUser
// @Tags HkActivityUser
// @Summary 创建HkActivityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivityUser true "创建HkActivityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityUser/createHkActivityUser [post]
func (hkActivityUserApi *HkActivityUserApi) CreateHkActivityUser(c *gin.Context) {
	var hkActivityUser community.HkActivityUser
	err := c.ShouldBindJSON(&hkActivityUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityUserService.CreateHkActivityUser(hkActivityUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkActivityUser 删除HkActivityUser
// @Tags HkActivityUser
// @Summary 删除HkActivityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivityUser true "删除HkActivityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivityUser/deleteHkActivityUser [delete]
func (hkActivityUserApi *HkActivityUserApi) DeleteHkActivityUser(c *gin.Context) {
	var hkActivityUser community.HkActivityUser
	err := c.ShouldBindJSON(&hkActivityUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityUserService.DeleteHkActivityUser(hkActivityUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkActivityUserByIds 批量删除HkActivityUser
// @Tags HkActivityUser
// @Summary 批量删除HkActivityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkActivityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkActivityUser/deleteHkActivityUserByIds [delete]
func (hkActivityUserApi *HkActivityUserApi) DeleteHkActivityUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityUserService.DeleteHkActivityUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkActivityUser 更新HkActivityUser
// @Tags HkActivityUser
// @Summary 更新HkActivityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivityUser true "更新HkActivityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkActivityUser/updateHkActivityUser [put]
func (hkActivityUserApi *HkActivityUserApi) UpdateHkActivityUser(c *gin.Context) {
	var hkActivityUser community.HkActivityUser
	err := c.ShouldBindJSON(&hkActivityUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkActivityUserService.UpdateHkActivityUser(hkActivityUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkActivityUser 用id查询HkActivityUser
// @Tags HkActivityUser
// @Summary 用id查询HkActivityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkActivityUser true "用id查询HkActivityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkActivityUser/findHkActivityUser [get]
func (hkActivityUserApi *HkActivityUserApi) FindHkActivityUser(c *gin.Context) {
	var hkActivityUser community.HkActivityUser
	err := c.ShouldBindQuery(&hkActivityUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkActivityUser, err := hkActivityUserService.GetHkActivityUser(hkActivityUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkActivityUser": rehkActivityUser}, c)
	}
}

// GetHkActivityUserList 分页获取HkActivityUser列表
// @Tags HkActivityUser
// @Summary 分页获取HkActivityUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkActivityUserSearch true "分页获取HkActivityUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityUser/getHkActivityUserList [get]
func (hkActivityUserApi *HkActivityUserApi) GetHkActivityUserList(c *gin.Context) {
	var pageInfo communityReq.HkActivityUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkActivityUserService.GetHkActivityUserInfoList(pageInfo); err != nil {
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
