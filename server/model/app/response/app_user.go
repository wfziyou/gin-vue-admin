package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/common"

type UserBaseInfoResponse struct {
	User common.UserBaseInfo `json:"user"`
}
