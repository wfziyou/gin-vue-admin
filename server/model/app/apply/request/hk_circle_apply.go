package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CircleApplySearch struct {
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
	request.PageInfo
}

type CircleApplySearchAll struct {
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
	ShowName string `json:"showName" form:"showName"` //名称别名
}

type CircleHotApplySearch struct {
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
}

type ParamAddCircleApply struct {
	CircleId        uint64 `json:"circleId" form:"circleId"`                                                                       //圈子_编号
	ApplyGroupId    uint64 `json:"applyGroupId" form:"applyGroupId" gorm:"type:bigint(20);column:apply_group_id;comment:应用分组_编号;"` //应用分组_编号
	Power           int    `json:"power" form:"power" gorm:"column:power;comment:权限：0 所有可见 1 圈内成员可见;size:10;"`                     //权限：0 所有可见 1 圈内成员可见
	Name            string `json:"name" form:"name" gorm:"column:name;comment:名称;size:80;"`                                        //名称
	Icon            string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                                       //图标
	Type            int    `json:"type" form:"type" gorm:"column:type;comment:类型(0小程序、1第三方链接);size:10;"`                           //类型(0小程序、1第三方链接)
	MiniProgramId   uint64 `json:"miniProgramId" form:"miniProgramId" gorm:"type:bigint(20);column:mini_program_id;comment:小程序id"` //小程序id
	ApplyAddress    string `json:"applyAddress" form:"applyAddress" gorm:"column:apply_address;comment:访问地址;size:256;"`            //访问地址
	ApplyParameters string `json:"applyParameters" form:"applyParameters" gorm:"column:apply_parameters;comment:访问参数;size:256;"`   //访问参数
}
type ParamUpdateCircleApply struct {
	CircleId        uint64 `json:"circleId" form:"circleId"`                                                                       //圈子_编号
	ApplyGroupId    uint64 `json:"applyGroupId" form:"applyGroupId" gorm:"type:bigint(20);column:apply_group_id;comment:应用分组_编号;"` //应用分组_编号
	Power           int    `json:"power" form:"power" gorm:"column:power;comment:权限：0 所有可见 1 圈内成员可见;size:10;"`                     //权限：0 所有可见 1 圈内成员可见
	Name            string `json:"name" form:"name" gorm:"column:name;comment:名称;size:80;"`                                        //名称
	Icon            string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                                       //图标
	Type            int    `json:"type" form:"type" gorm:"column:type;comment:类型(0小程序、1第三方链接);size:10;"`                           //类型(0小程序、1第三方链接)
	MiniProgramId   uint64 `json:"miniProgramId" form:"miniProgramId" gorm:"type:bigint(20);column:mini_program_id;comment:小程序id"` //小程序id
	ApplyAddress    string `json:"applyAddress" form:"applyAddress" gorm:"column:apply_address;comment:访问地址;size:256;"`            //访问地址
	ApplyParameters string `json:"applyParameters" form:"applyParameters" gorm:"column:apply_parameters;comment:访问参数;size:256;"`   //访问参数
}
type ParamDeleteCircleApply struct {
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
	ApplyId  uint64 `json:"applyId" form:"applyId"`   //应用_编号
}
