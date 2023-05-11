package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type UserCoverImageSearch struct {
	Type int `json:"type" gorm:"comment:类型"` //类型
	request.PageInfo
}
