package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/app/community"

type ThirdPlatInfo struct {
	Platform string `json:"platform"` // 平台
	Name     string `json:"name"`     // 平台名称
	Icon     string `json:"icon"`     // 平台图标
	Appid    string `json:"appid"`    // 平台的应用编号
}

type LoginResponse struct {
	User      community.User `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type RegisterResponse struct {
	User community.User `json:"user"`
}

type LocalTelephone struct {
	Phone string `json:"phone" form:"phone" gorm:"column:phone;comment:手机;size:45;"` //手机
}
