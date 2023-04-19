package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CircleSearch struct {
	Type             *int   `json:"type" form:"type" `                        //类型：0官方圈子、1用户圈子、2小区
	Name             string `json:"name" form:"name" `                        //搜索名字：圈子名称或圈子简介
	CircleClassifyId uint64 `json:"circleClassifyId" form:"circleClassifyId"` //圈子分类_编号

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type ChildCircleSearch struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	request.PageInfo
}
