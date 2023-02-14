// 自动生成模板UserCollect
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserCollect 结构体
type UserCollect struct {
	global.GvaModelApp
	UserId  uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户编号;"`    //用户编号
	PostsId uint64 `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"` //帖子编号
}

// TableName UserCollect 表名
func (UserCollect) TableName() string {
	return "hk_user_collect"
}
