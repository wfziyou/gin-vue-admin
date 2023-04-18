package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RecordBrowsingCircleHomepageRouter struct {
}

// InitRecordBrowsingCircleHomepageRouter 初始化 RecordBrowsingCircleHomepage 路由信息
func (s *RecordBrowsingCircleHomepageRouter) InitRecordBrowsingCircleHomepageRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	router := appRouter.Group("hkRecordBrowsingCircleHomepage").Use(middleware.OperationRecord())
	routerWithoutRecord := appRouter.Group("hkRecordBrowsingCircleHomepage")
	var api = v1.ApiGroupApp.AppApiGroup.RecordBrowsingCircleHomepageApi
	{
		router.POST("createRecordBrowsingCircleHomepage", api.CreateRecordBrowsingCircleHomepage)             // 新建RecordBrowsingCircleHomepage
		router.DELETE("deleteRecordBrowsingCircleHomepage", api.DeleteRecordBrowsingCircleHomepage)           // 删除RecordBrowsingCircleHomepage
		router.DELETE("deleteRecordBrowsingCircleHomepageByIds", api.DeleteRecordBrowsingCircleHomepageByIds) // 批量删除RecordBrowsingCircleHomepage
		router.PUT("updateRecordBrowsingCircleHomepage", api.UpdateRecordBrowsingCircleHomepage)              // 更新RecordBrowsingCircleHomepage
	}
	{
		routerWithoutRecord.GET("findRecordBrowsingCircleHomepage", api.FindRecordBrowsingCircleHomepage)       // 根据ID获取RecordBrowsingCircleHomepage
		routerWithoutRecord.GET("getRecordBrowsingCircleHomepageList", api.GetRecordBrowsingCircleHomepageList) // 获取RecordBrowsingCircleHomepage列表
	}
}
