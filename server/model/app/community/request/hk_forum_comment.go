package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ForumCommentSearch struct {
	PostsId  uint64 `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"`     //帖子编号
	ParentId uint64 `json:"parentId" form:"parentId" gorm:"type:bigint(20);column:parent_id;comment:父评论编号;"` //父评论编号

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type CreateForumComment struct {
	UserId         uint64 `json:"-"`                                                                                 //用户ID
	PostsId        uint64 `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"`       //帖子编号
	CommentId      uint64 `json:"commentId" form:"commentId" gorm:"type:bigint(20);column:comment_id;comment:评论编号;"` //评论编号
	CommentContent string `json:"commentContent" form:"commentContent" gorm:"column:comment_content;comment:评论内容;"`  //评论内容
}

type DeleteForumThumbsUp struct {
	UserId  uint64 `json:"-"`                                                                           //用户ID
	PostsId uint64 `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"` //帖子编号
}
type DeleteCommentThumbsUp struct {
	UserId    uint64 `json:"-"`                                                                                 //用户ID
	CommentId uint64 `json:"commentId" form:"commentId" gorm:"type:bigint(20);column:comment_id;comment:评论编号;"` //评论编号
}
