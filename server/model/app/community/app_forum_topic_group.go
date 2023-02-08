// 自动生成模板ForumTopicGroup
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ForumTopicGroup 结构体
type ForumTopicGroup struct {
	global.GvaModelApp
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"` //名称
	Sort *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"` //排序
}

// TableName ForumTopicGroup 表名
func (ForumTopicGroup) TableName() string {
	return "hk_forum_topic_group"
}
