// 自动生成模板HkCircleUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkCircleUser 结构体
type HkCircleUser struct {
	global.GVA_MODEL

	UserId   *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:19;"`
	CircleId *int `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`
	Sort     *int `json:"sort" form:"sort" gorm:"column:sort;comment:用户的圈子排序;size:10;"`
}

// TableName HkCircleUser 表名
func (HkCircleUser) TableName() string {
	return "hk_circle_user"
}
