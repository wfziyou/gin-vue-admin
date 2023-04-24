// 自动生成模板Report
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

//举报类型:1用户、2圈子、3群、4帖子、5帖子评论
const (
	ReportTypeUser    = 1
	ReportTypeCircle  = 2
	ReportTypeGroup   = 3
	ReportTypePosts   = 4
	ReportTypeComment = 5
)

// Report 结构体
type Report struct {
	global.GvaModelApp
	UserId            uint64  `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:举报用户编号;"`                            //举报用户编号
	ReasonId          uint64  `json:"reasonId" form:"reasonId" gorm:"type:bigint(20);column:reason_id;comment:举报原因_编号;"`                     //举报原因_编号
	ReportType        int     `json:"reportType" form:"reportType" gorm:"column:report_type;comment:举报类型:1用户、2圈子、3群、4帖子、5帖子评论;size:10;"`     //举报类型:1用户、2圈子、3群、4帖子、5帖子评论
	ReportId          uint64  `json:"reportId" form:"reportId" gorm:"type:bigint(20);column:report_id;comment:被举报编号;"`                       //被举报编号
	Content           string  `json:"content" form:"content" gorm:"column:content;comment:举报内容;size:200;"`                                   //举报内容
	ContentAttachment string  `json:"contentAttachment" form:"contentAttachment" gorm:"column:content_attachment;comment:内容附件;size:400;"`    //内容附件
	CurStatus         *int    `json:"curStatus" form:"curStatus" gorm:"column:cur_status;comment:处理状态：0 未处理、1 处理中、2 拒绝、3 完成;size:10;"`       //处理状态：0 未处理、1 处理中、2 拒绝、3 完成
	HandleUserId      *uint64 `json:"handleUserId" form:"handleUserId" gorm:"type:bigint(20);column:handle_user_id;comment:操作用户_编号;"`        //操作用户_编号
	HandleType        *uint64 `json:"handleType" form:"handleType" gorm:"type:bigint(20);column:handle_type;comment:处理方式:0无效处理（不予处理）、账号禁言;"` //处理方式:0无效处理（不予处理）、账号禁言
	DurationId        *uint64 `json:"durationId" form:"durationId" gorm:"type:bigint(20);column:duration_id;comment:禁言时长_编号;"`               //禁言时长_编号
	HandleScore       *int    `json:"handleScore" form:"handleScore" gorm:"column:handle_score;comment:积分处理:0不扣分、1扣分;size:10;"`              //积分处理:0不扣分、1扣分
	ScoreExperience   *int    `json:"scoreExperience" form:"scoreExperience" gorm:"column:score_experience;comment:经验;size:10;"`             //经验
	ScoreCommunity    *int    `json:"scoreCommunity" form:"scoreCommunity" gorm:"column:score_community;comment:社区积分;size:10;"`              //社区积分
	ScoreBuy          *int    `json:"scoreBuy" form:"scoreBuy" gorm:"column:score_buy;comment:购物积分;size:10;"`                                //购物积分
	ScoreDownload     *int    `json:"scoreDownload" form:"scoreDownload" gorm:"column:score_download;comment:下载值;size:10;"`                  //下载值
	Unlock            *int    `json:"unlock" form:"unlock" gorm:"column:unlock;comment:是否解除：0 否、是;size:10;"`                                 //是否解除：0 否、是
}

// TableName Report 表名
func (Report) TableName() string {
	return "hk_report"
}
