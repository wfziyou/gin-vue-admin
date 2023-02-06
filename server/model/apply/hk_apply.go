// 自动生成模板HkApply
package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkApply 结构体
type HkApply struct {
	global.GVA_MODEL
	OwerType        *int   `json:"owerType" form:"owerType" gorm:"column:ower_type;comment:拥有者：0平台、1圈子、2个人;size:10;"` //拥有者：0平台、1圈子、2个人
	CircleId        *int   `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`           //圈子_编号
	UserId          *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户_编号;size:19;"`                 //用户_编号
	Name            string `json:"name" form:"name" gorm:"column:name;comment:名称;size:80;"`                           //名称
	Icon            string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                          //图标
	Sort            *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                           //排序
	Type            *int   `json:"type" form:"type" gorm:"column:type;comment:类型(0小程序、1第三方链接);size:10;"`
	MiniProgramId   *int   `json:"miniProgramId" form:"miniProgramId" gorm:"column:mini_program_id;comment:小程序id;size:19;"`      //类型(0小程序、1第三方链接)
	ApplyAddress    string `json:"applyAddress" form:"applyAddress" gorm:"column:apply_address;comment:访问地址;size:256;"`          //访问地址
	ApplyParameters string `json:"applyParameters" form:"applyParameters" gorm:"column:apply_parameters;comment:访问参数;size:256;"` //访问参数
}

// TableName HkApply 表名
func (HkApply) TableName() string {
	return "hk_apply"
}
