package common

import "time"

// UserBaseInfo 用户基本信息
type UserBaseInfo struct {
	ID        uint64     `json:"id" form:"id" gorm:"primarykey"`            // 主键ID
	NickName  string     `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	Phone     string     `json:"phone"  gorm:"comment:用户手机号"`               // 用户手机号
	Email     string     `json:"email"  gorm:"comment:用户邮箱"`                // 用户邮箱
	HeaderImg string     `json:"headerImg" gorm:"comment:用户头像"`             // 用户头像
	Birthday  *time.Time `json:"birthday" gorm:"comment:生日"`                //生日
	Sex       *int       `json:"sex" gorm:"comment:性别 1男 2女"`               //性别 1男 2女
}

type CacheCaptcha struct {
	Code     string    `json:"code"`     // 验证码
	Overtime time.Time `json:"overtime"` //过期时间
}
