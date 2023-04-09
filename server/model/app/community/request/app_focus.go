package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type FindFocusUser struct {
	UserId uint64 `json:"userId" form:"userId" ` //用户编号
}

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
	UserId uint64 `json:"userId" form:"userId"` //用户ID
}
type FansSearch struct {
	request.PageInfo
}
