package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RecordBrowsingUserHomepageApi struct {
}

// CreateRecordBrowsingUserHomepage 创建RecordBrowsingUserHomepage
// @Tags RecordBrowsingUserHomepage
// @Summary 创建RecordBrowsingUserHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.RecordBrowsingUserHomepage true "创建RecordBrowsingUserHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkRecordBrowsingUserHomepage/createRecordBrowsingUserHomepage [post]
func (hkRecordBrowsingUserHomepageApi *RecordBrowsingUserHomepageApi) CreateRecordBrowsingUserHomepage(c *gin.Context) {
	var hkRecordBrowsingUserHomepage community.RecordBrowsingUserHomepage
	err := c.ShouldBindJSON(&hkRecordBrowsingUserHomepage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkRecordBrowsingUserHomepageService.CreateRecordBrowsingUserHomepage(&hkRecordBrowsingUserHomepage); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRecordBrowsingUserHomepage 删除RecordBrowsingUserHomepage
// @Tags RecordBrowsingUserHomepage
// @Summary 删除RecordBrowsingUserHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.RecordBrowsingUserHomepage true "删除RecordBrowsingUserHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkRecordBrowsingUserHomepage/deleteRecordBrowsingUserHomepage [delete]
func (hkRecordBrowsingUserHomepageApi *RecordBrowsingUserHomepageApi) DeleteRecordBrowsingUserHomepage(c *gin.Context) {
	var hkRecordBrowsingUserHomepage community.RecordBrowsingUserHomepage
	err := c.ShouldBindJSON(&hkRecordBrowsingUserHomepage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkRecordBrowsingUserHomepageService.DeleteRecordBrowsingUserHomepage(hkRecordBrowsingUserHomepage); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRecordBrowsingUserHomepageByIds 批量删除RecordBrowsingUserHomepage
// @Tags RecordBrowsingUserHomepage
// @Summary 批量删除RecordBrowsingUserHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RecordBrowsingUserHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkRecordBrowsingUserHomepage/deleteRecordBrowsingUserHomepageByIds [delete]
func (hkRecordBrowsingUserHomepageApi *RecordBrowsingUserHomepageApi) DeleteRecordBrowsingUserHomepageByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkRecordBrowsingUserHomepageService.DeleteRecordBrowsingUserHomepageByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRecordBrowsingUserHomepage 更新RecordBrowsingUserHomepage
// @Tags RecordBrowsingUserHomepage
// @Summary 更新RecordBrowsingUserHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.RecordBrowsingUserHomepage true "更新RecordBrowsingUserHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkRecordBrowsingUserHomepage/updateRecordBrowsingUserHomepage [put]
func (hkRecordBrowsingUserHomepageApi *RecordBrowsingUserHomepageApi) UpdateRecordBrowsingUserHomepage(c *gin.Context) {
	var hkRecordBrowsingUserHomepage community.RecordBrowsingUserHomepage
	err := c.ShouldBindJSON(&hkRecordBrowsingUserHomepage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkRecordBrowsingUserHomepageService.UpdateRecordBrowsingUserHomepage(hkRecordBrowsingUserHomepage); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRecordBrowsingUserHomepage 用id查询RecordBrowsingUserHomepage
// @Tags RecordBrowsingUserHomepage
// @Summary 用id查询RecordBrowsingUserHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.RecordBrowsingUserHomepage true "用id查询RecordBrowsingUserHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkRecordBrowsingUserHomepage/findRecordBrowsingUserHomepage [get]
func (hkRecordBrowsingUserHomepageApi *RecordBrowsingUserHomepageApi) FindRecordBrowsingUserHomepage(c *gin.Context) {
	var hkRecordBrowsingUserHomepage community.RecordBrowsingUserHomepage
	err := c.ShouldBindQuery(&hkRecordBrowsingUserHomepage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkRecordBrowsingUserHomepage, err := hkRecordBrowsingUserHomepageService.GetRecordBrowsingUserHomepage(hkRecordBrowsingUserHomepage.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkRecordBrowsingUserHomepage": rehkRecordBrowsingUserHomepage}, c)
	}
}

// GetRecordBrowsingUserHomepageList 分页获取RecordBrowsingUserHomepage列表
// @Tags RecordBrowsingUserHomepage
// @Summary 分页获取RecordBrowsingUserHomepage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.RecordBrowsingUserHomepageSearch true "分页获取RecordBrowsingUserHomepage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkRecordBrowsingUserHomepage/getRecordBrowsingUserHomepageList [get]
func (hkRecordBrowsingUserHomepageApi *RecordBrowsingUserHomepageApi) GetRecordBrowsingUserHomepageList(c *gin.Context) {
	var pageInfo communityReq.RecordBrowsingUserHomepageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkRecordBrowsingUserHomepageService.GetRecordBrowsingUserHomepageInfoList(pageInfo); err != nil {
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
