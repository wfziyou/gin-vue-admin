// 自动生成模板Order
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

//订单状态:0待发货 1已完成 2申请退款 3已退货 4已退款 5拒绝退款
const (
	//OrderStatusWait 0待发货
	OrderStatusWait = 0
	//OrderStatusFinish 1已完成
	OrderStatusFinish = 1
	//OrderStatusRefund 2申请退款
	OrderStatusRefund = 2
	//OrderStatusReturned 3已退货
	OrderStatusReturned = 3
	//OrderStatusRefunded 4已退款
	OrderStatusRefunded = 4
	//OrderStatusRefusalRefund 5拒绝退款
	OrderStatusRefusalRefund = 5
)

//支付状态: 0未支付 1已支付
const (
	OrderPaidFalse = 0
	OrderPaidTrue  = 1
)

// Order 结构体
type Order struct {
	global.GvaModelApp
	OrderId                string     `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单号;size:36;"`
	ExtendOrderId          string     `json:"extendOrderId" form:"extendOrderId" gorm:"column:extend_order_id;comment:额外订单号;size:36;"`
	UserId                 uint64     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
	ProductId              uint64     `json:"productId" form:"productId" gorm:"column:product_id;comment:商品ID;size:19;"`
	CartNum                int        `json:"cartNum" form:"cartNum" gorm:"column:cart_num;comment:商品数量;"`
	TotalPrice             uint64     `json:"totalPrice" form:"totalPrice" gorm:"column:total_price;comment:订单总价;size:19;"`
	PayPrice               uint64     `json:"payPrice" form:"payPrice" gorm:"column:pay_price;comment:实际支付金额;size:19;"`
	Cost                   uint64     `json:"cost" form:"cost" gorm:"column:cost;comment:成本价;size:19;"`
	Paid                   int        `json:"paid" form:"paid" gorm:"column:paid;comment:支付状态: 0未支付 1已支付;"`
	PayTime                *time.Time `json:"payTime" form:"payTime" gorm:"column:pay_time;comment:支付时间;"`
	PayType                string     `json:"payType" form:"payType" gorm:"column:pay_type;comment:支付方式;size:32;"`
	RefundStatus           int        `json:"refundStatus" form:"refundStatus" gorm:"column:refund_status;comment:0 未退款 1 申请中 2 已退款;"`
	RefundReasonWapImg     string     `json:"refundReasonWapImg" form:"refundReasonWapImg" gorm:"column:refund_reason_wap_img;comment:退款图片;size:255;"`
	RefundReasonWapExplain string     `json:"refundReasonWapExplain" form:"refundReasonWapExplain" gorm:"column:refund_reason_wap_explain;comment:退款用户说明;size:255;"`
	RefundReasonTime       *time.Time `json:"refundReasonTime" form:"refundReasonTime" gorm:"column:refund_reason_time;comment:退款时间;"`
	RefundReasonWap        string     `json:"refundReasonWap" form:"refundReasonWap" gorm:"column:refund_reason_wap;comment:前台退款原因;size:255;"`
	RefundReason           string     `json:"refundReason" form:"refundReason" gorm:"column:refund_reason;comment:不退款的理由;size:255;"`
	RefundPrice            uint64     `json:"refundPrice" form:"refundPrice" gorm:"column:refund_price;comment:退款金额;size:19;"`
	GainIntegral           uint64     `json:"gainIntegral" form:"gainIntegral" gorm:"column:gain_integral;comment:消费赚取积分;size:19;"`
	BackIntegral           uint64     `json:"backIntegral" form:"backIntegral" gorm:"column:back_integral;comment:给用户退了多少积分;size:19;"`
	Mark                   string     `json:"mark" form:"mark" gorm:"column:mark;comment:备注;size:512;"`
	Remark                 string     `json:"remark" form:"remark" gorm:"column:remark;comment:管理员备注;size:512;"`
	VerifyCode             string     `json:"verifyCode" form:"verifyCode" gorm:"column:verify_code;comment:核销码;size:50;"`
}

// TableName Order 表名
func (Order) TableName() string {
	return "hk_order"
}
