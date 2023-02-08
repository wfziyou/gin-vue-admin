// 自动生成模板ForbiddenSpeakDuration
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ForbiddenSpeakDuration 结构体
type ForbiddenSpeakDuration struct {
	global.GvaModelApp
	Count *int `json:"count" form:"count" gorm:"column:count;comment:时长;size:10;"`        //时长
	Type  *int `json:"type" form:"type" gorm:"column:type;comment:时间类型：0 小时、1天;size:10;"` //时间类型：0 小时、1天
	Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`           //排序
}

// TableName ForbiddenSpeakDuration 表名
func (ForbiddenSpeakDuration) TableName() string {
	return "hk_forbidden_speak_duration"
}
