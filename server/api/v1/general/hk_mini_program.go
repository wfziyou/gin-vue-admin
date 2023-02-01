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

type HkMiniProgramApi struct {
}

var hkMiniProgramService = service.ServiceGroupApp.GeneralServiceGroup.HkMiniProgramService


// CreateHkMiniProgram 创建HkMiniProgram
// @Tags HkMiniProgram
// @Summary 创建HkMiniProgram
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkMiniProgram true "创建HkMiniProgram"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkMiniProgram/createHkMiniProgram [post]
func (hkMiniProgramApi *HkMiniProgramApi) CreateHkMiniProgram(c *gin.Context) {
	var hkMiniProgram general.HkMiniProgram
	err := c.ShouldBindJSON(&hkMiniProgram)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkMiniProgramService.CreateHkMiniProgram(hkMiniProgram); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkMiniProgram 删除HkMiniProgram
// @Tags HkMiniProgram
// @Summary 删除HkMiniProgram
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkMiniProgram true "删除HkMiniProgram"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkMiniProgram/deleteHkMiniProgram [delete]
func (hkMiniProgramApi *HkMiniProgramApi) DeleteHkMiniProgram(c *gin.Context) {
	var hkMiniProgram general.HkMiniProgram
	err := c.ShouldBindJSON(&hkMiniProgram)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkMiniProgramService.DeleteHkMiniProgram(hkMiniProgram); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkMiniProgramByIds 批量删除HkMiniProgram
// @Tags HkMiniProgram
// @Summary 批量删除HkMiniProgram
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkMiniProgram"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkMiniProgram/deleteHkMiniProgramByIds [delete]
func (hkMiniProgramApi *HkMiniProgramApi) DeleteHkMiniProgramByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkMiniProgramService.DeleteHkMiniProgramByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkMiniProgram 更新HkMiniProgram
// @Tags HkMiniProgram
// @Summary 更新HkMiniProgram
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkMiniProgram true "更新HkMiniProgram"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkMiniProgram/updateHkMiniProgram [put]
func (hkMiniProgramApi *HkMiniProgramApi) UpdateHkMiniProgram(c *gin.Context) {
	var hkMiniProgram general.HkMiniProgram
	err := c.ShouldBindJSON(&hkMiniProgram)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkMiniProgramService.UpdateHkMiniProgram(hkMiniProgram); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkMiniProgram 用id查询HkMiniProgram
// @Tags HkMiniProgram
// @Summary 用id查询HkMiniProgram
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query general.HkMiniProgram true "用id查询HkMiniProgram"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkMiniProgram/findHkMiniProgram [get]
func (hkMiniProgramApi *HkMiniProgramApi) FindHkMiniProgram(c *gin.Context) {
	var hkMiniProgram general.HkMiniProgram
	err := c.ShouldBindQuery(&hkMiniProgram)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkMiniProgram, err := hkMiniProgramService.GetHkMiniProgram(hkMiniProgram.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkMiniProgram": rehkMiniProgram}, c)
	}
}

// GetHkMiniProgramList 分页获取HkMiniProgram列表
// @Tags HkMiniProgram
// @Summary 分页获取HkMiniProgram列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.HkMiniProgramSearch true "分页获取HkMiniProgram列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkMiniProgram/getHkMiniProgramList [get]
func (hkMiniProgramApi *HkMiniProgramApi) GetHkMiniProgramList(c *gin.Context) {
	var pageInfo generalReq.HkMiniProgramSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkMiniProgramService.GetHkMiniProgramInfoList(pageInfo); err != nil {
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
