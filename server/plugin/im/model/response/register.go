package response

type RegisterRsp struct {
	Name  string `json:"name"`
	Accid string `json:"accid"`
	Token string `json:"token"`
}
