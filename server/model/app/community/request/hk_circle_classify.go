package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CircleClassifySearch struct {
	request.PageInfo
}
type CircleClassifyAllSearch struct {
	Keyword string `json:"keyword" form:"keyword"` //关键字
}
