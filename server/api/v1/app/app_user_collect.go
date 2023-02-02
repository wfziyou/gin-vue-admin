package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserCollectApi struct {
}

// CreateUserCollect 收藏帖子
// @Tags App_UserCollect
// @Summary 收藏帖子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body appReq.UserCollectReq true "收藏帖子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/userCollect/createUserCollect [post]
func (userCollectApi *UserCollectApi) CreateUserCollect(c *gin.Context) {
	var hkUserCollect appReq.UserCollectReq
	err := c.ShouldBindJSON(&hkUserCollect)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := hkUserCollectService.CreateHkUserCollect(hkUserCollect); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
}

// DeleteUserCollect 删除UserCollect
// @Tags App_UserCollect
// @Summary 删除UserCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除UserCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/userCollect/deleteUserCollect [delete]
func (userCollectApi *UserCollectApi) DeleteUserCollect(c *gin.Context) {
	var hkUserCollect request.IdDelete
	err := c.ShouldBindJSON(&hkUserCollect)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := hkUserCollectService.DeleteHkUserCollect(hkUserCollect); err != nil {
	//	global.GVA_LOG.Error("删除失败!", zap.Error(err))
	//	response.FailWithMessage("删除失败", c)
	//} else {
	//	response.OkWithMessage("删除成功", c)
	//}
}

// DeleteUserCollectByIds 批量删除UserCollect
// @Tags App_UserCollect
// @Summary 批量删除UserCollect
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserCollect"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/userCollect/deleteUserCollectByIds [delete]
func (userCollectApi *UserCollectApi) DeleteUserCollectByIds(c *gin.Context) {
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

// GetUserCollectList 分页获取UserCollect列表
// @Tags App_UserCollect
// @Summary 分页获取UserCollect列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.UserCollectSearch true "分页获取UserCollect列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/userCollect/getUserCollectList [get]
func (userCollectApi *UserCollectApi) GetUserCollectList(c *gin.Context) {
	var pageInfo appReq.UserCollectSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := hkUserCollectService.GetHkUserCollectInfoList(pageInfo); err != nil {
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("获取失败", c)
	//} else {
	//	response.OkWithDetailed(response.PageResult{
	//		List:     list,
	//		Total:    total,
	//		Page:     pageInfo.Page,
	//		PageSize: pageInfo.PageSize,
	//	}, "获取成功", c)
	//}
}
