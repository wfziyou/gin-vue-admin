package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"time"
)

type HkCircleAddRequestSearch struct {
	community.HkCircleAddRequest
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
