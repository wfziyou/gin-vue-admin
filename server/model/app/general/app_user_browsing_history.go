// 自动生成模板UserBrowsingHistory
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"time"
)

// UserBrowsingHistory 结构体
type UserBrowsingHistory struct {
	global.GvaModelApp
	UserId   uint64     `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户编号;"`                        //用户编号
	PostsId  uint64     `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"`                     //帖子编号
	Category int        `json:"category" form:"category" gorm:"column:category;comment:类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动;size:10;"` //类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动
	Time     *time.Time `json:"time" form:"time" gorm:"column:time;comment:浏览时间;"`                                               //浏览时间
}

// TableName UserBrowsingHistory 表名
func (UserBrowsingHistory) TableName() string {
	return "hk_user_browsing_history"
}

type UserBrowsingHistoryInfo struct {
	ID        uint64                        `json:"id" form:"id" gorm:"primarykey"`                                                                  // 主键ID
	PostsId   uint64                        `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"`                     //帖子编号
	Category  int                           `json:"category" form:"category" gorm:"column:category;comment:类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动;size:10;"` //类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动
	Time      *time.Time                    `json:"time" form:"time" gorm:"column:time;comment:浏览时间;"`
	PostsInfo *community.ForumPostsBaseInfo `json:"postsInfo" gorm:"foreignKey:ID;references:PostsId;comment:帖子信息"`
}

func (UserBrowsingHistoryInfo) TableName() string {
	return "hk_user_browsing_history"
}
