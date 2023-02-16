package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CircleApplyGroupSearch struct {
	CircleId *uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //圈子_编号
	Name     string  `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                         //名称
	ParentId *uint64 `json:"parentId" form:"parentId" gorm:"type:bigint(20);column:parent_id;comment:父节点编号;"` //父节点编号

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type CircleApplyGroupSearchAll struct {
	CircleId *uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"` //圈子_编号
	Name     string  `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                         //名称
	ParentId *uint64 `json:"parentId" form:"parentId" gorm:"type:bigint(20);column:parent_id;comment:父节点编号;"` //父节点编号

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
