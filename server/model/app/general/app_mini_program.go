// 自动生成模板HkMiniProgram
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkMiniProgram 结构体
type HkMiniProgram struct {
	global.GvaModelApp
	Name        string `json:"name" form:"name" gorm:"column:name;comment:名称;size:80;"`                           //名称
	Icon        string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                          //图标
	CompanyName string `json:"companyName" form:"companyName" gorm:"column:company_name;comment:公司名称;size:256;"`  //公司名称
	CurVersion  string `json:"curVersion" form:"curVersion" gorm:"column:cur_version;comment:当前版本;size:12;"`      //当前版本
	CurPacketId string `json:"curPacketId" form:"curPacketId" gorm:"column:cur_packet_id;comment:当前包id;size:12;"` //当前包id
	Hidden      *int   `json:"hidden" form:"hidden" gorm:"column:hidden;comment:隐藏(0否 1是);size:10;"`              //隐藏(0否 1是)
}

// TableName HkMiniProgram 表名
func (HkMiniProgram) TableName() string {
	return "hk_mini_program"
}
