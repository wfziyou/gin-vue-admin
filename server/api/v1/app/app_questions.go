package app

import (
	"github.com/gin-gonic/gin"
)

type QuestionsApi struct {
}

// GetGlobalQuestionsList 获取全局问题列表
// @Tags questions
// @Summary 获取全局问题列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "获取全局问题列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /questions/getGlobalQuestionsList [post]
func (questionsApi *QuestionsApi) GetGlobalQuestionsList(c *gin.Context) {

}

// GetCircleQuestionsList 获取圈子问题列表
// @Tags questions
// @Summary 获取圈子问题列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "获取圈子问题列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /questions/getCircleQuestionsList [post]
func (questionsApi *QuestionsApi) GetCircleQuestionsList(c *gin.Context) {

}

// CreateQuestions 创建问题
// @Tags questions
// @Summary 创建问题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivity true "创建问题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questions/createQuestions [post]
func (hkActivityApi *HkActivityApi) CreateQuestions(c *gin.Context) {
	//var hkActivity community.HkActivity
	//err := c.ShouldBindJSON(&hkActivity)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//if err := hkActivityService.CreateHkActivity(hkActivity); err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//} else {
	//	response.OkWithMessage("创建成功", c)
	//}
}

// CloseQuestions 关闭问题
// @Tags questions
// @Summary 关闭问题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivity true "关闭问题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questions/closeQuestions [post]
func (hkActivityApi *HkActivityApi) CloseQuestions(c *gin.Context) {

}

// GetAnswerList 获取回答列表
// @Tags questions
// @Summary 获取回答列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "获取回答列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /questions/getAnswerList [post]
func (questionsApi *QuestionsApi) GetAnswerList(c *gin.Context) {

}

// AnswerQuestions 回答问题
// @Tags questions
// @Summary 回答问题
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivity true "回答问题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questions/answerQuestions [post]
func (hkActivityApi *HkActivityApi) AnswerQuestions(c *gin.Context) {

}

// DelSelfAnswer 删除自己的回答
// @Tags questions
// @Summary 删除自己的回答
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivity true "删除自己的回答"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questions/delSelfAnswer [post]
func (hkActivityApi *HkActivityApi) DelSelfAnswer(c *gin.Context) {

}

// SetBestAnswer 设置最佳答案
// @Tags questions
// @Summary 设置最佳答案
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivity true "设置最佳答案"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questions/setBestAnswer [post]
func (hkActivityApi *HkActivityApi) SetBestAnswer(c *gin.Context) {

}

// SetQuestionsScore 给问题打分（1-5星）
// @Tags questions
// @Summary 给问题打分（1-5星）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.HkActivity true "给问题打分（1-5星）"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questions/setQuestionsScore [post]
func (hkActivityApi *HkActivityApi) SetQuestionsScore(c *gin.Context) {

}
