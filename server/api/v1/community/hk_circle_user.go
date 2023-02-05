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

type HkCircleUserApi struct {
}

var hkCircleUserService = service.ServiceGroupApp.CommunityServiceGroup.HkCircleUserService

// CreateHkCircleUser 创建HkCircleUser
// @Tags HkCircleUser
// @Summary 创建HkCircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleUser true "创建HkCircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleUser/createHkCircleUser [post]
func (hkCircleUserApi *HkCircleUserApi) CreateHkCircleUser(c *gin.Context) {
	var hkCircleUser community.HkCircleUser
	err := c.ShouldBindJSON(&hkCircleUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleUserService.CreateHkCircleUser(hkCircleUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkCircleUser 删除HkCircleUser
// @Tags HkCircleUser
// @Summary 删除HkCircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleUser true "删除HkCircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleUser/deleteHkCircleUser [delete]
func (hkCircleUserApi *HkCircleUserApi) DeleteHkCircleUser(c *gin.Context) {
	var hkCircleUser community.HkCircleUser
	err := c.ShouldBindJSON(&hkCircleUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleUserService.DeleteHkCircleUser(hkCircleUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkCircleUserByIds 批量删除HkCircleUser
// @Tags HkCircleUser
// @Summary 批量删除HkCircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkCircleUser/deleteHkCircleUserByIds [delete]
func (hkCircleUserApi *HkCircleUserApi) DeleteHkCircleUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleUserService.DeleteHkCircleUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkCircleUser 更新HkCircleUser
// @Tags HkCircleUser
// @Summary 更新HkCircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleUser true "更新HkCircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleUser/updateHkCircleUser [put]
func (hkCircleUserApi *HkCircleUserApi) UpdateHkCircleUser(c *gin.Context) {
	var hkCircleUser community.HkCircleUser
	err := c.ShouldBindJSON(&hkCircleUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleUserService.UpdateHkCircleUser(hkCircleUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkCircleUser 用id查询HkCircleUser
// @Tags HkCircleUser
// @Summary 用id查询HkCircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkCircleUser true "用id查询HkCircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleUser/findHkCircleUser [get]
func (hkCircleUserApi *HkCircleUserApi) FindHkCircleUser(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleUser, err := hkCircleUserService.GetHkCircleUser(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircleUser": rehkCircleUser}, c)
	}
}

// GetHkCircleUserList 分页获取HkCircleUser列表
// @Tags HkCircleUser
// @Summary 分页获取HkCircleUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkCircleUserSearch true "分页获取HkCircleUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleUser/getHkCircleUserList [get]
func (hkCircleUserApi *HkCircleUserApi) GetHkCircleUserList(c *gin.Context) {
	var pageInfo communityReq.HkCircleUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkCircleUserService.GetHkCircleUserInfoList(pageInfo); err != nil {
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
