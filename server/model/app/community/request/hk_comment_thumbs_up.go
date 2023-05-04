package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CommentThumbsUpSearch struct {
	community.CommentThumbsUp
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type CommentThumbsUpReq struct {
	CommentId uint64 `json:"commentId" form:"commentId" gorm:"type:bigint(20);column:comment_id;comment:评论编号;"` //评论编号
}
