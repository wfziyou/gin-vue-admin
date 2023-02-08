package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type UserBrowsingHistorySearch struct {
	general.UserBrowsingHistory
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
