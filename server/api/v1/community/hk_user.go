package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HkUserApi struct {
}

// CreateHkUser 创建HkUser
// @Tags HkUser
// @Summary 创建HkUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUser true "创建HkUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUser/createHkUser [post]
func (hkUserApi *HkUserApi) CreateHkUser(c *gin.Context) {
	var hkUser community.HkUser
	err := c.ShouldBindJSON(&hkUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserService.CreateHkUser(hkUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkUser 删除HkUser
// @Tags HkUser
// @Summary 删除HkUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUser true "删除HkUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUser/deleteHkUser [delete]
func (hkUserApi *HkUserApi) DeleteHkUser(c *gin.Context) {
	var hkUser community.HkUser
	err := c.ShouldBindJSON(&hkUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserService.DeleteHkUser(hkUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkUserByIds 批量删除HkUser
// @Tags HkUser
// @Summary 批量删除HkUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkUser/deleteHkUserByIds [delete]
func (hkUserApi *HkUserApi) DeleteHkUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserService.DeleteHkUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkUser 更新HkUser
// @Tags HkUser
// @Summary 更新HkUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUser true "更新HkUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUser/updateHkUser [put]
func (hkUserApi *HkUserApi) UpdateHkUser(c *gin.Context) {
	var hkUser community.HkUser
	err := c.ShouldBindJSON(&hkUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserService.UpdateHkUser(hkUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkUser 用id查询HkUser
// @Tags HkUser
// @Summary 用id查询HkUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkUser true "用id查询HkUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUser/findHkUser [get]
func (hkUserApi *HkUserApi) FindHkUser(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkUser, err := hkUserService.GetHkUser(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(gin.H{"rehkUser": rehkUser}, c)
	}
}

// GetHkUserList 分页获取HkUser列表
// @Tags HkUser
// @Summary 分页获取HkUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkUserSearch true "分页获取HkUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUser/getHkUserList [get]
func (hkUserApi *HkUserApi) GetHkUserList(c *gin.Context) {
	var pageInfo communityReq.HkUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkUserService.GetHkUserInfoList(pageInfo); err != nil {
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
