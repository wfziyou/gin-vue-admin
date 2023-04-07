// 自动生成模板ActivityUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ActivityUser 结构体
type ActivityUser struct {
	global.GVA_MODEL
	TenantId   string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	ActivityId *int   `json:"activityId" form:"activityId" gorm:"column:activity_id;comment:活动编号;size:19;"`
	UserId     *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
	Power      *int   `json:"power" form:"power" gorm:"column:power;comment:权限：0普通 1管理员 2发起者;size:10;"`
	Pay        *int   `json:"pay" form:"pay" gorm:"column:pay;comment:是否付费：0 否 1 是;size:10;"`
	OrderId    *int   `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单编号;size:19;"`
	CreateUser *int   `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept *int   `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser *int   `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	Status     *int   `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	IsDel      *int   `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}

// TableName ActivityUser 表名
func (ActivityUser) TableName() string {
	return "hk_activity_user"
}
