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

type HkCircleClassifyApi struct {
}

var hkCircleClassifyService = service.ServiceGroupApp.CommunityServiceGroup.HkCircleClassifyService

// CreateHkCircleClassify 创建HkCircleClassify
// @Tags HkCircleClassify
// @Summary 创建HkCircleClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleClassify true "创建HkCircleClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleClassify/createHkCircleClassify [post]
func (hkCircleClassifyApi *HkCircleClassifyApi) CreateHkCircleClassify(c *gin.Context) {
	var hkCircleClassify community.HkCircleClassify
	err := c.ShouldBindJSON(&hkCircleClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleClassifyService.CreateHkCircleClassify(hkCircleClassify); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkCircleClassify 删除HkCircleClassify
// @Tags HkCircleClassify
// @Summary 删除HkCircleClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleClassify true "删除HkCircleClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleClassify/deleteHkCircleClassify [delete]
func (hkCircleClassifyApi *HkCircleClassifyApi) DeleteHkCircleClassify(c *gin.Context) {
	var hkCircleClassify community.HkCircleClassify
	err := c.ShouldBindJSON(&hkCircleClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleClassifyService.DeleteHkCircleClassify(hkCircleClassify); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkCircleClassifyByIds 批量删除HkCircleClassify
// @Tags HkCircleClassify
// @Summary 批量删除HkCircleClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkCircleClassify/deleteHkCircleClassifyByIds [delete]
func (hkCircleClassifyApi *HkCircleClassifyApi) DeleteHkCircleClassifyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleClassifyService.DeleteHkCircleClassifyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkCircleClassify 更新HkCircleClassify
// @Tags HkCircleClassify
// @Summary 更新HkCircleClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleClassify true "更新HkCircleClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleClassify/updateHkCircleClassify [put]
func (hkCircleClassifyApi *HkCircleClassifyApi) UpdateHkCircleClassify(c *gin.Context) {
	var hkCircleClassify community.HkCircleClassify
	err := c.ShouldBindJSON(&hkCircleClassify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleClassifyService.UpdateHkCircleClassify(hkCircleClassify); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkCircleClassify 用id查询HkCircleClassify
// @Tags HkCircleClassify
// @Summary 用id查询HkCircleClassify
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkCircleClassify true "用id查询HkCircleClassify"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleClassify/findHkCircleClassify [get]
func (hkCircleClassifyApi *HkCircleClassifyApi) FindHkCircleClassify(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleClassify, err := hkCircleClassifyService.GetHkCircleClassify(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircleClassify": rehkCircleClassify}, c)
	}
}

// GetHkCircleClassifyList 分页获取HkCircleClassify列表
// @Tags HkCircleClassify
// @Summary 分页获取HkCircleClassify列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkCircleClassifySearch true "分页获取HkCircleClassify列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleClassify/getHkCircleClassifyList [get]
func (hkCircleClassifyApi *HkCircleClassifyApi) GetHkCircleClassifyList(c *gin.Context) {
	var pageInfo communityReq.HkCircleClassifySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkCircleClassifyService.GetHkCircleClassifyInfoList(pageInfo); err != nil {
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
