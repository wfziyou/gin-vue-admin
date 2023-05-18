package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type GoldBillSearch struct {
	Pm             *int       `json:"pm" form:"pm" `     //0 = 支出 1 = 获得
	Type           *int       `json:"type" form:"type" ` //账单类型 1充值 2付费 3红包 4打赏 5其他
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
