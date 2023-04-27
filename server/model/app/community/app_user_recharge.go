// 自动生成模板UserRecharge
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// UserRecharge 结构体
type UserRecharge struct {
	global.GVA_MODEL
	UserId       *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
	Nickname     string     `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;size:50;"`
	OrderId      string     `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单号;size:32;"`
	Price        *float64   `json:"price" form:"price" gorm:"column:price;comment:充值金额;size:8;"`
	GivePrice    *float64   `json:"givePrice" form:"givePrice" gorm:"column:give_price;comment:购买赠送金额;size:8;"`
	RechargeType string     `json:"rechargeType" form:"rechargeType" gorm:"column:recharge_type;comment:充值类型;size:32;"`
	Paid         *bool      `json:"paid" form:"paid" gorm:"column:paid;comment:是否充值;"`
	PayTime      *time.Time `json:"payTime" form:"payTime" gorm:"column:pay_time;comment:充值支付时间;"`
	RefundPrice  *float64   `json:"refundPrice" form:"refundPrice" gorm:"column:refund_price;comment:退款金额;size:10;"`
	CreateUser   *int       `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept   *int       `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser   *int       `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	Status       *int       `json:"status" form:"status" gorm:"column:status;comment:状态:0带确定 1有效 -1无效;size:10;"`
	IsDel        *int       `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}

// TableName UserRecharge 表名
func (UserRecharge) TableName() string {
	return "hk_user_recharge"
}
