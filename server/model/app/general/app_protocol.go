// 自动生成模板Protocol
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Protocol 结构体
type Protocol struct {
	global.GvaModelApp
	Name    string `json:"name" form:"name" gorm:"column:name;comment:协议名称;size:20;"`            //协议名称
	Des     string `json:"des" form:"des" gorm:"column:des;comment:协议说明;size:20;"`               //协议说明
	Content string `json:"content" form:"content" gorm:"column:content;comment:协议内容;size:2000;"` //协议内容
	Module  string `json:"module" form:"module" gorm:"column:module;comment:所属模块/插件;size:20;"`   //所属模块/插件
	Pos     string `json:"pos" form:"pos" gorm:"column:pos;comment:显示位置;size:20;"`               //显示位置
}

// TableName Protocol 表名
func (Protocol) TableName() string {
	return "hk_protocol"
}
