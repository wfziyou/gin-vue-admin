// 自动生成模板UserRecharge
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ExtractType 结构体
type ExtractType struct {
	global.GvaModelAppEx
	Icon string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                                          //图标
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:40;"`                                           //名称
	Code string `json:"code" form:"code" gorm:"column:code;comment:标识: h5=网页 alipay = 支付宝 wx=微信 paypal apple=苹果;size:20;"` //标识: h5=网页 alipay = 支付宝 wx=微信 paypal apple=苹果
}

// TableName UserRecharge 表名
func (ExtractType) TableName() string {
	return "hk_extract_type"
}
