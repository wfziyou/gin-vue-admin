// 自动生成模板HkUserCollect
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkUserCollect 结构体
type HkUserCollect struct {
	global.GvaModelApp
	UserId  *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`    //用户编号
	PostsId *int       `json:"postsId" form:"postsId" gorm:"column:posts_id;comment:帖子编号;size:19;"` //帖子编号
	Time    *time.Time `json:"time" form:"time" gorm:"column:time;comment:收藏时间;"`                   //收藏时间
}

// TableName HkUserCollect 表名
func (HkUserCollect) TableName() string {
	return "hk_user_collect"
}
