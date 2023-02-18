// 自动生成模板ForumComment
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ForumComment 结构体
type ForumComment struct {
	global.GvaModelApp
	PostsId        uint64     `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"`                   //帖子编号
	UserId         uint64     `json:"userId" form:"userId" gorm:"type:bigint(20);column:circle_id;comment:评论者;"`                     //评论者
	CommentContent string     `json:"commentContent" form:"commentContent" gorm:"column:comment_content;comment:评论内容;"`              //评论内容
	LikeTimes      int        `json:"likeTimes" form:"likeTimes" gorm:"column:like_times;comment:点赞数;size:10;"`                      //点赞数
	ParentId       uint64     `json:"parentId" form:"parentId" gorm:"type:bigint(20);column:parent_id;comment:父评论编号;"`               //父评论编号
	CheckUser      uint64     `json:"checkUser" form:"checkUser" gorm:"type:bigint(20);column:check_user;comment:审核人;"`              //审核人
	CheckTime      *time.Time `json:"checkTime" form:"checkTime" gorm:"column:check_time;comment:审核时间;"`                             //审核时间
	CheckStatus    int        `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态(0未审批 1通过 2拒绝);size:10;"` //审核状态(0未审批 1通过 2拒绝)
	ThumbsUp       int        `json:"thumbsUp"`                                                                                      //是否点赞：0否、1是
}

// TableName ForumComment 表名
func (ForumComment) TableName() string {
	return "hk_forum_comment"
}
