package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CreateUserCircleApplyReq struct {
	CircleId *int `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //CircleId
	ApplyId  *int `json:"applyId" form:"applyId" gorm:"column:apply_id;comment:应用_编号;size:19;"`            //ApplyId
	Sort     *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                         //Sort
}

type ApplySearchReq struct {
	OwerType *int   `json:"owerType" form:"owerType" gorm:"column:ower_type;comment:拥有者：0平台、1圈子、2个人;size:10;"` //拥有者：0平台、1圈子、2个人
	CircleId uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`   //圈子_编号
	UserId   uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户_编号;"`         //用户_编号
	Name     string `json:"name" form:"name" gorm:"column:name;comment:名称;size:80;"`                           //名称

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
