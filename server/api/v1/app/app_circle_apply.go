package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CircleApplyApi struct {
}

/*************************************
用户圈子应用
**************************************/

// CreateUserCircleApply 创建UserCircleApply
// @Tags CircleApply
// @Summary 创建UserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body appReq.CreateUserCircleApplyReq true "创建UserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circleApply/createUserCircleApply [post]
func (circleApplyApi *CircleApplyApi) CreateUserCircleApply(c *gin.Context) {
	var hkUserCircleApply appReq.CreateUserCircleApplyReq
	err := c.ShouldBindJSON(&hkUserCircleApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := hkUserCircleApplyService.CreateHkUserCircleApply(hkUserCircleApply); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
}

// DeleteUserCircleApply 删除UserCircleApply
// @Tags CircleApply
// @Summary 删除UserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除UserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/circleApply/deleteUserCircleApply [delete]
func (circleApplyApi *CircleApplyApi) DeleteUserCircleApply(c *gin.Context) {
	var hkUserCircleApply request.IdDelete

	err := c.ShouldBindJSON(&hkUserCircleApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := hkUserCircleApplyService.DeleteHkUserCircleApply(hkUserCircleApply); err != nil {
	//	global.GVA_LOG.Error("删除失败!", zap.Error(err))
	//	response.FailWithMessage("删除失败", c)
	//} else {
	//	response.OkWithMessage("删除成功", c)
	//}
}

// DeleteUserCircleApplyByIds 批量删除UserCircleApply
// @Tags CircleApply
// @Summary 批量删除UserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/circleApply/deleteUserCircleApplyByIds [delete]
func (circleApplyApi *CircleApplyApi) DeleteUserCircleApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkUserCircleApplyService.DeleteHkUserCircleApplyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// GetUserCircleApplyListALL 获取UserCircleApply列表
// @Tags CircleApply
// @Summary 获取UserCircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circleApply/getUserCircleApplyListALL [get]
func (circleApplyApi *CircleApplyApi) GetUserCircleApplyListALL(c *gin.Context) {
	//if list, total, err := hkUserCircleApplyService.GetHkUserCircleApplyInfoListALL(pageInfo); err != nil {
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("获取失败", c)
	//} else {
	//	response.OkWithDetailed(response.PageResult{
	//		List:  list,
	//		Total: total,
	//	}, "获取成功", c)
	//}
}

/*************************************
应用
**************************************/

// FindApply 用id查询Apply
// @Tags CircleApply
// @Summary 用id查询Apply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询Apply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/circleApply/FindApply [get]
func (circleApplyApi *CircleApplyApi) FindApply(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkApply, err := hkApplyService.GetHkApply(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkApply": rehkApply}, c)
	}
}

// GetApplyList 分页获取Apply列表
// @Tags CircleApply
// @Summary 分页获取Apply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.ApplySearchReq true "分页获取Apply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circleApply/GetApplyList [get]
func (circleApplyApi *CircleApplyApi) GetApplyList(c *gin.Context) {
	var pageInfo appReq.ApplySearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := hkApplyService.GetHkApplyInfoList(pageInfo); err != nil {
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

// GetApplyListAll 获取Apply列表
// @Tags CircleApply
// @Summary 获取Apply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.ApplySearchReq true "获取Apply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circleApply/getApplyListAll [get]
func (circleApplyApi *CircleApplyApi) GetApplyListAll(c *gin.Context) {
	var pageInfo appReq.ApplySearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := hkApplyService.GetHkApplyInfoList(pageInfo); err != nil {
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

// FindCircleApply 用id查询CircleApply
// @Tags CircleApply
// @Summary 用id查询CircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询CircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/circleApply/findCircleApply [get]
func (circleApplyApi *CircleApplyApi) FindCircleApply(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleApply, err := hkCircleApplyService.GetHkCircleApply(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircleApply": rehkCircleApply}, c)
	}
}

// GetCircleApplyList 分页获取CircleApply列表
// @Tags CircleApply
// @Summary 分页获取CircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.CircleApplySearchReq true "分页获取CircleApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circleApply/getCircleApplyList [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyList(c *gin.Context) {
	var pageInfo appReq.CircleApplySearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := hkCircleApplyService.GetHkCircleApplyInfoList(pageInfo); err != nil {
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

// GetCircleApplyListAll 获取CircleApply列表
// @Tags CircleApply
// @Summary 获取CircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.CircleApplySearchReq true "获取CircleApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circleApply/getCircleApplyListAll [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyListAll(c *gin.Context) {
	var pageInfo appReq.CircleApplySearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := hkCircleApplyService.GetHkCircleApplyInfoListAll(pageInfo); err != nil {
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("获取失败", c)
	//} else {
	//	response.OkWithDetailed(response.PageResult{
	//		List:  list,
	//		Total: total,
	//	}, "获取成功", c)
	//}
}

// GetCircleApplyGroupList 分页获取CircleApplyGroup列表
// @Tags CircleApply
// @Summary 分页获取CircleApplyGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.CircleApplyGroupSearchReq true "分页获取CircleApplyGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circleApply/getCircleApplyGroupList [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyGroupList(c *gin.Context) {
	var pageInfo appReq.CircleApplyGroupSearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := hkCircleApplyGroupService.GetHkCircleApplyGroupInfoList(pageInfo); err != nil {
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

// GetCircleApplyGroupListAll 获取CircleApplyGroup列表
// @Tags CircleApply
// @Summary 获取CircleApplyGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.CircleApplyGroupSearchReq true "获取CircleApplyGroup列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circleApply/getCircleApplyGroupListAll [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyGroupListAll(c *gin.Context) {
	var pageInfo appReq.CircleApplyGroupSearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := hkCircleApplyGroupService.GetHkCircleApplyGroupInfoList(pageInfo); err != nil {
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
