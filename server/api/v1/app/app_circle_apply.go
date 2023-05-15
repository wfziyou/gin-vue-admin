package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CircleApplyApi struct {
}

/*************************************
应用
**************************************/

// FindApply 用id查询Apply
// @Tags 圈子应用
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
		response.OkWithData(rehkApply, c)
	}
}

// GetApplyList 分页获取Apply列表
// @Tags 圈子应用
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
// @Tags 圈子应用
// @Summary 获取Apply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.ApplyAllSearchReq true "获取Apply列表"
// @Success 200 {object}  response.PageResult{List=[]apply.Apply,msg=string} "返回apply.Apply"
// @Router /app/circleApply/getApplyListAll [get]
func (circleApplyApi *CircleApplyApi) GetApplyListAll(c *gin.Context) {
	var pageInfo applyReq.ApplyAllSearchReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := appApplyService.GetApplyInfoListAll(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// FindCircleApply 用id查询CircleApply
// @Tags 圈子应用
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
		return
	} else {
		userId := utils.GetUserID(c)
		var isMember = true
		if _, err := appCircleUserService.GetCircleUser(rehkCircleApply.CircleId, userId); err != nil {
			isMember = false
		}
		if isMember == false && rehkCircleApply.Power == 1 {
			response.FailWithMessage("查询失败", c)
			return
		}
		response.OkWithData(rehkCircleApply, c)
	}
}

// GetCircleApplyList 分页获取CircleApply列表
// @Tags 圈子应用
// @Summary 分页获取CircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.CircleApplySearch true "分页获取CircleApply列表"
// @Success 200 {object}  response.PageResult{List=[]apply.CircleApply,msg=string} "返回apply.CircleApply"
// @Router /app/circleApply/getCircleApplyList [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyList(c *gin.Context) {
	var req applyReq.CircleApplySearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	var isMember = true
	if _, err := appCircleUserService.GetCircleUser(req.CircleId, userId); err != nil {
		isMember = false
	}
	if list, total, err := appCircleApplyService.GetCircleApplyInfoList(req, isMember); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// GetCircleApplyListAll 获取CircleApply列表
// @Tags 圈子应用
// @Summary 获取CircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.CircleApplySearchAll true "获取CircleApply列表"
// @Success 200 {object}  response.Response{data=[]apply.CircleApply,msg=string} "返回apply.CircleApply"
// @Router /app/circleApply/getCircleApplyListAll [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyListAll(c *gin.Context) {
	var req applyReq.CircleApplySearchAll
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	var isMember = true
	if _, err := appCircleUserService.GetCircleUser(req.CircleId, userId); err != nil {
		isMember = false
	}
	if list, err := appCircleApplyService.GetCircleApplyInfoListAll(req, isMember); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// GetCircleHotApplyList 获取圈子热门应用列表
// @Tags 圈子应用
// @Summary 获取圈子热门应用列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.CircleHotApplySearch true "获取圈子热门应用列表"
// @Success 200 {object}  response.Response{data=[]apply.CircleApply,msg=string} "返回apply.CircleApply"
// @Router /app/circleApply/getCircleHotApplyList [get]
func (circleApplyApi *CircleApplyApi) GetCircleHotApplyList(c *gin.Context) {
	var req applyReq.CircleHotApplySearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)
	var isMember = true
	if _, err := appCircleUserService.GetCircleUser(req.CircleId, userId); err != nil {
		isMember = false
	}
	if list, _, err := appCircleApplyService.GetCircleHotApplyList(req.CircleId, isMember); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// GetCircleApplyGroupList 分页获取CircleApplyGroup列表
// @Tags 圈子应用
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
// @Tags 圈子应用
// @Summary 获取CircleApplyGroup列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.CircleApplyGroupSearchAll true "获取CircleApplyGroup列表"
// @Success 200 {object}  response.Response{data=[]apply.CircleApplyGroup,msg=string} "返回apply.CircleApplyGroup"
// @Router /app/circleApply/getCircleApplyGroupListAll [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyGroupListAll(c *gin.Context) {
	var pageInfo applyReq.CircleApplyGroupSearchAll
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := appCircleApplyGroupService.GetCircleApplyGroupInfoListAll(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// AddCircleApply (圈子管理者)添加圈子应用
// @Tags 圈子应用
// @Summary (圈子管理者)添加圈子应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.ParamAddCircleApply true "(圈子管理者)添加圈子应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /app/circleApply/addCircleApply [post]
func (circleApplyApi *CircleApplyApi) AddCircleApply(c *gin.Context) {
	var pageInfo applyReq.ParamAddCircleApply
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

}

// UpdateCircleApply (圈子管理者)更新圈子应用
// @Tags 圈子应用
// @Summary (圈子管理者)更新圈子应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.ParamUpdateCircleApply true "(圈子管理者)更新圈子应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /app/circleApply/updateCircleApply [post]
func (circleApplyApi *CircleApplyApi) UpdateCircleApply(c *gin.Context) {
	var pageInfo applyReq.ParamUpdateCircleApply
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

}

// DeleteCircleApply (圈子管理者)删除圈子应用
// @Tags 圈子应用
// @Summary (圈子管理者)删除圈子应用
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.ParamUpdateCircleApply true "(圈子管理者)删除圈子应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /app/circleApply/deleteCircleApply [post]
func (circleApplyApi *CircleApplyApi) DeleteCircleApply(c *gin.Context) {
	var pageInfo applyReq.ParamUpdateCircleApply
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

}
