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

type HkUserExtendApi struct {
}

var hkUserExtendService = service.ServiceGroupApp.CommunityServiceGroup.HkUserExtendService

// CreateHkUserExtend 创建HkUserExtend
// @Tags HkUserExtend
// @Summary 创建HkUserExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUserExtend true "创建HkUserExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserExtend/createHkUserExtend [post]
func (hkUserExtendApi *HkUserExtendApi) CreateHkUserExtend(c *gin.Context) {
	var hkUserExtend community.HkUserExtend
	err := c.ShouldBindJSON(&hkUserExtend)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserExtendService.CreateHkUserExtend(hkUserExtend); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkUserExtend 删除HkUserExtend
// @Tags HkUserExtend
// @Summary 删除HkUserExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUserExtend true "删除HkUserExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserExtend/deleteHkUserExtend [delete]
func (hkUserExtendApi *HkUserExtendApi) DeleteHkUserExtend(c *gin.Context) {
	var hkUserExtend community.HkUserExtend
	err := c.ShouldBindJSON(&hkUserExtend)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserExtendService.DeleteHkUserExtend(hkUserExtend); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkUserExtendByIds 批量删除HkUserExtend
// @Tags HkUserExtend
// @Summary 批量删除HkUserExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkUserExtend/deleteHkUserExtendByIds [delete]
func (hkUserExtendApi *HkUserExtendApi) DeleteHkUserExtendByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserExtendService.DeleteHkUserExtendByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkUserExtend 更新HkUserExtend
// @Tags HkUserExtend
// @Summary 更新HkUserExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkUserExtend true "更新HkUserExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserExtend/updateHkUserExtend [put]
func (hkUserExtendApi *HkUserExtendApi) UpdateHkUserExtend(c *gin.Context) {
	var hkUserExtend community.HkUserExtend
	err := c.ShouldBindJSON(&hkUserExtend)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserExtendService.UpdateHkUserExtend(hkUserExtend); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkUserExtend 用id查询HkUserExtend
// @Tags HkUserExtend
// @Summary 用id查询HkUserExtend
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkUserExtend true "用id查询HkUserExtend"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserExtend/findHkUserExtend [get]
func (hkUserExtendApi *HkUserExtendApi) FindHkUserExtend(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkUserExtend, err := hkUserExtendService.GetHkUserExtend(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkUserExtend": rehkUserExtend}, c)
	}
}

// GetHkUserExtendList 分页获取HkUserExtend列表
// @Tags HkUserExtend
// @Summary 分页获取HkUserExtend列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkUserExtendSearch true "分页获取HkUserExtend列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserExtend/getHkUserExtendList [get]
func (hkUserExtendApi *HkUserExtendApi) GetHkUserExtendList(c *gin.Context) {
	var pageInfo communityReq.HkUserExtendSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkUserExtendService.GetHkUserExtendInfoList(pageInfo); err != nil {
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
