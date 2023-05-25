package request

type ParamCheckAppUpdate struct {
	Platform   int    `json:"platform"`   // 平台：1IOS,2Android,3Windows,4OSX,5Web,6MiniWeb,7Linux,8APad,9IPad,10Admin
	CurVersion string `json:"curVersion"` //当前版本
}
