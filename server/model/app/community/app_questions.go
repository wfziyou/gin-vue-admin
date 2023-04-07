package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// QuestionAnswer 结构体
type QuestionAnswer struct {
	global.GvaModelApp
	QuestionId uint64 `json:"questionId" form:"questionId" gorm:"type:bigint(20);column:question_id;comment:问题编号;"` //问题编号
	UserId     uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:评论者;"`              //回答者
	Content    string `json:"content" form:"content" gorm:"column:content;comment:回答内容;size:2000;"`                 //回答内容
	LikeNum    int    `json:"likeNum" form:"likeNum" gorm:"column:like_num;comment:点赞数;size:10;"`                   //点赞数
	ThumbsUp   int    `json:"thumbsUp"`                                                                             //是否点赞：0否、1是
}

// TableName ForumComment 表名
func (QuestionAnswer) TableName() string {
	return "hk_question_answer"
}
