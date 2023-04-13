package request

// LoginPwd 登录（账号密码）
type LoginPwd struct {
	Account  string `json:"account"`  // 用户名
	Password string `json:"password"` // 密码
}

// LoginTelephone 登录（手机）
type LoginTelephone struct {
	Telephone string `json:"telephone"` // 电话号码
	Captcha   string `json:"captcha"`   // 验证码
}

// LoginThird 登录（第三方授权）
type LoginThird struct {
	Plat string `json:"plat"` //平台
	Code string `json:"code"` //平台code
}

// LoginOneClick 一键登录
type LoginOneClick struct {
	Token string `json:"token"` //令牌
}

// CaptchaReq  验证码请求
type CaptchaReq struct {
	Telephone string `json:"telephone" example:"12345678901"` // 电话号码
	Type      int    `json:"type" example:"0"`                //类型：0 测试，1注册，2修改密码，3绑定电话，4忘记密码，5绑定银行
}

// Register User register structure
type Register struct {
	Account  string `json:"account" example:"admin"`   // 用户名
	Password string `json:"passWord" example:"123456"` // 密码
	NickName string `json:"nickName" example:"昵称"`     // 昵称
}

// ResetPasswordReq 重置密码
type ResetPasswordReq struct {
	Captcha  string `json:"captcha" example:"666666"` // 验证码
	Password string `json:"password" example:"pwd"`   // 密码
}
