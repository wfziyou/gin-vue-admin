package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HkPlatApplyApi struct {
}

var hkPlatApplyService = service.ServiceGroupApp.ApplyServiceGroup.HkPlatApplyService

// CreateHkPlatApply 创建HkPlatApply
// @Tags HkPlatApply
// @Summary 创建HkPlatApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkPlatApply true "创建HkPlatApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkPlatApply/createHkPlatApply [post]
func (hkPlatApplyApi *HkPlatApplyApi) CreateHkPlatApply(c *gin.Context) {
	var hkPlatApply apply.HkPlatApply
	err := c.ShouldBindJSON(&hkPlatApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkPlatApplyService.CreateHkPlatApply(hkPlatApply); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkPlatApply 删除HkPlatApply
// @Tags HkPlatApply
// @Summary 删除HkPlatApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkPlatApply true "删除HkPlatApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkPlatApply/deleteHkPlatApply [delete]
func (hkPlatApplyApi *HkPlatApplyApi) DeleteHkPlatApply(c *gin.Context) {
	var hkPlatApply apply.HkPlatApply
	err := c.ShouldBindJSON(&hkPlatApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkPlatApplyService.DeleteHkPlatApply(hkPlatApply); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkPlatApplyByIds 批量删除HkPlatApply
// @Tags HkPlatApply
// @Summary 批量删除HkPlatApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkPlatApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkPlatApply/deleteHkPlatApplyByIds [delete]
func (hkPlatApplyApi *HkPlatApplyApi) DeleteHkPlatApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkPlatApplyService.DeleteHkPlatApplyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkPlatApply 更新HkPlatApply
// @Tags HkPlatApply
// @Summary 更新HkPlatApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkPlatApply true "更新HkPlatApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkPlatApply/updateHkPlatApply [put]
func (hkPlatApplyApi *HkPlatApplyApi) UpdateHkPlatApply(c *gin.Context) {
	var hkPlatApply apply.HkPlatApply
	err := c.ShouldBindJSON(&hkPlatApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkPlatApplyService.UpdateHkPlatApply(hkPlatApply); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkPlatApply 用id查询HkPlatApply
// @Tags HkPlatApply
// @Summary 用id查询HkPlatApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkPlatApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkPlatApply/findHkPlatApply [get]
func (hkPlatApplyApi *HkPlatApplyApi) FindHkPlatApply(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkPlatApply, err := hkPlatApplyService.GetHkPlatApply(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkPlatApply": rehkPlatApply}, c)
	}
}

// GetHkPlatApplyList 分页获取HkPlatApply列表
// @Tags HkPlatApply
// @Summary 分页获取HkPlatApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.HkPlatApplySearch true "分页获取HkPlatApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkPlatApply/getHkPlatApplyList [get]
func (hkPlatApplyApi *HkPlatApplyApi) GetHkPlatApplyList(c *gin.Context) {
	var pageInfo applyReq.HkPlatApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkPlatApplyService.GetHkPlatApplyInfoList(pageInfo); err != nil {
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
