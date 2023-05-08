package response

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/im-yunxin/model"

type UserGetUinfosActionRsp struct {
	Code   int              `json:"code"`
	Uinfos []model.UserInfo `json:"uinfos"`
	Desc   string           `json:"desc"`
}
