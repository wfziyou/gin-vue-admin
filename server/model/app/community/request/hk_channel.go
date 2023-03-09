package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HkChannelSearch struct {
	community.HkChannel
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type ParamSetUserChannel struct {
	ChannelIds string `json:"channelIds" form:"channelIds"`
}
