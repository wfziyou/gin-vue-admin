package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FeedbackSearch struct {
	request.PageInfo
}

type CreateFeedbackReq struct {
	TypeId     uint64 `json:"typeId" form:"typeId"`          //类型编号
	Des        string `json:"des" form:"des" `               //描述
	Attachment string `json:"attachment" form:"attachment" ` //附件
}

type UpdateFeedbackReq struct {
	Id          uint64 `json:"id" form:"id"`                    //反馈编号
	CheckStatus *int   `json:"checkStatus" form:"checkStatus" ` //处理状态：0 未处理、1 处理中、2 拒绝、3 完成
	Process     string `json:"process" form:"process" `         //处理描述
}
