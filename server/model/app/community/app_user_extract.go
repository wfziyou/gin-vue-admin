// 自动生成模板HkUserExtract
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkUserExtract 结构体
type HkUserExtract struct {
      global.GVA_MODEL
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
      RealName  string `json:"realName" form:"realName" gorm:"column:real_name;comment:名称;size:64;"`
      ExtractType  string `json:"extractType" form:"extractType" gorm:"column:extract_type;comment:bank = 银行卡 alipay = 支付宝wx=微信;size:32;"`
      BankCode  string `json:"bankCode" form:"bankCode" gorm:"column:bank_code;comment:银行卡;size:32;"`
      BankAddress  string `json:"bankAddress" form:"bankAddress" gorm:"column:bank_address;comment:开户地址;size:256;"`
      AlipayCode  string `json:"alipayCode" form:"alipayCode" gorm:"column:alipay_code;comment:支付宝账号;size:64;"`
      Wechat  string `json:"wechat" form:"wechat" gorm:"column:wechat;comment:微信号;size:15;"`
      ExtractPrice  *float64 `json:"extractPrice" form:"extractPrice" gorm:"column:extract_price;comment:提现金额;size:8;"`
      Mark  string `json:"mark" form:"mark" gorm:"column:mark;comment:;size:512;"`
      Balance  *float64 `json:"balance" form:"balance" gorm:"column:balance;comment:;size:8;"`
      FailMsg  string `json:"failMsg" form:"failMsg" gorm:"column:fail_msg;comment:无效原因;size:128;"`
      FailTime  *time.Time `json:"failTime" form:"failTime" gorm:"column:fail_time;comment:;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态:0审核中 1已提现 2未通过;size:10;"`
      IsDel  *int `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}


// TableName HkUserExtract 表名
func (HkUserExtract) TableName() string {
  return "hk_user_extract"
}

