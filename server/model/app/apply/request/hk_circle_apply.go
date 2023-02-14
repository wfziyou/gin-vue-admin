package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CircleApplySearch struct {
	CircleId     uint64 `json:"circleId" form:"circleId"`         //圈子_编号
	ApplyGroupId uint64 `json:"applyGroupId" form:"applyGroupId"` //应用分组_编号
	ShowName     string `json:"showName" form:"showName"`         //名称别名

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type CircleApplySearchAll struct {
	CircleId     uint64 `json:"circleId" form:"circleId"`         //圈子_编号
	ApplyGroupId uint64 `json:"applyGroupId" form:"applyGroupId"` //应用分组_编号
	ShowName     string `json:"showName" form:"showName"`         //名称别名

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
}
