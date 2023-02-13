// 自动生成模板UserBrowsingHistory
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// UserBrowsingHistory 结构体
type UserBrowsingHistory struct {
	global.GvaModelApp
	UserId  uint64     `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户编号;"`    //用户编号
	PostsId uint64     `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"` //帖子编号
	Time    *time.Time `json:"time" form:"time" gorm:"column:time;comment:浏览时间;"`                           //浏览时间
}

// TableName UserBrowsingHistory 表名
func (UserBrowsingHistory) TableName() string {
	return "hk_user_browsing_history"
}
