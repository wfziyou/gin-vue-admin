// 自动生成模板RecordBrowsingUserHomepage
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// RecordBrowsingUserHomepage 结构体
type RecordBrowsingUserHomepage struct {
	global.GvaModelApp
	UserId       uint64       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`
	Browser      uint64       `json:"browser" form:"browser" gorm:"column:browser;comment:浏览者编号;size:19;"`
	BrowseTime   time.Time    `json:"browseTime" form:"browseTime" gorm:"column:browse_time;comment:浏览时间;"`
	BrowseNum    int          `json:"browseNum" form:"browseNum" gorm:"column:browse_num;comment:浏览次数;size:10;"`
	BrowseData   string       `json:"browseData" form:"browseData" gorm:"column:browse_data;comment:浏览数据;size:2048;"`
	UserBaseInfo UserBaseInfo `json:"userBaseInfo" gorm:"foreignKey:ID;references:UserId;comment:用户基本信息"` //用户基本信息
}

// TableName RecordBrowsingUserHomepage 表名
func (RecordBrowsingUserHomepage) TableName() string {
	return "hk_record_browsing_user_homepage"
}
