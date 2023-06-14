package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// UserCollectReq 结构体
type UserCollectReq struct {
	PostsId uint64 `json:"postsId" form:"postsId" gorm:"type:bigint(20);column:posts_id;comment:帖子编号;"` //帖子编号
}
type UserCollectSearch struct {
	Category int `json:"category" form:"category" gorm:"column:category;comment:类别：1视频、2动态、5文章、6问答、7活动;size:10;"` //类别：1视频、2动态、5文章、6问答、7活动
	request.PageInfo
}

type DeleteAllUserCollectReq struct {
	Category int `json:"category" form:"category" gorm:"column:category;comment:类别：1视频、2动态、5文章、6问答、7活动;size:10;"` //类别：1视频、2动态、5文章、6问答、7活动
}
