// 自动生成模板Feedback
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Feedback 结构体
type Feedback struct {
	global.GvaModelApp
	TenantId    string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	UserId      uint64 `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
	TypeId      uint64 `json:"typeId" form:"typeId" gorm:"column:type_id;comment:类型编号;size:19;"`
	Type        string `json:"type" form:"type" gorm:"column:type;comment:类型;size:20;"`
	Des         string `json:"des" form:"des" gorm:"column:des;comment:描述;size:1024;"`
	Attachment  string `json:"attachment" form:"attachment" gorm:"column:attachment;comment:附件;size:2000;"`
	CheckStatus int    `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:处理状态：0 未处理、1 处理中、2 拒绝、3 完成;size:10;"`
	Process     string `json:"process" form:"process" gorm:"column:process;comment:处理描述;size:2000;"`
}

// TableName Feedback 表名
func (Feedback) TableName() string {
	return "hk_feedback"
}
