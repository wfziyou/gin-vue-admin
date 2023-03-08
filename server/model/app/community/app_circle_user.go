// 自动生成模板CircleUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CircleUser 结构体
type CircleUser struct {
	global.GvaModelApp
	UserId   uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户ID;"`          //用户ID
	CircleId uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //圈子_编号
	Power    int    `json:"power" form:"power" gorm:"column:power;comment:权限：0普通 1圈主;"`                    //权限：0普通 1圈主
	Sort     int    `json:"sort" form:"sort" gorm:"column:sort;comment:用户的圈子排序;size:10;"`                 //用户的圈子排序

	//UserBaseInfo *User  `json:"userBaseInfo" gorm:"foreignKey:UserId;references:hk_user.id;comment:用户基本信息"` //用户基本信息
	//CircleBaseInfo *CircleBaseInfo      `json:"userBaseInfo" gorm:"foreignKey:CircleId;references:hk_circle.id;comment:圈子基本信息"` //圈子基本信息
}

// TableName CircleUser 表名
func (CircleUser) TableName() string {
	return "hk_circle_user"
}

// CircleUser 结构体
type CircleUserInfo struct {
	global.GvaModelApp
	UserId       uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户ID;"`          //用户ID
	CircleId     uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //圈子_编号
	Power        int    `json:"power" form:"power" gorm:"column:power;comment:权限：0普通 1圈主;"`                    //权限：0普通 1圈主
	Sort         int    `json:"sort" form:"sort" gorm:"column:sort;comment:用户的圈子排序;size:10;"`                 //用户的圈子排序
	UserBaseInfo User   `json:"userBaseInfo" gorm:"foreignKey:UserId;references:id;comment:用户基本信息"`            //用户基本信息
}

// TableName CircleUser 表名
func (CircleUserInfo) TableName() string {
	return "hk_circle_user"
}
