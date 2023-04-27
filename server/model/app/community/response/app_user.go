package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
)

type UserBaseInfoResponse struct {
	User community.UserBaseInfo `json:"user"`
}
type SysUserResponse struct {
	User community.User `json:"user"`
}
