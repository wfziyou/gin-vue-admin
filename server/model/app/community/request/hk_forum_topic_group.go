package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ForumTopicGroupSearch struct {
	request.PageInfo
}
type ForumTopicGroupListAllSearch struct {
	Keyword string `json:"keyword" form:"keyword"` //关键字
}
