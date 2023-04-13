package model

type LoginTokenValidateRequest struct {
	Version             string `json:"version"`
	Msgid               string `json:"msgid"`
	Systemtime          string `json:"systemtime"`
	Strictcheck         string `json:"strictcheck"`
	Appid               string `json:"appid"`
	Token               string `json:"token"`
	Sign                string `json:"sign"`
	Encryptionalgorithm string `json:"encryptionalgorithm"`
}
type LoginTokenValidateRespone struct {
	Inresponseto string `json:"inresponseto"`
	Systemtime   string `json:"systemtime"`
	ResultCode   string `json:"resultCode"`
	Msisdn       string `json:"msisdn"`
	TaskId       string `json:"taskId"`
	Telephone    string
}
