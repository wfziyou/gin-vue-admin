// 自动生成模板HkUserCollect
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkUserCollect 结构体
type HkUserCollect struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
      PostsId  *int `json:"postsId" form:"postsId" gorm:"column:posts_id;comment:帖子编号;size:19;"`
      Time  *time.Time `json:"time" form:"time" gorm:"column:time;comment:收藏时间;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDeleted  *int `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
}


// TableName HkUserCollect 表名
func (HkUserCollect) TableName() string {
  return "hk_user_collect"
}

