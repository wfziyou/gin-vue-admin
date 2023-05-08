package request

type UserGetUinfosActionReq struct {
	Accids []string `json:"accids"` //用户帐号（例如：JSONArray对应的accid串：\["zhangsan"\]，如果解析出错，会报414） 一次最多查询200个
}
