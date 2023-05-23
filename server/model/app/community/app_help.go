// 自动生成模板HelpType
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Help 结构体
type Help struct {
	global.GvaModelAppEx
	Title   string `json:"title" form:"title" gorm:"column:title;comment:标题;size:20;"`             //标题
	Content string `json:"content" form:"content" gorm:"type:longText;column:content;comment:内容;"` //内容
	TypeId  uint64 `json:"typeId" form:"typeId" gorm:"column:type_id;comment:类型编号;size:19;"`       //类型编号
	GoodNum int    `json:"goodNum" form:"goodNum" gorm:"column:good_num;comment:好评数;size:10;"`     //好评数
	BadNum  int    `json:"badNum" form:"badNum" gorm:"column:bad_num;comment:差评数;size:10;"`        //差评数
	Top     int    `json:"top" form:"top" gorm:"column:top;comment:是否置顶：0不 1是;size:10;"`           //是否置顶：0不 1是
	Sort    int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                //排序
}
type HelpBaseInfo struct {
	global.GvaModelAppEx
	Title   string `json:"title" form:"title" gorm:"column:title;comment:标题;size:20;"`         //标题
	TypeId  uint64 `json:"typeId" form:"typeId" gorm:"column:type_id;comment:类型编号;size:19;"`   //类型编号
	GoodNum int    `json:"goodNum" form:"goodNum" gorm:"column:good_num;comment:好评数;size:10;"` //好评数
	BadNum  int    `json:"badNum" form:"badNum" gorm:"column:bad_num;comment:差评数;size:10;"`    //差评数
	Sort    int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`            //排序
}

// TableName HelpType 表名
func (Help) TableName() string {
	return "hk_help"
}

type MainHelp struct {
	Help     []HelpBaseInfo `json:"help"`     //帮助
	HelpType []HelpType     `json:"helpType"` //帮助类型
}
