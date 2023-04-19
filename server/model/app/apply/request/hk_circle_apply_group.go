package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CircleApplyGroupSearch struct {
	CircleId uint64  `json:"circleId" form:"circleId"`                                                        //圈子_编号
	ParentId *uint64 `json:"parentId" form:"parentId" gorm:"type:bigint(20);column:parent_id;comment:父节点编号;"` //父节点编号
	request.PageInfo
}

type CircleApplyGroupSearchAll struct {
	CircleId uint64  `json:"circleId" form:"circleId"`                                                        //圈子_编号
	Name     string  `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                         //名称
	ParentId *uint64 `json:"parentId" form:"parentId" gorm:"type:bigint(20);column:parent_id;comment:父节点编号;"` //父节点编号
}
