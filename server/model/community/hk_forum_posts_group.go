// 自动生成模板HkForumPostsGroup
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkForumPostsGroup 结构体
type HkForumPostsGroup struct {
	global.GVA_MODEL
	TenantId string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	CircleId *int   `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`
	Name     string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`
	ParentId *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父节点编号;size:19;"`
	Sort     *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`

	Status *int `json:"status" form:"status" gorm:"column:status;comment:审核状态(0未审批 1通过 2拒绝);size:10;"`
}

// TableName HkForumPostsGroup 表名
func (HkForumPostsGroup) TableName() string {
	return "hk_forum_posts_group"
}
