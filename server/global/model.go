package global

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

type GVA_MODEL struct {
	ID         uint64                `json:"id" form:"id" gorm:"not null;unique;primary_key;"`                               // 主键ID
	TenantId   string                `json:"-" form:"tenantId" gorm:"column:tenant_id;default:000000;comment:租户ID;size:12;"` //租户ID
	CreatedAt  time.Time             // 创建时间
	UpdatedAt  time.Time             // 更新时间
	DeletedAt  gorm.DeletedAt        `gorm:"index" json:"-"`                                                                    // 删除时间                         // 删除时间
	IsDel      soft_delete.DeletedAt `json:"-" gorm:"softDelete:flag,DeletedAtField:DeletedAt;default:0;comment:刪除标志;size:10;"` // 刪除标志
	Status     int                   `json:"status" form:"status" gorm:"column:status;default:0;comment:状态;size:10;"`
	CreateUser *int                  `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept *int                  `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser *int                  `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
}

type GvaModelApp struct {
	ID         uint64                `json:"id" form:"id" gorm:"primarykey"`                                                 // 主键ID
	TenantId   string                `json:"-" form:"tenantId" gorm:"column:tenant_id;default:000000;comment:租户ID;size:12;"` //租户ID
	CreatedAt  time.Time             // 创建时间
	UpdatedAt  time.Time             // 更新时间
	DeletedAt  gorm.DeletedAt        `json:"-" gorm:"index"`                                                                    // 删除时间                         // 删除时间
	IsDel      soft_delete.DeletedAt `json:"-" gorm:"softDelete:flag,DeletedAtField:DeletedAt;default:0;comment:刪除标志;size:10;"` // 刪除标志
	Status     int                   `json:"-" form:"status" gorm:"column:status;default:0;comment:状态;size:10;"`
	CreateUser *int                  `json:"-" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept *int                  `json:"-" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser *int                  `json:"-" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
}
