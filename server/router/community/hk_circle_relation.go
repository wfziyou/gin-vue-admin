package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkCircleRelationRouter struct {
}

// InitHkCircleRelationRouter 初始化 HkCircleRelation 路由信息
func (s *HkCircleRelationRouter) InitHkCircleRelationRouter(Router *gin.RouterGroup) {
	hkCircleRelationRouter := Router.Group("hkCircleRelation").Use(middleware.OperationRecord())
	hkCircleRelationRouterWithoutRecord := Router.Group("hkCircleRelation")
	var hkCircleRelationApi = v1.ApiGroupApp.CommunityApiGroup.HkCircleRelationApi
	{
		hkCircleRelationRouter.POST("createHkCircleRelation", hkCircleRelationApi.CreateHkCircleRelation)   // 新建HkCircleRelation
		hkCircleRelationRouter.DELETE("deleteHkCircleRelation", hkCircleRelationApi.DeleteHkCircleRelation) // 删除HkCircleRelation
		hkCircleRelationRouter.DELETE("deleteHkCircleRelationByIds", hkCircleRelationApi.DeleteHkCircleRelationByIds) // 批量删除HkCircleRelation
		hkCircleRelationRouter.PUT("updateHkCircleRelation", hkCircleRelationApi.UpdateHkCircleRelation)    // 更新HkCircleRelation
	}
	{
		hkCircleRelationRouterWithoutRecord.GET("findHkCircleRelation", hkCircleRelationApi.FindHkCircleRelation)        // 根据ID获取HkCircleRelation
		hkCircleRelationRouterWithoutRecord.GET("getHkCircleRelationList", hkCircleRelationApi.GetHkCircleRelationList)  // 获取HkCircleRelation列表
	}
}
