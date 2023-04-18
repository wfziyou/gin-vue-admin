// 自动生成模板RecordBrowsingCircleHomepage
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// RecordBrowsingCircleHomepage 结构体
type RecordBrowsingCircleHomepage struct {
	global.GVA_MODEL
	TenantId   string     `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	CircleId   *int       `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子编号;size:19;"`
	Browser    *int       `json:"browser" form:"browser" gorm:"column:browser;comment:浏览者编号;size:19;"`
	BrowseTime *time.Time `json:"browseTime" form:"browseTime" gorm:"column:browse_time;comment:浏览时间;"`
	BrowseNum  *int       `json:"browseNum" form:"browseNum" gorm:"column:browse_num;comment:浏览次数;size:10;"`
	CreateUser *int       `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept *int       `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser *int       `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	Status     *int       `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	IsDel      *int       `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}

// TableName RecordBrowsingCircleHomepage 表名
func (RecordBrowsingCircleHomepage) TableName() string {
	return "hk_record_browsing_circle_homepage"
}
