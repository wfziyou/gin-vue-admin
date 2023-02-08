// 自动生成模板ReportReason
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ReportReason 结构体
type ReportReason struct {
	global.GvaModelApp
	Reason string `json:"reason" form:"reason" gorm:"column:reason;comment:举报理由;size:20;"` //举报理由
	Sort   *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`         //排序
}

// TableName ReportReason 表名
func (ReportReason) TableName() string {
	return "hk_report_reason"
}
