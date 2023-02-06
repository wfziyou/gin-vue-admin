// 自动生成模板HkCircleUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkCircleUser 结构体
type HkCircleUser struct {
	global.GVA_MODEL
	UserId   *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:19;"`        //用户ID
	CircleId *int `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"` //圈子_编号
	Sort     *int `json:"sort" form:"sort" gorm:"column:sort;comment:用户的圈子排序;size:10;"`            //用户的圈子排序
}

// TableName HkCircleUser 表名
func (HkCircleUser) TableName() string {
	return "hk_circle_user"
}
