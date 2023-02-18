package request

type CreateUserCircleApplyReq struct {
	CircleId uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //CircleId
	ApplyId  uint64 `json:"applyId" form:"applyId" gorm:"type:bigint(20);column:apply_id;comment:应用_编号;"`    //ApplyId
	Sort     int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                         //Sort
}
