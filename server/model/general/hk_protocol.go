// 自动生成模板HkProtocol
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkProtocol 结构体
type HkProtocol struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:协议名称;size:20;"`
      Des  string `json:"des" form:"des" gorm:"column:des;comment:协议说明;size:20;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:协议内容;size:2000;"`
      Module  string `json:"module" form:"module" gorm:"column:module;comment:所属模块/插件;size:20;"`
      Pos  string `json:"pos" form:"pos" gorm:"column:pos;comment:显示位置;size:20;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDeleted  *int `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
}


// TableName HkProtocol 表名
func (HkProtocol) TableName() string {
  return "hk_protocol"
}

