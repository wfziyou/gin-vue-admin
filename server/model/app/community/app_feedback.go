// 自动生成模板Feedback
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Feedback 结构体
type Feedback struct {
	global.GvaModelApp
	UserId      uint64 `json:"userId" form:"userId" gorm:"column:user_id;comment:用户编号;size:19;"`                                                //用户编号
	Des         string `json:"des" form:"des" gorm:"column:des;comment:描述;size:1024;"`                                                            //描述
	Attachment  string `json:"attachment" form:"attachment" gorm:"column:attachment;comment:附件;size:2000;"`                                       //附件
	CheckStatus int    `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:处理状态：0 未处理、1 处理中、2 拒绝、3 完成;size:10;"` //处理状态：0 未处理、1 处理中、2 拒绝、3 完成
	Process     string `json:"process" form:"process" gorm:"type:longText;column:process;comment:处理描述;"`                                        //处理描述
}

// TableName Feedback 表名
func (Feedback) TableName() string {
	return "hk_feedback"
}
