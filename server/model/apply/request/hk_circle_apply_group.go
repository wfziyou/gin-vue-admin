package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/apply"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HkCircleApplyGroupSearch struct {
	apply.HkCircleApplyGroup
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type HkCircleApplyGroupSearchAll struct {
	apply.HkCircleApplyGroup
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
}
