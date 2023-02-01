package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/general"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HkBugReportSearch struct{
    general.HkBugReport
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
