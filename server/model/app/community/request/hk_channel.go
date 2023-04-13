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
	ChannelIds string `json:"channelIds" form:"channelIds" example:"1,2"` //频道编号，通过逗号分割
}

type GlobalTopInfoSearch struct {
	request.PageInfo
}
type GlobalRecommendInfoSearch struct {
	request.PageInfo
}

type GlobalRecommendDynamicSearch struct {
	request.PageInfo
}
