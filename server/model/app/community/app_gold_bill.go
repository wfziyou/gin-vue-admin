// 自动生成模板GoldBill
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

const (
	GoldBillTypeIncrease = 0
	GoldBillTypeDecrease = 1
)

//账单类型 1充值 2付费 3红包 4打赏 5其他
const (
	//BillTypeRecharge 1充值
	BillTypeRecharge = 1
	//BillTypeBuy 2付费
	BillTypeBuy = 2
	//BillTypeRedEnvelope 3红包
	BillTypeRedEnvelope = 3
	//BillTypeReward 4打赏
	BillTypeReward = 4
	//BillTypeOther 5其他
	BillTypeOther = 5
)
const (
	//BillChildTypeDefault 0 默认
	BillChildTypeDefault = 0
	//BillChildTypeCreateQuestion 1发布问题
	BillChildTypeCreateQuestion = 1
	//BillChildTypeAddActivity 2加入活动
	BillChildTypeAddActivity = 2
	//BillChildTypeAnswerQuestion 3回答问题
	BillChildTypeAnswerQuestion = 3
)

// GoldBill 结构体
type GoldBill struct {
	global.GvaModelApp
	UserId       uint64 `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`                   //用户编号
	LinkId       uint64 `json:"linkId" form:"linkId" gorm:"column:link_id;comment:关联id;size:19;"`                   //关联id
	Pm           int    `json:"pm" form:"pm" gorm:"column:pm;comment:0 = 支出 1 = 获得;"`                               //0 = 支出 1 = 获得
	Title        string `json:"title" form:"title" gorm:"column:title;comment:账单标题;size:64;"`                       //账单标题
	Description  string `json:"description" form:"description" gorm:"column:description;comment:描述;size:45;"`       //描述
	Type         int    `json:"type" form:"type" gorm:"column:type;comment:账单类型 1充值 2付费 3红包 4打赏 5其他;size:10;"`      //账单类型 1充值 2付费 3红包 4打赏 5其他
	ChildType    int    `json:"childType" form:"childType" gorm:"column:child_type;comment:子类型：;size:10;"`          //子类型： 1充值 2付费 3红包 4打赏 5其他
	BeforeNumber uint64 `json:"beforeNumber" form:"beforeNumber" gorm:"column:before_number;comment:之前金额;size:20;"` //之前金额
	ChangeNumber uint64 `json:"changeNumber" form:"changeNumber" gorm:"column:change_number;comment:改变金额;size:20;"` //改变金额
	Balance      uint64 `json:"balance" form:"balance" gorm:"column:balance;comment:余额;size:20;"`                   //余额
	Mark         string `json:"mark" form:"mark" gorm:"column:mark;comment:备注;size:512;"`                           //备注
}

// TableName GoldBill 表名
func (GoldBill) TableName() string {
	return "hk_gold_bill"
}
