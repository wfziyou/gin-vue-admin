package app

import (
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type QuestionApi struct {
}

// GetGlobalRecommendQuestionList 分页获取全局推荐问题列表
// @Tags 问答
// @Summary 分页获取全局推荐问题列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.GlobalRecommendQuestionSearch true "分页获取全局推荐问题列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPosts,msg=string} "返回community.ForumPosts"
// @Router /app/question/getGlobalRecommendQuestionList [get]
func (questionApi *QuestionApi) GetGlobalRecommendQuestionList(c *gin.Context) {
	var pageInfo communityReq.GlobalRecommendQuestionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
}

// GetCircleQuestionList 分页获取圈子问题列表
// @Tags 问答
// @Summary 分页获取圈子问题列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.CircleQuestionSearch true "分页获取圈子问题列表"
// @Success 200 {object} response.PageResult{List=[]community.ForumPosts,msg=string} "返回community.ForumPosts"
// @Router /app/question/getCircleQuestionList [get]
func (questionApi *QuestionApi) GetCircleQuestionList(c *gin.Context) {
	var pageInfo communityReq.CircleQuestionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
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
	var param communityReq.CreateQuestion
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if err := hkActivityService.CreateActivity(hkActivity); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
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
	var param communityReq.CloseQuestion
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
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
	var pageInfo communityReq.QuestionAnswerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
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
