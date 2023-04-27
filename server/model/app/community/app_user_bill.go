// 自动生成模板UserBill
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserBill 结构体
type UserBill struct {
	global.GVA_MODEL
	UserId     *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
	LinkId     string   `json:"linkId" form:"linkId" gorm:"column:link_id;comment:关联id;size:32;"`
	Pm         *bool    `json:"pm" form:"pm" gorm:"column:pm;comment:0 = 支出 1 = 获得;"`
	Title      string   `json:"title" form:"title" gorm:"column:title;comment:账单标题;size:64;"`
	Category   string   `json:"category" form:"category" gorm:"column:category;comment:明细种类;size:64;"`
	Type       string   `json:"type" form:"type" gorm:"column:type;comment:明细类型;size:64;"`
	Number     *float64 `json:"number" form:"number" gorm:"column:number;comment:明细数字;size:8;"`
	Balance    *float64 `json:"balance" form:"balance" gorm:"column:balance;comment:剩余;size:8;"`
	Mark       string   `json:"mark" form:"mark" gorm:"column:mark;comment:备注;size:512;"`
	CreateUser *int     `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept *int     `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser *int     `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	Status     *int     `json:"status" form:"status" gorm:"column:status;comment:状态:0带确定 1有效 -1无效;size:10;"`
	IsDel      *int     `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}

// TableName UserBill 表名
func (UserBill) TableName() string {
	return "hk_user_bill"
}
