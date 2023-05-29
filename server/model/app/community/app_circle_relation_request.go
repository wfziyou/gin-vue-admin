// 自动生成模板CircleRelationRequest
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

//审批状态：0未审批 1同意 2拒绝;size
const (
	CircleRelationStatusUnCheck = 0
	CircleRelationStatusAgree   = 1
	CircleRelationStatusRefuse  = 2
)

// CircleRelationRequest 结构体
type CircleRelationRequest struct {
	global.GvaModelApp
	RelationType    int    `json:"relationType" form:"relationType" gorm:"column:relation_type;comment:关系类型：0父子 1友好;size:10;"`     //关系类型：0父子 1友好
	CircleId        uint64 `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`                        //圈子_编号
	RequestCircleId uint64 `json:"requestCircleId" form:"requestCircleId" gorm:"column:request_circle_id;comment:请求圈子编号;size:19;"` //请求圈子编号
	RequestDes      string `json:"requestDes" form:"requestDes" gorm:"column:request_des;comment:请求描述;size:256;"`                  //请求描述
	CheckStatus     int    `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审批状态：0未审批 1同意 2拒绝;size:10;"`   //审批状态：0未审批 1同意 2拒绝
}

// TableName CircleRelationRequest 表名
func (CircleRelationRequest) TableName() string {
	return "hk_circle_relation_request"
}
