package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/community"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type HkForbiddenSpeakReasonApi struct {
}

var hkForbiddenSpeakReasonService = service.ServiceGroupApp.CommunityServiceGroup.HkForbiddenSpeakReasonService


// CreateHkForbiddenSpeakReason 创建HkForbiddenSpeakReason
// @Tags HkForbiddenSpeakReason
// @Summary 创建HkForbiddenSpeakReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForbiddenSpeakReason true "创建HkForbiddenSpeakReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForbiddenSpeakReason/createHkForbiddenSpeakReason [post]
func (hkForbiddenSpeakReasonApi *HkForbiddenSpeakReasonApi) CreateHkForbiddenSpeakReason(c *gin.Context) {
	var hkForbiddenSpeakReason community.HkForbiddenSpeakReason
	err := c.ShouldBindJSON(&hkForbiddenSpeakReason)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForbiddenSpeakReasonService.CreateHkForbiddenSpeakReason(hkForbiddenSpeakReason); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkForbiddenSpeakReason 删除HkForbiddenSpeakReason
// @Tags HkForbiddenSpeakReason
// @Summary 删除HkForbiddenSpeakReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForbiddenSpeakReason true "删除HkForbiddenSpeakReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForbiddenSpeakReason/deleteHkForbiddenSpeakReason [delete]
func (hkForbiddenSpeakReasonApi *HkForbiddenSpeakReasonApi) DeleteHkForbiddenSpeakReason(c *gin.Context) {
	var hkForbiddenSpeakReason community.HkForbiddenSpeakReason
	err := c.ShouldBindJSON(&hkForbiddenSpeakReason)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForbiddenSpeakReasonService.DeleteHkForbiddenSpeakReason(hkForbiddenSpeakReason); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkForbiddenSpeakReasonByIds 批量删除HkForbiddenSpeakReason
// @Tags HkForbiddenSpeakReason
// @Summary 批量删除HkForbiddenSpeakReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForbiddenSpeakReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkForbiddenSpeakReason/deleteHkForbiddenSpeakReasonByIds [delete]
func (hkForbiddenSpeakReasonApi *HkForbiddenSpeakReasonApi) DeleteHkForbiddenSpeakReasonByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForbiddenSpeakReasonService.DeleteHkForbiddenSpeakReasonByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkForbiddenSpeakReason 更新HkForbiddenSpeakReason
// @Tags HkForbiddenSpeakReason
// @Summary 更新HkForbiddenSpeakReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForbiddenSpeakReason true "更新HkForbiddenSpeakReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForbiddenSpeakReason/updateHkForbiddenSpeakReason [put]
func (hkForbiddenSpeakReasonApi *HkForbiddenSpeakReasonApi) UpdateHkForbiddenSpeakReason(c *gin.Context) {
	var hkForbiddenSpeakReason community.HkForbiddenSpeakReason
	err := c.ShouldBindJSON(&hkForbiddenSpeakReason)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForbiddenSpeakReasonService.UpdateHkForbiddenSpeakReason(hkForbiddenSpeakReason); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkForbiddenSpeakReason 用id查询HkForbiddenSpeakReason
// @Tags HkForbiddenSpeakReason
// @Summary 用id查询HkForbiddenSpeakReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkForbiddenSpeakReason true "用id查询HkForbiddenSpeakReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForbiddenSpeakReason/findHkForbiddenSpeakReason [get]
func (hkForbiddenSpeakReasonApi *HkForbiddenSpeakReasonApi) FindHkForbiddenSpeakReason(c *gin.Context) {
	var hkForbiddenSpeakReason community.HkForbiddenSpeakReason
	err := c.ShouldBindQuery(&hkForbiddenSpeakReason)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForbiddenSpeakReason, err := hkForbiddenSpeakReasonService.GetHkForbiddenSpeakReason(hkForbiddenSpeakReason.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForbiddenSpeakReason": rehkForbiddenSpeakReason}, c)
	}
}

// GetHkForbiddenSpeakReasonList 分页获取HkForbiddenSpeakReason列表
// @Tags HkForbiddenSpeakReason
// @Summary 分页获取HkForbiddenSpeakReason列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkForbiddenSpeakReasonSearch true "分页获取HkForbiddenSpeakReason列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForbiddenSpeakReason/getHkForbiddenSpeakReasonList [get]
func (hkForbiddenSpeakReasonApi *HkForbiddenSpeakReasonApi) GetHkForbiddenSpeakReasonList(c *gin.Context) {
	var pageInfo communityReq.HkForbiddenSpeakReasonSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForbiddenSpeakReasonService.GetHkForbiddenSpeakReasonInfoList(pageInfo); err != nil {
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
