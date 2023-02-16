// 自动生成模板ForumTag
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ForumTag 结构体
type ForumTag struct {
	global.GvaModelApp
	Name    string  `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                     //名称
	GroupId *uint64 `json:"groupId" form:"groupId" gorm:"type:bigint(20);column:group_id;comment:分组id;"` //分组id
}

// TableName ForumTag 表名
func (ForumTag) TableName() string {
	return "hk_forum_tag"
}
