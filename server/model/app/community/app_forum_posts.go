// 自动生成模板ForumPosts
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// ForumPosts 结构体
type ForumPosts struct {
	global.GvaModelApp
	CircleId              *int       `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`                                                 //圈子_编号
	Category              *int       `json:"category" form:"category" gorm:"column:category;comment:类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）;size:10;"`                        //类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）
	GroupId               *int       `json:"groupId" form:"groupId" gorm:"column:group_id;comment:帖子分类编号;size:19;"`                                                   //帖子分类编号
	Title                 string     `json:"title" form:"title" gorm:"column:title;comment:标题;size:80;"`                                                              //标题
	SeoKey                string     `json:"seoKey" form:"seoKey" gorm:"column:seo_key;comment:SEO关键词;size:500;"`                                                     //SEO关键词
	SeoIntroduce          string     `json:"seoIntroduce" form:"seoIntroduce" gorm:"column:seo_introduce;comment:简介(SEO简介);size:150;"`                                //简介(SEO简介)
	CoverImage            string     `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:封面;size:500;"`                                             //封面
	Source                string     `json:"source" form:"source" gorm:"column:source;comment:来源;size:40;"`                                                           //来源
	Time                  *time.Time `json:"time" form:"time" gorm:"column:time;comment:发布时间;"`                                                                       //发布时间
	Content               string     `json:"content" form:"content" gorm:"column:content;comment:内容;"`                                                                //内容
	ContentType           *int       `json:"contentType" form:"contentType" gorm:"column:content_type;comment:内容类型:0 markdown,1 html;size:10;"`                       //内容类型:0 markdown,1 html
	ContentMarkdown       string     `json:"contentMarkdown" form:"contentMarkdown" gorm:"column:content_markdown;comment:markdown内容;"`                               //markdown内容
	ContentHtml           string     `json:"contentHtml" form:"contentHtml" gorm:"column:content_html;comment:html内容;"`                                               //html内容
	Video                 string     `json:"video" form:"video" gorm:"column:video;comment:视频地址;size:500;"`                                                           //视频地址
	Attachment            string     `json:"attachment" form:"attachment" gorm:"column:attachment;comment:附件;size:2000;"`                                             //附件
	Tag                   string     `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:400;"`                                                                   //标签
	Top                   *int       `json:"top" form:"top" gorm:"column:top;comment:置顶(0否 1是);size:10;"`                                                             //置顶(0否 1是)
	Marrow                *int       `json:"marrow" form:"marrow" gorm:"column:marrow;comment:精华(0否 1是);size:10;"`                                                    //精华(0否 1是)
	CommentId             *int       `json:"commentId" form:"commentId" gorm:"column:comment_id;comment:问答最佳答案ID;size:19;"`                                           //问答最佳答案ID
	Anonymity             *int       `json:"anonymity" form:"anonymity" gorm:"column:anonymity;comment:匿名发布(0否 1是);size:10;"`                                         //匿名发布(0否 1是)
	UserId                *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:发布者编号;size:19;"`                                                       //发布者编号
	ReadNum               *int       `json:"readNum" form:"readNum" gorm:"column:read_num;comment:阅读次数;size:10;"`                                                     //阅读次数
	CommentNum            *int       `json:"commentNum" form:"commentNum" gorm:"column:comment_num;comment:评论次数;size:10;"`                                            //评论次数
	ShareNum              *int       `json:"shareNum" form:"shareNum" gorm:"column:share_num;comment:分享次数;size:10;"`                                                  //分享次数
	CollectNum            *int       `json:"collectNum" form:"collectNum" gorm:"column:collect_num;comment:收藏次数;size:10;"`                                            //收藏次数
	LikeNum               *int       `json:"likeNum" form:"likeNum" gorm:"column:like_num;comment:点赞次数;size:10;"`                                                     //点赞次数
	CheckUser             *int       `json:"checkUser" form:"checkUser" gorm:"column:check_user;comment:审核人;size:19;"`                                                //审核人
	CheckTime             *time.Time `json:"checkTime" form:"checkTime" gorm:"column:check_time;comment:审核时间;"`                                                       //审核时间
	CheckStatus           *int       `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态(0草稿 1未审批 2通过 3拒绝);size:10;"`                       //审核状态(0草稿 1未审批 2通过 3拒绝)
	PowerComment          *int       `json:"powerComment" form:"powerComment" gorm:"column:power_comment;comment:评论权限(0关闭 1开启);size:10;"`                             //评论权限(0关闭 1开启)
	PowerCommentAnonymity *int       `json:"powerCommentAnonymity" form:"powerCommentAnonymity" gorm:"column:power_comment_anonymity;comment:匿名评论(0关闭 1开启);size:10;"` //匿名评论(0关闭 1开启)
	Pay                   *int       `json:"pay" form:"pay" gorm:"column:pay;comment:付费：0 否，1是;size:10;"`                                                             //付费：0 否，1是
	PayContent            *int       `json:"payContent" form:"payContent" gorm:"column:pay_content;comment:内容付费：0 否，1是;size:10;"`                                     //内容付费：0 否，1是
	PayContentLook        *int       `json:"payContentLook" form:"payContentLook" gorm:"column:pay_content_look;comment:内容付费可查看百分比例;size:10;"`                        //内容付费可查看百分比例
	PayAttachment         *int       `json:"payAttachment" form:"payAttachment" gorm:"column:pay_attachment;comment:附件付费：0 否，1是;size:10;"`                            //附件付费：0 否，1是
	PayCurrency           *int       `json:"payCurrency" form:"payCurrency" gorm:"column:pay_currency;comment:付费货币：0 人民、1积分、2代币;size:10;"`                            //付费货币：0 人民、1积分、2代币
	PayNum                *int       `json:"payNum" form:"payNum" gorm:"column:pay_num;comment:付费金额;size:10;"`                                                        //付费金额
}

