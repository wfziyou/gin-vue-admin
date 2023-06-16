package request

type GetOpenScreenAdvertisingReq struct {
	Platform   int    `json:"platform"`                     // 平台：1IOS,2Android,3Windows,4OSX,5Web,6MiniWeb,7Linux,8APad,9IPad,10Admin
	Address    string `json:"address" form:"address"`       //当前位置
	OperatorId uint64 `json:"operatorId" form:"operatorId"` //操作者编号（上次登录的用户编号，没有则不赋值）
}
