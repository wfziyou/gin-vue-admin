package system

import (
	"time"
)

type SysApiGroup struct {
	CreatedAt time.Time  // 创建时间
	UpdatedAt time.Time  // 更新时间
	DeletedAt *time.Time `sql:"index"`
	ID        uint       `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID

	Name        string `json:"name" gorm:"comment:名称"`        // 名称
	Description string `json:"description" gorm:"comment:描述"` // 描述
	Sort        int    `json:"sort" gorm:"comment:api组"`      // 排序
}

func (SysApiGroup) TableName() string {
	return "sys_api_group"
}
