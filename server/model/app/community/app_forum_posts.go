// 自动生成模板ForumPosts
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

const (
	//PostsCategoryVideo 1视频
	PostsCategoryVideo = 1
	//PostsCategoryDynamic 2动态
	PostsCategoryDynamic = 2
	//PostsCategoryNews 3资讯
	PostsCategoryNews = 3
	//PostsCategoryNotice 4公告
	PostsCategoryNotice = 4
	//PostsCategoryArticle 5文章
	PostsCategoryArticle = 5
	//PostsCategoryQuestion 6问答
	PostsCategoryQuestion = 6
	//PostsCategoryActivity 7活动
	PostsCategoryActivity = 7
)
const (
	//PostsCheckStatusDraft 1草稿
	PostsCheckStatusDraft = 1
	//PostsCheckStatusNotProcessor 2未审批
	PostsCheckStatusNotProcessor = 2
	//PostsCheckStatusPass 3通过
	PostsCheckStatusPass = 3
	//PostsCheckStatusRefuse 4拒绝
	PostsCheckStatusRefuse = 4
)
const (
	DraftFalse = 0
	DraftTrue  = 1
)

//内容类型：1markdown、2html
const (
	ContentTypeMarkdown = 1
	ContentTypeHtml     = 2
)

//是否公开：0 否 1 是
const (
	ForumPostsIsPublicFalse = 0
	ForumPostsIsPublicTrue  = 1
)

//评论权限：0关闭、1开启;size:10;
const (
	ForumPostsPowerCommentClose = 0
	ForumPostsPowerCommentOpen  = 1
)

