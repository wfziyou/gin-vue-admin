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

type HkCircleRelationApi struct {
}

var hkCircleRelationService = service.ServiceGroupApp.CommunityServiceGroup.HkCircleRelationService


// CreateHkCircleRelation 创建HkCircleRelation
// @Tags HkCircleRelation
// @Summary 创建HkCircleRelation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleRelation true "创建HkCircleRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleRelation/createHkCircleRelation [post]
func (hkCircleRelationApi *HkCircleRelationApi) CreateHkCircleRelation(c *gin.Context) {
	var hkCircleRelation community.HkCircleRelation
	err := c.ShouldBindJSON(&hkCircleRelation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleRelationService.CreateHkCircleRelation(hkCircleRelation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkCircleRelation 删除HkCircleRelation
// @Tags HkCircleRelation
// @Summary 删除HkCircleRelation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleRelation true "删除HkCircleRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleRelation/deleteHkCircleRelation [delete]
func (hkCircleRelationApi *HkCircleRelationApi) DeleteHkCircleRelation(c *gin.Context) {
	var hkCircleRelation community.HkCircleRelation
	err := c.ShouldBindJSON(&hkCircleRelation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleRelationService.DeleteHkCircleRelation(hkCircleRelation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkCircleRelationByIds 批量删除HkCircleRelation
// @Tags HkCircleRelation
// @Summary 批量删除HkCircleRelation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkCircleRelation/deleteHkCircleRelationByIds [delete]
func (hkCircleRelationApi *HkCircleRelationApi) DeleteHkCircleRelationByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleRelationService.DeleteHkCircleRelationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkCircleRelation 更新HkCircleRelation
// @Tags HkCircleRelation
// @Summary 更新HkCircleRelation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircleRelation true "更新HkCircleRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleRelation/updateHkCircleRelation [put]
func (hkCircleRelationApi *HkCircleRelationApi) UpdateHkCircleRelation(c *gin.Context) {
	var hkCircleRelation community.HkCircleRelation
	err := c.ShouldBindJSON(&hkCircleRelation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleRelationService.UpdateHkCircleRelation(hkCircleRelation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkCircleRelation 用id查询HkCircleRelation
// @Tags HkCircleRelation
// @Summary 用id查询HkCircleRelation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkCircleRelation true "用id查询HkCircleRelation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleRelation/findHkCircleRelation [get]
func (hkCircleRelationApi *HkCircleRelationApi) FindHkCircleRelation(c *gin.Context) {
	var hkCircleRelation community.HkCircleRelation
	err := c.ShouldBindQuery(&hkCircleRelation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleRelation, err := hkCircleRelationService.GetHkCircleRelation(hkCircleRelation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircleRelation": rehkCircleRelation}, c)
	}
}

// GetHkCircleRelationList 分页获取HkCircleRelation列表
// @Tags HkCircleRelation
// @Summary 分页获取HkCircleRelation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkCircleRelationSearch true "分页获取HkCircleRelation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleRelation/getHkCircleRelationList [get]
func (hkCircleRelationApi *HkCircleRelationApi) GetHkCircleRelationList(c *gin.Context) {
	var pageInfo communityReq.HkCircleRelationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkCircleRelationService.GetHkCircleRelationInfoList(pageInfo); err != nil {
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
