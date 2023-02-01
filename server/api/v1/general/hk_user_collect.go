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

type HkUserCollectApi struct {
}

var hkUserCollectService = service.ServiceGroupApp.GeneralServiceGroup.HkUserCollectService


// CreateHkUserCollect 创建HkUserCollect
// @Tags HkUserCollect
// @Summary 创建HkUserCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkUserCollect true "创建HkUserCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserCollect/createHkUserCollect [post]
func (hkUserCollectApi *HkUserCollectApi) CreateHkUserCollect(c *gin.Context) {
	var hkUserCollect general.HkUserCollect
	err := c.ShouldBindJSON(&hkUserCollect)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserCollectService.CreateHkUserCollect(hkUserCollect); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkUserCollect 删除HkUserCollect
// @Tags HkUserCollect
// @Summary 删除HkUserCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkUserCollect true "删除HkUserCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserCollect/deleteHkUserCollect [delete]
func (hkUserCollectApi *HkUserCollectApi) DeleteHkUserCollect(c *gin.Context) {
	var hkUserCollect general.HkUserCollect
	err := c.ShouldBindJSON(&hkUserCollect)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserCollectService.DeleteHkUserCollect(hkUserCollect); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkUserCollectByIds 批量删除HkUserCollect
// @Tags HkUserCollect
// @Summary 批量删除HkUserCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkUserCollect/deleteHkUserCollectByIds [delete]
func (hkUserCollectApi *HkUserCollectApi) DeleteHkUserCollectByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserCollectService.DeleteHkUserCollectByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkUserCollect 更新HkUserCollect
// @Tags HkUserCollect
// @Summary 更新HkUserCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body general.HkUserCollect true "更新HkUserCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserCollect/updateHkUserCollect [put]
func (hkUserCollectApi *HkUserCollectApi) UpdateHkUserCollect(c *gin.Context) {
	var hkUserCollect general.HkUserCollect
	err := c.ShouldBindJSON(&hkUserCollect)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserCollectService.UpdateHkUserCollect(hkUserCollect); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkUserCollect 用id查询HkUserCollect
// @Tags HkUserCollect
// @Summary 用id查询HkUserCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query general.HkUserCollect true "用id查询HkUserCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserCollect/findHkUserCollect [get]
func (hkUserCollectApi *HkUserCollectApi) FindHkUserCollect(c *gin.Context) {
	var hkUserCollect general.HkUserCollect
	err := c.ShouldBindQuery(&hkUserCollect)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkUserCollect, err := hkUserCollectService.GetHkUserCollect(hkUserCollect.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkUserCollect": rehkUserCollect}, c)
	}
}

// GetHkUserCollectList 分页获取HkUserCollect列表
// @Tags HkUserCollect
// @Summary 分页获取HkUserCollect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query generalReq.HkUserCollectSearch true "分页获取HkUserCollect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserCollect/getHkUserCollectList [get]
func (hkUserCollectApi *HkUserCollectApi) GetHkUserCollectList(c *gin.Context) {
	var pageInfo generalReq.HkUserCollectSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkUserCollectService.GetHkUserCollectInfoList(pageInfo); err != nil {
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
