package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/apply"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HkPlatApplySearch struct{
    apply.HkPlatApply
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
