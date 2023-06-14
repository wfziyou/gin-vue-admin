package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type DraftSearch struct {
	Category int `json:"category" form:"category" gorm:"column:category;comment:类别：1视频、2动态、5文章、6问答、7活动;size:10;"` //类别：1视频、2动态、5文章、6问答、7活动
	request.PageInfo
}
type DeleteAllDraftReq struct {
	Category int `json:"category" form:"category" gorm:"column:category;comment:类别：1视频、2动态、5文章、6问答、7活动;size:10;"` //类别：1视频、2动态、5文章、6问答、7活动
}

type UpdateDraftReq struct {
	Id                    uint64  `json:"id" form:"id"`                                                                                                                         //反馈编号
	CircleId              *uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`                                                  //圈子_编号
	Category              *int    `json:"category" form:"category" gorm:"column:category;comment:类别：1视频、2动态、5文章、6问答、7活动;size:10;"`                                  //类别：1视频、2动态、5文章、6问答、7活动
	ChannelId             *uint64 `json:"channelId" form:"channelId" gorm:"type:bigint(20);column:channel_id;comment:频道_编号;"`                                               //频道_编号
	Title                 string  `json:"title" form:"title" gorm:"column:title;comment:标题;size:80;"`                                                                         //标题
	CoverImage            string  `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:封面;size:500;"`                                                        //封面
	ContentHtml           string  `json:"contentHtml" form:"contentHtml" gorm:"type:longText;column:content_html;comment:html内容;"`                                            //html内容
	Video                 string  `json:"video" form:"video" gorm:"column:video;comment:视频地址;size:500;"`                                                                    //视频地址
	Attachment            string  `json:"attachment" form:"attachment" gorm:"column:attachment;comment:附件;size:2000;"`                                                        //附件
	Tag                   string  `json:"tag" form:"tag" gorm:"column:tag;comment:标签;size:400;"`                                                                              //标签
	Anonymity             *int    `json:"anonymity" form:"anonymity" gorm:"column:anonymity;comment:匿名发布：0否、1是;size:10;"`                                                 //匿名发布：0否、1是
	PowerComment          *int    `json:"powerComment" form:"powerComment" gorm:"column:power_comment;comment:评论权限：0关闭、1开启;size:10;"`                                   //评论权限：0关闭、1开启
	PowerCommentAnonymity *int    `json:"powerCommentAnonymity" form:"powerCommentAnonymity" gorm:"column:power_comment_anonymity;comment:匿名评论：0关闭、1开启;size:10;"`       //匿名评论：0关闭、1开启
	PayContent            *int    `json:"payContent" form:"payContent" gorm:"column:pay_content;comment:内容付费：0否、1是;size:10;"`                                             //内容付费：0否、1是
	PayContentLook        *int    `json:"payContentLook" form:"payContentLook" gorm:"column:pay_content_look;comment:内容付费可查看百分比例;size:10;"`                          //内容付费可查看百分比例
	PayAttachment         *int    `json:"payAttachment" form:"payAttachment" gorm:"column:pay_attachment;comment:附件付费：0否、1是;size:10;"`                                    //附件付费：0否、1是
	PayNum                *uint64 `json:"payNum" form:"payNum" gorm:"column:pay_num;comment:付费金额;size:19;"`                                                                 //付费金额
	ActivityStartAt       string  `json:"activityStartAt" form:"activityStartAt" gorm:"column:activity_start_at;comment:活动开始时间;size:20;"`                                 //活动开始时间
	ActivityEndAt         string  `json:"activityEndAt" form:"activityEndAt" gorm:"column:activity_end_at;comment:活动结束时间;size:20;"`                                       //活动结束时间
	ActivityAddress       string  `json:"activityAddress" form:"activityAddress" gorm:"column:activity_address;comment:活动地址;size:500;"`                                     //活动地址
	ActivityUserNum       *int    `json:"activityUserNum" form:"activityUserNum" gorm:"column:activity_user_num;comment:活动用户人数（0不限制活动人数，否则为活动人数）;size:10;"` //活动用户人数（0不限制活动人数，否则为活动人数）
	ActivityAddApprove    *int    `json:"activityAddApprove" form:"activityAddApprove" gorm:"column:activity_add_approve;comment:参加活动是否需要审批：0不审批、1审批;size:10;"`  //参加活动是否需要审批：0不审批、1审批
	IsPublic              *int    `json:"isPublic" form:"isPublic" gorm:"column:is_public;comment:是否公开：0 否 1 是;size:10;"`                                                 //是否公开：0 否 1 是
	TopicId               string  `json:"topicId" form:"topicId" example:"1,2"`                                                                                                 //话题_编号,通过逗号分割
	Draft                 *int    `json:"draft" form:"draft"`                                                                                                                   //是否是草稿：0不是，1是
}
