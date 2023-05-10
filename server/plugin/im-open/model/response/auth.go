package response

type UserTokenInfo struct {
	UserID      string `json:"userID"`
	Token       string `json:"token"`
	ExpiredTime int64  `json:"expiredTime"`
}
type UserRegisterResp struct {
	CommResp
	UserToken UserTokenInfo `json:"data"`
}

type UserTokenResp struct {
	CommResp
	UserToken UserTokenInfo `json:"data"`
}
type ExpireTime struct {
	ExpireTimeSeconds uint32 `json:"expireTimeSeconds" `
}
type ParseTokenResp struct {
	CommResp
	Data       map[string]interface{} `json:"data" swaggerignore:"true"`
	ExpireTime ExpireTime             `json:"-"`
}
