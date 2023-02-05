// 自动生成模板HkForbiddenSpeakDuration
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkForbiddenSpeakDuration 结构体
type HkForbiddenSpeakDuration struct {
	global.GVA_MODEL
	TenantId string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	Count    *int   `json:"count" form:"count" gorm:"column:count;comment:时长;size:10;"`
	Type     *int   `json:"type" form:"type" gorm:"column:type;comment:时间类型：0 小时、1天;size:10;"`
	Sort     *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
}

// TableName HkForbiddenSpeakDuration 表名
func (HkForbiddenSpeakDuration) TableName() string {
	return "hk_forbidden_speak_duration"
}
