// 自动生成模板ActivityUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ActivityUser 结构体
type ActivityUser struct {
	global.GvaModelApp
	ActivityId uint64 `json:"activityId" form:"activityId" gorm:"column:activity_id;comment:活动编号;size:19;"` //活动编号
	UserId     uint64 `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`             //用户编号
	Remark     string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:40;"`                //备注
	Tag        string `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:512;"`                        //标签
	Power      int    `json:"power" form:"power" gorm:"column:power;comment:权限：0普通 1管理员 2发起者;size:10;"`     //权限：0普通 1管理员 2发起者
	Pay        int    `json:"pay" form:"pay" gorm:"column:pay;comment:是否付费：0 否 1 是;size:10;"`               //是否付费：0 否 1 是
	OrderId    uint64 `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单编号;size:19;"`          //订单编号
}

// TableName ActivityUser 表名
func (ActivityUser) TableName() string {
	return "hk_activity_user"
}
