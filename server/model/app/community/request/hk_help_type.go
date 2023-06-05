package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HelpTypeSearch struct {
	community.HelpType
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type HelpThumbsUpReq struct {
	HelpId   uint64 `json:"helpId" form:"helpId"`     //帮助编号
	ThumbsUp int    `json:"thumbsUp" form:"thumbsUp"` //是否好评：0未点赞 1有用 2无用
}
