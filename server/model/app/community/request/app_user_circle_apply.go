package request

type UserCircleApplySearch struct {
	UserId   uint64 `json:"-" form:"userId" `         //用户编号
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
}
