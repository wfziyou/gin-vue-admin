package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
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
	//if err := appUserCircleApplyService.CreateUserCircleApply(hkUserCircleApply); err != nil {
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
	var req request.IdDelete

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appUserCircleApplyService.DeleteUserCircleApply(community.UserCircleApply{GvaModelApp: global.GvaModelApp{ID: req.ID}}); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
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
	if err := appUserCircleApplyService.DeleteUserCircleApplyByIds(IDS); err != nil {
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
// @Success 200 {object}  response.PageResult{List=[]communityReq.UserCircleApplySearch,msg=string} "返回communityReq.UserCircleApplySearch"
// @Router /app/circleApply/getUserCircleApplyListALL [get]
func (circleApplyApi *CircleApplyApi) GetUserCircleApplyListALL(c *gin.Context) {
	var req communityReq.UserCircleApplySearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	req.UserId = utils.GetUserID(c)
	if _, err := appCircleUserService.GetCircleUserEx(req.CircleId, req.UserId); err != nil {
		response.FailWithMessage("用户没有在圈子中", c)
		return
	}

	if list, total, err := appUserCircleApplyService.GetUserCircleApplyInfoListALL(req); err != nil {
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
// @Success 200 {object}  response.Response{data=apply.Apply,msg=string}  "返回apply.Apply"
// @Router /app/circleApply/findApply [get]
func (circleApplyApi *CircleApplyApi) FindApply(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//var aa apply.Apply
	if rehkApply, err := appApplyService.GetApply(idSearch.ID); err != nil {
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
// @Success 200 {object}  response.PageResult{List=[]apply.Apply,msg=string} "返回apply.Apply"
// @Router /app/circleApply/getApplyList [get]
func (circleApplyApi *CircleApplyApi) GetApplyList(c *gin.Context) {
	var pageInfo applyReq.ApplySearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appApplyService.GetApplyInfoList(pageInfo); err != nil {
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

// GetApplyListAll 获取Apply列表
// @Tags App_CircleApply
// @Summary 获取Apply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.ApplySearchReq true "获取Apply列表"
// @Success 200 {object}  response.PageResult{List=[]apply.Apply,msg=string} "返回apply.Apply"
// @Router /app/circleApply/getApplyListAll [get]
func (circleApplyApi *CircleApplyApi) GetApplyListAll(c *gin.Context) {
	var pageInfo applyReq.ApplySearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appApplyService.GetApplyInfoListAll(pageInfo); err != nil {
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

// FindCircleApply 用id查询CircleApply
// @Tags App_CircleApply
// @Summary 用id查询CircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询CircleApply"
// @Success 200 {object}  response.Response{data=apply.CircleApply,msg=string}  "返回apply.CircleApply"
// @Router /app/circleApply/findCircleApply [get]
func (circleApplyApi *CircleApplyApi) FindCircleApply(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleApply, err := appCircleApplyService.GetCircleApply(idSearch.ID); err != nil {
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
// @Param data query applyReq.CircleApplySearch true "分页获取CircleApply列表"
// @Success 200 {object}  response.PageResult{List=[]apply.CircleApply,msg=string} "返回apply.CircleApply"
// @Router /app/circleApply/getCircleApplyList [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyList(c *gin.Context) {
	var pageInfo applyReq.CircleApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appCircleApplyService.GetCircleApplyInfoList(pageInfo); err != nil {
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

// GetCircleApplyListAll 获取CircleApply列表
// @Tags App_CircleApply
// @Summary 获取CircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.CircleApplySearchAll true "获取CircleApply列表"
// @Success 200 {object}  response.PageResult{List=[]apply.CircleApply,msg=string} "返回apply.CircleApply"
// @Router /app/circleApply/getCircleApplyListAll [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyListAll(c *gin.Context) {
	var pageInfo applyReq.CircleApplySearchAll
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appCircleApplyService.GetCircleApplyInfoListAll(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}

// GetCircleApplyGroupList 分页获取CircleApplyGroup列表
// @Tags App_CircleApply
// @Summary 分页获取CircleApplyGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.CircleApplyGroupSearch true "分页获取CircleApplyGroup列表"
// @Success 200 {object}  response.PageResult{List=[]apply.CircleApplyGroup,msg=string} "返回apply.CircleApplyGroup"
// @Router /app/circleApply/getCircleApplyGroupList [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyGroupList(c *gin.Context) {
	var pageInfo applyReq.CircleApplyGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appCircleApplyGroupService.GetCircleApplyGroupInfoList(pageInfo); err != nil {
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

// GetCircleApplyGroupListAll 获取CircleApplyGroup列表
// @Tags App_CircleApply
// @Summary 获取CircleApplyGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.CircleApplyGroupSearchAll true "获取CircleApplyGroup列表"
// @Success 200 {object}  response.PageResult{List=[]apply.CircleApplyGroup,msg=string} "返回apply.CircleApplyGroup"
// @Router /app/circleApply/getCircleApplyGroupListAll [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyGroupListAll(c *gin.Context) {
	var pageInfo applyReq.CircleApplyGroupSearchAll
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appCircleApplyGroupService.GetCircleApplyGroupInfoListAll(pageInfo); err != nil {
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
