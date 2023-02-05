// 自动生成模板HkForumTopicGroup
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkForumTopicGroup 结构体
type HkForumTopicGroup struct {
	global.GVA_MODEL

	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`
	Sort *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
}

// TableName HkForumTopicGroup 表名
func (HkForumTopicGroup) TableName() string {
	return "hk_forum_topic_group"
}
