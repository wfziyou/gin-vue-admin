// 自动生成模板PlatApplyGroup
package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PlatApplyGroup 结构体
type PlatApplyGroup struct {
	global.GvaModelApp
	Name     string  `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`
	Icon     string  `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`
	ParentId *uint64 `json:"parentId" form:"parentId" gorm:"type:bigint(20);column:parent_id;comment:父节点编号;"`
	Sort     *int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
}

// TableName PlatApplyGroup 表名
func (PlatApplyGroup) TableName() string {
	return "hk_plat_apply_group"
}
