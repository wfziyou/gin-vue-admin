// 自动生成模板FocusUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

//type FocusUserInfo struct {
//	Id          uint64 `json:"id" form:"id" gorm:"not null;unique;primary_key;"` //编号
//	FocusUserId uint64 `json:"focusUserId" form:"focusUserId" gorm:"column:focus_user_id;comment:关注用户ID;size:19;"`
//	Remark      string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:40;"`                //备注
//	Tag         string `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:512;"`                        //标签
//	Account     string `json:"account" form:"account" gorm:"column:account;comment:账号;size:45;"`             //账号
//	NickName    string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:昵称;size:20;"`         //昵称
//	RealName    string `json:"realName" form:"realName" gorm:"column:real_name;comment:真名;size:10;"`         //真名
//	HeaderImg   string `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:头像;size:500;"`     //头像
//	Sex         int    `json:"sex" form:"sex" gorm:"column:sex;comment:性别： 0未知、1男、2女;size:10;"`              //性别： 0未知、1男、2女
//	Description string `json:"description" form:"description" gorm:"column:description;comment:描述;size:45;"` //描述
//}

// FocusUser 结构体
type FocusUser struct {
	global.GvaModelApp
	UserId        uint64 `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:19;"`                         //用户ID
	NickName      string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:20;"`                   //用户昵称
	Remark        string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:40;"`                            //备注
	Tag           string `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:512;"`                                    //标签
	FocusUserId   uint64 `json:"focusUserId" form:"focusUserId" gorm:"column:focus_user_id;comment:关注用户ID;size:19;"`       //关注用户ID
	FocusNickName string `json:"focusNickName" form:"focusNickName" gorm:"column:focus_nick_name;comment:关注用户昵称;size:20;"` //关注用户昵称
	Mutual        int    `json:"mutual" form:"mutual" gorm:"column:mutual;comment:是否相互关注:0否 1是:10;"`                       //是否相互关注:0否 1是
}

type FocusUserInfo struct {
	UserInfo
	Mutual int `json:"mutual" form:"mutual" gorm:"column:mutual;comment:是否相互关注:0否 1是:10;"` //是否相互关注:0否 1是
}

// TableName FocusUser 表名
func (FocusUser) TableName() string {
	return "hk_focus_user"
}
