package global

import (
	"time"
)

type GVA_MODEL struct {
	ID        uint      `json:"id" form:"id" gorm:"primarykey"`                                                 // 主键ID
	TenantId  string    `json:"-" form:"tenantId" gorm:"column:tenant_id;default:000000;comment:租户ID;size:12;"` //租户ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	DeletedAt time.Time // gorm.DeletedAt        `gorm:"index" json:"-"`                           // 删除时间
	//IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"` // 使用 `1` `0` 标识

	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
	UpdateUser *int       `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	Status     int        `json:"status" form:"status" gorm:"column:status;default:0;comment:状态;size:10;"`
	IsDeleted  *int       `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
	CreateUser *int       `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept *int       `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
}
