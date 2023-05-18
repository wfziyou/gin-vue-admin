// 自动生成模板UserRecharge
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PayType 结构体
type PayType struct {
	global.GvaModelApp
	Icon         string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                                          //图标
	Name         string `json:"name" form:"name" gorm:"column:name;comment:名称;size:40;"`                                           //名称
	Code         string `json:"code" form:"code" gorm:"column:code;comment:标识: h5=网页 alipay = 支付宝 wx=微信 paypal apple=苹果;size:20;"` //标识: h5=网页 alipay = 支付宝 wx=微信 paypal apple=苹果
	PlatAppid    string `json:"platAppid" form:"platAppid" gorm:"column:plat_appid;comment:支付平台appid;size:80;"`                    //支付平台appid
	PlatKey      string `json:"platKey" form:"platKey" gorm:"column:plat_key;comment:支付平台key;size:80;"`                            //支付平台key
	PlatSecurity string `json:"platSecurity" form:"platSecurity" gorm:"column:plat_security;comment:支付平台security;size:80;"`        //支付平台security
}

// TableName UserRecharge 表名
func (PayType) TableName() string {
	return "hk_pay_type"
}

type PayTypeInfo struct {
	global.GvaModelAppEx
	Icon string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                                          //图标
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:40;"`                                           //名称
	Code string `json:"code" form:"code" gorm:"column:code;comment:标识: h5=网页 alipay = 支付宝 wx=微信 paypal apple=苹果;size:20;"` //标识: h5=网页 alipay = 支付宝 wx=微信 paypal apple=苹果
}
