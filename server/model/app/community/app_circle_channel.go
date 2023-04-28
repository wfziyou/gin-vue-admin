// 自动生成模板HkChannel
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type CircleChannelInfo struct {
	ID        uint64 `json:"id" gorm:"not null;unique;primary_key;"`                                          //编号
	Name      string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                         //名称
	Code      int    `json:"code" form:"code" gorm:"column:code;comment:标识：0常规、1资讯、2关注、3附近、4活动、5问答;size:10;"` //标识：0常规、1资讯、2关注、3附近、4活动、5问答
	FixedType int    `json:"fixedType" form:"fixedType" gorm:"column:fixed_type;comment:固定标识(不固定:0，固定:1);"`   //固定标识(不固定:0，固定:1)
	Sort      int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                         //排序
}

func (CircleChannelInfo) TableName() string {
	return "hk_circle_channel"
}

// CircleChannel 结构体
type CircleChannel struct {
	global.GvaModelApp
	CircleId  uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //圈子_编号
	Name      string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                         //名称
	Icon      string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                        //图标
	Code      int    `json:"code" form:"code" gorm:"column:code;comment:标识：0常规、1资讯、2关注、3附近、4活动、5问答;size:10;"` //标识：0常规、1资讯、2关注、3附近、4活动、5问答
	FixedType int    `json:"fixedType" form:"fixedType" gorm:"column:fixed_type;comment:固定标识(不固定:0，固定:1);"`   //固定标识(不固定:0，固定:1)
	Sort      int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                         //排序
}

// TableName HkCircleChannel 表名
func (CircleChannel) TableName() string {
	return "hk_circle_channel"
}
