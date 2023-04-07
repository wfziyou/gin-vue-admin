package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

// BindTelephoneReq 绑定手机
type BindTelephoneReq struct {
	Telephone string `json:"telephone"` // 电话号码
	Captcha   string `json:"captcha"`   // 验证码
}

// BindEmailReq 绑定邮箱
type BindEmailReq struct {
	Email   string `json:"email"`   // 邮箱
	Captcha string `json:"captcha"` // 验证码
}

// SetSelfBaseInfoReq 设置用户基本信息
type SetSelfBaseInfoReq struct {
	ID          uint64 `json:"-" gorm:"primarykey"`                                                          // 主键ID
	NickName    string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                    // 用户昵称
	HeaderImg   string `json:"headerImg" gorm:"comment:用户头像"`                                                // 用户头像
	Birthday    string `json:"birthday" gorm:"comment:生日"`                                                   //生日
	Sex         *int   `json:"sex" gorm:"comment:性别： 0未知、1男、2女"`                                             //性别： 0未知、1男、2女
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述;size:45;"` //描述
}

type UserSearch struct {
	Account  string `json:"account" gorm:"comment:账号"`        //账号
	NickName string `json:"nickName" gorm:"comment:用户昵称"`     // 用户昵称
	Sex      *int   `json:"sex" gorm:"comment:性别： 0未知、1男、2女"` //性别： 0未知、1男、2女

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
