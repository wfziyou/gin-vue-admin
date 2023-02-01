package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/general"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HkProtocolSearch struct{
    general.HkProtocol
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
