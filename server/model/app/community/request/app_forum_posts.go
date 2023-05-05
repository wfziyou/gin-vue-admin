package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ForumPostsSearch struct {
	UserId   uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:发布者编号;"` //发布者编号
	CircleId uint64 `json:"circleId" form:"circleId" "`                                                //圈子_编号
	Category *int   `json:"category" form:"category" "`                                                //类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动
	GroupId  *int   `json:"groupId" form:"groupId" "`                                                  //帖子分类编号
	Title    string `json:"title" form:"title" gorm:"column:title;size:80;"`                           //标题
	Top      *int   `json:"top" form:"top" `                                                           //置顶：0否、1是
	Marrow   *int   `json:"marrow" form:"marrow" `                                                     //精华：0否、1是

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

// CreateForumPostsReq 结构体
type CreateForumPostsReq struct {
	UserId uint64 `json:"-" ` //发布者编号

	TopicId         string `json:"topicId" form:"topicId"`                 //话题_编号
	CircleId        uint64 `json:"circleId" form:"circleId"`               //圈子_编号
	Category        int    `json:"category" form:"category"`               //类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动
	Title           string `json:"title" form:"title"`                     //标题
	CoverImage      string `json:"coverImage" form:"coverImage"`           //封面
	ContentType     int    `json:"contentType" form:"contentType"`         //内容类型：1markdown、2html
	ContentMarkdown string `json:"contentMarkdown" form:"contentMarkdown"` //markdown内容
	ContentHtml     string `json:"contentHtml" form:"contentHtml"`         //html内容
	Video           string `json:"video" form:"video"`                     //视频地址
	Attachment      string `json:"attachment" form:"attachment"`           //附件
	Anonymity       int    `json:"anonymity" form:"anonymity"`             //匿名发布：0否、1是
	Draft           int    `json:"draft" form:"draft"`                     //是否是草稿：0不是，1是
}

type GetRecommendPostsSearch struct {
	ChannelId uint64 `json:"channelId" form:"channelId"` //频道_编号
	request.PageInfo
}
type UserForumPostsSearch struct {
	UserId   uint64 `json:"userId" form:"userId"`      //用户编号
	Category int    `json:"category" form:"category" ` //类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动
	request.PageInfo
}
