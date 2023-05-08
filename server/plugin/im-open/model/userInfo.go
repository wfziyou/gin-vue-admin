package model

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
type ApiUserInfo struct {
	UserID      string `json:"userID" binding:"required,min=1,max=64" swaggo:"true,用户ID,"`
	Nickname    string `json:"nickname" binding:"omitempty,min=1,max=64" swaggo:"true,my id,19"`
	FaceURL     string `json:"faceURL" binding:"omitempty,max=1024"`
	Gender      int    `json:"gender" binding:"omitempty,oneof=0 1 2"`
	PhoneNumber string `json:"phoneNumber" binding:"omitempty,max=32"`
	Birth       uint32 `json:"birth" binding:"omitempty"`
	Email       string `json:"email" binding:"omitempty,max=64"`
	CreateTime  int64  `json:"createTime"`
	LoginLimit  int    `json:"loginLimit" binding:"omitempty"`
	Ex          string `json:"ex" binding:"omitempty,max=1024"`
	BirthStr    string `json:"birthStr" binding:"omitempty"`
}
