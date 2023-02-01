// 自动生成模板HkCircleAddRequest
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkCircleAddRequest 结构体
type HkCircleAddRequest struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      CircleId  *int `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`
      Reason  string `json:"reason" form:"reason" gorm:"column:reason;comment:申请理由;size:500;"`
      CheckUser  *int `json:"checkUser" form:"checkUser" gorm:"column:check_user;comment:审核人;size:19;"`
      CheckTime  *time.Time `json:"checkTime" form:"checkTime" gorm:"column:check_time;comment:审核时间;"`
      CheckStatus  *int `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态：0 未处理 1 通过，2驳回;size:10;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDeleted  *int `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
}


// TableName HkCircleAddRequest 表名
func (HkCircleAddRequest) TableName() string {
  return "hk_circle_add_request"
}

