// 自动生成模板UserCircleApply
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserCircleApply 结构体
type UserCircleApply struct {
	global.GvaModelApp
	UserId   uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户编号;"`        //用户编号
	CircleId uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //圈子_编号
	ApplyId  uint64 `json:"applyId" form:"applyId" gorm:"type:bigint(20);column:apply_id;comment:应用_编号;"`    //应用_编号
	Sort     *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                         //排序
}

// TableName UserCircleApply 表名
func (UserCircleApply) TableName() string {
	return "hk_user_circle_apply"
}
