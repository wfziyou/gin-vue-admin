// 自动生成模板HkForumTopicPostsMapping
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkForumTopicPostsMapping 结构体
type ForumTopicPostsMapping struct {
	global.GVA_MODEL
	TopicId uint `json:"topicId" form:"topicId" gorm:"column:topic_id;comment:话题id;size:19;"`
	PostsId uint `json:"postsId" form:"postsId" gorm:"column:posts_id;comment:帖子id;size:19;"`
}

// TableName HkForumTopicPostsMapping 表名
func (ForumTopicPostsMapping) TableName() string {
	return "hk_forum_topic_posts_mapping"
}
