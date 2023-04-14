// 自动生成模板HkChannel
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ChannelInfo struct {
	ID        uint64 `json:"id" gorm:"not null;unique;primary_key;"`                                          //编号
	Name      string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                         //名称
	Code      int    `json:"code" form:"code" gorm:"column:code;comment:标识：0常规、1资讯、2关注、3附近、4活动、5问答;size:10;"` //标识：0常规、1资讯、2关注、3附近、4活动、5问答
	FixedType int    `json:"fixedType" form:"fixedType" gorm:"column:fixed_type;comment:固定标识(不固定:0，固定:1);"`   //固定标识(不固定:0，固定:1)
	Sort      int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                         //排序
}

func (ChannelInfo) TableName() string {
	return "hk_channel"
}

// HkChannel 结构体
type HkChannel struct {
	global.GVA_MODEL
	TenantId   string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`
	Code       int    `json:"code" form:"code" gorm:"column:code;comment:标识：0常规、1资讯、2关注、3附近、4活动、5问答;size:10;"`
	FixedType  int    `json:"fixedType" form:"fixedType" gorm:"column:fixed_type;comment:固定标识(不固定:0，固定:1);"`
	Type       int    `json:"type" form:"type" gorm:"column:type;comment:类型：0系统 1自定义;"`
	Sort       int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
	CreateUser *int   `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept *int   `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser *int   `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	Status     *int   `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	IsDel      *int   `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}

// TableName HkChannel 表名
func (HkChannel) TableName() string {
	return "hk_channel"
}
