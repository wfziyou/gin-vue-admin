// 自动生成模板ActivityAddRequest
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ActivityAddRequest 结构体
type ActivityAddRequest struct {
	global.GvaModelApp
	UserId      uint64     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:19;"`                              //用户ID
	ActivityId  uint64     `json:"activityId" form:"activityId" gorm:"column:activity_id;comment:活动_编号;size:19;"`                 //活动_编号
	Reason      string     `json:"reason" form:"reason" gorm:"column:reason;comment:申请理由;size:500;"`                              //申请理由
	CheckUser   uint64     `json:"checkUser" form:"checkUser" gorm:"column:check_user;comment:审核人;size:19;"`                      //审核人
	CheckTime   *time.Time `json:"checkTime" form:"checkTime" gorm:"column:check_time;comment:审核时间;"`                             //审核时间
	CheckStatus int        `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态：0 未处理、1通过、2驳回;size:10;"` //审核状态：0 未处理、1通过、2驳回
}

// TableName ActivityAddRequest 表名
func (ActivityAddRequest) TableName() string {
	return "hk_activity_add_request"
}
