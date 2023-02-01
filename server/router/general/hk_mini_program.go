package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkMiniProgramRouter struct {
}

// InitHkMiniProgramRouter 初始化 HkMiniProgram 路由信息
func (s *HkMiniProgramRouter) InitHkMiniProgramRouter(Router *gin.RouterGroup) {
	hkMiniProgramRouter := Router.Group("hkMiniProgram").Use(middleware.OperationRecord())
	hkMiniProgramRouterWithoutRecord := Router.Group("hkMiniProgram")
	var hkMiniProgramApi = v1.ApiGroupApp.GeneralApiGroup.HkMiniProgramApi
	{
		hkMiniProgramRouter.POST("createHkMiniProgram", hkMiniProgramApi.CreateHkMiniProgram)   // 新建HkMiniProgram
		hkMiniProgramRouter.DELETE("deleteHkMiniProgram", hkMiniProgramApi.DeleteHkMiniProgram) // 删除HkMiniProgram
		hkMiniProgramRouter.DELETE("deleteHkMiniProgramByIds", hkMiniProgramApi.DeleteHkMiniProgramByIds) // 批量删除HkMiniProgram
		hkMiniProgramRouter.PUT("updateHkMiniProgram", hkMiniProgramApi.UpdateHkMiniProgram)    // 更新HkMiniProgram
	}
	{
		hkMiniProgramRouterWithoutRecord.GET("findHkMiniProgram", hkMiniProgramApi.FindHkMiniProgram)        // 根据ID获取HkMiniProgram
		hkMiniProgramRouterWithoutRecord.GET("getHkMiniProgramList", hkMiniProgramApi.GetHkMiniProgramList)  // 获取HkMiniProgram列表
	}
}
