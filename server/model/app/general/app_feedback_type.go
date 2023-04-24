// 自动生成模板FeedbackType
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// FeedbackType 结构体
type FeedbackType struct {
	global.GVA_MODEL
	TenantId   string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`
	Sort       *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
	CreateUser *int   `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept *int   `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser *int   `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	Status     *int   `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	IsDel      *int   `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}

// TableName FeedbackType 表名
func (FeedbackType) TableName() string {
	return "hk_feedback_type"
}
