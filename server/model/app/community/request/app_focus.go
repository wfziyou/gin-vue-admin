package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type FrequentBrowsingUserSearch struct {
	request.PageInfo
}

type FocusUserDynamicSearch struct {
	request.PageInfo
}

type FocusUser struct {
	UserId uint64 `json:"userId" form:"userId"` //用户ID
}
type FocusUserCancel struct {
	Ids []uint64 `json:"ids" form:"ids"`
}
type FansSearch struct {
	request.PageInfo
}
