// 自动生成模板HkPlatApplyGroup
package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkPlatApplyGroup 结构体
type HkPlatApplyGroup struct {
	global.GVA_MODEL
	TenantId string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	Name     string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`
	Icon     string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`
	ParentId *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父节点编号;size:19;"`
	Sort     *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
}

// TableName HkPlatApplyGroup 表名
func (HkPlatApplyGroup) TableName() string {
	return "hk_plat_apply_group"
}
