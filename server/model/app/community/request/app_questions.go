package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GlobalRecommendQuestionSearch struct {
	CurPos string `json:"curPos" form:"curPos"` //当前位置
	request.PageInfo
}
type CircleQuestionSearch struct {
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
	request.PageInfo
}

type CreateQuestion struct {
	Title      string `json:"title" form:"title"`           //标题
	Content    string `json:"content" form:"content"`       //内容
	Attachment string `json:"attachment" form:"attachment"` //附件
	PayNum     int    `json:"payNum" form:"payNum"`         //付费金额
}

type CloseQuestion struct {
	Id uint64 `json:"id" form:"id"` //问题_编号

}

type QuestionAnswerSearch struct {
	QuestionId uint64 `json:"questionId" form:"questionId"` //问题_编号
	request.PageInfo
}

type AnswerQuestion struct {
	QuestionId uint64 `json:"questionId" form:"questionId"` //问题_编号
	Content    string `json:"content" form:"content"`       //内容
}

type SetBestAnswer struct {
	QuestionId uint64 `json:"questionId" form:"questionId"` //问题_编号
	AnswerId   uint64 `json:"answerId" form:"answerId"`     //答案_编号
}
type SetQuestionScore struct {
	QuestionId uint64 `json:"questionId" form:"questionId"` //问题_编号
	Score      int    `json:"score" form:"score"`           //分数（1-5星）
}
