package config

type OpenIm struct {
	AppKey    string `mapstructure:"appKey" json:"appKey" yaml:"appKey"`          // 云信IM AppKey
	AppSecret string `mapstructure:"appSecret" json:"appSecret" yaml:"appSecret"` // 云信IM AppSecret
	Url       string `mapstructure:"url" json:"url" yaml:"url"`                   // 云信IM 地址
}
