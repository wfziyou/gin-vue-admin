package oneLogin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/oneLogin/global"
	"github.com/gin-gonic/gin"
)

type OneLoginPlugin struct{}

func CreateOneLoginPlug(Appid, StrictCheck, RsaPrivateKey string) *OneLoginPlugin {
	global.GlobalConfig.Appid = Appid
	global.GlobalConfig.StrictCheck = StrictCheck
	global.GlobalConfig.RsaPrivateKey = RsaPrivateKey
	return &OneLoginPlugin{}
}

func (*OneLoginPlugin) Register(group *gin.RouterGroup) {

}

func (*OneLoginPlugin) RouterPath() string {
	return "oneLogin"
}
