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

type HkForumTopicGroupApi struct {
}

var hkForumTopicGroupService = service.ServiceGroupApp.CommunityServiceGroup.HkForumTopicGroupService


// CreateHkForumTopicGroup 创建HkForumTopicGroup
// @Tags HkForumTopicGroup
// @Summary 创建HkForumTopicGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumTopicGroup true "创建HkForumTopicGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTopicGroup/createHkForumTopicGroup [post]
func (hkForumTopicGroupApi *HkForumTopicGroupApi) CreateHkForumTopicGroup(c *gin.Context) {
	var hkForumTopicGroup community.HkForumTopicGroup
	err := c.ShouldBindJSON(&hkForumTopicGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTopicGroupService.CreateHkForumTopicGroup(hkForumTopicGroup); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkForumTopicGroup 删除HkForumTopicGroup
// @Tags HkForumTopicGroup
// @Summary 删除HkForumTopicGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumTopicGroup true "删除HkForumTopicGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTopicGroup/deleteHkForumTopicGroup [delete]
func (hkForumTopicGroupApi *HkForumTopicGroupApi) DeleteHkForumTopicGroup(c *gin.Context) {
	var hkForumTopicGroup community.HkForumTopicGroup
	err := c.ShouldBindJSON(&hkForumTopicGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTopicGroupService.DeleteHkForumTopicGroup(hkForumTopicGroup); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkForumTopicGroupByIds 批量删除HkForumTopicGroup
// @Tags HkForumTopicGroup
// @Summary 批量删除HkForumTopicGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumTopicGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkForumTopicGroup/deleteHkForumTopicGroupByIds [delete]
func (hkForumTopicGroupApi *HkForumTopicGroupApi) DeleteHkForumTopicGroupByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTopicGroupService.DeleteHkForumTopicGroupByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkForumTopicGroup 更新HkForumTopicGroup
// @Tags HkForumTopicGroup
// @Summary 更新HkForumTopicGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumTopicGroup true "更新HkForumTopicGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumTopicGroup/updateHkForumTopicGroup [put]
func (hkForumTopicGroupApi *HkForumTopicGroupApi) UpdateHkForumTopicGroup(c *gin.Context) {
	var hkForumTopicGroup community.HkForumTopicGroup
	err := c.ShouldBindJSON(&hkForumTopicGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTopicGroupService.UpdateHkForumTopicGroup(hkForumTopicGroup); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkForumTopicGroup 用id查询HkForumTopicGroup
// @Tags HkForumTopicGroup
// @Summary 用id查询HkForumTopicGroup
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkForumTopicGroup true "用id查询HkForumTopicGroup"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumTopicGroup/findHkForumTopicGroup [get]
func (hkForumTopicGroupApi *HkForumTopicGroupApi) FindHkForumTopicGroup(c *gin.Context) {
	var hkForumTopicGroup community.HkForumTopicGroup
	err := c.ShouldBindQuery(&hkForumTopicGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumTopicGroup, err := hkForumTopicGroupService.GetHkForumTopicGroup(hkForumTopicGroup.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForumTopicGroup": rehkForumTopicGroup}, c)
	}
}

// GetHkForumTopicGroupList 分页获取HkForumTopicGroup列表
// @Tags HkForumTopicGroup
// @Summary 分页获取HkForumTopicGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkForumTopicGroupSearch true "分页获取HkForumTopicGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTopicGroup/getHkForumTopicGroupList [get]
func (hkForumTopicGroupApi *HkForumTopicGroupApi) GetHkForumTopicGroupList(c *gin.Context) {
	var pageInfo communityReq.HkForumTopicGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForumTopicGroupService.GetHkForumTopicGroupInfoList(pageInfo); err != nil {
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
