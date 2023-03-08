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

type HkFocusUserApi struct {
}

var hkFocusUserService = service.ServiceGroupApp.AppServiceGroup.Community.HkFocusUserService

// CreateHkFocusUser 创建HkFocusUser
// @Tags HkFocusUser
// @Summary 创建HkFocusUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkFocusUser true "创建HkFocusUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkFocusUser/createHkFocusUser [post]
func (hkFocusUserApi *HkFocusUserApi) CreateHkFocusUser(c *gin.Context) {
	var hkFocusUser community.HkFocusUser
	err := c.ShouldBindJSON(&hkFocusUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkFocusUserService.CreateHkFocusUser(hkFocusUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkFocusUser 删除HkFocusUser
// @Tags HkFocusUser
// @Summary 删除HkFocusUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkFocusUser true "删除HkFocusUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkFocusUser/deleteHkFocusUser [delete]
func (hkFocusUserApi *HkFocusUserApi) DeleteHkFocusUser(c *gin.Context) {
	var hkFocusUser community.HkFocusUser
	err := c.ShouldBindJSON(&hkFocusUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkFocusUserService.DeleteHkFocusUser(hkFocusUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkFocusUserByIds 批量删除HkFocusUser
// @Tags HkFocusUser
// @Summary 批量删除HkFocusUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkFocusUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkFocusUser/deleteHkFocusUserByIds [delete]
func (hkFocusUserApi *HkFocusUserApi) DeleteHkFocusUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkFocusUserService.DeleteHkFocusUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkFocusUser 更新HkFocusUser
// @Tags HkFocusUser
// @Summary 更新HkFocusUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkFocusUser true "更新HkFocusUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkFocusUser/updateHkFocusUser [put]
func (hkFocusUserApi *HkFocusUserApi) UpdateHkFocusUser(c *gin.Context) {
	var hkFocusUser community.HkFocusUser
	err := c.ShouldBindJSON(&hkFocusUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkFocusUserService.UpdateHkFocusUser(hkFocusUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkFocusUser 用id查询HkFocusUser
// @Tags HkFocusUser
// @Summary 用id查询HkFocusUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkFocusUser true "用id查询HkFocusUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkFocusUser/findHkFocusUser [get]
func (hkFocusUserApi *HkFocusUserApi) FindHkFocusUser(c *gin.Context) {
	var hkFocusUser community.HkFocusUser
	err := c.ShouldBindQuery(&hkFocusUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkFocusUser, err := hkFocusUserService.GetHkFocusUser(hkFocusUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkFocusUser": rehkFocusUser}, c)
	}
}

// GetHkFocusUserList 分页获取HkFocusUser列表
// @Tags HkFocusUser
// @Summary 分页获取HkFocusUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkFocusUserSearch true "分页获取HkFocusUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkFocusUser/getHkFocusUserList [get]
func (hkFocusUserApi *HkFocusUserApi) GetHkFocusUserList(c *gin.Context) {
	var pageInfo communityReq.HkFocusUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkFocusUserService.GetHkFocusUserInfoList(pageInfo); err != nil {
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