// ForumPosts 结构体
type ForumPosts struct {
	global.GvaModelApp
	CircleId              uint64     `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`                                                  //圈子_编号
	ActivityId            uint64     `json:"activityId" form:"activityId" gorm:"type:bigint(20);column:activity_id;comment:活动_编号;"`                                            //活动_编号
	Category              int        `json:"category" form:"category" gorm:"column:category;comment:类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动;size:10;"`                      //类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动
	ChannelId             uint64     `json:"channelId" form:"channelId" gorm:"type:bigint(20);column:channel_id;comment:频道_编号;"`                                               //频道_编号
	Title                 string     `json:"title" form:"title" gorm:"column:title;comment:标题;size:80;"`                                                                         //标题
	SeoKey                string     `json:"seoKey" form:"seoKey" gorm:"column:seo_key;comment:SEO关键词;size:500;"`                                                               //SEO关键词
	SeoIntroduce          string     `json:"seoIntroduce" form:"seoIntroduce" gorm:"column:seo_introduce;comment:简介(SEO简介);size:150;"`                                         //简介(SEO简介)
	CoverImage            string     `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:封面;size:500;"`                                                        //封面
	Source                string     `json:"source" form:"source" gorm:"column:source;comment:来源;size:40;"`                                                                      //来源
	ContentType           int        `json:"contentType" form:"contentType" gorm:"column:content_type;comment:内容类型：1markdown、2html;size:10;"`                                  //内容类型：1markdown、2html
	ContentMarkdown       string     `json:"contentMarkdown" form:"contentMarkdown" gorm:"type:longText;column:content_markdown;comment:markdown内容;"`                            //markdown内容
	ContentHtml           string     `json:"contentHtml" form:"contentHtml" gorm:"type:longText;column:content_html;comment:html内容;"`                                            //html内容
	Video                 string     `json:"video" form:"video" gorm:"column:video;comment:视频地址;size:500;"`                                                                    //视频地址
	Attachment            string     `json:"attachment" form:"attachment" gorm:"column:attachment;comment:附件;size:2000;"`                                                        //附件
	Tag                   string     `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:400;"`                                                                              //标签
	Top                   int        `json:"top" form:"top" gorm:"column:top;comment:置顶：0否、1是;size:10;"`                                                                       //置顶：0否、1是
	Marrow                int        `json:"marrow" form:"marrow" gorm:"column:marrow;comment:精华：0否、1是;size:10;"`                                                              //精华：0否、1是
	CommentId             uint64     `json:"commentId" form:"commentId" gorm:"type:bigint(20);column:comment_id;comment:问答最佳答案ID 0未解答，否则已解答;"`                       //问答最佳答案ID 0未解答，否则已解答
	Anonymity             int        `json:"anonymity" form:"anonymity" gorm:"column:anonymity;comment:匿名发布：0否、1是;size:10;"`                                                 //匿名发布：0否、1是
	UserId                uint64     `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:发布者编号;"`                                                       //发布者编号
	ReadNum               int        `json:"readNum" form:"readNum" gorm:"column:read_num;comment:阅读次数;size:10;"`                                                              //阅读次数
	CommentNum            int        `json:"commentNum" form:"commentNum" gorm:"column:comment_num;comment:评论次数;size:10;"`                                                     //评论次数
	ShareNum              int        `json:"shareNum" form:"shareNum" gorm:"column:share_num;comment:分享次数;size:10;"`                                                           //分享次数
	CollectNum            int        `json:"collectNum" form:"collectNum" gorm:"column:collect_num;comment:收藏次数;size:10;"`                                                     //收藏次数
	LikeNum               int        `json:"likeNum" form:"likeNum" gorm:"column:like_num;comment:点赞次数;size:10;"`                                                              //点赞次数
	CheckUser             uint64     `json:"checkUser" form:"checkUser" gorm:"type:bigint(20);column:check_user;comment:审核人;"`                                                  //审核人
	CheckTime             *time.Time `json:"checkTime" form:"checkTime" gorm:"column:check_time;comment:审核时间;"`                                                                //审核时间
	CheckStatus           int        `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态：1草稿、2未审批、3通过、4拒绝;size:10;"`                        //审核状态：1草稿、2未审批、3通过、4拒绝
	PowerComment          int        `json:"powerComment" form:"powerComment" gorm:"column:power_comment;comment:评论权限：0关闭、1开启;size:10;"`                                   //评论权限：0关闭、1开启
	PowerCommentAnonymity int        `json:"powerCommentAnonymity" form:"powerCommentAnonymity" gorm:"column:power_comment_anonymity;comment:匿名评论：0关闭、1开启;size:10;"`       //匿名评论：0关闭、1开启
	Pay                   int        `json:"pay" form:"pay" gorm:"column:pay;comment:付费：0否、1是;size:10;"`                                                                       //付费：0否、1是
	PayContent            int        `json:"payContent" form:"payContent" gorm:"column:pay_content;comment:内容付费：0否、1是;size:10;"`                                             //内容付费：0否、1是
	PayContentLook        int        `json:"payContentLook" form:"payContentLook" gorm:"column:pay_content_look;comment:内容付费可查看百分比例;size:10;"`                          //内容付费可查看百分比例
	PayAttachment         int        `json:"payAttachment" form:"payAttachment" gorm:"column:pay_attachment;comment:附件付费：0否、1是;size:10;"`                                    //附件付费：0否、1是
	PayCurrency           int        `json:"payCurrency" form:"payCurrency" gorm:"column:pay_currency;comment:付费货币：1人民、2积分、3金币;size:10;"`                                //付费货币：1人民、2积分、3金币
	PayNum                uint64     `json:"payNum" form:"payNum" gorm:"column:pay_num;comment:付费金额;size:19;"`                                                                 //付费金额
	ActivityStartAt       string     `json:"activityStartAt" form:"activityStartAt" gorm:"column:activity_start_at;comment:活动开始时间;size:20;"`                                 //活动开始时间
	ActivityEndAt         string     `json:"activityEndAt" form:"activityEndAt" gorm:"column:activity_end_at;comment:活动结束时间;size:20;"`                                       //活动结束时间
	ActivityAddress       string     `json:"activityAddress" form:"activityAddress" gorm:"column:activity_address;comment:活动地址;size:500;"`                                     //活动地址
	ActivityUserNum       int        `json:"activityUserNum" form:"activityUserNum" gorm:"column:activity_user_num;comment:活动用户人数（0不限制活动人数，否则为活动人数）;size:10;"` //活动用户人数（0不限制活动人数，否则为活动人数）
	ActivityCurUserNum    int        `json:"activityCurUserNum" form:"activityCurUserNum" gorm:"column:activity_cur_user_num;comment:活动参加人数;size:10;"`                       //活动参加人数
	ActivityChatGroupId   uint64     `json:"activityChatGroupId" form:"activityChatGroupId" gorm:"type:bigint(20);column:activity_chat_group_id;comment:活动聊天群编号;"`          //活动聊天群编号
	ActivityAddApprove    int        `json:"activityAddApprove" form:"activityAddApprove" gorm:"column:activity_add_approve;comment:参加活动是否需要审批：0不审批、1审批;size:10;"`  //参加活动是否需要审批：0不审批、1审批

	IsPublic      int                  `json:"isPublic" form:"isPublic" gorm:"column:is_public;comment:是否公开：0 否 1 是;size:10;"` //是否公开：0 否 1 是
	TopicInfo     []ForumTopicBaseInfo `json:"topicInfo" gorm:"many2many:hk_forum_topic_posts_mapping;foreignKey:ID;joinForeignKey:PostsId;References:ID;joinReferences:TopicId"`
	CircleInfo    CircleBaseInfo       `json:"circleInfo" gorm:"foreignKey:ID;references:CircleId;comment:用户基本信息"` //圈子基本信息
	UserInfo      UserBaseInfo         `json:"userInfo" gorm:"foreignKey:ID;references:UserId;comment:用户基本信息"`     //用户基本信息
	ThumbsUp      int                  `json:"thumbsUp" `                                                                //是否点赞：0否、1是
	Collect       int                  `json:"collect" `                                                                 //是否收藏：0否、1是
	IsAddActivity int                  `json:"isAddActivity"`                                                            //是否参加活动：0否、1是
}

