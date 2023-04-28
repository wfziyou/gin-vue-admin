// 自动生成模板CircleUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CircleTag 结构体
type CircleTag struct {
	global.GvaModelAppEx
	Name string `json:"name" form:"name" gorm:"column:name;comment:标签名称;size:20;"` //标签名称
	Sort int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`     //排序
}

// TableName CircleTag 表名
func (CircleTag) TableName() string {
	return "hk_circle_tag"
}
