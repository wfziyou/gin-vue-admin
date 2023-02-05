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

type HkForumThumbsUpApi struct {
}

var hkForumThumbsUpService = service.ServiceGroupApp.CommunityServiceGroup.HkForumThumbsUpService

// CreateHkForumThumbsUp 创建HkForumThumbsUp
// @Tags HkForumThumbsUp
// @Summary 创建HkForumThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumThumbsUp true "创建HkForumThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumThumbsUp/createHkForumThumbsUp [post]
func (hkForumThumbsUpApi *HkForumThumbsUpApi) CreateHkForumThumbsUp(c *gin.Context) {
	var hkForumThumbsUp community.HkForumThumbsUp
	err := c.ShouldBindJSON(&hkForumThumbsUp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumThumbsUpService.CreateHkForumThumbsUp(hkForumThumbsUp); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHkForumThumbsUp 删除HkForumThumbsUp
// @Tags HkForumThumbsUp
// @Summary 删除HkForumThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumThumbsUp true "删除HkForumThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForumThumbsUp/deleteHkForumThumbsUp [delete]
func (hkForumThumbsUpApi *HkForumThumbsUpApi) DeleteHkForumThumbsUp(c *gin.Context) {
	var hkForumThumbsUp community.HkForumThumbsUp
	err := c.ShouldBindJSON(&hkForumThumbsUp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumThumbsUpService.DeleteHkForumThumbsUp(hkForumThumbsUp); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHkForumThumbsUpByIds 批量删除HkForumThumbsUp
// @Tags HkForumThumbsUp
// @Summary 批量删除HkForumThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForumThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkForumThumbsUp/deleteHkForumThumbsUpByIds [delete]
func (hkForumThumbsUpApi *HkForumThumbsUpApi) DeleteHkForumThumbsUpByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumThumbsUpService.DeleteHkForumThumbsUpByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHkForumThumbsUp 更新HkForumThumbsUp
// @Tags HkForumThumbsUp
// @Summary 更新HkForumThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkForumThumbsUp true "更新HkForumThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForumThumbsUp/updateHkForumThumbsUp [put]
func (hkForumThumbsUpApi *HkForumThumbsUpApi) UpdateHkForumThumbsUp(c *gin.Context) {
	var hkForumThumbsUp community.HkForumThumbsUp
	err := c.ShouldBindJSON(&hkForumThumbsUp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkForumThumbsUpService.UpdateHkForumThumbsUp(hkForumThumbsUp); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHkForumThumbsUp 用id查询HkForumThumbsUp
// @Tags HkForumThumbsUp
// @Summary 用id查询HkForumThumbsUp
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询HkForumThumbsUp"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForumThumbsUp/findHkForumThumbsUp [get]
func (hkForumThumbsUpApi *HkForumThumbsUpApi) FindHkForumThumbsUp(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkForumThumbsUp, err := hkForumThumbsUpService.GetHkForumThumbsUp(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkForumThumbsUp": rehkForumThumbsUp}, c)
	}
}

// GetHkForumThumbsUpList 分页获取HkForumThumbsUp列表
// @Tags HkForumThumbsUp
// @Summary 分页获取HkForumThumbsUp列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.HkForumThumbsUpSearch true "分页获取HkForumThumbsUp列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForumThumbsUp/getHkForumThumbsUpList [get]
func (hkForumThumbsUpApi *HkForumThumbsUpApi) GetHkForumThumbsUpList(c *gin.Context) {
	var pageInfo communityReq.HkForumThumbsUpSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForumThumbsUpService.GetHkForumThumbsUpInfoList(pageInfo); err != nil {
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
