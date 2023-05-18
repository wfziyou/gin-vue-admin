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

type TokenValidateRequest struct {
	openType      string `json:"openType"`      //否 运营商类型：1：移动;2：联通;3：电信;0：未知
	requesterType string `json:"requesterType"` //是 请求方类型：0：APP；1：WAP
	message       string `json:"message"`       //否 接入方预留参数，该参数会透传给通知接口，此参数需 urlencode 编码
	expandParams  string `json:"expandParams"`  //否 扩展参数格式：param1=value1|param2=value2 方式传递，参数以竖线 | 间隔方式传递，此参数需 urlencode 编码。
	//加密算法类型，其中，
	//1）keyType=0 时，加密类型为 SHA；
	//2）keyType=1 时，加密类型为 RSA；
	//3）keyType=2 时，加密类型为国密
	//算法；受影响的字段包括 phoneNum、 sign 和 respSign，具体算法可阅读 对应的字段说明。（注：keyType 不 等于其它值或空时，按 keyType=0 处理）
	keyType string `json:"keyType"` //否
	//加密的手机号码，其中，
	//1）keyType=0 时，使用社区生成的appkey 对手机号码进行摘要处理，加密算法是 SHA256，字母大写，参数组装格式为（待校验的手机号码 +appKey+ timestamp）。
	//2）keyType=1 时，使用社区提供的RSA 公钥（即“平台公钥”）对手机号码进行加密，加密算法是 RSA，参数组装格式为（待校验的手机号码 +appKey+ timestamp）
	//3）keyType=2 时，使用社区提供的国密签名公钥作为 SM3 哈希加密的key，加密算法是 SM3，参数组装格式为（待校验的手机号码 + appKey+timestamp）。
	//具体签名算法可参考 2.5 提供的算法示例（注：建议开发者对用户输入的手机号码的格式进行校验，增加校验通过的概率）
	phoneNum string `json:"phoneNum"` //是
	Token    string `json:"token"`    //是 业务凭证
	//1）keyType=0 时，使用社区生成的appkey 作为秘钥进行签名，加密算法是 HMACSHA256，输出 64 位大写字母。参数名做自然排序，参数组装格式为(appId + msgId + phoneNum +timestamp + token + version)。
	//2）keyType=1 时，使用社区中的客户公钥所对应的私钥加签，加密算法是 RSA，参数组装格式为(appId +msgId + phoneNum + timestamp +token + version)。
	//3）keyType=2 时，使用开发者社区中的国密签名公钥所对应的私钥加签，加密算法是 SM2，参数组装格式为(appId + msgId + phoneNum +timestamp + token + version)。具体签名算法可参考 2.5 提供的算法示例
	Sign string `json:"sign"` //是
}
