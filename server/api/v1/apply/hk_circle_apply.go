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

type HkCircleApplyApi struct {
}

var hkCircleApplyService = service.ServiceGroupApp.ApplyServiceGroup.HkCircleApplyService


// CreateHkCircleApply 创建HkCircleApply
// @Tags HkCircleApply
// @Summary 创建HkCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkCircleApply true "创建HkCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleApply/createHkCircleApply [post]
func (hkCircleApplyApi *HkCircleApplyApi) CreateHkCircleApply(c *gin.Context) {
	var hkCircleApply apply.HkCircleApply
	err := c.ShouldBindJSON(&hkCircleApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleApplyService.CreateHkCircleApply(hkCircleApply); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkCircleApply 删除HkCircleApply
// @Tags HkCircleApply
// @Summary 删除HkCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkCircleApply true "删除HkCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleApply/deleteHkCircleApply [delete]
func (hkCircleApplyApi *HkCircleApplyApi) DeleteHkCircleApply(c *gin.Context) {
	var hkCircleApply apply.HkCircleApply
	err := c.ShouldBindJSON(&hkCircleApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleApplyService.DeleteHkCircleApply(hkCircleApply); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkCircleApplyByIds 批量删除HkCircleApply
// @Tags HkCircleApply
// @Summary 批量删除HkCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkCircleApply/deleteHkCircleApplyByIds [delete]
func (hkCircleApplyApi *HkCircleApplyApi) DeleteHkCircleApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleApplyService.DeleteHkCircleApplyByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkCircleApply 更新HkCircleApply
// @Tags HkCircleApply
// @Summary 更新HkCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apply.HkCircleApply true "更新HkCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleApply/updateHkCircleApply [put]
func (hkCircleApplyApi *HkCircleApplyApi) UpdateHkCircleApply(c *gin.Context) {
	var hkCircleApply apply.HkCircleApply
	err := c.ShouldBindJSON(&hkCircleApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleApplyService.UpdateHkCircleApply(hkCircleApply); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkCircleApply 用id查询HkCircleApply
// @Tags HkCircleApply
// @Summary 用id查询HkCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query apply.HkCircleApply true "用id查询HkCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleApply/findHkCircleApply [get]
func (hkCircleApplyApi *HkCircleApplyApi) FindHkCircleApply(c *gin.Context) {
	var hkCircleApply apply.HkCircleApply
	err := c.ShouldBindQuery(&hkCircleApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleApply, err := hkCircleApplyService.GetHkCircleApply(hkCircleApply.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircleApply": rehkCircleApply}, c)
	}
}

// GetHkCircleApplyList 分页获取HkCircleApply列表
// @Tags HkCircleApply
// @Summary 分页获取HkCircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.HkCircleApplySearch true "分页获取HkCircleApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleApply/getHkCircleApplyList [get]
func (hkCircleApplyApi *HkCircleApplyApi) GetHkCircleApplyList(c *gin.Context) {
	var pageInfo applyReq.HkCircleApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkCircleApplyService.GetHkCircleApplyInfoList(pageInfo); err != nil {
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
