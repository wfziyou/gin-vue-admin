package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type UserBrowsingHistorySearch struct {
	StartUpdatedAt *time.Time `json:"startUpdatedAt" form:"startUpdatedAt"`
	EndUpdatedAt   *time.Time `json:"endUpdatedAt" form:"endUpdatedAt"`
	request.PageInfo
}
