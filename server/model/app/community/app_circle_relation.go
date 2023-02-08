// 自动生成模板CircleRelation
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CircleRelation 结构体
type CircleRelation struct {
	global.GvaModelApp
	RelationType  *int `json:"relationType" form:"relationType" gorm:"column:relation_type;comment:关系类型：0父子节点 1关注;size:10;"`                          //关系类型：0父子节点 1关注
	CircleId      *int `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`                                               //圈子_编号
	OtherCircleId *int `json:"otherCircleId" form:"otherCircleId" gorm:"column:other_circle_id;comment:关系圈子_编号（ 关系类型0：父节点编号； 关系类型1：关注圈子编号）;size:19;"` //关系圈子_编号（ 关系类型0：父节点编号； 关系类型1：关注圈子编号）
}

// TableName CircleRelation 表名
func (CircleRelation) TableName() string {
	return "hk_circle_relation"
}
