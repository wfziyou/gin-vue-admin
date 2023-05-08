package imOpen

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open/global"
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

}

func (*ImPlugin) RouterPath() string {
	return "yunXinIm"
}
