package config

// CacheCaptcha
type CacheCaptcha struct {
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 缓存过期时间
	SmsTime     string `mapstructure:"sms-time" json:"sms-time" yaml:"sms-time"`             // 短信有效时间
}
