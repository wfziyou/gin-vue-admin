package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/apply"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HkCircleApplySearch struct {
	apply.HkCircleApply
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type HkCircleApplySearchAll struct {
	apply.HkCircleApply
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
}
