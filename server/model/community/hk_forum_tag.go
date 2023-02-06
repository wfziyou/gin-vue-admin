// 自动生成模板HkForumTag
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkForumTag 结构体
type HkForumTag struct {
	global.GVA_MODEL
	Name    string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`             //名称
	GroupId *int   `json:"groupId" form:"groupId" gorm:"column:group_id;comment:分组id;size:19;"` //分组id
}

// TableName HkForumTag 表名
func (HkForumTag) TableName() string {
	return "hk_forum_tag"
}
