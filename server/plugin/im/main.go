package imYunXin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im/router"
	"github.com/gin-gonic/gin"
)

type YunXinImPlugin struct {
}

func CreateYunXinImPlug(Url, AppKey, AppSecret string) *YunXinImPlugin {
	global.GlobalConfig.Url = Url
	global.GlobalConfig.AppKey = AppKey
	global.GlobalConfig.AppSecret = AppSecret
	return &YunXinImPlugin{}
}

func (*YunXinImPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitYunXinImRouter(group)
}

func (*YunXinImPlugin) RouterPath() string {
	return "yunXinIm"
}
