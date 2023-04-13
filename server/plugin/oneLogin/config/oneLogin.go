package config

type OneLogin struct {
	StrictCheck   string `mapstructure:"strictcheck" json:"strictcheck" yaml:"strictcheck"`       //是 暂时填写"0"，填写“1”时，将对服务器 IP 白名单进行强校验（后续将强制要求 IP 强校验）
	Appid         string `mapstructure:"appid" json:"appid" yaml:"appid"`                         //是 业务在统一认证申请的应用id
	RsaPrivateKey string `mapstructure:"rsaPrivateKey" json:"rsaPrivateKey" yaml:"rsaPrivateKey"` //rsa 私有key
}
