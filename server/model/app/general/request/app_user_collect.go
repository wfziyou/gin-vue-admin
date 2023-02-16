package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

// UserCollectReq 结构体
type UserCollectReq struct {
	PostsId uint64 `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"` //帖子编号
}
type UserCollectSearch struct {
	UserId uint64 `json:"-" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户编号;"` //用户编号

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
