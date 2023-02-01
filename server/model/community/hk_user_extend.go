// 自动生成模板HkUserExtend
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkUserExtend 结构体
type HkUserExtend struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:19;"`
      Github  string `json:"github" form:"github" gorm:"column:github;comment:github;size:64;"`
      Weibo  string `json:"weibo" form:"weibo" gorm:"column:weibo;comment:微博;size:32;"`
      Weixin  string `json:"weixin" form:"weixin" gorm:"column:weixin;comment:微信;size:32;"`
      Qq  string `json:"qq" form:"qq" gorm:"column:qq;comment:qq;size:32;"`
      Blog  string `json:"blog" form:"blog" gorm:"column:blog;comment:博客;size:500;"`
      NumCircle  *int `json:"numCircle" form:"numCircle" gorm:"column:num_circle;comment:圈子数;size:19;"`
      NumFocus  *int `json:"numFocus" form:"numFocus" gorm:"column:num_focus;comment:关注数;size:19;"`
      NumFan  *int `json:"numFan" form:"numFan" gorm:"column:num_fan;comment:粉丝数;size:19;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDeleted  *int `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
}


// TableName HkUserExtend 表名
func (HkUserExtend) TableName() string {
  return "hk_user_extend"
}

