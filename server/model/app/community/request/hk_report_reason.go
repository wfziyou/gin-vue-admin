package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ReportReasonSearch struct {
	Reason string `json:"reason" form:"reason" gorm:"column:reason;comment:举报理由;size:20;"` //举报理由

	request.PageInfo
}
