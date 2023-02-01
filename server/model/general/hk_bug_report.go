// 自动生成模板HkBugReport
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkBugReport 结构体
type HkBugReport struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;size:80;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:操作步骤;size:400;"`
      ContentAttachment  string `json:"contentAttachment" form:"contentAttachment" gorm:"column:content_attachment;comment:操作步骤附件;size:400;"`
      ExpectedResult  string `json:"expectedResult" form:"expectedResult" gorm:"column:expected_result;comment:预期结果;size:400;"`
      ActualResult  string `json:"actualResult" form:"actualResult" gorm:"column:actual_result;comment:实际结果;size:400;"`
      ActualResultAttachment  string `json:"actualResultAttachment" form:"actualResultAttachment" gorm:"column:actual_result_attachment;comment:实际结果附件;size:2000;"`
      OtherInfo  string `json:"otherInfo" form:"otherInfo" gorm:"column:other_info;comment:其他信息;size:400;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDeleted  *int `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
}


// TableName HkBugReport 表名
func (HkBugReport) TableName() string {
  return "hk_bug_report"
}

