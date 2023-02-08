// 自动生成模板CommentThumbsUp
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// CommentThumbsUp 结构体
type CommentThumbsUp struct {
	global.GvaModelApp
	UserId    *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:19;"`          //用户id
	Time      *time.Time `json:"time" form:"time" gorm:"column:time;comment:点赞时间;"`                         //点赞时间
	CommentId *int       `json:"commentId" form:"commentId" gorm:"column:comment_id;comment:评论编号;size:19;"` //评论编号
}

// TableName CommentThumbsUp 表名
func (CommentThumbsUp) TableName() string {
	return "hk_comment_thumbs_up"
}
