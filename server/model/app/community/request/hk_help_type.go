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

type HelpThumbsUpGoodReq struct {
	HelpId uint64 `json:"helpId" form:"helpId"` //帮助编号
	Good   int    `json:"good" form:"good"`     //是否好评：0否 1是
}
type HelpThumbsUpBadReq struct {
	HelpId uint64 `json:"helpId" form:"helpId"` //帮助编号
	Bad    int    `json:"bad" form:"bad"`       //是否好评：0否 1是
}
