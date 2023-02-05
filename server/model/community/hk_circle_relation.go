// 自动生成模板HkCircleRelation
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkCircleRelation 结构体
type HkCircleRelation struct {
	global.GVA_MODEL
	TenantId      string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	RelationType  *int   `json:"relationType" form:"relationType" gorm:"column:relation_type;comment:关系类型：0父子节点 1关注;size:10;"`
	CircleId      *int   `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`
	OtherCircleId *int   `json:"otherCircleId" form:"otherCircleId" gorm:"column:other_circle_id;comment:关系圈子_编号（ 关系类型0：父节点编号； 关系类型1：关注圈子编号）;size:19;"`
	CreateUser    *int   `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept    *int   `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
}

// TableName HkCircleRelation 表名
func (HkCircleRelation) TableName() string {
	return "hk_circle_relation"
}
