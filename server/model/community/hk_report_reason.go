// 自动生成模板HkReportReason
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkReportReason 结构体
type HkReportReason struct {
	global.GVA_MODEL
	TenantId string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	Reason   string `json:"reason" form:"reason" gorm:"column:reason;comment:举报理由;size:20;"`
	Sort     *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
}

// TableName HkReportReason 表名
func (HkReportReason) TableName() string {
	return "hk_report_reason"
}
