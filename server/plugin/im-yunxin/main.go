package imYunXin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im-yunxin/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im-yunxin/router"
	"github.com/gin-gonic/gin"
)

type ImPlugin struct {
}

func CreateImPlug(Url, AppKey, AppSecret string) *ImPlugin {
	global.GlobalConfig.Url = Url
	global.GlobalConfig.AppKey = AppKey
	global.GlobalConfig.AppSecret = AppSecret
	return &ImPlugin{}
}

func (*ImPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitYunXinImRouter(group)
}

func (*ImPlugin) RouterPath() string {
	return "yunXinIm"
}
