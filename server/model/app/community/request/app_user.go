package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

// CaptchaReq  验证码请求
type CaptchaReq struct {
	Telephone string `json:"telephone"` // 电话号码
	Type      int    `json:"type" `     //类型：0 测试，1注册，2修改密码，3绑定电话，4忘记密码，5绑定银行
}

// Register User register structure
type Register struct {
	Account  string `json:"account" example:"用户名"` // 用户名
	Password string `json:"passWord" example:"密码"` // 密码
	NickName string `json:"nickName" example:"昵称"` // 昵称
}

// LoginPwd 登录（账号密码）
type LoginPwd struct {
	Account  string `json:"account"`  // 用户名
	Password string `json:"password"` // 密码
}

// LoginTelephone 登录（手机）
type LoginTelephone struct {
	Telephone string `json:"telephone"` // 电话号码
	Captcha   string `json:"captcha"`   // 验证码
}

// LoginThird 登录（第三方授权）
type LoginThird struct {
}

// ResetPasswordReq 重置密码
type ResetPasswordReq struct {
	Captcha  string `json:"captcha"`  // 验证码
	Password string `json:"password"` // 密码
}

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
	ID        uint64     `gorm:"primarykey"`                                // 主键ID
	NickName  string     `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	HeaderImg string     `json:"headerImg" gorm:"comment:用户头像"`             // 用户头像
	Birthday  *time.Time `json:"birthday" gorm:"comment:生日"`                //生日
	Sex       *int       `json:"sex" gorm:"comment:性别"`                     //性别
}

type UserSearch struct {
	Account  string `json:"account" gorm:"comment:账号"`    //账号
	NickName string `json:"nickName" gorm:"comment:用户昵称"` // 用户昵称
	Sex      *int   `json:"sex" gorm:"comment:性别"`        //性别

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
