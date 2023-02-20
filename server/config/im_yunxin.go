package config

type YunXinIm struct {
	AppKey string `mapstructure:"appKey" json:"appKey" yaml:"appKey"` // 云信IM AppKey
	Url    string `mapstructure:"url" json:"url" yaml:"url"`          // 云信IM 地址
}
