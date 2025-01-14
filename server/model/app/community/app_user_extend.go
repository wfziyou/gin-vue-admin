// 自动生成模板UserExtend
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// UserExtend 结构体
type UserExtend struct {
	global.GvaModelApp
	Github               string     `json:"github" form:"github" gorm:"column:github;comment:github;size:64;"`                              //github
	Weibo                string     `json:"weibo" form:"weibo" gorm:"column:weibo;comment:微博;size:32;"`                                     //微博
	Weixin               string     `json:"weixin" form:"weixin" gorm:"column:weixin;comment:微信;size:32;"`                                  //微信
	Qq                   string     `json:"qq" form:"qq" gorm:"column:qq;comment:qq;size:32;"`                                              //qq
	Blog                 string     `json:"blog" form:"blog" gorm:"column:blog;comment:博客;size:500;"`                                       //博客
	CoverImage           string     `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:主页封面;size:256;"`                  //主页封面
	NumCircle            int64      `json:"numCircle" form:"numCircle" gorm:"type:bigint(20);column:num_circle;comment:圈子数;"`               //圈子数
	NumFocus             int64      `json:"numFocus" form:"numFocus" gorm:"type:bigint(20);column:num_focus;comment:关注数;"`                  //关注数
	NumFan               int64      `json:"numFan" form:"numFan" gorm:"type:bigint(20);column:num_fan;comment:粉丝数;"`                        //粉丝数
	ChannelId            string     `json:"channelId" form:"channelId" gorm:"column:channel_id;comment:频道_编号;"`                             //频道_编号
	CurrencyMoney        uint64     `json:"currencyMoney" form:"currencyMoney" gorm:"type:bigint(20);column:currency_money;comment:货币_零钱;"` //货币_零钱
	CurrencyGold         uint64     `json:"currencyGold" form:"currencyGold" gorm:"type:bigint(20);column:currency_gold;comment:货币_金币;"`    //货币_金币
	UpdateForumPostsTime *time.Time `json:"-" gorm:"column:update_forum_posts_time;comment:发布帖子时间;"`                                        //发布帖子时间
}

// TableName UserExtend 表名
func (UserExtend) TableName() string {
	return "hk_user_extend"
}
