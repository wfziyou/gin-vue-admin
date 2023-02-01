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

type HkForbiddenSpeakDurationApi struct {
}

var hkForbiddenSpeakDurationService = service.ServiceGroupApp.CommunityServiceGroup.HkForbiddenSpeakDurationService


// CreateHkForbiddenSpeakDuration 创建HkForbiddenSpeakDuration
// @Tags HkForbiddenSpeakDuration
// @Summary 创建HkForbiddenSpeakDuration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForbiddenSpeakDuration true "创建HkForbiddenSpeakDuration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForbiddenSpeakDuration/createHkForbiddenSpeakDuration [post]
func (hkForbiddenSpeakDurationApi *HkForbiddenSpeakDurationApi) CreateHkForbiddenSpeakDuration(c *gin.Context) {
	var hkForbiddenSpeakDuration community.HkForbiddenSpeakDuration
	err := c.ShouldBindJSON(&hkForbiddenSpeakDuration)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForbiddenSpeakDurationService.CreateHkForbiddenSpeakDuration(hkForbiddenSpeakDuration); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkForbiddenSpeakDuration 删除HkForbiddenSpeakDuration
// @Tags HkForbiddenSpeakDuration
// @Summary 删除HkForbiddenSpeakDuration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForbiddenSpeakDuration true "删除HkForbiddenSpeakDuration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForbiddenSpeakDuration/deleteHkForbiddenSpeakDuration [delete]
func (hkForbiddenSpeakDurationApi *HkForbiddenSpeakDurationApi) DeleteHkForbiddenSpeakDuration(c *gin.Context) {
	var hkForbiddenSpeakDuration community.HkForbiddenSpeakDuration
	err := c.ShouldBindJSON(&hkForbiddenSpeakDuration)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForbiddenSpeakDurationService.DeleteHkForbiddenSpeakDuration(hkForbiddenSpeakDuration); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkForbiddenSpeakDurationByIds 批量删除HkForbiddenSpeakDuration
// @Tags HkForbiddenSpeakDuration
// @Summary 批量删除HkForbiddenSpeakDuration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForbiddenSpeakDuration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkForbiddenSpeakDuration/deleteHkForbiddenSpeakDurationByIds [delete]
func (hkForbiddenSpeakDurationApi *HkForbiddenSpeakDurationApi) DeleteHkForbiddenSpeakDurationByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForbiddenSpeakDurationService.DeleteHkForbiddenSpeakDurationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkForbiddenSpeakDuration 更新HkForbiddenSpeakDuration
// @Tags HkForbiddenSpeakDuration
// @Summary 更新HkForbiddenSpeakDuration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForbiddenSpeakDuration true "更新HkForbiddenSpeakDuration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForbiddenSpeakDuration/updateHkForbiddenSpeakDuration [put]
func (hkForbiddenSpeakDurationApi *HkForbiddenSpeakDurationApi) UpdateHkForbiddenSpeakDuration(c *gin.Context) {
	var hkForbiddenSpeakDuration community.HkForbiddenSpeakDuration
	err := c.ShouldBindJSON(&hkForbiddenSpeakDuration)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForbiddenSpeakDurationService.UpdateHkForbiddenSpeakDuration(hkForbiddenSpeakDuration); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkForbiddenSpeakDuration 用id查询HkForbiddenSpeakDuration
// @Tags HkForbiddenSpeakDuration
// @Summary 用id查询HkForbiddenSpeakDuration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkForbiddenSpeakDuration true "用id查询HkForbiddenSpeakDuration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForbiddenSpeakDuration/findHkForbiddenSpeakDuration [get]
func (hkForbiddenSpeakDurationApi *HkForbiddenSpeakDurationApi) FindHkForbiddenSpeakDuration(c *gin.Context) {
	var hkForbiddenSpeakDuration community.HkForbiddenSpeakDuration
	err := c.ShouldBindQuery(&hkForbiddenSpeakDuration)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForbiddenSpeakDuration, err := hkForbiddenSpeakDurationService.GetHkForbiddenSpeakDuration(hkForbiddenSpeakDuration.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForbiddenSpeakDuration": rehkForbiddenSpeakDuration}, c)
	}
}

// GetHkForbiddenSpeakDurationList 分页获取HkForbiddenSpeakDuration列表
// @Tags HkForbiddenSpeakDuration
// @Summary 分页获取HkForbiddenSpeakDuration列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkForbiddenSpeakDurationSearch true "分页获取HkForbiddenSpeakDuration列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForbiddenSpeakDuration/getHkForbiddenSpeakDurationList [get]
func (hkForbiddenSpeakDurationApi *HkForbiddenSpeakDurationApi) GetHkForbiddenSpeakDurationList(c *gin.Context) {
	var pageInfo communityReq.HkForbiddenSpeakDurationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForbiddenSpeakDurationService.GetHkForbiddenSpeakDurationInfoList(pageInfo); err != nil {
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
