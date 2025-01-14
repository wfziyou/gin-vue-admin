package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/im-yunxin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/oneLogin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/plugin"
	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==》", PrivateGroup)
	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
	))
	PluginInit(PrivateGroup, smsAliyun.CreateAliSmsPlug(
		global.GVA_CONFIG.AliyunSms.AccessKeyId,
		global.GVA_CONFIG.AliyunSms.AccessSecret,
		global.GVA_CONFIG.AliyunSms.SignName,
	))

	PluginInit(PrivateGroup, imYunXin.CreateImPlug(
		global.GVA_CONFIG.YunXinIm.Url,
		global.GVA_CONFIG.YunXinIm.AppKey,
		global.GVA_CONFIG.YunXinIm.AppSecret,
	))
	PluginInit(PrivateGroup, imOpen.CreateImPlug(
		global.GVA_CONFIG.OpenIm.Url,
		global.GVA_CONFIG.OpenIm.AppKey,
		global.GVA_CONFIG.OpenIm.AppSecret,
	))
	PluginInit(PrivateGroup, oneLogin.CreateOneLoginPlug(
		global.GVA_CONFIG.OneLogin.Appid,
		global.GVA_CONFIG.OneLogin.StrictCheck,
		global.GVA_CONFIG.OneLogin.RsaPrivateKey,
	))
}
