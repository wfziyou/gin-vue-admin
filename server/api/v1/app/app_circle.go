package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CircleApi struct {
}

// GetCircleForumPostsList 分页获取圈子的帖子列表
// @Tags 圈子
// @Summary 分页获取圈子的帖子列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.CircleForumPostsSearch true "分页获取圈子的帖子列表"
// @Success 200 {object}  response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/circle/getCircleForumPostsList [get]
func (circleApi *CircleApi) GetCircleForumPostsList(c *gin.Context) {
	var pageInfo communityReq.CircleForumPostsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pageInfo.CircleId == 0 {
		response.FailWithMessage("请给circleId赋值", c)
		return
	}
	if list, total, err := appForumPostsService.GetCircleForumPostsList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		var userId = utils.GetUserID(c)
		appForumThumbsUpService.GetUserForumThumbsUp(userId, list)
		appUserCollectService.GetUserIsCollect(userId, list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// DeleteCircleForumPosts 删除圈子的帖子
// @Tags 圈子
// @Summary 删除圈子的帖子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.DeleteCircleForumPostsReq true "删除圈子的帖子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/circle/deleteCircleForumPosts [delete]
func (circleApi *CircleApi) DeleteCircleForumPosts(c *gin.Context) {
	var req communityReq.DeleteCircleForumPostsReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

}

// GetUserCircleForumPostsList 用户加入圈子的所有动态列表
// @Tags 圈子
// @Summary 用户加入圈子的所有动态列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.UserCircleForumPostsSearch true "用户加入圈子的所有动态列表"
// @Success 200 {object}  response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/circle/getUserCircleForumPostsList [get]
func (circleApi *CircleApi) GetUserCircleForumPostsList(c *gin.Context) {
	var pageInfo communityReq.UserCircleForumPostsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = utils.GetUserID(c)
	if list, total, err := appForumPostsService.GetUserForumPostsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		var userId = utils.GetUserID(c)
		appForumThumbsUpService.GetUserForumThumbsUp(userId, list)
		appUserCollectService.GetUserIsCollect(userId, list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetSelfCircleList 分页获取用户加入的圈子列表
// @Tags 圈子
// @Summary 分页获取用户加入的圈子列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.SelfCircleSearch true "分页获取用户加入的圈子列表"
// @Success 200 {object}  response.PageResult{List=[]community.CircleBaseInfo,msg=string} "返回community.CircleBaseInfo"
// @Router /app/circle/getSelfCircleList [get]
func (circleApi *CircleApi) GetSelfCircleList(c *gin.Context) {
	var pageInfo communityReq.SelfCircleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.UserId = utils.GetUserID(c)
	if list, total, err := appCircleService.GetSelfCircleList(pageInfo); err != nil {
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
// @Tags 圈子
// @Summary 用id查询Circle
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询Circle"
// @Success 200  {object}  response.Response{data=community.Circle,msg=string}  "返回community.Circle"
// @Router /app/circle/findCircle [get]
func (circleApi *CircleApi) FindCircle(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var userId = utils.GetUserID(c)
	if rehkCircle, err := appCircleService.GetCircle(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		if _, num, err := appCircleUserService.GetUserHaveCircles(userId, []uint64{idSearch.ID}); err == nil && num > 0 {
			rehkCircle.HaveCircle = 1
		}
		response.OkWithData(rehkCircle, c)
	}
}

// GetCircleList 分页获取圈子列表
// @Tags 圈子
// @Summary 分页获取圈子列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.CircleSearch true "分页获取圈子列表"
// @Success 200 {object}  response.PageResult{List=[]community.CircleBaseInfo,msg=string} "返回community.CircleBaseInfo"
// @Router /app/circle/getCircleList [get]
func (circleApi *CircleApi) GetCircleList(c *gin.Context) {
	var pageInfo communityReq.CircleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appCircleService.GetCircleInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		var userId = utils.GetUserID(c)
		appCircleUserService.GetUserHaveCircle(userId, list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// UpdateCircle (圈子管理者)更新圈子
// @Tags 圈子
// @Summary (圈子管理者)更新圈子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.UpdateCircleReq true "(圈子管理者)更新圈子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/circle/updateCircle [put]
func (circleApi *CircleApi) UpdateCircle(c *gin.Context) {
	var req communityReq.UpdateCircleReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appCircleService.UpdateCircle(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// SetUserCurCircle 设置用户当前圈子
// @Tags 圈子
// @Summary 设置用户当前圈子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.SetUserCurCircleReq true "设置用户当前圈子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /app/circle/setUserCurCircle [post]
func (circleApi *CircleApi) SetUserCurCircle(c *gin.Context) {
	var req communityReq.SetUserCurCircleReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if _, err := appCircleService.GetCircle(req.CircleId); err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}

	var userId = utils.GetUserID(c)
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

	if err = appCircleUserService.SetUserCurCircle(userId, req.CircleId); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// EnterCircle 加入圈子
// @Tags 圈子
// @Summary 加入圈子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.EnterCircleReq true "加入圈子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /app/circle/enterCircle [post]
func (circleApi *CircleApi) EnterCircle(c *gin.Context) {
	var req communityReq.EnterCircleReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	circleInfo, err := appCircleService.GetCircle(req.CircleId)
	if err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}

	var userId = utils.GetUserID(c)
	if data, _, err := appCircleUserService.GetCircleUserInfoList(communityReq.CircleUserSearch{
		CircleId: req.CircleId,
		UserId:   userId,
		PageInfo: request.PageInfo{Page: 1, PageSize: 2},
	}); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else if len(data) > 0 {
		response.FailWithMessage("用户已在圈子中", c)
		return
	}
	//圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）
	if circleInfo.Property == 0 {
		appCircleUserService.UpdateCircleUser(community.CircleUser{
			UserId:   userId,
			CircleId: req.CircleId,
			Sort:     1,
		})
		response.OkWithMessage("加入圈子成功", c)
		return
	} else {
		if data, _, err := appCircleAddRequestService.GetCircleAddRequestInfoList(communityReq.CircleAddRequestSearch{
			CircleId: req.CircleId,
			UserId:   userId,
			PageInfo: request.PageInfo{Page: 1, PageSize: 2},
		}); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else if len(data) > 0 {
			appCircleAddRequestService.UpdateCircleAddRequest(data[0])
			response.OkWithMessage("申请成功", c)
			return
		} else {
			appCircleAddRequestService.UpdateCircleAddRequest(community.CircleAddRequest{
				CircleId: req.CircleId,
				UserId:   userId,
			})
			response.OkWithMessage("申请成功", c)
			return
		}
	}
}

// ExitCircle 退出圈子
// @Tags 圈子
// @Summary 退出圈子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ExitCircleReq true "退出圈子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /app/circle/exitCircle [post]
func (circleApi *CircleApi) ExitCircle(c *gin.Context) {
	var req communityReq.ExitCircleReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	_, err = appCircleService.GetCircle(req.CircleId)
	if err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}

	var userId = utils.GetUserID(c)
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
	} else {
		//如果是圈主
		if data[0].Power == 1 {
			response.FailWithMessage("圈主不能退出", c)
			return
		} else {
			err = appCircleUserService.DeleteCircleUserInfo(data[0])
			if err == nil {
				response.OkWithMessage("退出成功", c)
				return
			}
			response.FailWithMessage("退出失败", c)
		}
	}
}

// ApplyEnterCircle 申请加入圈子
// @Tags 圈子
// @Summary 申请加入圈子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ApplyEnterCircleReq true "申请加入圈子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /app/circle/applyEnterCircle [post]
func (circleApi *CircleApi) ApplyEnterCircle(c *gin.Context) {
	var req communityReq.ApplyEnterCircleReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	circleInfo, err := appCircleService.GetCircle(req.CircleId)
	if err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}

	var userId = utils.GetUserID(c)
	if data, _, err := appCircleUserService.GetCircleUserInfoList(communityReq.CircleUserSearch{
		CircleId: req.CircleId,
		UserId:   userId,
		PageInfo: request.PageInfo{Page: 1, PageSize: 2},
	}); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else if len(data) > 0 {
		response.FailWithMessage("用户已在圈子中", c)
		return
	}
	//圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）
	if circleInfo.Property == 0 {
		appCircleUserService.UpdateCircleUser(community.CircleUser{
			UserId:   userId,
			CircleId: req.CircleId,
			Sort:     1,
		})
		response.OkWithMessage("加入圈子成功", c)
		return
	} else {
		if data, _, err := appCircleAddRequestService.GetCircleAddRequestInfoList(communityReq.CircleAddRequestSearch{
			CircleId: req.CircleId,
			UserId:   userId,
			PageInfo: request.PageInfo{Page: 1, PageSize: 2},
		}); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else if len(data) > 0 {
			data[0].Reason = req.Reason
			appCircleAddRequestService.UpdateCircleAddRequest(data[0])
			response.OkWithMessage("申请成功", c)
			return
		} else {
			appCircleAddRequestService.UpdateCircleAddRequest(community.CircleAddRequest{
				CircleId: req.CircleId,
				UserId:   userId,
				Reason:   req.Reason,
			})
			response.OkWithMessage("申请成功", c)
			return
		}
	}
}

// EnterCircleApplyList 分页获取加入圈子申请
// @Tags 圈子
// @Summary 分页获取加入圈子申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.CircleAddRequestSearch true "分页获取加入圈子申请"
// @Success 200 {object}  response.PageResult{List=[]community.CircleAddRequest,msg=string} "返回community.CircleUserInfo"
// @Router /app/circle/enterCircleApplyList [get]
func (circleApi *CircleApi) EnterCircleApplyList(c *gin.Context) {
	var pageInfo communityReq.CircleAddRequestSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if list, total, err := appCircleAddRequestService.GetCircleAddRequestInfoList(pageInfo); err != nil {
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

// ApproveEnterCircleRequest 审批加入圈子申请
// @Tags 圈子
// @Summary 审批加入圈子申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ApproveEnterCircleReq true "审批加入圈子申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /app/circle/approveEnterCircleRequest [post]
func (circleApi *CircleApi) ApproveEnterCircleRequest(c *gin.Context) {
	var req communityReq.ApproveEnterCircleReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//if _, err := appCircleService.GetCircle(req.CircleId); err != nil {
	//	response.FailWithMessage("圈子不存在", c)
	//	return
	//}
	//
	//var userId = utils.GetUserID(c)
	//if data, _, err := appCircleUserService.GetCircleUserInfoList(communityReq.CircleUserSearch{
	//	CircleId: req.CircleId,
	//	UserId:   userId,
	//	PageInfo: request.PageInfo{Page: 1, PageSize: 2},
	//}); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//} else if len(data) == 0 {
	//	response.FailWithMessage("用户不在圈子中", c)
	//	return
	//}
	//
	//if err = appCircleUserService.SetUserCurCircle(userId, req.CircleId); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//response.OkWithMessage("设置成功", c)
}

/*************************************
圈子成员
**************************************/

// DeleteCircleUser 删除圈子用户
// @Tags 圈子
// @Summary 删除圈子用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.DeleteCircleUserReq true "删除圈子用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/circle/deleteCircleUser [delete]
func (circleApi *CircleApi) DeleteCircleUser(c *gin.Context) {
	var req communityReq.DeleteCircleUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var pageInfo communityReq.CircleUserSearch
	pageInfo.Page = 1
	pageInfo.PageSize = 10
	pageInfo.CircleId = req.CircleId
	pageInfo.UserId = req.UserId
	if list, total, err := appCircleUserService.GetCircleUserInfoList(pageInfo); err == nil {
		if total == 0 {
			response.FailWithMessage("圈子成员不存在", c)
			return
		}

		if err := appCircleUserService.DeleteCircleUserInfo(list[0]); err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithMessage("删除成功", c)
		}
	} else {
		response.FailWithMessage("圈子成员不存在", c)
	}
}

// DeleteCircleUsers 批量删除圈子用户
// @Tags 圈子
// @Summary 批量删除圈子用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除圈子用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/circle/deleteCircleUsers [delete]
func (circleApi *CircleApi) DeleteCircleUsers(c *gin.Context) {
	var req communityReq.DeleteCircleUsersReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appCircleUserService.DeleteCircleUsers(req.CircleId, req.UserIds); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCircleUser 更新圈子用户
// @Tags 圈子
// @Summary 更新圈子用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.UpdateCircleUserReq true "更新圈子用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/circle/updateCircleUser [put]
func (circleApi *CircleApi) UpdateCircleUser(c *gin.Context) {
	var hkCircleUser communityReq.UpdateCircleUserReq
	err := c.ShouldBindJSON(&hkCircleUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := appCircleUserService.UpdateCircleUser(hkCircleUser); err != nil {
	//	global.GVA_LOG.Error("更新失败!", zap.Error(err))
	//	response.FailWithMessage("更新失败", c)
	//} else {
	//	response.OkWithMessage("更新成功", c)
	//}
}

// FindCircleUser 用id查询圈子用户信息
// @Tags 圈子
// @Summary 用id查询圈子用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询圈子用户信息"
// @Success 200  {object}  response.Response{data=community.CircleUserInfo,msg=string}  "返回community.CircleUserInfo"
// @Router /app/circle/findCircleUser [get]
func (circleApi *CircleApi) FindCircleUser(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleUser, err := appCircleUserService.GetCircleUserInfo(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rehkCircleUser, c)
	}
}

// GetCircleUserList 分页获取圈子的用户信息列表
// @Tags 圈子
// @Summary 分页获取圈子的用户信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.CircleUserSearch true "分页获取圈子的用户信息列表"
// @Success 200 {object}  response.PageResult{List=[]community.CircleUserInfo,msg=string} "返回community.CircleUserInfo"
// @Router /app/circle/getCircleUserList [get]
func (circleApi *CircleApi) GetCircleUserList(c *gin.Context) {
	var pageInfo communityReq.CircleUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appCircleUserService.GetCircleUserInfoList(pageInfo); err != nil {
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

// CreateCircleRequest 创建圈子请求
// @Tags 圈子
// @Summary 创建圈子请求
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.CreateCircleRequestReq true "创建圈子请求"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/circle/createCircleRequest [post]
func (circleApi *CircleApi) CreateCircleRequest(c *gin.Context) {
	var hkCircleRequest communityReq.CreateCircleRequestReq
	err := c.ShouldBindJSON(&hkCircleRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if _, err := appCircleClassifyService.GetCircleClassify(hkCircleRequest.CircleClassifyId); err != nil {
		response.FailWithMessage("没有找到圈子分类", c)
		return
	}

	//hkCircleRequest community.CircleRequest
	if err := appCircleRequestService.CreateCircleRequest(community.CircleRequest{
		Type:             1,
		Name:             hkCircleRequest.Name,
		Logo:             hkCircleRequest.Logo,
		CircleClassifyId: hkCircleRequest.CircleClassifyId,
		Slogan:           hkCircleRequest.Slogan,
		Des:              hkCircleRequest.Des,
		Protocol:         hkCircleRequest.Protocol,
		BackImage:        hkCircleRequest.BackImage,
	}); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// FindCircleRequest 用id查询创建圈子请求
// @Tags 圈子
// @Summary 用id查询创建圈子请求
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询创建圈子请求"
// @Success 200 {object}  response.Response{data=community.CircleRequest,msg=string}  "返回community.CircleUser"
// @Router /app/circle/findCircleRequest [get]
func (circleApi *CircleApi) FindCircleRequest(c *gin.Context) {
	var idSearch request.IdSearch
	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehkCircleRequest, err := appCircleRequestService.GetCircleRequest(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rehkCircleRequest, c)
	}
}

// GetCircleRequestList 分页获取创建圈子请求列表
// @Tags 圈子
// @Summary 分页获取创建圈子请求列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.CircleRequestSearch true "分页获取创建圈子请求列表"
// @Success 200 {object}  response.PageResult{List=[]community.CircleRequest,msg=string} "返回community.CircleRequest"
// @Router /app/circle/getCircleRequestList [get]
func (circleApi *CircleApi) GetCircleRequestList(c *gin.Context) {
	var pageInfo communityReq.CircleRequestSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appCircleRequestService.GetCircleRequestInfoList(pageInfo); err != nil {
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

// GetCircleClassifyList 分页获取圈子分类列表
// @Tags 圈子
// @Summary 分页获取圈子分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.CircleClassifySearch true "分页获取圈子分类列表"
// @Success 200 {object}  response.PageResult{List=[]community.CircleClassify,msg=string} "返回community.CircleClassify"
// @Router /app/circle/getCircleClassifyList [get]
func (circleApi *CircleApi) GetCircleClassifyList(c *gin.Context) {
	var pageInfo communityReq.CircleClassifySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appCircleClassifyService.GetCircleClassifyInfoList(pageInfo); err != nil {
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

// GetCircleClassifyListAll 获取圈子分类列表
// @Tags 圈子
// @Summary 获取圈子分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.CircleClassifySearch true "获取圈子分类列表"
// @Success 200 {object} response.Response{data=[]community.CircleClassify,msg=string} "返回community.CircleClassify"
// @Router /app/circle/getCircleClassifyListAll [get]
func (circleApi *CircleApi) GetCircleClassifyListAll(c *gin.Context) {
	var pageInfo communityReq.CircleClassifySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, _, err := appCircleClassifyService.GetCircleClassifyInfoListAll(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// SetCircleChannel 设置圈子频道
// @Tags 频道
// @Summary 设置圈子频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetCircleChannel true "设置圈子频道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /app/circle/setCircleChannel [post]
func (circleApi *CircleApi) SetCircleChannel(c *gin.Context) {
	var req communityReq.ParamSetCircleChannel
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//userId := utils.GetUserID(c)
	//if err := appUserService.UpdateUserChannel(userId, req.ChannelIds); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
}

// GetCircleChannelList 获取圈子频道
// @Tags 频道
// @Summary 获取圈子频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object}  response.Response{data=[]community.ChannelInfo,msg=string} "返回[]community.ChannelInfo"
// @Router /app/circle/getCircleChannelList [get]
func (circleApi *CircleApi) GetCircleChannelList(c *gin.Context) {
	//userId := utils.GetUserID(c)
	//channelIds, err := appUserService.GetUserChannel(userId)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//
	//if list, _, err := hkChannelService.GetChannelInfoListById(channelIds); err != nil {
	//	tmp := utils.SplitToUint64List(channelIds, ",")
	//	for index, id := range tmp {
	//		for _, obj := range list {
	//			if obj.ID == id {
	//				obj.Sort = index
	//				break
	//			}
	//		}
	//	}
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("获取失败", c)
	//} else {
	//	response.OkWithDetailed(list, "获取成功", c)
	//}
}

// GetChildCircleList 分页获取子圈子列表
// @Tags 圈子
// @Summary 分页获取子圈子列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ChildCircleSearch true "分页获取子圈子列表"
// @Success 200 {object}  response.PageResult{List=[]community.CircleBaseInfo,msg=string} "返回[]community.CircleBaseInfo"
// @Router /app/circle/getChildCircleList [get]
func (circleApi *CircleApi) GetChildCircleList(c *gin.Context) {
	var pageInfo communityReq.ChildCircleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
}
