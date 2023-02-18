// 自动生成模板Apply
package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Apply 结构体
type Apply struct {
	global.GvaModelApp
	OwerType        int    `json:"owerType" form:"owerType" gorm:"column:ower_type;comment:拥有者：0平台、1圈子、2个人;size:10;"` //拥有者：0平台、1圈子、2个人
	CircleId        uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`   //圈子_编号
	UserId          uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户_编号;"`         //用户_编号
	Name            string `json:"name" form:"name" gorm:"column:name;comment:名称;size:80;"`                           //名称
	Icon            string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                          //图标
	Sort            int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                           //排序
	Type            int    `json:"type" form:"type" gorm:"column:type;comment:类型(0小程序、1第三方链接);size:10;"`
	MiniProgramId   uint64 `json:"miniProgramId" form:"miniProgramId" gorm:"type:bigint(20);column:mini_program_id;comment:小程序id"` //类型(0小程序、1第三方链接)
	ApplyAddress    string `json:"applyAddress" form:"applyAddress" gorm:"column:apply_address;comment:访问地址;size:256;"`            //访问地址
	ApplyParameters string `json:"applyParameters" form:"applyParameters" gorm:"column:apply_parameters;comment:访问参数;size:256;"`   //访问参数
}

// TableName Apply 表名
func (Apply) TableName() string {
	return "hk_apply"
}
