// 自动生成模板UserCollect
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserCollect 结构体
type UserCollect struct {
	global.GvaModelApp
	UserId   uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户编号;"`                        //用户编号
	PostsId  uint64 `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"`                     //帖子编号
	Category int    `json:"category" form:"category" gorm:"column:category;comment:类别：1视频、2动态、3资讯、4公告、5文章、6问答、7建议;size:10;"` //类别：1视频、2动态、3资讯、4公告、5文章、6问答、7建议
}

// TableName UserCollect 表名
func (UserCollect) TableName() string {
	return "hk_user_collect"
}
