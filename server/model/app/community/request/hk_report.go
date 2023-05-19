package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ReportSearch struct {
	UserId         *uint64    `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:举报用户编号;"`                                      //举报用户编号
	ReasonId       *uint64    `json:"reasonId" form:"reasonId" gorm:"type:bigint(20);column:reason_id;comment:举报原因_编号;"`                               //举报原因_编号
	ReportType     *int       `json:"reportType" form:"reportType" gorm:"column:report_type;comment:举报类型:1用户、2圈子、3群、4帖子、5帖子评论;size:10;"`      //举报类型:1用户、2圈子、3群、4帖子、5帖子评论
	CurStatus      *int       `json:"curStatus" form:"curStatus" gorm:"column:cur_status;comment:处理状态：0 未处理、1 处理中、2 拒绝、3 完成;size:10;"`         //处理状态：0 未处理、1 处理中、2 拒绝、3 完成
	HandleUserId   *uint64    `json:"handleUserId" form:"handleUserId" gorm:"type:bigint(20);column:handle_user_id;comment:操作用户_编号;"`                  //操作用户_编号
	HandleType     *uint64    `json:"handleType" form:"handleType" gorm:"type:bigint(20);column:handle_type;comment:处理方式:0无效处理（不予处理）、账号禁言;"` //处理方式:0无效处理（不予处理）、账号禁言
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`                                                                                  //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`                                                                                      //创建时间（结束）
	request.PageInfo
}

type CreateReportReq struct {
	ReportType        int    `json:"reportType" form:"reportType"`                //举报类型:1用户、2圈子、3群、4帖子、5帖子评论
	ReportId          uint64 `json:"reportId" form:"reportId"`                    //被举报编号
	Reason            string `json:"reason" form:"reason" `                       //举报原因
	Content           string `json:"content" form:"content" `                     //举报内容
	ContentAttachment string `json:"contentAttachment" form:"contentAttachment" ` //内容附件
}
