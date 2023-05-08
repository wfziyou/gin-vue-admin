package response

type RegisterRsp struct {
	Code int             `json:"code"`
	Info RegisterRspInfo `json:"info"`
	Desc string          `json:"desc"`
}

type RegisterRspInfo struct {
	Name  string `json:"name"`
	Accid string `json:"accid"`
	Token string `json:"token"`
}
