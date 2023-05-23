package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GeneralRouter struct {
}

func (s *GeneralRouter) InitGeneralRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	router := appRouter.Group("general").Use(middleware.OperationRecord())
	withoutRecord := appRouter.Group("general")
	var api = v1.ApiGroupApp.AppApiGroup.GeneralApi
	{
		withoutRecord.GET("findMiniProgram", api.FindMiniProgram)             //用id查询小程序
		withoutRecord.GET("getUserCoverImageList", api.GetUserCoverImageList) // 获取用户主页封面列表
	}
	{
		withoutRecord.GET("findDraft", api.FindDraft)           // 用id查询草稿
		withoutRecord.GET("getDraftList", api.GetDraftList)     // 分页获取草稿列表
		router.DELETE("deleteAllDraft", api.DeleteAllDraft)     // 删除所有草稿
		router.DELETE("deleteDraft", api.DeleteDraft)           // 删除草稿
		router.DELETE("deleteDraftByIds", api.DeleteDraftByIds) // 批量删除草稿
		router.POST("updateDraft", api.UpdateDraft)             // 更新草稿
	}
	return router
}
