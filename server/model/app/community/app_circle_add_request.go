// 自动生成模板CircleAddRequest
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// CircleAddRequest 结构体
type CircleAddRequest struct {
	global.GvaModelApp
	CircleId    *uint64    `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`                //圈子_编号
	Reason      string     `json:"reason" form:"reason" gorm:"column:reason;comment:申请理由;size:500;"`                               //申请理由
	CheckUser   *uint64    `json:"checkUser" form:"checkUser" gorm:"type:bigint(20);column:check_user;comment:审核人;"`               //审核人
	CheckTime   *time.Time `json:"checkTime" form:"checkTime" gorm:"column:check_time;comment:审核时间;"`                              //审核时间
	CheckStatus *int       `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态：0 未处理 1 通过，2驳回;size:10;"` //审核状态：0 未处理 1 通过，2驳回
}

// TableName CircleAddRequest 表名
func (CircleAddRequest) TableName() string {
	return "hk_circle_add_request"
}
