package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/apply/request"
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

// GetUserCircleApplyListALL 获取UserCircleApply列表
// @Tags 圈子应用
// @Summary 获取UserCircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.UserCircleApplySearch true "获取UserCircleApply列表"
// @Success 200 {object}  response.Response{data=[]community.UserCircleApply,msg=string} "返回communityReq.UserCircleApplySearch"
// @Router /app/circleApply/getUserCircleApplyListAll [get]
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

	if list, _, err := appUserCircleApplyService.GetUserCircleApplyInfoListALL(req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// SetUserCircleApply 设置UserCircleApply
// @Tags 圈子应用
// @Summary 设置UserCircleApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.UserCircleApplyUpdate true "设置UserCircleApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circleApply/setUserCircleApply [post]
func (circleApplyApi *CircleApplyApi) SetUserCircleApply(c *gin.Context) {
	var req communityReq.UserCircleApplyUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var userId = utils.GetUserID(c)
	req.UserId = userId
	if _, err := appCircleService.GetCircle(req.CircleId); err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}

	if data, _, err := appCircleUserService.GetCircleUserInfoList(communityReq.CircleUserSearch{
		CircleId: req.CircleId,
		UserId:   userId,
		PageInfo: request.PageInfo{Page: 1, PageSize: 2},
	}); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else if len(data) == 0 {
		response.FailWithMessage("用户不在圈子中", c)
		return
	}

	if err = appUserCircleApplyService.SetUserCircleApply(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
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
	} else {
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
// @Tags 圈子应用
// @Summary 获取CircleApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query applyReq.CircleApplySearchAll true "获取CircleApply列表"
// @Success 200 {object}  response.Response{data=[]apply.CircleApply,msg=string} "返回apply.CircleApply"
// @Router /app/circleApply/getCircleApplyListAll [get]
func (circleApplyApi *CircleApplyApi) GetCircleApplyListAll(c *gin.Context) {
	var pageInfo applyReq.CircleApplySearchAll
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, _, err := appCircleApplyService.GetCircleApplyInfoListAll(pageInfo); err != nil {
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
