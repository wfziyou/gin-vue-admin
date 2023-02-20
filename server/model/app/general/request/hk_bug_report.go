package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BugReportSearch struct {
	UserId      uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户编号;"`                                        //用户编号
	Title       string `json:"title" form:"title" gorm:"column:title;comment:标题;size:80;"`                                                        //标题
	CheckStatus *int   `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:处理状态：0 未处理、1 处理中、2 拒绝、3 完成;size:10;"` //处理状态：0 未处理、1 处理中、2 拒绝、3 完成

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type BugReportCreate struct {
	Title                  string `json:"title" form:"title" gorm:"column:title;comment:标题;size:80;"`                                                                //标题
	Content                string `json:"content" form:"content" gorm:"column:content;comment:操作步骤;size:400;"`                                                     //操作步骤
	ContentAttachment      string `json:"contentAttachment" form:"contentAttachment" gorm:"column:content_attachment;comment:操作步骤附件;size:400;"`                  //操作步骤附件
	ExpectedResult         string `json:"expectedResult" form:"expectedResult" gorm:"column:expected_result;comment:预期结果;size:400;"`                               //预期结果
	ActualResult           string `json:"actualResult" form:"actualResult" gorm:"column:actual_result;comment:实际结果;size:400;"`                                     //实际结果
	ActualResultAttachment string `json:"actualResultAttachment" form:"actualResultAttachment" gorm:"column:actual_result_attachment;comment:实际结果附件;size:2000;"` //实际结果附件
	OtherInfo              string `json:"otherInfo" form:"otherInfo" gorm:"column:other_info;comment:其他信息;size:400;"`                                              //其他信息
}
