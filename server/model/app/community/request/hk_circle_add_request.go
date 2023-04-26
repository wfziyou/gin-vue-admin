package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CircleAddRequestSearch struct {
	CircleId    uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //圈子_编号
	CheckStatus *int   `json:"checkStatus" form:"checkStatus"`                                                  //审核状态：0 未处理 1 通过，2驳回

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
