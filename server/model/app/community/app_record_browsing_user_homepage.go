// 自动生成模板RecordBrowsingUserHomepage
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// RecordBrowsingUserHomepage 结构体
type RecordBrowsingUserHomepage struct {
	global.GvaModelApp
	TenantId     string       `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	UserId       uint64       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
	Browser      uint64       `json:"browser" form:"browser" gorm:"column:browser;comment:浏览者编号;size:19;"`
	BrowseTime   *time.Time   `json:"browseTime" form:"browseTime" gorm:"column:browse_time;comment:浏览时间;"`
	BrowseNum    *int         `json:"browseNum" form:"browseNum" gorm:"column:browse_num;comment:浏览次数;size:10;"`
	CreateUser   *int         `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept   *int         `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser   *int         `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	Status       *int         `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	IsDel        *int         `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
	UserBaseInfo UserBaseInfo `json:"userBaseInfo" gorm:"foreignKey:ID;references:Browser;comment:用户基本信息"` //用户基本信息
}

// TableName RecordBrowsingUserHomepage 表名
func (RecordBrowsingUserHomepage) TableName() string {
	return "hk_record_browsing_user_homepage"
}
