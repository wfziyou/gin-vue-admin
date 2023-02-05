// 自动生成模板HkCircleClassify
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkCircleClassify 结构体
type HkCircleClassify struct {
	global.GVA_MODEL

	ParentId *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父节点编号;size:19;"`
	Name     string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`
	Des      string `json:"des" form:"des" gorm:"column:des;comment:描述;size:20;"`
	Property *int   `json:"property" form:"property" gorm:"column:property;comment:属性：0公开 ，1受限;size:10;"`
	Sort     *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
}

// TableName HkCircleClassify 表名
func (HkCircleClassify) TableName() string {
	return "hk_circle_classify"
}
