// 自动生成模板ProductGold
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ProductGold 结构体
type ProductGold struct {
	global.GvaModelApp
	Count int    `json:"count" form:"count" gorm:"column:count;comment:数量;size:19;"`   //数量
	Price uint64 `json:"price" form:"price" gorm:"column:price;comment:商品价格;size:20;"` //商品价格
	Cost  uint64 `json:"cost" form:"cost" gorm:"column:cost;comment:成本价;size:20;"`     //成本价
}

// TableName ProductGold 表名
func (ProductGold) TableName() string {
	return "hk_product_gold"
}

type ProductGoldInfo struct {
	global.GvaModelAppEx
	Count int `json:"count" form:"count" gorm:"column:count;comment:数量;size:19;"`   //数量
	Price int `json:"price" form:"price" gorm:"column:price;comment:商品价格;size:20;"` //商品价格
}
