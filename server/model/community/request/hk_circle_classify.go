package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HkCircleClassifySearch struct{
    community.HkCircleClassify
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
