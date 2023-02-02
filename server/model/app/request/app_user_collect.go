package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

// UserCollectReq 结构体
type UserCollectReq struct {
	PostsId *int `json:"postsId" form:"postsId" gorm:"column:posts_id;comment:帖子编号;size:19;"` //帖子编号
}
type UserCollectSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
