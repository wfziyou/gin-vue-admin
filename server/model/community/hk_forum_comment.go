// 自动生成模板HkForumComment
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkForumComment 结构体
type HkForumComment struct {
	global.GVA_MODEL
	PostsId        *int       `json:"postsId" form:"postsId" gorm:"column:posts_id;comment:帖子编号;size:19;"`                           //帖子编号
	UserId         *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:评论者;size:19;"`                               //评论者
	CommentTime    *time.Time `json:"commentTime" form:"commentTime" gorm:"column:comment_time;comment:评论时间;"`                       //评论时间
	CommentContent string     `json:"commentContent" form:"commentContent" gorm:"column:comment_content;comment:评论内容;"`              //评论内容
	LikeTimes      *int       `json:"likeTimes" form:"likeTimes" gorm:"column:like_times;comment:点赞数;size:10;"`                      //点赞数
	ParentId       *int       `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父评论编号;size:19;"`                       //父评论编号
	CheckUser      *int       `json:"checkUser" form:"checkUser" gorm:"column:check_user;comment:审核人;size:19;"`                      //审核人
	CheckTime      *time.Time `json:"checkTime" form:"checkTime" gorm:"column:check_time;comment:审核时间;"`                             //审核时间
	CheckStatus    *int       `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态(0未审批 1通过 2拒绝);size:10;"` //审核状态(0未审批 1通过 2拒绝)
}

// TableName HkForumComment 表名
func (HkForumComment) TableName() string {
	return "hk_forum_comment"
}
