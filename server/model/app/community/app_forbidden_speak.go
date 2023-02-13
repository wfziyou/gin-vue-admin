// 自动生成模板ForbiddenSpeak
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ForbiddenSpeak 结构体
type ForbiddenSpeak struct {
	global.GvaModelApp
	UserId       *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`                       //用户编号
	CircleId     *int `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`        //圈子_编号
	DurationId   *int `json:"durationId" form:"durationId" gorm:"column:duration_id;comment:禁言时长_编号;size:19;"`        //禁言时长_编号
	ReasonId     *int `json:"reasonId" form:"reasonId" gorm:"column:reason_id;comment:禁言原因_编号;size:19;"`              //禁言原因_编号
	CurStatus    *int `json:"curStatus" form:"curStatus" gorm:"column:cur_status;comment:当前状态：0 未解锁、1已解锁;size:10;"`   //当前状态：0 未解锁、1已解锁
	UnlockUserId *int `json:"unlockUserId" form:"unlockUserId" gorm:"column:unlock_user_id;comment:解锁用户_编号;size:19;"` //解锁用户_编号
}

// TableName ForbiddenSpeak 表名
func (ForbiddenSpeak) TableName() string {
	return "hk_forbidden_speak"
}
