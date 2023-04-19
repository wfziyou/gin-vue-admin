package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CircleSearch struct {
	Type             *int   `json:"type" form:"type" `                        //类型：0官方圈子、1用户圈子、2小区
	CircleClassifyId uint64 `json:"circleClassifyId" form:"circleClassifyId"` //圈子分类_编号
	request.PageInfo
}

type ChildCircleSearch struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	request.PageInfo
}
