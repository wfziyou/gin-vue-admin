package request

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open/model"

type RegisterReq struct {
	Secret   string `json:"secret" binding:"required,max=32"`
	Platform int32  `json:"platform" binding:"required,min=1,max=12"`
	model.ApiUserInfo
	OperationID string `json:"operationID" binding:"required"`
}
