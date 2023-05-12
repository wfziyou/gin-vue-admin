// 自动生成模板ProductGold
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ProductGold 结构体
type ProductGold struct {
	global.GVA_MODEL
	Image      string `json:"image" form:"image" gorm:"column:image;comment:商品图片;size:256;"`
	Num        *int   `json:"num" form:"num" gorm:"column:num;comment:数量;size:10;"`
	Price      *int   `json:"price" form:"price" gorm:"column:price;comment:商品价格;size:20;"`
	VipPrice   *int   `json:"vipPrice" form:"vipPrice" gorm:"column:vip_price;comment:会员价格;size:20;"`
	CreateUser *int   `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept *int   `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	UpdateUser *int   `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	Status     *int   `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	IsDel      *int   `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}

// TableName ProductGold 表名
func (ProductGold) TableName() string {
	return "hk_product_gold"
}
