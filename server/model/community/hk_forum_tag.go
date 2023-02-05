// 自动生成模板HkForumTag
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkForumTag 结构体
type HkForumTag struct {
	global.GVA_MODEL
	TenantId string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	Name     string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`
	GroupId  *int   `json:"groupId" form:"groupId" gorm:"column:group_id;comment:分组id;size:19;"`
}

// TableName HkForumTag 表名
func (HkForumTag) TableName() string {
	return "hk_forum_tag"
}
