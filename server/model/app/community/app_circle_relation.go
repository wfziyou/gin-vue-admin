// 自动生成模板CircleRelation
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CircleRelation 结构体
type CircleRelation struct {
	global.GvaModelApp
	RelationType int    `json:"relationType" form:"relationType" gorm:"column:relation_type;comment:关系类型：0父子 1友好;size:10;"` //关系类型：0父子 1友好
	CircleId     uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`            //圈子_编号
	CircleName   string `json:"circleName" form:"circleId" gorm:"column:circle_name;comment:圈子名称;size:20;"`                 //圈子名称

	OtherCircleId   uint64 `json:"otherCircleId" form:"otherCircleId" gorm:"type:bigint(20);column:other_circle_id;comment:关系圈子_编号(关系类型0_父节点编号)"` //关系圈子_编号(关系类型0_父节点编号)
	OtherCircleName string `json:"otherCircleName" form:"otherCircleName" gorm:"column:other_circle_name;comment:关系圈子名称;size:20;"`                //关系圈子名称

}

// TableName CircleRelation 表名
func (CircleRelation) TableName() string {
	return "hk_circle_relation"
}
