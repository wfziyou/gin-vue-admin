package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CircleApi struct {
}

// GetCircleForumPostsList 分页获取圈子ForumPosts列表
// @Tags App_Circle
// @Summary 分页获取圈子ForumPosts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.CircleForumPostsSearch true "分页获取圈子ForumPosts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circle/getCircleForumPostsList [get]
func (circleApi *CircleApi) GetCircleForumPostsList(c *gin.Context) {
	var pageInfo appReq.CircleForumPostsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForumPostsService.GetForumPostsInfoList(pageInfo); err != nil {
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

// GetUserCircleForumPostsList 分页获取用户圈子ForumPosts列表
// @Tags App_Circle
// @Summary 分页获取用户圈子ForumPosts列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.UserCircleForumPostsSearch true "分页获取用户圈子ForumPosts列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circle/getUserCircleForumPostsList [get]
func (circleApi *CircleApi) GetUserCircleForumPostsList(c *gin.Context) {
	var pageInfo appReq.UserCircleForumPostsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkForumPostsService.GetUserForumPostsInfoList(pageInfo); err != nil {
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

// GetSelfCircleList 分页获取用户加入的Circle列表
// @Tags App_Circle
// @Summary 分页获取用户加入的Circle列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.SelfCircleSearch true "分页获取用户加入的Circle列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circle/getSelfCircleList [get]
func (circleApi *CircleApi) GetSelfCircleList(c *gin.Context) {
	var pageInfo appReq.SelfCircleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkCircleService.GetSelfCircleList(pageInfo); err != nil {
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

// FindCircle 用id查询Circle
// @Tags App_Circle
// @Summary 用id查询Circle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询Circle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/circle/findCircle [get]
func (circleApi *CircleApi) FindCircle(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircle, err := hkCircleService.GetHkCircle(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircle": rehkCircle}, c)
	}
}

// GetCircleList 分页获取Circle列表
// @Tags App_Circle
// @Summary 分页获取Circle列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.CircleSearch true "分页获取Circle列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circle/getCircleList [get]
func (circleApi *CircleApi) GetCircleList(c *gin.Context) {
	var pageInfo appReq.CircleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkCircleService.AppGetHkCircleInfoList(pageInfo); err != nil {
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

// UpdateCircle (圈子管理者)更新Circle
// @Tags App_Circle
// @Summary (圈子管理者)更新Circle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkCircle true "(圈子管理者)更新Circle"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/circle/updateCircle [put]
func (circleApi *CircleApi) UpdateCircle(c *gin.Context) {
	var hkCircle community.HkCircle
	err := c.ShouldBindJSON(&hkCircle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleService.UpdateHkCircle(hkCircle); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// SetUserCurCircle 设置用户当前圈子
// @Tags App_Circle
// @Summary 设置用户当前圈子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body appReq.SetUserCurCircleReq true "设置用户当前圈子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/circle/setUserCurCircle [post]
func (circleApi *CircleApi) SetUserCurCircle(c *gin.Context) {
	var hkCircle appReq.SetUserCurCircleReq
	err := c.ShouldBindJSON(&hkCircle)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
}

/*************************************
圈子成员
**************************************/

// DeleteCircleUser 删除CircleUser
// @Tags App_Circle
// @Summary 删除CircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body appReq.DeleteCircleUserReq true "删除CircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/circle/deleteHkCircleUser [delete]
func (circleApi *CircleApi) DeleteCircleUser(c *gin.Context) {
	var req appReq.DeleteCircleUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var pageInfo communityReq.HkCircleUserSearch
	pageInfo.CircleId = req.CircleId
	pageInfo.UserId = req.UserId
	pageInfo.Page = 1
	pageInfo.PageSize = 10
	if list, total, err := hkCircleUserService.GetHkCircleUserInfoList(pageInfo); err != nil {
		if total == 0 {
			response.OkWithMessage("圈子成员不存在", c)
			return
		}

		if err := hkCircleUserService.DeleteHkCircleUser(list[0]); err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			response.FailWithMessage("删除失败", c)
		} else {
			response.OkWithMessage("删除成功", c)
		}
	} else {
		response.OkWithMessage("圈子成员不存在", c)
	}
}

// DeleteCircleUserByIds 批量删除CircleUser
// @Tags App_Circle
// @Summary 批量删除CircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/circle/deleteCircleUserByIds [delete]
func (circleApi *CircleApi) DeleteCircleUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hkCircleUserService.DeleteHkCircleUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCircleUser 更新CircleUser
// @Tags App_Circle
// @Summary 更新CircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body appReq.UpdateCircleUserReq true "更新CircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/circle/updateCircleUser [put]
func (circleApi *CircleApi) UpdateCircleUser(c *gin.Context) {
	var hkCircleUser appReq.UpdateCircleUserReq
	err := c.ShouldBindJSON(&hkCircleUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := hkCircleUserService.UpdateHkCircleUser(hkCircleUser); err != nil {
	//	global.GVA_LOG.Error("更新失败!", zap.Error(err))
	//	response.FailWithMessage("更新失败", c)
	//} else {
	//	response.OkWithMessage("更新成功", c)
	//}
}

// FindCircleUser 用id查询CircleUser
// @Tags App_Circle
// @Summary 用id查询CircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询CircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/circle/findCircleUser [get]
func (circleApi *CircleApi) FindCircleUser(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleUser, err := hkCircleUserService.GetHkCircleUser(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircleUser": rehkCircleUser}, c)
	}
}

// GetCircleUserList 分页获取CircleUser列表
// @Tags App_Circle
// @Summary 分页获取CircleUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.CircleUserSearch true "分页获取CircleUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circle/getCircleUserList [get]
func (circleApi *CircleApi) GetCircleUserList(c *gin.Context) {
	var pageInfo appReq.CircleUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkCircleUserService.AppGetHkCircleUserInfoList(pageInfo); err != nil {
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

// CreateCircleRequest 创建HkCircleRequest
// @Tags App_Circle
// @Summary 创建HkCircleRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body appReq.CreateCircleRequestReq true "创建HkCircleRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circle/createCircleRequest [post]
func (circleApi *CircleApi) CreateCircleRequest(c *gin.Context) {
	var hkCircleRequest appReq.CreateCircleRequestReq
	err := c.ShouldBindJSON(&hkCircleRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := hkCircleRequestService.CreateHkCircleRequest(hkCircleRequest); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
}

// FindCircleRequest 用id查询CircleRequest
// @Tags App_Circle
// @Summary 用id查询CircleRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询CircleRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/circle/findCircleRequest [get]
func (circleApi *CircleApi) FindCircleRequest(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleRequest, err := hkCircleRequestService.GetHkCircleRequest(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehkCircleRequest": rehkCircleRequest}, c)
	}
}

// GetCircleRequestList 分页获取CircleRequest列表
// @Tags App_Circle
// @Summary 分页获取CircleRequest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.CircleRequestSearch true "分页获取CircleRequest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circle/getCircleRequestList [get]
func (circleApi *CircleApi) GetCircleRequestList(c *gin.Context) {
	var pageInfo appReq.CircleRequestSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := hkCircleRequestService.GetHkCircleRequestInfoList(pageInfo); err != nil {
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

// GetCircleClassifyList 分页获取CircleClassify列表
// @Tags App_Circle
// @Summary 分页获取CircleClassify列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.CircleClassifySearch true "分页获取CircleClassify列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circle/getCircleClassifyList [get]
func (circleApi *CircleApi) GetCircleClassifyList(c *gin.Context) {
	var pageInfo appReq.CircleClassifySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := hkCircleClassifyService.GetHkCircleClassifyInfoList(pageInfo); err != nil {
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

// GetCircleClassifyListAll 获取CircleClassify列表
// @Tags App_Circle
// @Summary 获取CircleClassify列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.CircleClassifySearch true "获取CircleClassify列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circle/getCircleClassifyListAll [get]
func (circleApi *CircleApi) GetCircleClassifyListAll(c *gin.Context) {
	var pageInfo appReq.CircleClassifySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := hkCircleClassifyService.GetHkCircleClassifyInfoList(pageInfo); err != nil {
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
