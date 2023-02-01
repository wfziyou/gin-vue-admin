// 自动生成模板HkUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	uuid "github.com/satori/go.uuid"
	"time"
)

// HkUser 结构体
type HkUser struct {
	global.GVA_MODEL
	TenantId    string                `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	Uuid        uuid.UUID             `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户编号;size:50;"`
	UserType    *int                  `json:"userType" form:"userType" gorm:"column:user_type;comment:用户平台;size:10;"`
	Account     string                `json:"account" form:"account" gorm:"column:account;comment:账号;size:45;"`
	Password    string                `json:"password" form:"password" gorm:"column:password;comment:密码;size:80;"`
	NickName    string                `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:昵称;size:20;"`
	RealName    string                `json:"realName" form:"realName" gorm:"column:real_name;comment:真名;size:10;"`
	HeaderImg   string                `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:头像;size:500;"`
	Email       string                `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:45;"`
	Phone       string                `json:"phone" form:"phone" gorm:"column:phone;comment:手机;size:45;"`
	Birthday    *time.Time            `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日;"`
	Sex         *int                  `json:"sex" form:"sex" gorm:"column:sex;comment:性别;size:10;"`
	RoleId      uint                  `json:"roleId" form:"roleId" gorm:"column:role_id;comment:用户角色ID;size:20;"`
	CreateUser  *int                  `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept  *int                  `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	CreateTime  *time.Time            `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
	UpdateUser  *int                  `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	UpdateTime  *time.Time            `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
	Status      int                   `json:"status" form:"status" gorm:"column:status;comment:状态(用户是否被冻结) 1正常 2冻结;size:10;"`
	IsDeleted   uint                  `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
	Authority   system.SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:RoleId;comment:用户角色"`
	Authorities []system.SysAuthority `json:"authorities" gorm:"many2many:hk_user_authority;"`
}

// TableName HkUser 表名
func (HkUser) TableName() string {
	return "hk_user"
}
