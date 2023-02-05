package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HkForumTopicApi struct {
}

var hkForumTopicService = service.ServiceGroupApp.CommunityServiceGroup.HkForumTopicService

// CreateHkForumTopic 创建HkForumTopic
// @Tags HkForumTopic
// @Summary 创建HkForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumTopic true "创建HkForumTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTopic/createHkForumTopic [post]
func (hkForumTopicApi *HkForumTopicApi) CreateHkForumTopic(c *gin.Context) {
	var hkForumTopic community.HkForumTopic
	err := c.ShouldBindJSON(&hkForumTopic)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTopicService.CreateHkForumTopic(hkForumTopic); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkForumTopic 删除HkForumTopic
// @Tags HkForumTopic
// @Summary 删除HkForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumTopic true "删除HkForumTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumTopic/deleteHkForumTopic [delete]
func (hkForumTopicApi *HkForumTopicApi) DeleteHkForumTopic(c *gin.Context) {
	var hkForumTopic community.HkForumTopic
	err := c.ShouldBindJSON(&hkForumTopic)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTopicService.DeleteHkForumTopic(hkForumTopic); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkForumTopicByIds 批量删除HkForumTopic
// @Tags HkForumTopic
// @Summary 批量删除HkForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkForumTopic/deleteHkForumTopicByIds [delete]
func (hkForumTopicApi *HkForumTopicApi) DeleteHkForumTopicByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTopicService.DeleteHkForumTopicByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkForumTopic 更新HkForumTopic
// @Tags HkForumTopic
// @Summary 更新HkForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumTopic true "更新HkForumTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumTopic/updateHkForumTopic [put]
func (hkForumTopicApi *HkForumTopicApi) UpdateHkForumTopic(c *gin.Context) {
	var hkForumTopic community.HkForumTopic
	err := c.ShouldBindJSON(&hkForumTopic)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumTopicService.UpdateHkForumTopic(hkForumTopic); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkForumTopic 用id查询HkForumTopic
// @Tags HkForumTopic
// @Summary 用id查询HkForumTopic
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.HkForumTopic true "用id查询HkForumTopic"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumTopic/findHkForumTopic [get]
func (hkForumTopicApi *HkForumTopicApi) FindHkForumTopic(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumTopic, err := hkForumTopicService.GetHkForumTopic(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForumTopic": rehkForumTopic}, c)
	}
}

// GetHkForumTopicList 分页获取HkForumTopic列表
// @Tags HkForumTopic
// @Summary 分页获取HkForumTopic列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkForumTopicSearch true "分页获取HkForumTopic列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumTopic/getHkForumTopicList [get]
func (hkForumTopicApi *HkForumTopicApi) GetHkForumTopicList(c *gin.Context) {
	var pageInfo communityReq.HkForumTopicSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForumTopicService.GetHkForumTopicInfoList(pageInfo); err != nil {
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
