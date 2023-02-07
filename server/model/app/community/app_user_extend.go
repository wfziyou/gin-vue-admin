// 自动生成模板HkUserExtend
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkUserExtend 结构体
type HkUserExtend struct {
	global.GvaModelApp
	UserId    *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:19;"`         //用户ID
	Github    string `json:"github" form:"github" gorm:"column:github;comment:github;size:64;"`        //github
	Weibo     string `json:"weibo" form:"weibo" gorm:"column:weibo;comment:微博;size:32;"`               //微博
	Weixin    string `json:"weixin" form:"weixin" gorm:"column:weixin;comment:微信;size:32;"`            //微信
	Qq        string `json:"qq" form:"qq" gorm:"column:qq;comment:qq;size:32;"`                        //qq
	Blog      string `json:"blog" form:"blog" gorm:"column:blog;comment:博客;size:500;"`                 //博客
	NumCircle *int   `json:"numCircle" form:"numCircle" gorm:"column:num_circle;comment:圈子数;size:19;"` //圈子数
	NumFocus  *int   `json:"numFocus" form:"numFocus" gorm:"column:num_focus;comment:关注数;size:19;"`    //关注数
	NumFan    *int   `json:"numFan" form:"numFan" gorm:"column:num_fan;comment:粉丝数;size:19;"`          //粉丝数
}

// TableName HkUserExtend 表名
func (HkUserExtend) TableName() string {
	return "hk_user_extend"
}
