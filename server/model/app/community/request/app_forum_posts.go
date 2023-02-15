package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ForumPostsSearch struct {
	UserId   uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:发布者编号;"` //发布者编号
	CircleId uint64 `json:"circleId" form:"circleId" "`                                                //圈子_编号
	Category *int   `json:"category" form:"category" "`                                                //类别（1视频、2动态、3资讯、4公告、5文章、6问答、7建议）
	GroupId  *int   `json:"groupId" form:"groupId" "`                                                  //帖子分类编号
	Title    string `json:"title" form:"title" gorm:"column:title;size:80;"`                           //标题
	Top      *int   `json:"top" form:"top" `                                                           //置顶(0否 1是)
	Marrow   *int   `json:"marrow" form:"marrow" `                                                     //精华(0否 1是)

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

// CreateForumPostsReq 结构体
type CreateForumPostsReq struct {
	UserId uint64 `json:"-" ` //发布者编号

	TopicId         uint64 `json:"topicId" form:"topicId" gorm:"column:topic_id;comment:话题_编号;size:19;"`                              //话题_编号
	CircleId        uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`                   //圈子_编号
	Category        *int   `json:"category" form:"category" gorm:"column:category;comment:类别（1视频、2动态、3资讯、4公告、5文章、6问答、7建议）;size:10;"`  //类别（1视频、2动态、3资讯、4公告、5文章、6问答、7建议）
	Title           string `json:"title" form:"title" gorm:"column:title;comment:标题;size:80;"`                                        //标题
	CoverImage      string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:封面;size:500;"`                       //封面
	ContentType     int    `json:"contentType" form:"contentType" gorm:"column:content_type;comment:内容类型:1 markdown,2 html;size:10;"` //内容类型:1 markdown,2 html
	ContentMarkdown string `json:"contentMarkdown" form:"contentMarkdown" gorm:"column:content_markdown;comment:markdown内容;"`         //markdown内容
	ContentHtml     string `json:"contentHtml" form:"contentHtml" gorm:"column:content_html;comment:html内容;"`                         //html内容
	Video           string `json:"video" form:"video" gorm:"column:video;comment:视频地址;size:500;"`                                     //视频地址
	Attachment      string `json:"attachment" form:"attachment" gorm:"column:attachment;comment:附件;size:2000;"`                       //附件
	Anonymity       *int   `json:"anonymity" form:"anonymity" gorm:"column:anonymity;comment:匿名发布(0否 1是);size:10;"`                   //匿名发布(0否 1是)
}
