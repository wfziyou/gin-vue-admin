package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysApi struct {
	global.GVA_MODEL
	Path        string      `json:"path" gorm:"comment:api路径"`                     // api路径
	Description string      `json:"description" gorm:"comment:api中文描述"`            // api中文描述
	GroupId     uint        `json:"apiGroup" gorm:"column:api_group;comment:api组"` // api组
	Method      string      `json:"method" gorm:"default:POST;comment:方法"`         // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
	SysApiGroup SysApiGroup `json:"sysApiGroup" gorm:"foreignKey:ID;references:GroupId;comment:分组"`
}

func (SysApi) TableName() string {
	return "sys_apis"
}
