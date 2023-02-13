package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type UserBrowsingHistorySearch struct {
	UserId uint64 `json:"-" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户编号;"` //用户编号

	StartUpdatedAt *time.Time `json:"startUpdatedAt" form:"startUpdatedAt"`
	EndUpdatedAt   *time.Time `json:"endUpdatedAt" form:"endUpdatedAt"`
	request.PageInfo
}
