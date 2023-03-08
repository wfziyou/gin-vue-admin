// 自动生成模板HkActivity
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkActivity 结构体
type HkActivity struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;size:80;"`
      Des  string `json:"des" form:"des" gorm:"column:des;comment:描述;size:200;"`
      ContentHtml  string `json:"contentHtml" form:"contentHtml" gorm:"column:content_html;comment:内容;"`
      ClassifyId  *int `json:"classifyId" form:"classifyId" gorm:"column:classify_id;comment:分类_编号;size:19;"`
      CircleId  *int `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:创建者编号;size:19;"`
      StartAt  *time.Time `json:"startAt" form:"startAt" gorm:"column:start_at;comment:开始时间;"`
      EndAt  *time.Time `json:"endAt" form:"endAt" gorm:"column:end_at;comment:介绍时间;"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:活动地址;size:50;"`
      AddressZone  string `json:"addressZone" form:"addressZone" gorm:"column:address_zone;comment:地址所在区域;size:50;"`
      AddressPos  string `json:"addressPos" form:"addressPos" gorm:"column:address_pos;comment:地址经纬度;size:50;"`
      UserNum  *int `json:"userNum" form:"userNum" gorm:"column:user_num;comment:用户人数;size:10;"`
      CurUserNum  *int `json:"curUserNum" form:"curUserNum" gorm:"column:cur_user_num;comment:参加人数;size:10;"`
      Public  *int `json:"public" form:"public" gorm:"column:public;comment:是否公开：0 否 1 是;size:10;"`
      Pay  *int `json:"pay" form:"pay" gorm:"column:pay;comment:是否付费：0 否 1 是;size:10;"`
      PayCurrency  *int `json:"payCurrency" form:"payCurrency" gorm:"column:pay_currency;comment:付费货币：1人民、2积分、3代币;"`
      PayNum  *int `json:"payNum" form:"payNum" gorm:"column:pay_num;comment:付费金额;"`
      DynamicNum  *int `json:"dynamicNum" form:"dynamicNum" gorm:"column:dynamic_num;comment:动态数;size:10;"`
      CollectNum  *int `json:"collectNum" form:"collectNum" gorm:"column:collect_num;comment:收藏次数;"`
      Collect  *int `json:"collect" form:"collect" gorm:"column:collect;comment:是否收藏：0否、1是;"`
      ChatGroupId  *int `json:"chatGroupId" form:"chatGroupId" gorm:"column:chat_group_id;comment:聊天群编号;size:19;"`
      CheckStatus  *int `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态：0 未处理 1 通过，2驳回;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDel  *int `json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否已删除;size:10;"`
}


// TableName HkActivity 表名
func (HkActivity) TableName() string {
  return "hk_activity"
}

