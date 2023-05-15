package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CircleUserSearch struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	Power    *int   `json:"power" form:"power" `       //权限：0普通 1圈主
	request.PageInfo
}
