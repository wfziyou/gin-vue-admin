// 自动生成模板HkForbiddenSpeakReason
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkForbiddenSpeakReason 结构体
type HkForbiddenSpeakReason struct {
	global.GVA_MODEL
	Reason string `json:"reason" form:"reason" gorm:"column:reason;comment:禁言理由;size:20;"` //禁言理由
	Sort   *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`         //排序
}

// TableName HkForbiddenSpeakReason 表名
func (HkForbiddenSpeakReason) TableName() string {
	return "hk_forbidden_speak_reason"
}
