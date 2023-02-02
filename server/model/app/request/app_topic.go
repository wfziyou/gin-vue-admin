package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ForumTopicGroupSearch struct {
	Name string `json:"name" form:"name"` //名称

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

// CreateForumTopic 结构体
type CreateForumTopicReq struct {
	Name         string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                             //名称
	CoverImage   string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:封面;size:500;"`         //封面
	TopicGroupId *int   `json:"topicGroupId" form:"topicGroupId" gorm:"column:topic_group_id;comment:分组id;size:19;"` //分组id
	Type         *int   `json:"type" form:"type" gorm:"column:type;comment:话题类型：0 全局，1圈子;size:10;"`                  //话题类型：0 全局，1圈子
	CircleId     *int   `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`             //圈子_编号
}
