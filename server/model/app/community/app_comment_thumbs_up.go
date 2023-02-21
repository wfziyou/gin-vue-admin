// 自动生成模板CommentThumbsUp
package community

// CommentThumbsUp 结构体
type CommentThumbsUp struct {
	UserId    uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户id;"`          //用户id 	//点赞时间
	CommentId uint64 `json:"commentId" form:"commentId" gorm:"type:bigint(20);column:comment_id;comment:评论编号;"` //评论编号
}

// TableName CommentThumbsUp 表名
func (CommentThumbsUp) TableName() string {
	return "hk_comment_thumbs_up"
}
