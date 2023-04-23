package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ActivitySearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
type GlobalRecommendActivitySearch struct {
	request.PageInfo
}

type CircleRecommendActivitySearch struct {
	CircleId uint64 `json:"circleId" form:"circleId"` // 圈子_编号
	request.PageInfo
}

type ActivityUserSearch struct {
	request.PageInfo
}

type CreateActivityDynamicReq struct {
	Content    string `json:"content" form:"content"`       //活动内容
	Attachment string `json:"attachment" form:"attachment"` //附件
}
type ActivityDynamicSearch struct {
	request.PageInfo
}

type CreateActivityReq struct {
	Title           string     `json:"title" form:"title" `                    //标题
	CoverImage      string     `json:"coverImage" form:"coverImage"`           //封面
	Content         string     `json:"content" form:"content"`                 //活动内容
	ActivityStartAt *time.Time `json:"activityStartAt" form:"activityStartAt"` //活动开始时间
	ActivityEndAt   *time.Time `json:"activityEndAt" form:"activityEndAt"`     //活动结束时间
	ActivityAddress string     `json:"activityAddress" form:"activityAddress"` //活动地址
	ActivityUserNum int        `json:"activityUserNum" form:"activityUserNum"` //活动用户人数（0不限制活动人数，否则为活动人数）
	PayCurrency     int        `json:"payCurrency" form:"payCurrency"`         //付费货币：1人民、2积分、3金币
	PayNum          int        `json:"payNum" form:"payNum"`                   //付费金额
}
type KickOutActivityUsersReq struct {
	Id  uint64   `json:"id" form:"id"`   // 活动编号
	Ids []uint64 `json:"ids" form:"ids"` // 用户编号
}
type UpdateActivityReq struct {
	Id              uint64     `json:"id" form:"id"`                           // 活动编号
	Title           string     `json:"title" form:"title" `                    //标题
	Content         string     `json:"content" form:"content"`                 //活动内容
	ActivityStartAt *time.Time `json:"activityStartAt" form:"activityStartAt"` //活动开始时间
	ActivityEndAt   *time.Time `json:"activityEndAt" form:"activityEndAt"`     //活动结束时间
	ActivityAddress string     `json:"activityAddress" form:"activityAddress"` //活动地址
	ActivityUserNum int        `json:"activityUserNum" form:"activityUserNum"` //活动用户人数（0不限制活动人数，否则为活动人数）
	PayCurrency     int        `json:"payCurrency" form:"payCurrency"`         //付费货币：1人民、2积分、3金币
	PayNum          int        `json:"payNum" form:"payNum"`                   //付费金额
}

type ExitActivityReq struct {
	Id uint64 `json:"id" form:"id"` // 活动编号
}

type FindActivityUserReq struct {
	Id     uint64 `json:"id" form:"id"`         // 活动编号
	UserId uint64 `json:"userId" form:"userId"` // 用户_编号
}

type JoinActivityReq struct {
	Id     uint64 `json:"id" form:"id"`         // 活动编号
	Reason string `json:"reason" form:"reason"` // 申请理由
}

type ActivityAddRequestSearch struct {
	request.PageInfo
}
type UpdateActivityAddRequestReq struct {
	Id          uint64 `json:"id" form:"id"`                   // 活动申请编号
	CheckStatus int    `json:"checkStatus" form:"checkStatus"` //审核状态：0 未处理、1通过、2驳回
}
