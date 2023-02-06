// 自动生成模板HkUserCircleApply
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkUserCircleApply 结构体
type HkUserCircleApply struct {
	global.GVA_MODEL
	UserId   *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`        //用户编号
	CircleId *int `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"` //圈子_编号
	ApplyId  *int `json:"applyId" form:"applyId" gorm:"column:apply_id;comment:应用_编号;size:19;"`    //应用_编号
	Sort     *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                 //排序
}

// TableName HkUserCircleApply 表名
func (HkUserCircleApply) TableName() string {
	return "hk_user_circle_apply"
}
