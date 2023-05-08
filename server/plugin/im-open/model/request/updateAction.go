package request

type UpdateActionReq struct {
	Accid string `json:"accid"` //32	必选	"123456"	云信账号，必须保证唯一性。若涉及字母，传参时请一律小写处理。只允许字母、数字、半角下划线_、@、半角点以及半角-。请注意以此接口返回结果中的accid为准。
	Token string `json:"token"` //128	选填	"abcdef"	用户账号对应的登录密钥token。如果未指定，云信会自动生成token，并在创建成功后返回。
}
