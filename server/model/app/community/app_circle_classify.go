// 自动生成模板CircleClassify
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CircleClassify 结构体
type CircleClassify struct {
	global.GvaModelApp
	Name     string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                      //名称
	Des      string `json:"des" form:"des" gorm:"column:des;comment:描述;size:20;"`                         //描述
	Property *int   `json:"property" form:"property" gorm:"column:property;comment:属性：0公开 ，1受限;size:10;"` //属性：0公开 ，1受限
	Sort     *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                      //排序
}

// TableName CircleClassify 表名
func (CircleClassify) TableName() string {
	return "hk_circle_classify"
}
