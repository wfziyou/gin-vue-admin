// 自动生成模板HelpType
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HelpType 结构体
type HelpType struct {
	global.GvaModelAppEx
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"` //名称
	Sort int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"` //排序
}

// TableName HelpType 表名
func (HelpType) TableName() string {
	return "hk_help_type"
}
