// 自动生成模板ActivityAddRequest
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ActivityAddRequest 结构体
type ActivityAddRequest struct {
	global.GVA_MODEL
	TenantId    string     `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	UserId      *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:19;"`
	ActivityId  *int       `json:"activityId" form:"activityId" gorm:"column:activity_id;comment:活动_编号;size:19;"`
	Reason      string     `json:"reason" form:"reason" gorm:"column:reason;comment:申请理由;size:500;"`
	CheckUser   *int       `json:"checkUser" form:"checkUser" gorm:"column:check_user;comment:审核人;size:19;"`
	CheckTime   *time.Time `json:"checkTime" form:"checkTime" gorm:"column:check_time;comment:审核时间;"`
	CheckStatus *int       `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态：0 未处理、1通过、2驳回;size:10;"`
	CreateUser  *int       `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept  *int       `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser  *int       `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	Status      *int       `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	IsDel       *int       `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}

// TableName ActivityAddRequest 表名
func (ActivityAddRequest) TableName() string {
	return "hk_activity_add_request"
}
