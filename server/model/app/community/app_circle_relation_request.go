// 自动生成模板CircleRelationRequest
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CircleRelationRequest 结构体
type CircleRelationRequest struct {
	global.GVA_MODEL
	TenantId          string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	RelationType      *int   `json:"relationType" form:"relationType" gorm:"column:relation_type;comment:关系类型：0父子 1友好;size:10;"`
	CircleId          *int   `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`
	CircleName        string `json:"circleName" form:"circleName" gorm:"column:circle_name;comment:圈子名称;size:20;"`
	RequestCircleId   *int   `json:"requestCircleId" form:"requestCircleId" gorm:"column:request_circle_id;comment:请求圈子编号;size:19;"`
	RequestCircleName string `json:"requestCircleName" form:"requestCircleName" gorm:"column:request_circle_name;comment:请求圈子名称;size:20;"`
	RequestDes        string `json:"requestDes" form:"requestDes" gorm:"column:request_des;comment:请求描述;size:256;"`
	CheckStatus       *int   `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审批状态：0未审批 1同意 2拒绝;size:10;"`
	CreateUser        *int   `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept        *int   `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser        *int   `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	Status            *int   `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	IsDel             *int   `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}

// TableName CircleRelationRequest 表名
func (CircleRelationRequest) TableName() string {
	return "hk_circle_relation_request"
}
