package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CircleUserSearch struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   uint64 `json:"-" form:"userId" `          //用户ID

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
