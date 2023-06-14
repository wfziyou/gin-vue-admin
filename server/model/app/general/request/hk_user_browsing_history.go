package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type UserBrowsingHistorySearch struct {
	Category       int        `json:"category" form:"category" gorm:"column:category;comment:类别：1视频、2动态、5文章、6问答、7活动;size:10;"` //类别：1视频、2动态、5文章、6问答、7活动
	StartUpdatedAt *time.Time `json:"startUpdatedAt" form:"startUpdatedAt"`
	EndUpdatedAt   *time.Time `json:"endUpdatedAt" form:"endUpdatedAt"`
	request.PageInfo
}

type DeleteAllUserBrowsingHistoryReq struct {
	Category int `json:"category" form:"category" gorm:"column:category;comment:类别：1视频、2动态、5文章、6问答、7活动;size:10;"` //类别：1视频、2动态、5文章、6问答、7活动
}
