// 自动生成模板ForumThumbsUp
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ForumThumbsUp 结构体
type ForumThumbsUp struct {
	global.GvaModelApp
	UserId  *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:19;"`    //用户id
	Time    *time.Time `json:"time" form:"time" gorm:"column:time;comment:点赞时间;"`                   //点赞时间
	PostsId *int       `json:"postsId" form:"postsId" gorm:"column:posts_id;comment:帖子编号;size:19;"` //帖子编号
}

// TableName ForumThumbsUp 表名
func (ForumThumbsUp) TableName() string {
	return "hk_forum_thumbs_up"
}
