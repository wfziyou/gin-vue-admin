// 自动生成模板HkForumTopicPostsMapping
package community

// HkForumTopicPostsMapping 结构体
type ForumTopicPostsMapping struct {
	TopicId uint64 `gorm:"foreign key: topic_id, reference: hk_forum_topic.id;comment:话题id;"` //话题id
	PostsId uint64 `gorm:"foreign key: posts_id, reference: hk_forum_posts.id;comment:帖子id;"` //帖子id
}

// TableName HkForumTopicPostsMapping 表名
func (ForumTopicPostsMapping) TableName() string {
	return "hk_forum_topic_posts_mapping"
}
