package response

type GetSelfUserInfoResp struct {
	CommResp
	UserInfo *UserInfo              `json:"-"`
	Data     map[string]interface{} `json:"data" swaggerignore:"true"`
}
