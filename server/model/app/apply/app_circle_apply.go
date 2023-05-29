// 自动生成模板CircleApply
package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CircleApply 结构体
type CircleApply struct {
	global.GvaModelApp
	CircleId     uint64    `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`                  //圈子_编号
	ApplyGroupId uint64    `json:"applyGroupId" form:"applyGroupId" gorm:"type:bigint(20);column:apply_group_id;comment:应用分组_编号;"` //应用分组_编号
	ApplyId      uint64    `json:"applyId" form:"applyId" gorm:"type:bigint(20);column:apply_id;comment:应用_编号;"`                     //应用_编号
	BrowseNum    uint64    `json:"browseNum" form:"browseNum" gorm:"column:browse_num;comment:浏览次数;size:19;"`                        //浏览次数
	Power        int       `json:"power" form:"power" gorm:"column:power;comment:权限：0 所有可见 1 圈内成员可见;size:10;"`               //权限：0 所有可见 1 圈内成员可见
	Sort         int       `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                            //排序
	Apply        ApplyInfo `json:"apply" gorm:"foreignKey:ID;references:ApplyId;comment:应用"`                                           //应用
}

// TableName CircleApply 表名
func (CircleApply) TableName() string {
	return "hk_circle_apply"
}
