package config

type EmailTemplate struct {
	BindEmailSubject string `mapstructure:"bindEmailSubject" json:"bindEmailSubject" yaml:"bindEmailSubject"` // 绑定邮箱
	BindEmailBody    string `mapstructure:"bindEmailBody" json:"bindEmailBody" yaml:"bindEmailBody"`          // 绑定邮箱
}
