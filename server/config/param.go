package config

type Param struct {
	CreateCircleCheck bool `mapstructure:"createCircleCheck" json:"createCircleCheck" yaml:"createCircleCheck"` // 创建圈子是否需要审核

}
