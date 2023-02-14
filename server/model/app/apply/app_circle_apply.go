// 自动生成模板CircleApply
package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CircleApply 结构体
type CircleApply struct {
	global.GvaModelApp
	CircleId        *int   `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`                //圈子_编号
	ApplyGroupId    *int   `json:"applyGroupId" form:"applyGroupId" gorm:"type:bigint(20);column:apply_group_id;comment:应用分组_编号;"` //应用分组_编号
	ShowName        string `json:"showName" form:"showName" gorm:"column:show_name;comment:名称别名;size:80;"`                         //名称别名
	ApplyId         *int   `json:"applyId" form:"applyId" gorm:"column:apply_id;comment:应用_编号;size:19;"`                           //应用_编号
	ApplyParameters string `json:"applyParameters" form:"applyParameters" gorm:"column:apply_parameters;comment:访问参数;size:256;"`   //访问参数
	Sort            *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                        //排序
}

// TableName CircleApply 表名
func (CircleApply) TableName() string {
	return "hk_circle_apply"
}
