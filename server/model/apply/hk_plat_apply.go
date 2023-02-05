// 自动生成模板HkPlatApply
package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkPlatApply 结构体
type HkPlatApply struct {
	global.GVA_MODEL
	TenantId        string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	ApplyGroupId    *int   `json:"applyGroupId" form:"applyGroupId" gorm:"column:apply_group_id;comment:应用分组_编号;size:19;"`
	ShowName        string `json:"showName" form:"showName" gorm:"column:show_name;comment:名称别名;size:80;"`
	ApplyId         *int   `json:"applyId" form:"applyId" gorm:"column:apply_id;comment:应用_编号;size:19;"`
	ApplyParameters string `json:"applyParameters" form:"applyParameters" gorm:"column:apply_parameters;comment:访问参数;size:256;"`
	Sort            *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
	CreateUser      *int   `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept      *int   `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
}

// TableName HkPlatApply 表名
func (HkPlatApply) TableName() string {
	return "hk_plat_apply"
}
