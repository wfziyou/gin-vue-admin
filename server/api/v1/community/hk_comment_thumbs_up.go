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

type HkCommentThumbsUpApi struct {
}

var hkCommentThumbsUpService = service.ServiceGroupApp.CommunityServiceGroup.HkCommentThumbsUpService


// CreateHkCommentThumbsUp 创建HkCommentThumbsUp
// @Tags HkCommentThumbsUp
// @Summary 创建HkCommentThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCommentThumbsUp true "创建HkCommentThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCommentThumbsUp/createHkCommentThumbsUp [post]
func (hkCommentThumbsUpApi *HkCommentThumbsUpApi) CreateHkCommentThumbsUp(c *gin.Context) {
	var hkCommentThumbsUp community.HkCommentThumbsUp
	err := c.ShouldBindJSON(&hkCommentThumbsUp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCommentThumbsUpService.CreateHkCommentThumbsUp(hkCommentThumbsUp); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkCommentThumbsUp 删除HkCommentThumbsUp
// @Tags HkCommentThumbsUp
// @Summary 删除HkCommentThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCommentThumbsUp true "删除HkCommentThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCommentThumbsUp/deleteHkCommentThumbsUp [delete]
func (hkCommentThumbsUpApi *HkCommentThumbsUpApi) DeleteHkCommentThumbsUp(c *gin.Context) {
	var hkCommentThumbsUp community.HkCommentThumbsUp
	err := c.ShouldBindJSON(&hkCommentThumbsUp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCommentThumbsUpService.DeleteHkCommentThumbsUp(hkCommentThumbsUp); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkCommentThumbsUpByIds 批量删除HkCommentThumbsUp
// @Tags HkCommentThumbsUp
// @Summary 批量删除HkCommentThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCommentThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkCommentThumbsUp/deleteHkCommentThumbsUpByIds [delete]
func (hkCommentThumbsUpApi *HkCommentThumbsUpApi) DeleteHkCommentThumbsUpByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCommentThumbsUpService.DeleteHkCommentThumbsUpByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkCommentThumbsUp 更新HkCommentThumbsUp
// @Tags HkCommentThumbsUp
// @Summary 更新HkCommentThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCommentThumbsUp true "更新HkCommentThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCommentThumbsUp/updateHkCommentThumbsUp [put]
func (hkCommentThumbsUpApi *HkCommentThumbsUpApi) UpdateHkCommentThumbsUp(c *gin.Context) {
	var hkCommentThumbsUp community.HkCommentThumbsUp
	err := c.ShouldBindJSON(&hkCommentThumbsUp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCommentThumbsUpService.UpdateHkCommentThumbsUp(hkCommentThumbsUp); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkCommentThumbsUp 用id查询HkCommentThumbsUp
// @Tags HkCommentThumbsUp
// @Summary 用id查询HkCommentThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkCommentThumbsUp true "用id查询HkCommentThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCommentThumbsUp/findHkCommentThumbsUp [get]
func (hkCommentThumbsUpApi *HkCommentThumbsUpApi) FindHkCommentThumbsUp(c *gin.Context) {
	var hkCommentThumbsUp community.HkCommentThumbsUp
	err := c.ShouldBindQuery(&hkCommentThumbsUp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCommentThumbsUp, err := hkCommentThumbsUpService.GetHkCommentThumbsUp(hkCommentThumbsUp.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCommentThumbsUp": rehkCommentThumbsUp}, c)
	}
}

// GetHkCommentThumbsUpList 分页获取HkCommentThumbsUp列表
// @Tags HkCommentThumbsUp
// @Summary 分页获取HkCommentThumbsUp列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkCommentThumbsUpSearch true "分页获取HkCommentThumbsUp列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCommentThumbsUp/getHkCommentThumbsUpList [get]
func (hkCommentThumbsUpApi *HkCommentThumbsUpApi) GetHkCommentThumbsUpList(c *gin.Context) {
	var pageInfo communityReq.HkCommentThumbsUpSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkCommentThumbsUpService.GetHkCommentThumbsUpInfoList(pageInfo); err != nil {
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
