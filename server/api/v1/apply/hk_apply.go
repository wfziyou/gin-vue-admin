package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/apply"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/apply/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type HkApplyApi struct {
}

var hkApplyService = service.ServiceGroupApp.ApplyServiceGroup.HkApplyService


// CreateHkApply 创建HkApply
// @Tags HkApply
// @Summary 创建HkApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkApply true "创建HkApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkApply/createHkApply [post]
func (hkApplyApi *HkApplyApi) CreateHkApply(c *gin.Context) {
	var hkApply apply.HkApply
	err := c.ShouldBindJSON(&hkApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkApplyService.CreateHkApply(hkApply); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkApply 删除HkApply
// @Tags HkApply
// @Summary 删除HkApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkApply true "删除HkApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkApply/deleteHkApply [delete]
func (hkApplyApi *HkApplyApi) DeleteHkApply(c *gin.Context) {
	var hkApply apply.HkApply
	err := c.ShouldBindJSON(&hkApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkApplyService.DeleteHkApply(hkApply); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkApplyByIds 批量删除HkApply
// @Tags HkApply
// @Summary 批量删除HkApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkApply/deleteHkApplyByIds [delete]
func (hkApplyApi *HkApplyApi) DeleteHkApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkApplyService.DeleteHkApplyByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkApply 更新HkApply
// @Tags HkApply
// @Summary 更新HkApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkApply true "更新HkApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkApply/updateHkApply [put]
func (hkApplyApi *HkApplyApi) UpdateHkApply(c *gin.Context) {
	var hkApply apply.HkApply
	err := c.ShouldBindJSON(&hkApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkApplyService.UpdateHkApply(hkApply); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkApply 用id查询HkApply
// @Tags HkApply
// @Summary 用id查询HkApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query apply.HkApply true "用id查询HkApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkApply/findHkApply [get]
func (hkApplyApi *HkApplyApi) FindHkApply(c *gin.Context) {
	var hkApply apply.HkApply
	err := c.ShouldBindQuery(&hkApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkApply, err := hkApplyService.GetHkApply(hkApply.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkApply": rehkApply}, c)
	}
}

// GetHkApplyList 分页获取HkApply列表
// @Tags HkApply
// @Summary 分页获取HkApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.HkApplySearch true "分页获取HkApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkApply/getHkApplyList [get]
func (hkApplyApi *HkApplyApi) GetHkApplyList(c *gin.Context) {
	var pageInfo applyReq.HkApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkApplyService.GetHkApplyInfoList(pageInfo); err != nil {
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
