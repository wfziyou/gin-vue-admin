package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/general"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/general/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type HkProtocolApi struct {
}

var hkProtocolService = service.ServiceGroupApp.GeneralServiceGroup.HkProtocolService


// CreateHkProtocol 创建HkProtocol
// @Tags HkProtocol
// @Summary 创建HkProtocol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkProtocol true "创建HkProtocol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkProtocol/createHkProtocol [post]
func (hkProtocolApi *HkProtocolApi) CreateHkProtocol(c *gin.Context) {
	var hkProtocol general.HkProtocol
	err := c.ShouldBindJSON(&hkProtocol)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkProtocolService.CreateHkProtocol(hkProtocol); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkProtocol 删除HkProtocol
// @Tags HkProtocol
// @Summary 删除HkProtocol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkProtocol true "删除HkProtocol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkProtocol/deleteHkProtocol [delete]
func (hkProtocolApi *HkProtocolApi) DeleteHkProtocol(c *gin.Context) {
	var hkProtocol general.HkProtocol
	err := c.ShouldBindJSON(&hkProtocol)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkProtocolService.DeleteHkProtocol(hkProtocol); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkProtocolByIds 批量删除HkProtocol
// @Tags HkProtocol
// @Summary 批量删除HkProtocol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkProtocol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkProtocol/deleteHkProtocolByIds [delete]
func (hkProtocolApi *HkProtocolApi) DeleteHkProtocolByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkProtocolService.DeleteHkProtocolByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkProtocol 更新HkProtocol
// @Tags HkProtocol
// @Summary 更新HkProtocol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkProtocol true "更新HkProtocol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkProtocol/updateHkProtocol [put]
func (hkProtocolApi *HkProtocolApi) UpdateHkProtocol(c *gin.Context) {
	var hkProtocol general.HkProtocol
	err := c.ShouldBindJSON(&hkProtocol)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkProtocolService.UpdateHkProtocol(hkProtocol); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkProtocol 用id查询HkProtocol
// @Tags HkProtocol
// @Summary 用id查询HkProtocol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query general.HkProtocol true "用id查询HkProtocol"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkProtocol/findHkProtocol [get]
func (hkProtocolApi *HkProtocolApi) FindHkProtocol(c *gin.Context) {
	var hkProtocol general.HkProtocol
	err := c.ShouldBindQuery(&hkProtocol)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkProtocol, err := hkProtocolService.GetHkProtocol(hkProtocol.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkProtocol": rehkProtocol}, c)
	}
}

// GetHkProtocolList 分页获取HkProtocol列表
// @Tags HkProtocol
// @Summary 分页获取HkProtocol列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.HkProtocolSearch true "分页获取HkProtocol列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkProtocol/getHkProtocolList [get]
func (hkProtocolApi *HkProtocolApi) GetHkProtocolList(c *gin.Context) {
	var pageInfo generalReq.HkProtocolSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkProtocolService.GetHkProtocolInfoList(pageInfo); err != nil {
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
