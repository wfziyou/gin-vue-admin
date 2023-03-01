package response

type UserGetUinfosActionRsp struct {
	Code   int        `json:"code"`
	Uinfos []UserInfo `json:"uinfos"`
	Desc   string     `json:"desc"`
}

type UserInfo struct {
	Accid  string `json:"accid"`  ////账号
	Name   string `json:"name"`   ////昵称
	Icon   string `json:"icon"`   //头像
	Sign   string `json:"sign"`   //签名
	Email  string `json:"email"`  //邮箱
	Birth  string `json:"birth"`  //"2015-12-23",//生日
	Mobile string `json:"mobile"` //手机号
	Ex     string `json:"ex"`     //扩展信息
	Gender int    `json:"gender"` //性别：0表示未知，1表示男，2女表示女
	Valid  bool   `json:"valid"`  //账号是否有效
	Mute   bool   `json:"mute"`   //账号是否被禁言
}
