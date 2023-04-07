package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/app/community"

type ThirdPlatInfo struct {
	Plat  string `json:"Plat"`
	Name  string `json:"name"`
	Appid string `json:"appid"`
	Icon  string `json:"icon"`
}

type LoginResponse struct {
	User      community.User `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type RegisterResponse struct {
	User community.User `json:"user"`
}
