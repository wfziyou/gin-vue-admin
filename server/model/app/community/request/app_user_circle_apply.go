package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/app/community"

type UserCircleApplySearch struct {
	UserId   uint64 `json:"-" form:"userId" `         //用户编号
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
}
type UserCircleApplyUpdate struct {
	UserId   uint64                              `json:"-" form:"userId" `         //用户编号
	CircleId uint64                              `json:"circleId" form:"circleId"` //圈子_编号
	Apply    []community.UserCircleApplyBaseInfo `json:"apply" form:"apply" `      //应用列表
}
