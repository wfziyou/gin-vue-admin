// 自动生成模板HkThirdPlatform
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// HkThirdPlatform 结构体
type HkThirdPlatform struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      Platform  string `json:"platform" form:"platform" gorm:"column:platform;comment:平台;size:20;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:平台名称;size:20;"`
      Icon  string `json:"icon" form:"icon" gorm:"column:icon;comment:平台图标;size:256;"`
      Appid  string `json:"appid" form:"appid" gorm:"column:appid;comment:平台的应用编号;size:50;"`
      Key  string `json:"key" form:"key" gorm:"column:key;comment:秘钥;size:50;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDel  *int `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}


// TableName HkThirdPlatform 表名
func (HkThirdPlatform) TableName() string {
  return "hk_third_platform"
}

