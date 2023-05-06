package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type QuestionApi struct {
}

var hkQuestionService = service.ServiceGroupApp.AppServiceGroup.Community.QuestionService

// GetGlobalRecommendQuestionList 分页获取全局推荐问题列表
// @Tags 问答
// @Summary 分页获取全局推荐问题列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.GlobalRecommendQuestionSearch true "分页获取全局推荐问题列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/question/getGlobalRecommendQuestionList [get]
func (questionApi *QuestionApi) GetGlobalRecommendQuestionList(c *gin.Context) {
	var req communityReq.GlobalRecommendQuestionSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := appForumPostsService.GetGlobalRecommendQuestionList(userId, req.CurPos, req.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		var userId = utils.GetUserID(c)
		appForumThumbsUpService.GetUserForumThumbsUp(userId, list)
		appUserCollectService.GetUserIsCollect(userId, list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// GetCircleQuestionList 分页获取圈子问题列表
// @Tags 问答
// @Summary 分页获取圈子问题列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.CircleQuestionSearch true "分页获取圈子问题列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPostsBaseInfo,msg=string} "返回community.ForumPostsBaseInfo"
// @Router /app/question/getCircleQuestionList [get]
func (questionApi *QuestionApi) GetCircleQuestionList(c *gin.Context) {
	var req communityReq.CircleQuestionSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	circle, err := appCircleService.GetCircle(req.CircleId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage("圈子不存在", c)
	}
	userId := utils.GetUserID(c)
	if list, total, err := hkQuestionService.GetCircleQuestionList(userId, circle.ID, req.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		var userId = utils.GetUserID(c)
		appForumThumbsUpService.GetUserForumThumbsUp(userId, list)
		appUserCollectService.GetUserIsCollect(userId, list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// CreateQuestion 创建问题
// @Tags 问答
// @Summary 创建问题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.CreateQuestion true "创建问题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/question/createQuestion [post]
func (questionApi *QuestionApi) CreateQuestion(c *gin.Context) {
	var req communityReq.CreateQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	user, err := appUserService.GetUser(userId)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	if user.UserExtend.CurrencyGold < req.PayNum {
		response.FailWithMessage("金币不够", c)
		return
	}

	err = appUserService.DecreaseGold(userId, req.PayNum, "发布问题", "", "")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = hkQuestionService.CreateQuestion(userId, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	appUserService.UpdatePostsTime(userId)
	response.OkWithMessage("成功", c)
}

// CloseQuestion 关闭问题
// @Tags 问答
// @Summary 关闭问题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.CloseQuestion true "关闭问题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /app/question/closeQuestion [post]
func (questionApi *QuestionApi) CloseQuestion(c *gin.Context) {
	var req communityReq.CloseQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	err = hkQuestionService.CloseQuestion(userId, req.Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithMessage("成功", c)
}

// GetAnswerList 分页获取问题的回答列表
// @Tags 问答
// @Summary 分页获取问题的回答列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.QuestionAnswerSearch true "分页获取问题的回答列表"
// @Success 200 {object} response.PageResult{List=[]community.QuestionAnswer,msg=string} "返回community.QuestionAnswer"
// @Router /app/question/getAnswerList [get]
func (questionApi *QuestionApi) GetAnswerList(c *gin.Context) {
	var req communityReq.QuestionAnswerSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if list, total, err := hkQuestionService.GetAnswerList(userId, req.QuestionId, req.PageInfo); err != nil {
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

// AnswerQuestion 回答问题
// @Tags 问答
// @Summary 回答问题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.AnswerQuestion true "回答问题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/question/answerQuestion [post]
func (questionApi *QuestionApi) AnswerQuestion(c *gin.Context) {
	var req communityReq.AnswerQuestion
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	err = hkQuestionService.AnswerQuestion(userId, req.QuestionId, req.Content)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("成功", c)
}

// DelSelfAnswer 删除自己的回答
// @Tags 问答
// @Summary 删除自己的回答
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdDelete true "删除自己的回答"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/question/delSelfAnswer [delete]
func (questionApi *QuestionApi) DelSelfAnswer(c *gin.Context) {
	var req request.IdDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	err = hkQuestionService.DelSelfAnswer(userId, req.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("成功", c)
}

// SetBestAnswer 设置最佳答案
// @Tags 问答
// @Summary 设置最佳答案
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.SetBestAnswer true "设置最佳答案"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/question/setBestAnswer [post]
func (questionApi *QuestionApi) SetBestAnswer(c *gin.Context) {
	var req communityReq.SetBestAnswer
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)
	user, err := appUserService.GetUser(userId)
	question, err := hkQuestionService.GetQuestion(req.QuestionId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if question.Status == utils.QuestionStatusClose {
		response.FailWithMessage("问题已关闭", c)
		return
	}
	if question.UserId != userId {
		response.FailWithMessage("只能给自己问题设置答案", c)
		return
	}
	answer, err := hkQuestionService.SetBestAnswer(question, req.AnswerId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else if question.PayNum > 0 {
		appUserService.IncreaseGold(answer.UserId, question.PayNum, user.NickName, user.HeaderImg, "")
	}
}

// SetQuestionScore 给问题打分（1-5星）
// @Tags 问答
// @Summary 给问题打分（1-5星）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.SetQuestionScore true "给问题打分（1-5星）"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/question/setQuestionScore [post]
func (questionApi *QuestionApi) SetQuestionScore(c *gin.Context) {
	var req communityReq.SetQuestionScore
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return

	}
}
