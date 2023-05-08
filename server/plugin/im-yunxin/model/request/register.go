package request

type RegisterReq struct {
	Accid  string `json:"accid"`  //32	必选	"123456"	云信账号，必须保证唯一性。若涉及字母，传参时请一律小写处理。只允许字母、数字、半角下划线_、@、半角点以及半角-。请注意以此接口返回结果中的accid为准。
	Token  string `json:"token"`  //128	选填	"abcdef"	用户账号对应的登录密钥token。如果未指定，云信会自动生成token，并在创建成功后返回。
	Name   string `json:"name"`   //64	选填	"zhangsan"	用户昵称
	Icon   string `json:"icon"`   //1024	选填	"https://netease/xxx.png"	用户头像 URL
	Sign   string `json:"sign"`   //256	选填	"Hello World"	用户签名
	Email  string `json:"email"`  //64	选填	"xxx@163.com"	用户邮箱地址
	Birth  string `json:"birth"`  //16	选填	"xxxx-xx-xx"	用户生日
	Mobile string `json:"mobile"` //32	选填	"+852-xxxxxxxx"	用户手机号码，非中国大陆手机号码需要填写国家代码(如美国：+1-xxxxxxxxxx)或地区代码(如香港：+852-xxxxxxxx)
	Gender int    `json:"gender"` //选填	2	用户性别，0-未知，1-男，2-女。其它会报参数错误。
	Ex     string `json:"ex"`     //1024	选填	{"level":1}	用户资料扩展字段，建议封装成JSON。
}
