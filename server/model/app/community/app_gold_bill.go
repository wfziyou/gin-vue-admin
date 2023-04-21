// 自动生成模板HkGoldBill
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

const (
	GoldBillTypeIncrease = 0
	GoldBillTypeDecrease = 1
)

// HkGoldBill 结构体
type HkGoldBill struct {
	global.GvaModelApp
	UserId       uint64 `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
	LinkId       string `json:"linkId" form:"linkId" gorm:"column:link_id;comment:关联id;size:32;"`
	Pm           int    `json:"pm" form:"pm" gorm:"column:pm;comment:0 = 支出 1 = 获得;"`
	Title        string `json:"title" form:"title" gorm:"column:title;comment:账单标题;size:64;"`
	TitleIcon    string `json:"titleIcon" form:"titleIcon" gorm:"column:title_icon;comment:标题图标;size:256;"`
	Type         string `json:"type" form:"type" gorm:"column:type;comment:类型;size:64;"`
	BeforeNumber uint64 `json:"beforeNumber" form:"beforeNumber" gorm:"column:before_number;comment:之前金额;size:20;"`
	ChangeNumber uint64 `json:"changeNumber" form:"changeNumber" gorm:"column:change_number;comment:改变金额;size:20;"`
	Balance      uint64 `json:"balance" form:"balance" gorm:"column:balance;comment:余额;size:20;"`
	Mark         string `json:"mark" form:"mark" gorm:"column:mark;comment:备注;size:512;"`
}

// TableName HkGoldBill 表名
func (HkGoldBill) TableName() string {
	return "hk_gold_bill"
}
