package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CreateUserCircleApplyReq struct {
	CircleId *int `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"` //CircleId
	ApplyId  *int `json:"applyId" form:"applyId" gorm:"column:apply_id;comment:应用_编号;size:19;"`    //ApplyId
	Sort     *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                   //Sort
}

type ApplySearchReq struct {
	OwerType *int   `json:"owerType" form:"owerType" gorm:"column:ower_type;comment:拥有者：0平台、1圈子、2个人;size:10;"` //拥有者：0平台、1圈子、2个人
	CircleId *int   `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`                //圈子_编号
	UserId   *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户_编号;size:19;"`                      //用户_编号
	Name     string `json:"name" form:"name" gorm:"column:name;comment:名称;size:80;"`                                  //名称

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type CircleApplySearchReq struct {
	CircleId     *int   `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`                  //圈子_编号
	ApplyGroupId *int   `json:"applyGroupId" form:"applyGroupId" gorm:"column:apply_group_id;comment:应用分组_编号;size:19;"` //应用分组_编号
	ShowName     string `json:"showName" form:"showName" gorm:"column:show_name;comment:名称别名;size:80;"`                   //名称别名
	ApplyId      *int   `json:"applyId" form:"applyId" gorm:"column:apply_id;comment:应用_编号;size:19;"`                     //应用_编号

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type CircleApplyGroupSearchReq struct {
	CircleId *int   `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`  //圈子_编号
	Name     string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                    //名称
	ParentId *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父节点编号;size:19;"` //父节点编号
	
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}