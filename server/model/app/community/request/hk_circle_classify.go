package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CircleClassifySearch struct {
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"` //名称
	Des  string `json:"des" form:"des" gorm:"column:des;comment:描述;size:20;"`    //描述

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
