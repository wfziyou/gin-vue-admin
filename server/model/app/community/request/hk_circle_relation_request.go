package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CircleRelationRequestSearch struct {
	CircleId       uint64     `json:"circleId" form:"circleId" `                                                                  //圈子_编号
	RelationType   *int       `json:"relationType" form:"relationType" gorm:"column:relation_type;comment:关系类型：0父子 1友好;size:10;"` //关系类型：0父子 1友好
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`                                                       //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`                                                           //创建时间（结束）
	request.PageInfo
}
