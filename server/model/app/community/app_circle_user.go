// 自动生成模板CircleUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
)

// CircleUser 结构体
type CircleUser struct {
	global.GvaModelApp
	UserId   uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户ID;"`        //用户ID
	Remark   string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:40;"`                   //备注
	Tag      string `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:512;"`                           //标签
	CircleId uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //圈子_编号
	Power    int    `json:"power" form:"power" gorm:"column:power;comment:权限：0普通 1圈主;"`                      //权限：0普通 1圈主
	Sort     int    `json:"sort" form:"sort" gorm:"column:sort;comment:用户的圈子排序;size:10;"`                    //用户的圈子排序
}

// TableName CircleUser 表名
func (CircleUser) TableName() string {
	return "hk_circle_user"
}

// CircleUser 结构体
type CircleUserInfo struct {
	global.GvaModelApp
	UserId   uint64              `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户ID;"`        //用户ID
	Remark   string              `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:40;"`                   //备注
	Tag      string              `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:512;"`                           //标签
	CircleId uint64              `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //圈子_编号
	Power    int                 `json:"power" form:"power" gorm:"column:power;comment:权限：0普通 1圈主;"`                      //权限：0普通 1圈主
	Sort     int                 `json:"sort" form:"sort" gorm:"column:sort;comment:用户的圈子排序;size:10;"`                    //用户的圈子排序
	UserInfo common.UserBaseInfo `json:"UserInfo" gorm:"foreignKey:ID;references:UserId;comment:用户基本信息"`                  //用户基本信息
}

// TableName CircleUser 表名
func (CircleUserInfo) TableName() string {
	return "hk_circle_user"
}
