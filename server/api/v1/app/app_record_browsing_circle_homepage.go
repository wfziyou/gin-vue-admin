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

type RecordBrowsingCircleHomepageApi struct {
}

// CreateRecordBrowsingCircleHomepage 创建RecordBrowsingCircleHomepage
// @Tags RecordBrowsingCircleHomepage
// @Summary 创建RecordBrowsingCircleHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.RecordBrowsingCircleHomepage true "创建RecordBrowsingCircleHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkRecordBrowsingCircleHomepage/createRecordBrowsingCircleHomepage [post]
func (hkRecordBrowsingCircleHomepageApi *RecordBrowsingCircleHomepageApi) CreateRecordBrowsingCircleHomepage(c *gin.Context) {
	var hkRecordBrowsingCircleHomepage community.RecordBrowsingCircleHomepage
	err := c.ShouldBindJSON(&hkRecordBrowsingCircleHomepage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkRecordBrowsingCircleHomepageService.CreateRecordBrowsingCircleHomepage(&hkRecordBrowsingCircleHomepage); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRecordBrowsingCircleHomepage 删除RecordBrowsingCircleHomepage
// @Tags RecordBrowsingCircleHomepage
// @Summary 删除RecordBrowsingCircleHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.RecordBrowsingCircleHomepage true "删除RecordBrowsingCircleHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkRecordBrowsingCircleHomepage/deleteRecordBrowsingCircleHomepage [delete]
func (hkRecordBrowsingCircleHomepageApi *RecordBrowsingCircleHomepageApi) DeleteRecordBrowsingCircleHomepage(c *gin.Context) {
	var hkRecordBrowsingCircleHomepage community.RecordBrowsingCircleHomepage
	err := c.ShouldBindJSON(&hkRecordBrowsingCircleHomepage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkRecordBrowsingCircleHomepageService.DeleteRecordBrowsingCircleHomepage(hkRecordBrowsingCircleHomepage); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRecordBrowsingCircleHomepageByIds 批量删除RecordBrowsingCircleHomepage
// @Tags RecordBrowsingCircleHomepage
// @Summary 批量删除RecordBrowsingCircleHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RecordBrowsingCircleHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hkRecordBrowsingCircleHomepage/deleteRecordBrowsingCircleHomepageByIds [delete]
func (hkRecordBrowsingCircleHomepageApi *RecordBrowsingCircleHomepageApi) DeleteRecordBrowsingCircleHomepageByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkRecordBrowsingCircleHomepageService.DeleteRecordBrowsingCircleHomepageByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRecordBrowsingCircleHomepage 更新RecordBrowsingCircleHomepage
// @Tags RecordBrowsingCircleHomepage
// @Summary 更新RecordBrowsingCircleHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.RecordBrowsingCircleHomepage true "更新RecordBrowsingCircleHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkRecordBrowsingCircleHomepage/updateRecordBrowsingCircleHomepage [put]
func (hkRecordBrowsingCircleHomepageApi *RecordBrowsingCircleHomepageApi) UpdateRecordBrowsingCircleHomepage(c *gin.Context) {
	var hkRecordBrowsingCircleHomepage community.RecordBrowsingCircleHomepage
	err := c.ShouldBindJSON(&hkRecordBrowsingCircleHomepage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkRecordBrowsingCircleHomepageService.UpdateRecordBrowsingCircleHomepage(hkRecordBrowsingCircleHomepage); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRecordBrowsingCircleHomepage 用id查询RecordBrowsingCircleHomepage
// @Tags RecordBrowsingCircleHomepage
// @Summary 用id查询RecordBrowsingCircleHomepage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.RecordBrowsingCircleHomepage true "用id查询RecordBrowsingCircleHomepage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkRecordBrowsingCircleHomepage/findRecordBrowsingCircleHomepage [get]
func (hkRecordBrowsingCircleHomepageApi *RecordBrowsingCircleHomepageApi) FindRecordBrowsingCircleHomepage(c *gin.Context) {
	var hkRecordBrowsingCircleHomepage community.RecordBrowsingCircleHomepage
	err := c.ShouldBindQuery(&hkRecordBrowsingCircleHomepage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkRecordBrowsingCircleHomepage, err := hkRecordBrowsingCircleHomepageService.GetRecordBrowsingCircleHomepage(hkRecordBrowsingCircleHomepage.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkRecordBrowsingCircleHomepage": rehkRecordBrowsingCircleHomepage}, c)
	}
}

// GetRecordBrowsingCircleHomepageList 分页获取RecordBrowsingCircleHomepage列表
// @Tags RecordBrowsingCircleHomepage
// @Summary 分页获取RecordBrowsingCircleHomepage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.RecordBrowsingCircleHomepageSearch true "分页获取RecordBrowsingCircleHomepage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkRecordBrowsingCircleHomepage/getRecordBrowsingCircleHomepageList [get]
func (hkRecordBrowsingCircleHomepageApi *RecordBrowsingCircleHomepageApi) GetRecordBrowsingCircleHomepageList(c *gin.Context) {
	var pageInfo communityReq.RecordBrowsingCircleHomepageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkRecordBrowsingCircleHomepageService.GetRecordBrowsingCircleHomepageInfoList(pageInfo); err != nil {
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
