// 自动生成模板HkForumTagGroup
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkForumTagGroup 结构体
type HkForumTagGroup struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"` //名称
}

// TableName HkForumTagGroup 表名
func (HkForumTagGroup) TableName() string {
	return "hk_forum_tag_group"
}
