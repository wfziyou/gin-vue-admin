package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type ConfigParam struct {
	MiniProgram *MiniProgramBaseInfo `json:"miniProgram" yaml:"miniProgram"` //小程序
	ParamList   []system.SysParam    `json:"paramList" yaml:"paramList"`     //参数参数
}

type ResponseCheckAppUpdate struct {
	Version     AppVersion `json:"version"`      //最新版本
	ForceUpdate int        `json:"forceUpdate" ` //强制更新：0否 1是
}
