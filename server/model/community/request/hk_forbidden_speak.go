package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HkForbiddenSpeakSearch struct{
    community.HkForbiddenSpeak
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
