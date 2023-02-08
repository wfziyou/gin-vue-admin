// 自动生成模板ForumTagGroup
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ForumTagGroup 结构体
type ForumTagGroup struct {
	global.GvaModelApp
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"` //名称
}

// TableName ForumTagGroup 表名
func (ForumTagGroup) TableName() string {
	return "hk_forum_tag_group"
}
