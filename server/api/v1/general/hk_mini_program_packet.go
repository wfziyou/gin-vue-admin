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

type HkMiniProgramPacketApi struct {
}

var hkMiniProgramPacketService = service.ServiceGroupApp.GeneralServiceGroup.HkMiniProgramPacketService


// CreateHkMiniProgramPacket 创建HkMiniProgramPacket
// @Tags HkMiniProgramPacket
// @Summary 创建HkMiniProgramPacket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkMiniProgramPacket true "创建HkMiniProgramPacket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkMiniProgramPacket/createHkMiniProgramPacket [post]
func (hkMiniProgramPacketApi *HkMiniProgramPacketApi) CreateHkMiniProgramPacket(c *gin.Context) {
	var hkMiniProgramPacket general.HkMiniProgramPacket
	err := c.ShouldBindJSON(&hkMiniProgramPacket)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkMiniProgramPacketService.CreateHkMiniProgramPacket(hkMiniProgramPacket); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkMiniProgramPacket 删除HkMiniProgramPacket
// @Tags HkMiniProgramPacket
// @Summary 删除HkMiniProgramPacket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkMiniProgramPacket true "删除HkMiniProgramPacket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkMiniProgramPacket/deleteHkMiniProgramPacket [delete]
func (hkMiniProgramPacketApi *HkMiniProgramPacketApi) DeleteHkMiniProgramPacket(c *gin.Context) {
	var hkMiniProgramPacket general.HkMiniProgramPacket
	err := c.ShouldBindJSON(&hkMiniProgramPacket)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkMiniProgramPacketService.DeleteHkMiniProgramPacket(hkMiniProgramPacket); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkMiniProgramPacketByIds 批量删除HkMiniProgramPacket
// @Tags HkMiniProgramPacket
// @Summary 批量删除HkMiniProgramPacket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkMiniProgramPacket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkMiniProgramPacket/deleteHkMiniProgramPacketByIds [delete]
func (hkMiniProgramPacketApi *HkMiniProgramPacketApi) DeleteHkMiniProgramPacketByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkMiniProgramPacketService.DeleteHkMiniProgramPacketByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkMiniProgramPacket 更新HkMiniProgramPacket
// @Tags HkMiniProgramPacket
// @Summary 更新HkMiniProgramPacket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkMiniProgramPacket true "更新HkMiniProgramPacket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkMiniProgramPacket/updateHkMiniProgramPacket [put]
func (hkMiniProgramPacketApi *HkMiniProgramPacketApi) UpdateHkMiniProgramPacket(c *gin.Context) {
	var hkMiniProgramPacket general.HkMiniProgramPacket
	err := c.ShouldBindJSON(&hkMiniProgramPacket)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkMiniProgramPacketService.UpdateHkMiniProgramPacket(hkMiniProgramPacket); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkMiniProgramPacket 用id查询HkMiniProgramPacket
// @Tags HkMiniProgramPacket
// @Summary 用id查询HkMiniProgramPacket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query general.HkMiniProgramPacket true "用id查询HkMiniProgramPacket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkMiniProgramPacket/findHkMiniProgramPacket [get]
func (hkMiniProgramPacketApi *HkMiniProgramPacketApi) FindHkMiniProgramPacket(c *gin.Context) {
	var hkMiniProgramPacket general.HkMiniProgramPacket
	err := c.ShouldBindQuery(&hkMiniProgramPacket)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkMiniProgramPacket, err := hkMiniProgramPacketService.GetHkMiniProgramPacket(hkMiniProgramPacket.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkMiniProgramPacket": rehkMiniProgramPacket}, c)
	}
}

// GetHkMiniProgramPacketList 分页获取HkMiniProgramPacket列表
// @Tags HkMiniProgramPacket
// @Summary 分页获取HkMiniProgramPacket列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.HkMiniProgramPacketSearch true "分页获取HkMiniProgramPacket列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkMiniProgramPacket/getHkMiniProgramPacketList [get]
func (hkMiniProgramPacketApi *HkMiniProgramPacketApi) GetHkMiniProgramPacketList(c *gin.Context) {
	var pageInfo generalReq.HkMiniProgramPacketSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkMiniProgramPacketService.GetHkMiniProgramPacketInfoList(pageInfo); err != nil {
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
