// 自动生成模板HkCircleApply
package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkCircleApply 结构体
type HkCircleApply struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      CircleId  *int `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`
      ApplyGroupId  *int `json:"applyGroupId" form:"applyGroupId" gorm:"column:apply_group_id;comment:应用分组_编号;size:19;"`
      ShowName  string `json:"showName" form:"showName" gorm:"column:show_name;comment:名称别名;size:80;"`
      ApplyId  *int `json:"applyId" form:"applyId" gorm:"column:apply_id;comment:应用_编号;size:19;"`
      ApplyParameters  string `json:"applyParameters" form:"applyParameters" gorm:"column:apply_parameters;comment:访问参数;size:256;"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDeleted  *int `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
}


// TableName HkCircleApply 表名
func (HkCircleApply) TableName() string {
  return "hk_circle_apply"
}

