// 自动生成模板HkUserSign
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// HkUserSign 结构体
type HkUserSign struct {
      global.GVA_MODEL
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:签到说明;size:255;"`
      Number  *int `json:"number" form:"number" gorm:"column:number;comment:获得积分;size:10;"`
      Balance  *int `json:"balance" form:"balance" gorm:"column:balance;comment:剩余积分;size:10;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDel  *int `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}


// TableName HkUserSign 表名
func (HkUserSign) TableName() string {
  return "hk_user_sign"
}

