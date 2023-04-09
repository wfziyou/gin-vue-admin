// 自动生成模板FocusUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type FocusUserInfo struct {
	Id          uint64 `json:"id" form:"id" gorm:"not null;unique;primary_key;"` //编号
	FocusUserId uint64 `json:"focusUserId" form:"focusUserId" gorm:"column:focus_user_id;comment:关注用户ID;size:19;"`
	Remark      string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:40;"`                //备注
	Tag         string `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:512;"`                        //标签
	Account     string `json:"account" form:"account" gorm:"column:account;comment:账号;size:45;"`             //账号
	NickName    string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:昵称;size:20;"`         //昵称
	RealName    string `json:"realName" form:"realName" gorm:"column:real_name;comment:真名;size:10;"`         //真名
	HeaderImg   string `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:头像;size:500;"`     //头像
	Sex         int    `json:"sex" form:"sex" gorm:"column:sex;comment:性别： 0未知、1男、2女;size:10;"`              //性别： 0未知、1男、2女
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述;size:45;"` //描述
}

// FocusUser 结构体
type FocusUser struct {
	global.GVA_MODEL
	TenantId    string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	UserId      uint64 `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:19;"`
	Remark      string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:40;"`
	Tag         string `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:512;"`
	FocusUserId uint64 `json:"focusUserId" form:"focusUserId" gorm:"column:focus_user_id;comment:关注用户ID;size:19;"`
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
