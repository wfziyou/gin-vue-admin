package request

type ParamCreateOrder struct {
	ProductId uint64 `json:"productId"`              //商品ID
	CartNum   int    `json:"cartNum" form:"cartNum"` //商品数量
	PayType   string `json:"payType" form:"payType"` //支付方式
}

type ParamUserExtract struct {
	Code uint64 `json:"code"` //标识: bank = 银行卡 alipay = 支付宝wx=微信
	Num  uint64 `json:"num"`  //数量
	Bank string `json:"bank"` //银行卡
}
