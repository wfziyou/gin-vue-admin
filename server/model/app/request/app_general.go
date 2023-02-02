package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ProtocolSearch struct {
	Name string `json:"name" form:"name" ` //协议名称

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type BugReportSearch struct {
	Status *int `json:"status" form:"status"` //状态：0 未处理 1 处理中 2 拒绝 3 完成

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
