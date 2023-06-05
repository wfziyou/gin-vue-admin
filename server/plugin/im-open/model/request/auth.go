package request

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open/model"

type RegisterReq struct {
	Secret   string `json:"secret" binding:"required,max=32"`
	Platform int    `json:"platform" binding:"required,min=1,max=12"`
	model.ApiUserInfo
	OperationID string `json:"operationID" binding:"required"`
}
type UpdateSelfUserInfoReq struct {
	model.ApiUserInfo
	OperationID string `json:"operationID" binding:"required"`
}
type UserTokenReq struct {
	Secret      string `json:"secret" binding:"required,max=32"`
	Platform    int    `json:"platform" binding:"required,min=1,max=12"`
	UserID      string `json:"userID" binding:"required,min=1,max=64"`
	LoginIp     string `json:"loginIp"`
	OperationID string `json:"operationID" binding:"required"`
}

type ParseTokenReq struct {
	OperationID string `json:"operationID" binding:"required"`
}
