package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ForumThumbsUpSearch struct {
	community.ForumThumbsUp
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

// ForumThumbsUp 结构体
type ForumThumbsUpReq struct {
	UserId  uint64 `json:"-" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户id;"`         //用户id
	PostsId uint64 `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"` //帖子编号
}
