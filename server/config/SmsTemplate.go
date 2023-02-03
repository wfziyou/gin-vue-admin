package config

type SmsTemplate struct {
	Test      string `mapstructure:"test" json:"test" yaml:"test"`                // 测试
	Register  string `mapstructure:"register" json:"register" yaml:"register"`    // 注册
	ChangePwd string `mapstructure:"changePwd" json:"changePwd" yaml:"changePwd"` // 修改密码
	BindTel   string `mapstructure:"bindTel" json:"bindTel" yaml:"bindTel"`       // 绑定电话
	ResetPwd  string `mapstructure:"resetPwd" json:"resetPwd" yaml:"resetPwd"`    // 重置密码
	BindBank  string `mapstructure:"bindBank" json:"bindBank" yaml:"bindBank"`    // 绑定银行
}
