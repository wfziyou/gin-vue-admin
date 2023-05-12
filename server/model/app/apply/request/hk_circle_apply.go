package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CircleApplySearch struct {
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
	request.PageInfo
}

type CircleApplySearchAll struct {
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
	ShowName string `json:"showName" form:"showName"` //名称别名
}

type CircleHotApplySearch struct {
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
}

type ParamAddCircleApply struct {
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
}
type ParamUpdateCircleApply struct {
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
}
