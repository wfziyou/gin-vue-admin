package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RecordBrowsingUserHomepageRouter struct {
}

// InitRecordBrowsingUserHomepageRouter 初始化 RecordBrowsingUserHomepage 路由信息
func (s *RecordBrowsingUserHomepageRouter) InitRecordBrowsingUserHomepageRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	router := appRouter.Group("hkRecordBrowsingUserHomepage").Use(middleware.OperationRecord())
	routerWithoutRecord := appRouter.Group("hkRecordBrowsingUserHomepage")
	var api = v1.ApiGroupApp.AppApiGroup.RecordBrowsingUserHomepageApi
	{
		router.POST("createRecordBrowsingUserHomepage", api.CreateRecordBrowsingUserHomepage)             // 新建RecordBrowsingUserHomepage
		router.DELETE("deleteRecordBrowsingUserHomepage", api.DeleteRecordBrowsingUserHomepage)           // 删除RecordBrowsingUserHomepage
		router.DELETE("deleteRecordBrowsingUserHomepageByIds", api.DeleteRecordBrowsingUserHomepageByIds) // 批量删除RecordBrowsingUserHomepage
		router.PUT("updateRecordBrowsingUserHomepage", api.UpdateRecordBrowsingUserHomepage)              // 更新RecordBrowsingUserHomepage
	}
	{
		routerWithoutRecord.GET("findRecordBrowsingUserHomepage", api.FindRecordBrowsingUserHomepage)       // 根据ID获取RecordBrowsingUserHomepage
		routerWithoutRecord.GET("getRecordBrowsingUserHomepageList", api.GetRecordBrowsingUserHomepageList) // 获取RecordBrowsingUserHomepage列表
	}
}
