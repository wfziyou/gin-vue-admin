package config

type Param struct {
	CreateCircleCheck   bool   `mapstructure:"createCircleCheck" json:"createCircleCheck" yaml:"createCircleCheck"`       // 创建圈子是否需要审批
	UseSmsCheckCode     bool   `mapstructure:"useSmsCheckCode" json:"useSmsCheckCode" yaml:"useSmsCheckCode"`             // 使用短信验证码
	DefaultSmsCheckCode string `mapstructure:"defaultSmsCheckCode" json:"defaultSmsCheckCode" yaml:"defaultSmsCheckCode"` // 默认短信验证码

}
