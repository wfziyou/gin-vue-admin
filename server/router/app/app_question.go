package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type QuestionRouter struct {
}

// InitQuestionRouter 初始化 Question 路由信息
func (s *QuestionRouter) InitQuestionRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	router := appRouter.Group("question").Use(middleware.OperationRecord())
	routerWithoutRecord := appRouter.Group("question")
	var api = v1.ApiGroupApp.AppApiGroup.QuestionApi
	{
		router.POST("createQuestion", api.CreateQuestion) // 创建问题
		router.POST("closeQuestion", api.CloseQuestion)   // 关闭问题
		router.POST("answerQuestion", api.AnswerQuestion) // 回答问题
		router.DELETE("delSelfAnswer", api.DelSelfAnswer) // 删除自己的回答

		router.POST("setBestAnswer", api.SetBestAnswer)       // 设置最佳答案
		router.POST("setQuestionScore", api.SetQuestionScore) // 给问题打分（1-5星）
	}
	{
		routerWithoutRecord.GET("getGlobalRecommendQuestionList", api.GetGlobalRecommendQuestionList) // 分页获取全局推荐问题列表
		routerWithoutRecord.GET("getCircleQuestionList", api.GetCircleQuestionList)                   // 分页获取圈子问题列表
		routerWithoutRecord.GET("getAnswerList", api.GetAnswerList)                                   // 分页获取问题的回答列表
	}
}
