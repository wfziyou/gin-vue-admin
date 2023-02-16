// 自动生成模板CircleApplyGroup
package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CircleApplyGroup 结构体
type CircleApplyGroup struct {
	global.GvaModelApp
	CircleId *uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //圈子_编号
	Name     string  `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                         //名称
	Icon     string  `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                        //图标
	ParentId *uint64 `json:"parentId" form:"parentId" gorm:"type:bigint(20);column:parent_id;comment:父节点编号;"` //父节点编号
	Sort     *int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                         //排序
}

// TableName CircleApplyGroup 表名
func (CircleApplyGroup) TableName() string {
	return "hk_circle_apply_group"
}
