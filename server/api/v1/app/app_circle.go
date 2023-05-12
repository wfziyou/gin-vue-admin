package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	userId := utils.GetUserID(c)
	if list, total, err := appForumPostsService.GetCircleForumPostsList(userId, pageInfo); err != nil {
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

// DeleteCircleForumPosts (圈子管理员)删除圈子的帖子
// @Tags 圈子
// @Summary (圈子管理员)删除圈子的帖子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.DeleteCircleForumPostsReq true "(圈子管理员)删除圈子的帖子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/circle/deleteCircleForumPosts [delete]
func (circleApi *CircleApi) DeleteCircleForumPosts(c *gin.Context) {
	var req communityReq.DeleteCircleForumPostsReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	_, err = appCircleService.GetCircle(req.CircleId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("圈子不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)
	circleUser, err := appCircleUserService.GetCircleUserEx(req.CircleId, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("用户不在圈子中", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if circleUser.Power == community.CircleUserPowerGeneral {
		response.FailWithMessage("普通用户没有权限删除", c)
		return
	}

	err = appForumPostsService.DeleteCircleForumPostsByIds(req.CircleId, req.Ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("成功", c)
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

	userId := utils.GetUserID(c)
	if list, total, err := appForumPostsService.GetUserForumPostsInfoList(userId, pageInfo); err != nil {
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
	userId := utils.GetUserID(c)
	if list, total, err := appCircleService.GetSelfCircleList(userId, pageInfo); err != nil {
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
	if circle, err := appCircleService.GetCircle(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		if _, num, err := appCircleUserService.GetUserHaveCircles(userId, []uint64{idSearch.ID}); err == nil && num > 0 {
			circle.HaveCircle = 1
		}

		response.OkWithData(circle, c)
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
	userId := utils.GetUserID(c)
	circle, err := appCircleService.GetCircle(req.ID)
	if err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}
	if circle.ID != userId {
		response.FailWithMessage("用户没有权限修改圈子", c)
		return
	}
	if err := appCircleService.UpdateCircle(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
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

	userId := utils.GetUserID(c)
	if data, _, err := appCircleUserService.GetCircleUserInfoList(userId, communityReq.CircleUserSearch{
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
		appCircleUserService.CreateCircleUser(community.CircleUser{
			UserId:     userId,
			CircleId:   req.CircleId,
			CircleName: circleInfo.Name,
			Sort:       1,
		})
		response.OkWithMessage("加入圈子成功", c)
		return
	} else {
		if _, err := appCircleAddRequestService.GetCircleAddRequestByUserId(req.CircleId, userId); errors.Is(err, gorm.ErrRecordNotFound) {
			appCircleAddRequestService.UpdateCircleAddRequest(community.CircleAddRequest{
				CircleId: req.CircleId,
				UserId:   userId,
			})
			response.OkWithMessage("申请成功", c)
			return
		} else if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
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

	userId := utils.GetUserID(c)
	if data, _, err := appCircleUserService.GetCircleUserInfoList(userId, communityReq.CircleUserSearch{
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

	userId := utils.GetUserID(c)
	if data, _, err := appCircleUserService.GetCircleUserInfoList(userId, communityReq.CircleUserSearch{
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
		appCircleUserService.CreateCircleUser(community.CircleUser{
			UserId:     userId,
			CircleId:   req.CircleId,
			CircleName: circleInfo.Name,
			Sort:       1,
		})
		response.OkWithMessage("加入圈子成功", c)
		return
	} else {
		if data, err := appCircleAddRequestService.GetCircleAddRequestByUserId(req.CircleId, userId); errors.Is(err, gorm.ErrRecordNotFound) {
			appCircleAddRequestService.UpdateCircleAddRequest(community.CircleAddRequest{
				CircleId: req.CircleId,
				UserId:   userId,
			})
			response.OkWithMessage("申请成功", c)
			return
		} else if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			data.Reason = req.Reason
			appCircleAddRequestService.UpdateCircleAddRequest(data)
			response.OkWithMessage("申请成功", c)
			return
		}
	}
}

// EnterCircleApplyList (圈子管理员)分页获取加入圈子申请
// @Tags 圈子
// @Summary (圈子管理员)分页获取加入圈子申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.CircleAddRequestSearch true "(圈子管理员)分页获取加入圈子申请"
// @Success 200 {object}  response.PageResult{List=[]community.CircleAddRequest,msg=string} "返回community.CircleUserInfo"
// @Router /app/circle/enterCircleApplyList [get]
func (circleApi *CircleApi) EnterCircleApplyList(c *gin.Context) {
	var req communityReq.CircleAddRequestSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	_, err = appCircleService.GetCircle(req.CircleId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("圈子不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)
	circleUser, err := appCircleUserService.GetCircleUserEx(req.CircleId, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("用户不在圈子中", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if circleUser.Power == community.CircleUserPowerGeneral {
		response.FailWithMessage("没有权限获取", c)
		return
	}

	if list, total, err := appCircleAddRequestService.GetCircleAddRequestInfoList(req); err != nil {
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

// ApproveEnterCircleRequest (圈子管理员)审批加入圈子申请
// @Tags 圈子
// @Summary (圈子管理员)审批加入圈子申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ApproveEnterCircleReq true "(圈子管理员)审批加入圈子申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /app/circle/approveEnterCircleRequest [post]
func (circleApi *CircleApi) ApproveEnterCircleRequest(c *gin.Context) {
	var req communityReq.ApproveEnterCircleReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	_, err = appCircleService.GetCircle(req.CircleId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("圈子不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)
	circleUser, err := appCircleUserService.GetCircleUserEx(req.CircleId, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("用户不在圈子中", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if circleUser.Power == community.CircleUserPowerGeneral {
		response.FailWithMessage("没有权限获取", c)
		return
	}

	if _, err := appCircleService.GetCircle(req.CircleId); err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}

	if err = appCircleAddRequestService.ApproveEnterCircleRequest(userId, req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

/*************************************
圈子成员
**************************************/

// DeleteCircleUser (圈子管理者)删除圈子用户
// @Tags 圈子
// @Summary (圈子管理者)删除圈子用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.DeleteCircleUserReq true "(圈子管理者)删除圈子用户"
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
	userId := utils.GetUserID(c)
	if list, total, err := appCircleUserService.GetCircleUserInfoList(userId, pageInfo); err == nil {
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

// DeleteCircleUsers (圈子管理者)批量删除圈子用户
// @Tags 圈子
// @Summary (圈子管理者)批量删除圈子用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "(圈子管理者)批量删除圈子用户"
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
	var req communityReq.UpdateCircleUserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := appCircleUserService.UpdateCircleUser(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCircleUser 查询圈子用户信息
// @Tags 圈子
// @Summary 查询圈子用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.FindCircleUserReq true "查询圈子用户信息"
// @Success 200  {object}  response.Response{data=community.CircleUserInfo,msg=string}  "返回community.CircleUserInfo"
// @Router /app/circle/findCircleUser [get]
func (circleApi *CircleApi) FindCircleUser(c *gin.Context) {
	var req communityReq.FindCircleUserReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if rehkCircleUser, err := appCircleUserService.GetCircleUserInfo(userId, req.CircleId, req.UserId); err != nil {
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
	userId := utils.GetUserID(c)
	if list, total, err := appCircleUserService.GetCircleUserInfoList(userId, pageInfo); err != nil {
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
	var req communityReq.CreateCircleRequestReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user := utils.GetUserInfo(c)
	if _, err := appCircleClassifyService.GetCircleClassify(req.CircleClassifyId); err != nil {
		response.FailWithMessage("没有找到圈子分类", c)
		return
	}

	obj, err := systemConfigService.GetParam()
	if obj.CreateCircleCheck == true {
		if err := appCircleRequestService.CreateCircleRequest(community.CircleRequest{
			UserId:           user.GetUserId(),
			UserNickName:     user.NickName,
			Type:             community.CircleTypeUser,
			Name:             req.Name,
			Logo:             req.Logo,
			CircleClassifyId: req.CircleClassifyId,
			Slogan:           req.Slogan,
			Des:              req.Des,
			Protocol:         req.Protocol,
			CoverImage:       req.CoverImage,
		}); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithMessage("创建成功", c)
		}
		return
	} else {
		appCircleService.CreateCircle(community.Circle{
			Type:             community.CircleTypeUser,
			Name:             req.Name,
			Logo:             req.Logo,
			CircleClassifyId: req.CircleClassifyId,
			Slogan:           req.Slogan,
			Des:              req.Des,
			Protocol:         req.Protocol,
			CoverImage:       req.CoverImage,
		})
	}
}

// DestroyCircle (圈子管理者)注销圈子
// @Tags 圈子
// @Summary (圈子管理者)注销圈子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamDestroyCircle true "(圈子管理者)注销圈子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /app/circle/destroyCircle [post]
func (circleApi *CircleApi) DestroyCircle(c *gin.Context) {
	var req communityReq.ParamDestroyCircle
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	obj, err := appCircleService.GetCircle(req.CircleId)
	if err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}

	//need to do 删除圈子关联数据
	if err := appCircleService.DeleteCircle(obj); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		appCircleUserService.DeleteCircleAllUser(req.CircleId)
		response.OkWithMessage("成功", c)
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
// @Param data query communityReq.CircleClassifyAllSearch true "获取圈子分类列表"
// @Success 200 {object} response.Response{data=[]community.CircleClassify,msg=string} "返回community.CircleClassify"
// @Router /app/circle/getCircleClassifyListAll [get]
func (circleApi *CircleApi) GetCircleClassifyListAll(c *gin.Context) {
	var pageInfo communityReq.CircleClassifyAllSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := appCircleClassifyService.GetCircleClassifyInfoListAll(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// SetCircleChannel (圈子管理者)设置圈子频道
// @Tags 圈子
// @Summary (圈子管理者)设置圈子频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetCircleChannel true "(圈子管理者)设置圈子频道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /app/circle/setCircleChannel [post]
func (circleApi *CircleApi) SetCircleChannel(c *gin.Context) {
	var req communityReq.ParamSetCircleChannel
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	circle, err := appCircleService.GetCircle(req.CircleId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("圈子不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := appCircleService.UpdateCircleChannel(circle.ID, req.ChannelIds); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// AddCircleChannel (圈子管理者)添加圈子频道
// @Tags 圈子
// @Summary (圈子管理者)添加圈子频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamAddCircleChannel true "(圈子管理者)添加圈子频道"
// @Success 200  {object}  response.Response{data=community.CircleChannel,msg=string}  "返回community.CircleChannel"
// @Router /app/circle/addCircleChannel [post]
func (circleApi *CircleApi) AddCircleChannel(c *gin.Context) {
	var req communityReq.ParamAddCircleChannel
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	circle, err := appCircleService.GetCircle(req.CircleId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("圈子不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if obj, err := appCircleService.AddCircleChannel(circle.ID, req.Name); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(community.CircleChannelInfo{
			ID:   obj.ID,
			Name: obj.Name,
		}, "获取成功", c)
	}
}

// DeleteCircleChannel (圈子管理者)删除圈子频道
// @Tags 圈子
// @Summary (圈子管理者)删除圈子频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamDeleteCircleChannel true "(圈子管理者)删除圈子频道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /app/circle/deleteCircleChannel [delete]
func (circleApi *CircleApi) DeleteCircleChannel(c *gin.Context) {
	var req communityReq.ParamDeleteCircleChannel
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if _, err := appCircleService.GetCircle(req.CircleId); err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}

	if err := appCircleChannelService.DeleteCircleChannelByIds(req.CircleId, req.Ids); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

// GetCircleChannelList 获取圈子频道
// @Tags 圈子
// @Summary 获取圈子频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ParamGetCircleChannel true "获取圈子频道"
// @Success 200 {object}  response.Response{data=[]community.ChannelInfo,msg=string} "返回[]community.ChannelInfo"
// @Router /app/circle/getCircleChannelList [get]
func (circleApi *CircleApi) GetCircleChannelList(c *gin.Context) {
	var req communityReq.ParamGetCircleChannel
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	circle, err := appCircleService.GetCircle(req.CircleId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("圈子不存在", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(circle.ChannelId) == 0 {
		list, err := appCircleChannelService.GetDefaultChannelInfoList(circle.ID)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithDetailed(list, "获取成功", c)
		return
	}
	if list, err := appCircleChannelService.GetChannelInfoListById(circle.ChannelId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	} else if len(list) > 0 {
		newList := make([]community.CircleChannelInfo, 0, len(list))
		tmp := utils.SplitToUint64List(circle.ChannelId, ",")
		for _, id := range tmp {
			for _, obj := range list {
				if obj.ID == id {
					newList = append(newList, obj)
					break
				}
			}
		}
		response.OkWithDetailed(newList, "获取成功", c)
	} else {
		list, err := appCircleChannelService.GetDefaultChannelInfoList(circle.ID)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithDetailed(list, "获取成功", c)
		return
	}
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
	var req communityReq.ChildCircleSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appCircleRelationService.GetChildCircleList(req.CircleId, req.PageInfo); err != nil {
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

// AddCircleTag (圈子管理者)添加圈子标签
// @Tags 圈子
// @Summary (圈子管理者)添加圈子标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamCreateCircleTag true "(圈子管理者)添加圈子标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /app/circle/addCircleTag [post]
func (circleApi *CircleApi) AddCircleTag(c *gin.Context) {
	var req communityReq.ParamCreateCircleTag
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if _, err := appCircleService.GetCircle(req.CircleId); err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}

	if _, err := appCircleTagService.AddCircleTag(req.CircleId, req.Name); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCircleTags (圈子管理者)删除圈子标签
// @Tags 圈子
// @Summary (圈子管理者)删除圈子标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamDeleteCircleTags true "(圈子管理者)删除圈子标签"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /app/circle/deleteCircleTags [delete]
func (circleApi *CircleApi) DeleteCircleTags(c *gin.Context) {
	var req communityReq.ParamDeleteCircleTags
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if _, err := appCircleService.GetCircle(req.CircleId); err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}

	if err := appCircleTagService.DeleteCircleTags(req.CircleId, req.Names); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("成功", c)
	}
}

// GetCircleTagList (圈子管理者)获取圈子的标签
// @Tags 圈子
// @Summary (圈子管理者)获取圈子的标签
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ParamGetCircleTags true "(圈子管理者)获取圈子的标签"
// @Success 200 {object}  response.PageResult{List=[]community.CircleTag,msg=string} "返回[]community.CircleTag"
// @Router /app/circle/getCircleTagList [get]
func (circleApi *CircleApi) GetCircleTagList(c *gin.Context) {
	var req communityReq.ParamGetCircleTags
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := appCircleTagService.GetCircleTagList(req.CircleId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// RequestBecomeChildCircle (圈子管理者)申请成为子圈子
// @Tags 圈子
// @Summary (圈子管理者)申请成为子圈子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamRequestBecomeChildCircle true "(圈子管理者)申请成为子圈子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /app/circle/requestBecomeChildCircle [post]
func (circleApi *CircleApi) RequestBecomeChildCircle(c *gin.Context) {
	var req communityReq.ParamRequestBecomeChildCircle
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if _, err := appCircleService.GetCircle(req.CircleId); err != nil {
		response.FailWithMessage("圈子不存在", c)
		return
	}
	if _, err := appCircleService.GetCircle(req.ParentCircleId); err != nil {
		response.FailWithMessage("父圈子不存在", c)
		return
	}

	//
	//if err := appCircleTagService.CreateCircleTag(req.CircleId, req.Name); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage(err.Error(), c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
}

// ApproveChildCircleRequest (圈子管理者)审批子圈子申请
// @Tags 圈子
// @Summary (圈子管理者)审批子圈子申请
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamApproveChildCircleRequest true "(圈子管理者)审批子圈子申请"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /app/circle/approveChildCircleRequest [post]
func (circleApi *CircleApi) ApproveChildCircleRequest(c *gin.Context) {
	var req communityReq.ParamApproveChildCircleRequest
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
	//if err := appCircleTagService.CreateCircleTag(req.CircleId, req.Name); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage(err.Error(), c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
}

// CreateChildCircle (圈子管理者)创建子圈子
// @Tags 圈子
// @Summary (圈子管理者)创建子圈子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamCreateChildCircle true "(圈子管理者)创建子圈子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /app/circle/createChildCircle [post]
func (circleApi *CircleApi) CreateChildCircle(c *gin.Context) {
	var req communityReq.ParamCreateChildCircle
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
	//if err := appCircleTagService.CreateCircleTag(req.CircleId, req.Name); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage(err.Error(), c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
}

// DeleteChildCircle (圈子管理者)删除子圈子
// @Tags 圈子
// @Summary (圈子管理者)删除子圈子
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamDeleteChildCircle true "(圈子管理者)删除子圈子"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /app/circle/deleteChildCircle [delete]
func (circleApi *CircleApi) DeleteChildCircle(c *gin.Context) {
	var req communityReq.ParamDeleteChildCircle
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
	//if err := appCircleTagService.CreateCircleTag(req.CircleId, req.Name); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage(err.Error(), c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
}

// GetChildCircleRequestList (圈子管理者)获取子圈子申请列表
// @Tags 圈子
// @Summary (圈子管理者)获取子圈子申请列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.ChildCircleRequestSearch true "(圈子管理者)获取子圈子申请列表"
// @Success 200 {object}  response.PageResult{List=[]community.CircleTag,msg=string} "返回[]community.CircleTag"
// @Router /app/circle/getChildCircleRequestList [get]
func (circleApi *CircleApi) GetChildCircleRequestList(c *gin.Context) {
	var req communityReq.ChildCircleRequestSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//var pageInfo communityReq.CircleRelationRequestSearch
	//err := c.ShouldBindQuery(&pageInfo)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//if list, total, err := hkCircleRelationRequestService.GetCircleRelationRequestInfoList(pageInfo); err != nil {
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
