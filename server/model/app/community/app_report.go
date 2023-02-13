// 自动生成模板Report
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Report 结构体
type Report struct {
	global.GvaModelApp
	ReportUserId      *int   `json:"reportUserId" form:"reportUserId" gorm:"column:report_user_id;comment:被举报用户编号;size:19;"`                                       //被举报用户编号
	UserId            *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:举报用户编号;size:19;"`                                                           //举报用户编号
	CircleId          *int   `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`                                              //圈子_编号
	ReasonId          *int   `json:"reasonId" form:"reasonId" gorm:"column:reason_id;comment:举报原因_编号;size:19;"`                                                    //举报原因_编号
	ReportType        *int   `json:"reportType" form:"reportType" gorm:"column:report_type;comment:举报类型:0用户举报、1评论举报、2内容举报-帖子、3内容举报-视频、4内容举报-动态、5内容举报-话题;size:10;"` //举报类型:0用户举报、1评论举报、2内容举报-帖子、3内容举报-视频、4内容举报-动态、5内容举报-话题
	Content           string `json:"content" form:"content" gorm:"column:content;comment:举报内容;size:200;"`                                                          //举报内容
	ContentAttachment string `json:"contentAttachment" form:"contentAttachment" gorm:"column:content_attachment;comment:内容附件;size:400;"`                           //内容附件
	CurStatus         *int   `json:"curStatus" form:"curStatus" gorm:"column:cur_status;comment:处理状态：0 未处理、1已处理;size:10;"`                                         //处理状态：0 未处理、1已处理
	HandleUserId      *int   `json:"handleUserId" form:"handleUserId" gorm:"column:handle_user_id;comment:操作用户_编号;size:19;"`                                       //操作用户_编号
	HandleType        *int   `json:"handleType" form:"handleType" gorm:"column:handle_type;comment:处理方式:0无效处理（不予处理）、账号禁言;size:19;"`                                //处理方式:0无效处理（不予处理）、账号禁言
	DurationId        *int   `json:"durationId" form:"durationId" gorm:"column:duration_id;comment:禁言时长_编号;size:19;"`                                              //禁言时长_编号
	HandleScore       *int   `json:"handleScore" form:"handleScore" gorm:"column:handle_score;comment:积分处理:0不扣分、1扣分;size:10;"`                                     //积分处理:0不扣分、1扣分
	ScoreExperience   *int   `json:"scoreExperience" form:"scoreExperience" gorm:"column:score_experience;comment:经验;size:10;"`                                    //经验
	ScoreCommunity    *int   `json:"scoreCommunity" form:"scoreCommunity" gorm:"column:score_community;comment:社区积分;size:10;"`                                     //社区积分
	ScoreBuy          *int   `json:"scoreBuy" form:"scoreBuy" gorm:"column:score_buy;comment:购物积分;size:10;"`                                                       //购物积分
	ScoreDownload     *int   `json:"scoreDownload" form:"scoreDownload" gorm:"column:score_download;comment:下载值;size:10;"`                                         //下载值
	Unlock            *int   `json:"unlock" form:"unlock" gorm:"column:unlock;comment:是否解除：0 否、是;size:10;"`                                                        //是否解除：0 否、是
}

// TableName Report 表名
func (Report) TableName() string {
	return "hk_report"
}
