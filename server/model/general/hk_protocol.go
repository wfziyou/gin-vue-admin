// 自动生成模板HkProtocol
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkProtocol 结构体
type HkProtocol struct {
	global.GVA_MODEL

	Name    string `json:"name" form:"name" gorm:"column:name;comment:协议名称;size:20;"`
	Des     string `json:"des" form:"des" gorm:"column:des;comment:协议说明;size:20;"`
	Content string `json:"content" form:"content" gorm:"column:content;comment:协议内容;size:2000;"`
	Module  string `json:"module" form:"module" gorm:"column:module;comment:所属模块/插件;size:20;"`
	Pos     string `json:"pos" form:"pos" gorm:"column:pos;comment:显示位置;size:20;"`
}

// TableName HkProtocol 表名
func (HkProtocol) TableName() string {
	return "hk_protocol"
}
