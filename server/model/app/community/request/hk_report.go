package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ReportSearch struct {
	community.Report
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type CreateReportReq struct {
	ReportType        int    `json:"reportType" form:"reportType"`                //举报类型:1用户、2圈子、3群、4帖子、5帖子评论
	ReportId          uint64 `json:"reportId" form:"reportId"`                    //被举报编号
	Content           string `json:"content" form:"content" `                     //举报内容
	ContentAttachment string `json:"contentAttachment" form:"contentAttachment" ` //内容附件
}
