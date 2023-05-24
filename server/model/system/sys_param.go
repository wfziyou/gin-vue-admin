package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysParam struct {
	global.GvaModelAppEx
	ParamName  string `json:"paramName" gorm:"column:param_name;comment:参数名"`   // 参数名
	ParamKey   string `json:"paramKey" gorm:"column:param_key;comment:参数键"`     // 参数键
	ParamValue string `json:"paramValue" gorm:"column:param_value;comment:参数值"` // 参数值
	Remark     string `json:"remark" gorm:"column:remark;comment:备注"`           // 备注
}

func (SysParam) TableName() string {
	return "blade_param"
}
