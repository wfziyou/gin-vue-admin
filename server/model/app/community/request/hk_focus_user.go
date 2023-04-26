package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FocusUserSearch struct {
	request.PageInfo
}

type UpdateFocusUserReq struct {
	UserId uint64 `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:19;"` //用户ID
	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:40;"`    //备注
	Tag    string `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:512;"`            //标签
}
