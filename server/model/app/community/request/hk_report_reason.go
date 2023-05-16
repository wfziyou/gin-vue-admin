package request

type ReportReasonSearch struct {
	ReportType int `json:"reportType" form:"reportType"` //举报类型:1用户、2圈子、3群、4帖子、5帖子评论
}
