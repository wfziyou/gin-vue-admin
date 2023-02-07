// 自动生成模板HkCircleAddRequest
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkCircleAddRequest 结构体
type HkCircleAddRequest struct {
	global.GvaModelApp
	CircleId    *int       `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`                        //圈子_编号
	Reason      string     `json:"reason" form:"reason" gorm:"column:reason;comment:申请理由;size:500;"`                               //申请理由
	CheckUser   *int       `json:"checkUser" form:"checkUser" gorm:"column:check_user;comment:审核人;size:19;"`                       //审核人
	CheckTime   *time.Time `json:"checkTime" form:"checkTime" gorm:"column:check_time;comment:审核时间;"`                              //审核时间
	CheckStatus *int       `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态：0 未处理 1 通过，2驳回;size:10;"` //审核状态：0 未处理 1 通过，2驳回
}

// TableName HkCircleAddRequest 表名
func (HkCircleAddRequest) TableName() string {
	return "hk_circle_add_request"
}
