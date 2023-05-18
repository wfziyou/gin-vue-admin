// 自动生成模板UserBill
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserBill 结构体
type UserBill struct {
	global.GvaModelApp
	UserId   uint64 `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
	LinkId   string `json:"linkId" form:"linkId" gorm:"column:link_id;comment:关联id;size:32;"`
	Pm       int    `json:"pm" form:"pm" gorm:"column:pm;comment:0 = 支出 1 = 获得;"`
	Title    string `json:"title" form:"title" gorm:"column:title;comment:账单标题;size:64;"`
	Category string `json:"category" form:"category" gorm:"column:category;comment:明细种类;size:64;"`
	Type     string `json:"type" form:"type" gorm:"column:type;comment:账单类型 1收入 2创作收益 3提现 4消费 5其他;size:64;"`
	Number   uint64 `json:"number" form:"number" gorm:"column:number;comment:明细数字;size:8;"`
	Balance  uint64 `json:"balance" form:"balance" gorm:"column:balance;comment:剩余;size:8;"`
	Mark     string `json:"mark" form:"mark" gorm:"column:mark;comment:备注;size:512;"`
}

// TableName UserBill 表名
func (UserBill) TableName() string {
	return "hk_user_bill"
}
