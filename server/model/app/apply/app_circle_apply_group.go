// 自动生成模板CircleApplyGroup
package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CircleApplyGroup 结构体
type CircleApplyGroup struct {
	global.GvaModelApp
	CircleId *int   `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"` //圈子_编号
	Name     string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                 //名称
	Icon     string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                //图标
	ParentId *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父节点编号;size:19;"` //父节点编号
	Sort     *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                 //排序
}

// TableName CircleApplyGroup 表名
func (CircleApplyGroup) TableName() string {
	return "hk_circle_apply_group"
}
