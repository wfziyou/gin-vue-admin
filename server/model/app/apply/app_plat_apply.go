// 自动生成模板PlatApply
package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PlatApply 结构体
type PlatApply struct {
	global.GvaModelApp
	ApplyGroupId *uint64 `json:"applyGroupId" form:"applyGroupId" gorm:"type:bigint(20);column:apply_group_id;comment:应用分组_编号;"`
	ApplyId      *uint64 `json:"applyId" form:"applyId" gorm:"type:bigint(20);column:apply_id;comment:应用_编号"`
	Sort         *int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
}

// TableName PlatApply 表名
func (PlatApply) TableName() string {
	return "hk_plat_apply"
}