// TableName ForumPosts 表名
func (ForumPosts) TableName() string {
	return "hk_forum_posts"
}

// ForumPostsBaseInfo 结构体
type ForumPostsBaseInfo struct {
	global.GvaModelApp
	CircleId            uint64               `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`                                                  //圈子_编号
	Category            int                  `json:"category" form:"category" gorm:"column:category;comment:类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动;size:10;"`                      //类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动
	ChannelId           uint64               `json:"channelId" form:"channelId" gorm:"type:bigint(20);column:channel_id;comment:频道_编号;"`                                               //频道_编号
	Title               string               `json:"title" form:"title" gorm:"column:title;comment:标题;size:80;"`                                                                         //标题
	CoverImage          string               `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:封面;size:500;"`                                                        //封面
	Source              string               `json:"source" form:"source" gorm:"column:source;comment:来源;size:40;"`                                                                      //来源
	Video               string               `json:"video" form:"video" gorm:"column:video;comment:视频地址;size:500;"`                                                                    //视频地址
	Attachment          string               `json:"attachment" form:"attachment" gorm:"column:attachment;comment:附件;size:2000;"`                                                        //附件
	Tag                 string               `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:400;"`                                                                              //标签
	Top                 int                  `json:"top" form:"top" gorm:"column:top;comment:置顶：0否、1是;size:10;"`                                                                       //置顶：0否、1是
	Marrow              int                  `json:"marrow" form:"marrow" gorm:"column:marrow;comment:精华：0否、1是;size:10;"`                                                              //精华：0否、1是
	CommentId           uint64               `json:"commentId" form:"commentId" gorm:"type:bigint(20);column:comment_id;comment:问答最佳答案ID 0未解答，否则已解答;"`                       //问答最佳答案ID 0未解答，否则已解答
	Anonymity           int                  `json:"anonymity" form:"anonymity" gorm:"column:anonymity;comment:匿名发布：0否、1是;size:10;"`                                                 //匿名发布：0否、1是
	UserId              uint64               `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:发布者编号;"`                                                       //发布者编号
	ReadNum             int                  `json:"readNum" form:"readNum" gorm:"column:read_num;comment:阅读次数;size:10;"`                                                              //阅读次数
	CommentNum          int                  `json:"commentNum" form:"commentNum" gorm:"column:comment_num;comment:评论次数;size:10;"`                                                     //评论次数
	ShareNum            int                  `json:"shareNum" form:"shareNum" gorm:"column:share_num;comment:分享次数;size:10;"`                                                           //分享次数
	CollectNum          int                  `json:"collectNum" form:"collectNum" gorm:"column:collect_num;comment:收藏次数;size:10;"`                                                     //收藏次数
	LikeNum             int                  `json:"likeNum" form:"likeNum" gorm:"column:like_num;comment:点赞次数;size:10;"`                                                              //点赞次数
	PayNum              uint64               `json:"payNum" form:"payNum" gorm:"column:pay_num;comment:付费金额;size:19;"`                                                                 //付费金额
	ActivityStartAt     string               `json:"activityStartAt" form:"activityStartAt" gorm:"column:activity_start_at;comment:活动开始时间;size:20;"`                                 //活动开始时间
	ActivityEndAt       string               `json:"activityEndAt" form:"activityEndAt" gorm:"column:activity_end_at;comment:活动结束时间;size:20;"`                                       //活动结束时间
	ActivityAddress     string               `json:"activityAddress" form:"activityAddress" gorm:"column:activity_address;comment:活动地址;size:500;"`                                     //活动地址
	ActivityUserNum     int                  `json:"activityUserNum" form:"activityUserNum" gorm:"column:activity_user_num;comment:活动用户人数（0不限制活动人数，否则为活动人数）;size:10;"` //活动用户人数（0不限制活动人数，否则为活动人数）
	ActivityCurUserNum  int                  `json:"activityCurUserNum" form:"activityCurUserNum" gorm:"column:activity_cur_user_num;comment:活动参加人数;size:10;"`                       //活动参加人数
	ActivityChatGroupId uint64               `json:"activityChatGroupId" form:"activityChatGroupId" gorm:"type:bigint(20);column:activity_chat_group_id;comment:活动聊天群编号;"`          //活动聊天群编号
	ActivityAddApprove  int                  `json:"activityAddApprove" form:"activityAddApprove" gorm:"column:activity_add_approve;comment:参加活动是否需要审批：0不审批、1审批;size:10;"`  //参加活动是否需要审批：0不审批、1审批
	TopicInfo           []ForumTopicBaseInfo `json:"topicInfo" gorm:"many2many:hk_forum_topic_posts_mapping;foreignKey:ID;joinForeignKey:PostsId;References:ID;joinReferences:TopicId"`
	CircleInfo          *CircleBaseInfo      `json:"circleInfo" gorm:"foreignKey:ID;references:CircleId;comment:用户基本信息"`                                      //圈子基本信息
	UserInfo            UserBaseInfo         `json:"userInfo" gorm:"foreignKey:ID;references:UserId;comment:用户基本信息"`                                          //用户基本信息
	ThumbsUp            int                  `json:"thumbsUp"`                                                                                                      //是否点赞：0否、1是
	Collect             int                  `json:"collect"`                                                                                                       //是否收藏：0否、1是
	CheckStatus         int                  `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态：1草稿、2未审批、3通过、4拒绝;size:10;"` //审核状态：1草稿、2未审批、3通过、4拒绝

}

// TableName ForumPosts 表名
func (ForumPostsBaseInfo) TableName() string {
	return "hk_forum_posts"
}
