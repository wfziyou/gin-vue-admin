// 自动生成模板HkCircleRequest
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkCircleRequest 结构体
type HkCircleRequest struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:类型：0官方圈子 ，1用户圈子;size:10;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`
      Logo  string `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`
      CircleClassifyId  *int `json:"circleClassifyId" form:"circleClassifyId" gorm:"column:circle_classify_id;comment:圈子分类_编号;size:19;"`
      Slogan  string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`
      Des  string `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`
      Protocol  string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:圈子规约;size:1000;"`
      BackImage  string `json:"backImage" form:"backImage" gorm:"column:back_image;comment:圈子背景图;size:500;"`
      CheckStatus  *int `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态：0 未处理 1 通过，2驳回;size:10;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDeleted  *int `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
}


// TableName HkCircleRequest 表名
func (HkCircleRequest) TableName() string {
  return "hk_circle_request"
}

