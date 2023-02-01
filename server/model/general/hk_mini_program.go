// 自动生成模板HkMiniProgram
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkMiniProgram 结构体
type HkMiniProgram struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;size:80;"`
      Icon  string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`
      CompanyName  string `json:"companyName" form:"companyName" gorm:"column:company_name;comment:公司名称;size:256;"`
      CurVersion  string `json:"curVersion" form:"curVersion" gorm:"column:cur_version;comment:当前版本;size:12;"`
      CurPacketId  string `json:"curPacketId" form:"curPacketId" gorm:"column:cur_packet_id;comment:当前包id;size:12;"`
      Hidden  *int `json:"hidden" form:"hidden" gorm:"column:hidden;comment:隐藏(0否 1是);size:10;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDeleted  *int `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
}


// TableName HkMiniProgram 表名
func (HkMiniProgram) TableName() string {
  return "hk_mini_program"
}

