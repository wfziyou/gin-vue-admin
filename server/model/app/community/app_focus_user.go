// 自动生成模板FocusUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// FocusUser 结构体
type FocusUser struct {
	global.GVA_MODEL
	TenantId    string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	UserId      *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:19;"`
	FocusUserId *int   `json:"focusUserId" form:"focusUserId" gorm:"column:focus_user_id;comment:关注用户ID;size:19;"`
	CreateUser  *int   `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept  *int   `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser  *int   `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	Status      *int   `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	IsDel       *int   `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}

// TableName FocusUser 表名
func (FocusUser) TableName() string {
	return "hk_focus_user"
}
