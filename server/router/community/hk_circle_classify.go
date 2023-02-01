package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkCircleClassifyRouter struct {
}

// InitHkCircleClassifyRouter 初始化 HkCircleClassify 路由信息
func (s *HkCircleClassifyRouter) InitHkCircleClassifyRouter(Router *gin.RouterGroup) {
	hkCircleClassifyRouter := Router.Group("hkCircleClassify").Use(middleware.OperationRecord())
	hkCircleClassifyRouterWithoutRecord := Router.Group("hkCircleClassify")
	var hkCircleClassifyApi = v1.ApiGroupApp.CommunityApiGroup.HkCircleClassifyApi
	{
		hkCircleClassifyRouter.POST("createHkCircleClassify", hkCircleClassifyApi.CreateHkCircleClassify)   // 新建HkCircleClassify
		hkCircleClassifyRouter.DELETE("deleteHkCircleClassify", hkCircleClassifyApi.DeleteHkCircleClassify) // 删除HkCircleClassify
		hkCircleClassifyRouter.DELETE("deleteHkCircleClassifyByIds", hkCircleClassifyApi.DeleteHkCircleClassifyByIds) // 批量删除HkCircleClassify
		hkCircleClassifyRouter.PUT("updateHkCircleClassify", hkCircleClassifyApi.UpdateHkCircleClassify)    // 更新HkCircleClassify
	}
	{
		hkCircleClassifyRouterWithoutRecord.GET("findHkCircleClassify", hkCircleClassifyApi.FindHkCircleClassify)        // 根据ID获取HkCircleClassify
		hkCircleClassifyRouterWithoutRecord.GET("getHkCircleClassifyList", hkCircleClassifyApi.GetHkCircleClassifyList)  // 获取HkCircleClassify列表
	}
}
