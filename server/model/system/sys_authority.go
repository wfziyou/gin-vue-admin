package system

import (
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time       // 创建时间
	UpdatedAt       time.Time       // 更新时间
	DeletedAt       *time.Time      `sql:"index"`
	Id              uint64          `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName   string          `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	ParentId        *uint64         `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	DataAuthorityId []*SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`
	Children        []SysAuthority  `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu   `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users           []SysUser       `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter   string          `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`     // 默认菜单(默认dashboard)
	Sort            *int            `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"` //排序
}

func (SysAuthority) TableName() string {
	return "sys_authorities"
}
