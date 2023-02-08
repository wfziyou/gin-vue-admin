package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ProtocolSearch struct {
	Name string `json:"name" form:"name" gorm:"column:name;comment:协议名称;size:20;"` //协议名称

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type FindProtocolByName struct {
	Name string `json:"name" form:"name" gorm:"column:name;comment:协议名称;size:20;"` //协议名称
}
