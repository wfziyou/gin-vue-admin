package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
)

type UserBaseInfoResponse struct {
	User common.UserBaseInfo `json:"user"`
}
type SysUserResponse struct {
	User community.User `json:"user"`
}
