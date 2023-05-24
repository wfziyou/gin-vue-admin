package general

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type ConfigParam struct {
	MiniProgram *MiniProgramBaseInfo `json:"miniProgram" yaml:"miniProgram"` //小程序
	ParamList   []system.SysParam    `json:"paramList" yaml:"paramList"`     //参数参数
}
