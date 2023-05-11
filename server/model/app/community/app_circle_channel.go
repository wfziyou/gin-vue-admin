// 自动生成模板HkChannel
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type CircleChannelInfo struct {
	ID   uint64 `json:"id" gorm:"not null;unique;primary_key;"`                  //编号
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"` //名称
}

func (CircleChannelInfo) TableName() string {
	return "hk_circle_channel"
}

// CircleChannel 结构体
type CircleChannel struct {
	global.GvaModelApp
	CircleId uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //圈子_编号
	Name     string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                         //名称
}

// TableName HkCircleChannel 表名
func (CircleChannel) TableName() string {
	return "hk_circle_channel"
}
