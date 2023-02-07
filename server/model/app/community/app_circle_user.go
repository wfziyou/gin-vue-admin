// 自动生成模板HkCircleUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
)

// HkCircleUser 结构体
type HkCircleUser struct {
	global.GvaModelApp
	UserId       *int                `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:19;"`        //用户ID
	CircleId     *int                `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"` //圈子_编号
	Sort         *int                `json:"sort" form:"sort" gorm:"column:sort;comment:用户的圈子排序;size:10;"`            //用户的圈子排序
	UserBaseInfo common.UserBaseInfo `json:"userBaseInfo" gorm:"foreignKey:ID;references:UserId;comment:用户基本信息"`      //用户基本信息
}

// TableName HkCircleUser 表名
func (HkCircleUser) TableName() string {
	return "hk_circle_user"
}
