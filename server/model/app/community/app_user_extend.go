// 自动生成模板UserExtend
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserExtend 结构体
type UserExtend struct {
	global.GvaModelApp

	CircleId   uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:用户ID;"`   //当前圈子编号
	CircleName string `json:"circleName" `                                                                      //当前圈子名称
	Github     string `json:"github" form:"github" gorm:"column:github;comment:github;size:64;"`                //github
	Weibo      string `json:"weibo" form:"weibo" gorm:"column:weibo;comment:微博;size:32;"`                       //微博
	Weixin     string `json:"weixin" form:"weixin" gorm:"column:weixin;comment:微信;size:32;"`                    //微信
	Qq         string `json:"qq" form:"qq" gorm:"column:qq;comment:qq;size:32;"`                                //qq
	Blog       string `json:"blog" form:"blog" gorm:"column:blog;comment:博客;size:500;"`                         //博客
	NumCircle  uint64 `json:"numCircle" form:"numCircle" gorm:"type:bigint(20);column:num_circle;comment:圈子数;"` //圈子数
	NumFocus   uint64 `json:"numFocus" form:"numFocus" gorm:"type:bigint(20);column:num_focus;comment:关注数;"`    //关注数
	NumFan     uint64 `json:"numFan" form:"numFan" gorm:"type:bigint(20);column:num_fan;comment:粉丝数;"`          //粉丝数
}

// TableName UserExtend 表名
func (UserExtend) TableName() string {
	return "hk_user_extend"
}
