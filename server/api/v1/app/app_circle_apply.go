package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/apply/request"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
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
// @Tags App_CircleApply
// @Summary 创建UserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body applyReq.CreateUserCircleApplyReq true "创建UserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circleApply/createUserCircleApply [post]
func (circleApplyApi *CircleApplyApi) CreateUserCircleApply(c *gin.Context) {
	var hkUserCircleApply applyReq.CreateUserCircleApplyReq
	err := c.ShouldBindJSON(&hkUserCircleApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := appUserCircleApplyService.CreateHkUserCircleApply(hkUserCircleApply); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
}

// DeleteUserCircleApply 删除UserCircleApply
// @Tags App_CircleApply
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
	//if err := appUserCircleApplyService.DeleteHkUserCircleApply(hkUserCircleApply); err != nil {
	//	global.GVA_LOG.Error("删除失败!", zap.Error(err))
	//	response.FailWithMessage("删除失败", c)
	//} else {
	//	response.OkWithMessage("删除成功", c)
	//}
}

// DeleteUserCircleApplyByIds 批量删除UserCircleApply
// @Tags App_CircleApply
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
	if err := appUserCircleApplyService.DeleteHkUserCircleApplyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// GetUserCircleApplyListALL 获取UserCircleApply列表
// @Tags App_CircleApply
// @Summary 获取UserCircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object}  response.PageResult{List=[]communityReq.HkUserCircleApplySearch,msg=string} "返回communityReq.HkUserCircleApplySearch"
// @Router /app/circleApply/getUserCircleApplyListALL [get]
func (circleApplyApi *CircleApplyApi) GetUserCircleApplyListALL(c *gin.Context) {
	var pageInfo communityReq.HkUserCircleApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := appUserCircleApplyService.GetHkUserCircleApplyInfoListALL(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}

/*************************************
应用
**************************************/

// FindApply 用id查询Apply
// @Tags App_CircleApply
// @Summary 用id查询Apply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询Apply"
// @Success 200 {object}  response.Response{data=apply.HkApply,msg=string}  "返回apply.HkApply"
// @Router /app/circleApply/findApply [get]
func (circleApplyApi *CircleApplyApi) FindApply(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//var aa apply.HkApply
	if rehkApply, err := appApplyService.GetHkApply(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkApply": rehkApply}, c)
	}
}

// GetApplyList 分页获取Apply列表
// @Tags App_CircleApply
// @Summary 分页获取Apply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.ApplySearchReq true "分页获取Apply列表"
// @Success 200 {object}  response.PageResult{List=[]apply.HkApply,msg=string} "返回apply.HkApply"
// @Router /app/circleApply/getApplyList [get]
func (circleApplyApi *CircleApplyApi) GetApplyList(c *gin.Context) {
	var pageInfo applyReq.ApplySearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := appApplyService.GetHkApplyInfoList(pageInfo); err != nil {
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
// @Tags App_CircleApply
// @Summary 获取Apply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.ApplySearchReq true "获取Apply列表"
// @Success 200 {object}  response.PageResult{List=[]apply.HkApply,msg=string} "返回apply.HkApply"
// @Router /app/circleApply/getApplyListAll [get]
func (circleApplyApi *CircleApplyApi) GetApplyListAll(c *gin.Context) {
	var pageInfo applyReq.ApplySearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := appApplyService.GetHkApplyInfoList(pageInfo); err != nil {
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
// @Tags App_CircleApply
// @Summary 用id查询CircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询CircleApply"
// @Success 200 {object}  response.Response{data=apply.HkCircleApply,msg=string}  "返回apply.HkCircleApply"
// @Router /app/circleApply/findCircleApply [get]
func (circleApplyApi *CircleApplyApi) FindCircleApply(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleApply, err := appCircleApplyService.GetHkCircleApply(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircleApply": rehkCircleApply}, c)
	}
}

// GetCircleApplyList 分页获取CircleApply列表
// @Tags App_CircleApply
// @Summary 分页获取CircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.CircleApplySearchReq true "分页获取CircleApply列表"
// @Success 200 {object}  response.PageResult{List=[]apply.HkCircleApply,msg=string} "返回apply.HkCircleApply"
// @Router /app/circleApply/getCircleApplyList [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyList(c *gin.Context) {
	var pageInfo applyReq.CircleApplySearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := appCircleApplyService.GetHkCircleApplyInfoList(pageInfo); err != nil {
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
// @Tags App_CircleApply
// @Summary 获取CircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.CircleApplySearchReq true "获取CircleApply列表"
// @Success 200 {object}  response.PageResult{List=[]apply.HkCircleApply,msg=string} "返回apply.HkCircleApply"
// @Router /app/circleApply/getCircleApplyListAll [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyListAll(c *gin.Context) {
	var pageInfo applyReq.CircleApplySearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := appCircleApplyService.GetHkCircleApplyInfoListAll(pageInfo); err != nil {
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
// @Tags App_CircleApply
// @Summary 分页获取CircleApplyGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.CircleApplyGroupSearchReq true "分页获取CircleApplyGroup列表"
// @Success 200 {object}  response.PageResult{List=[]apply.HkCircleApplyGroup,msg=string} "返回apply.HkCircleApplyGroup"
// @Router /app/circleApply/getCircleApplyGroupList [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyGroupList(c *gin.Context) {
	var pageInfo applyReq.CircleApplyGroupSearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := appCircleApplyGroupService.GetHkCircleApplyGroupInfoList(pageInfo); err != nil {
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
// @Tags App_CircleApply
// @Summary 获取CircleApplyGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.CircleApplyGroupSearchReq true "获取CircleApplyGroup列表"
// @Success 200 {object}  response.PageResult{List=[]apply.HkCircleApplyGroup,msg=string} "返回apply.HkCircleApplyGroup"
// @Router /app/circleApply/getCircleApplyGroupListAll [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyGroupListAll(c *gin.Context) {
	var pageInfo applyReq.CircleApplyGroupSearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := appCircleApplyGroupService.GetHkCircleApplyGroupInfoList(pageInfo); err != nil {
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
