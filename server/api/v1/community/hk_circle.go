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

type HkCircleApi struct {
}

var hkCircleService = service.ServiceGroupApp.CommunityServiceGroup.HkCircleService

// CreateHkCircle 创建HkCircle
// @Tags HkCircle
// @Summary 创建HkCircle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircle true "创建HkCircle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircle/createHkCircle [post]
func (hkCircleApi *HkCircleApi) CreateHkCircle(c *gin.Context) {
	var hkCircle community.HkCircle
	err := c.ShouldBindJSON(&hkCircle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleService.CreateHkCircle(hkCircle); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkCircle 删除HkCircle
// @Tags HkCircle
// @Summary 删除HkCircle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircle true "删除HkCircle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircle/deleteHkCircle [delete]
func (hkCircleApi *HkCircleApi) DeleteHkCircle(c *gin.Context) {
	var hkCircle community.HkCircle
	err := c.ShouldBindJSON(&hkCircle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleService.DeleteHkCircle(hkCircle); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkCircleByIds 批量删除HkCircle
// @Tags HkCircle
// @Summary 批量删除HkCircle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkCircle/deleteHkCircleByIds [delete]
func (hkCircleApi *HkCircleApi) DeleteHkCircleByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleService.DeleteHkCircleByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkCircle 更新HkCircle
// @Tags HkCircle
// @Summary 更新HkCircle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircle true "更新HkCircle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircle/updateHkCircle [put]
func (hkCircleApi *HkCircleApi) UpdateHkCircle(c *gin.Context) {
	var hkCircle community.HkCircle
	err := c.ShouldBindJSON(&hkCircle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleService.UpdateHkCircle(hkCircle); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkCircle 用id查询HkCircle
// @Tags HkCircle
// @Summary 用id查询HkCircle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkCircle true "用id查询HkCircle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircle/findHkCircle [get]
func (hkCircleApi *HkCircleApi) FindHkCircle(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircle, err := hkCircleService.GetHkCircle(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircle": rehkCircle}, c)
	}
}

// GetHkCircleList 分页获取HkCircle列表
// @Tags HkCircle
// @Summary 分页获取HkCircle列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkCircleSearch true "分页获取HkCircle列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircle/getHkCircleList [get]
func (hkCircleApi *HkCircleApi) GetHkCircleList(c *gin.Context) {
	var pageInfo communityReq.HkCircleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkCircleService.GetHkCircleInfoList(pageInfo); err != nil {
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
