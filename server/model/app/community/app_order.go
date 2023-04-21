// 自动生成模板HkOrder
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkOrder 结构体
type HkOrder struct {
      global.GVA_MODEL
      OrderId  string `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单号;size:32;"`
      ExtendOrderId  string `json:"extendOrderId" form:"extendOrderId" gorm:"column:extend_order_id;comment:额外订单号;size:32;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
      ProductId  *int `json:"productId" form:"productId" gorm:"column:product_id;comment:商品ID;size:20;"`
      CartNum  *int `json:"cartNum" form:"cartNum" gorm:"column:cart_num;comment:商品数量;"`
      TotalPrice  *float64 `json:"totalPrice" form:"totalPrice" gorm:"column:total_price;comment:订单总价;size:8;"`
      PayPrice  *float64 `json:"payPrice" form:"payPrice" gorm:"column:pay_price;comment:实际支付金额;size:8;"`
      Paid  *bool `json:"paid" form:"paid" gorm:"column:paid;comment:支付状态;"`
      PayTime  *time.Time `json:"payTime" form:"payTime" gorm:"column:pay_time;comment:支付时间;"`
      PayType  string `json:"payType" form:"payType" gorm:"column:pay_type;comment:支付方式;size:32;"`
      RefundStatus  *bool `json:"refundStatus" form:"refundStatus" gorm:"column:refund_status;comment:0 未退款 1 申请中 2 已退款;"`
      RefundReasonWapImg  string `json:"refundReasonWapImg" form:"refundReasonWapImg" gorm:"column:refund_reason_wap_img;comment:退款图片;size:255;"`
      RefundReasonWapExplain  string `json:"refundReasonWapExplain" form:"refundReasonWapExplain" gorm:"column:refund_reason_wap_explain;comment:退款用户说明;size:255;"`
      RefundReasonTime  *time.Time `json:"refundReasonTime" form:"refundReasonTime" gorm:"column:refund_reason_time;comment:退款时间;"`
      RefundReasonWap  string `json:"refundReasonWap" form:"refundReasonWap" gorm:"column:refund_reason_wap;comment:前台退款原因;size:255;"`
      RefundReason  string `json:"refundReason" form:"refundReason" gorm:"column:refund_reason;comment:不退款的理由;size:255;"`
      RefundPrice  *float64 `json:"refundPrice" form:"refundPrice" gorm:"column:refund_price;comment:退款金额;size:8;"`
      GainIntegral  *float64 `json:"gainIntegral" form:"gainIntegral" gorm:"column:gain_integral;comment:消费赚取积分;size:8;"`
      BackIntegral  *float64 `json:"backIntegral" form:"backIntegral" gorm:"column:back_integral;comment:给用户退了多少积分;size:8;"`
      Mark  string `json:"mark" form:"mark" gorm:"column:mark;comment:备注;size:512;"`
      Unique  string `json:"unique" form:"unique" gorm:"column:unique;comment:唯一id(md5加密)类似id;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:管理员备注;size:512;"`
      Cost  *float64 `json:"cost" form:"cost" gorm:"column:cost;comment:成本价;size:8;"`
      VerifyCode  string `json:"verifyCode" form:"verifyCode" gorm:"column:verify_code;comment:核销码;size:50;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态:0待发货 1已完成 2申请退款 3已退货 4已退款 5拒绝退款;size:10;"`
      IsDel  *int `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}


// TableName HkOrder 表名
func (HkOrder) TableName() string {
  return "hk_order"
}

