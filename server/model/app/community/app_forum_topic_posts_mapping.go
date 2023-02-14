// 自动生成模板HkForumTopicPostsMapping
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkForumTopicPostsMapping 结构体
type ForumTopicPostsMapping struct {
	global.GvaModelApp
	TopicId uint64 `json:"topicId" form:"topicId" gorm:"type:bigint(20);column:topic_id;comment:话题id;"`
	PostsId uint64 `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子id;"`
}

// TableName HkForumTopicPostsMapping 表名
func (ForumTopicPostsMapping) TableName() string {
	return "hk_forum_topic_posts_mapping"
}
