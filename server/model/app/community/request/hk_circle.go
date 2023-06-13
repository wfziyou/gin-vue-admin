package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CircleSearch struct {
	Type             *int   `json:"type" form:"type" `                        //类型：0官方圈子、1用户圈子、2小区
	CircleClassifyId uint64 `json:"circleClassifyId" form:"circleClassifyId"` //圈子分类_编号
	request.PageInfo
}

type ChildCircleSearch struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	request.PageInfo
}

type UpdateNewsDraftReq struct {
	Id          uint64  `json:"id" form:"id"`                         //编号
	ChannelId   *uint64 `json:"channelId" form:"channelId"`           //频道_编号
	TopicId     string  `json:"topicId" form:"topicId" example:"1,2"` //话题_编号,通过逗号分割
	CircleId    uint64  `json:"circleId" form:"circleId"`             //圈子_编号
	Title       string  `json:"title" form:"title"`                   //标题
	CoverImage  string  `json:"coverImage" form:"coverImage"`         //封面
	ContentHtml string  `json:"contentHtml" form:"contentHtml"`       //html内容
	Draft       *int    `json:"draft" form:"draft"`                   //是否是草稿：0不是，1是
}
type DeleteNewsDraftReq struct {
	Id       uint64 `json:"id" form:"id"`             //编号
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
}
type DeleteNewsDraftByIdsReq struct {
	Ids      []uint64 `json:"ids" form:"ids"`
	CircleId uint64   `json:"circleId" form:"circleId"` //圈子_编号
}
type DeleteNewsReq struct {
	Id       uint64 `json:"id" form:"id"`             //编号
	CircleId uint64 `json:"circleId" form:"circleId"` //圈子_编号
}
type DeleteNewsByIdsReq struct {
	Ids      []uint64 `json:"ids" form:"ids"`
	CircleId uint64   `json:"circleId" form:"circleId"` //圈子_编号
}
