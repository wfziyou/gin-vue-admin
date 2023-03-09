// 自动生成模板User
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	uuid "github.com/satori/go.uuid"
	"time"
)

// User 结构体
type User struct {
	global.GvaModelApp
	Uuid        uuid.UUID  `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户编号;size:50;"`                                //用户编号
	UserType    int        `json:"userType" form:"userType" gorm:"column:user_type;comment:用户平台: 1web、2app、3other;size:10;"` //用户平台: 1web、2app、3other
	Account     string     `json:"account" form:"account" gorm:"column:account;comment:账号;size:45;"`                         //账号
	Password    string     `json:"password" form:"password" gorm:"column:password;comment:密码;size:80;"`                      //密码
	NickName    string     `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:昵称;size:20;"`                     //昵称
	RealName    string     `json:"realName" form:"realName" gorm:"column:real_name;comment:真名;size:10;"`                     //真名
	HeaderImg   string     `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:头像;size:500;"`                 //头像
	Email       string     `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:45;"`                               //邮箱
	Phone       string     `json:"phone" form:"phone" gorm:"column:phone;comment:手机;size:45;"`                               //手机
	Birthday    *time.Time `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日;"`
	Sex         int        `json:"sex" form:"sex" gorm:"column:sex;comment:性别： 0未知、1男、2女;size:10;"`              //性别： 0未知、1男、2女
	Description string     `json:"description" form:"description" gorm:"column:description;comment:描述;size:45;"` //描述

	AuthorityId uint64                `json:"roleId" form:"roleId" gorm:"column:role_id;default:888;comment:用户角色ID;size:20;"` // 用户角色ID
	Authority   system.SysAuthority   `json:"authority" gorm:"foreignKey:Id;references:AuthorityId;comment:用户角色"`             //用户角色
	Authorities []system.SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	UserExtend  UserExtend            `json:"userExtend" gorm:"foreignKey:ID;references:ID;comment:用户扩展"` //用户扩展
}

// TableName User 表名
func (User) TableName() string {
	return "hk_user"
}

type ForumPostsUser struct {
	ID          uint64 `json:"id" form:"id" gorm:"primarykey"`                                                           // 主键ID
	UserType    int    `json:"userType" form:"userType" gorm:"column:user_type;comment:用户平台: 1web、2app、3other;size:10;"` //用户平台: 1web、2app、3other
	Account     string `json:"account" form:"account" gorm:"column:account;comment:账号;size:45;"`                         //账号
	NickName    string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:昵称;size:20;"`                     //昵称
	RealName    string `json:"realName" form:"realName" gorm:"column:real_name;comment:真名;size:10;"`                     //真名
	HeaderImg   string `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:头像;size:500;"`                 //头像
	Sex         int    `json:"sex" form:"sex" gorm:"column:sex;comment:性别： 0未知、1男、2女;size:10;"`                          //性别： 0未知、1男、2女
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述;size:45;"`             //描述
}

// TableName User 表名
func (ForumPostsUser) TableName() string {
	return "hk_user"
}