// TableName ForumPosts 表名
func (ForumPosts) TableName() string {
	return "hk_forum_posts"
}

// ForumPostsBaseInfo 结构体
type ForumPostsBaseInfo struct {
	global.GvaModelApp
	CircleId     *int       `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`                          //圈子_编号
	Category     *int       `json:"category" form:"category" gorm:"column:category;comment:类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）;size:10;"` //类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）
	GroupId      *int       `json:"groupId" form:"groupId" gorm:"column:group_id;comment:帖子分类编号;size:19;"`                            //帖子分类编号
	Title        string     `json:"title" form:"title" gorm:"column:title;comment:标题;size:80;"`                                       //标题
	SeoKey       string     `json:"seoKey" form:"seoKey" gorm:"column:seo_key;comment:SEO关键词;size:500;"`                              //SEO关键词
	SeoIntroduce string     `json:"seoIntroduce" form:"seoIntroduce" gorm:"column:seo_introduce;comment:简介(SEO简介);size:150;"`         //简介(SEO简介)
	CoverImage   string     `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:封面;size:500;"`                      //封面
	Source       string     `json:"source" form:"source" gorm:"column:source;comment:来源;size:40;"`                                    //来源
	Time         *time.Time `json:"time" form:"time" gorm:"column:time;comment:发布时间;"`                                                //发布时间
	ContentHtml  string     `json:"contentHtml" form:"contentHtml" gorm:"column:content_html;comment:html内容;"`                        //html内容
	Video        string     `json:"video" form:"video" gorm:"column:video;comment:视频地址;size:500;"`                                    //视频地址
	Tag          string     `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:400;"`                                            //标签
	Top          *int       `json:"top" form:"top" gorm:"column:top;comment:置顶(0否 1是);size:10;"`                                      //置顶(0否 1是)
	Marrow       *int       `json:"marrow" form:"marrow" gorm:"column:marrow;comment:精华(0否 1是);size:10;"`                             //精华(0否 1是)
	Anonymity    *int       `json:"anonymity" form:"anonymity" gorm:"column:anonymity;comment:匿名发布(0否 1是);size:10;"`                  //匿名发布(0否 1是)
	UserId       *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:发布者编号;size:19;"`                                //发布者编号
	ReadNum      *int       `json:"readNum" form:"readNum" gorm:"column:read_num;comment:阅读次数;size:10;"`                              //阅读次数
	CommentNum   *int       `json:"commentNum" form:"commentNum" gorm:"column:comment_num;comment:评论次数;size:10;"`                     //评论次数
	ShareNum     *int       `json:"shareNum" form:"shareNum" gorm:"column:share_num;comment:分享次数;size:10;"`                           //分享次数
	CollectNum   *int       `json:"collectNum" form:"collectNum" gorm:"column:collect_num;comment:收藏次数;size:10;"`                     //收藏次数
	LikeNum      *int       `json:"likeNum" form:"likeNum" gorm:"column:like_num;comment:点赞次数;size:10;"`                              //点赞次数 	//付费金额
}

// TableName ForumPosts 表名
func (ForumPostsBaseInfo) TableName() string {
	return "hk_forum_posts"
}
