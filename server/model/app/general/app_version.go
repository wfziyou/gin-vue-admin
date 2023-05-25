// 自动生成模板Protocol
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppVersion 结构体
type AppVersion struct {
	global.GvaModelAppEx
	Platform    int    `json:"platform" form:"platform" gorm:"column:platform;comment:平台：1IOS,2Android,3Windows,4OSX,5Web,6MiniWeb,7Linux,8APad,9IPad,10Admin;"` //平台：1IOS,2Android,3Windows,4OSX,5Web,6MiniWeb,7Linux,8APad,9IPad,10Admin
	Version     string `json:"version" form:"version" gorm:"column:version;comment:版本;size:64;"`                                                                 //版本
	ForceUpdate int    `json:"forceUpdate" form:"forceUpdate" gorm:"column:force_update;comment:强制更新：0否 1是;"`                                                    //强制更新：0否 1是
	PacketName  string `json:"packetName" form:"packetName" gorm:"column:packet_name;comment:包名;size:256;"`                                                      //包名
	PacketUrl   string `json:"packetUrl" form:"packetUrl" gorm:"column:packet_url;comment:下载地址;size:256;"`                                                       //下载地址
	UpdateLog   string `json:"updateLog" form:"updateLog" gorm:"type:longText;column:update_log;comment:更新日志;"`                                                  //更新日志
}

// TableName Protocol 表名
func (AppVersion) TableName() string {
	return "hk_app_version"
}
