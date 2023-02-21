// 自动生成模板ForumThumbsUp
package community

// ForumThumbsUp 结构体
type ForumThumbsUp struct {
	UserId  uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户id;"`    //用户id
	PostsId uint64 `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"` //帖子编号
}

// TableName ForumThumbsUp 表名
func (ForumThumbsUp) TableName() string {
	return "hk_forum_thumbs_up"
}
