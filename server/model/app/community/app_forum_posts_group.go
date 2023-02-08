// 自动生成模板ForumPostsGroup
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ForumPostsGroup 结构体
type ForumPostsGroup struct {
	global.GvaModelApp
	CircleId    *int   `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`                       //圈子_编号
	Name        string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                                       //名称
	ParentId    *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父节点编号;size:19;"`                       //父节点编号
	CheckStatus *int   `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态(0未审批 1通过 2拒绝);size:10;"` //审核状态(0未审批 1通过 2拒绝)
	Sort        *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                       //排序
}

// TableName ForumPostsGroup 表名
func (ForumPostsGroup) TableName() string {
	return "hk_forum_posts_group"
}
