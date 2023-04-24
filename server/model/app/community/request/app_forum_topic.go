package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ForumTopicSearch struct {
	TopicGroupId *uint64 `json:"topicGroupId" form:"topicGroupId" gorm:"type:bigint(20);column:topic_group_id;comment:分组id;"` //分组id
	Hot          *int    `json:"hot" form:"hot" gorm:"column:hot;comment:是否热门：0 否 1 是;size:10;"`                              //是否热门：0 否 1 是
	request.PageInfo
}
type NearbyHotTopicSearch struct {
	CurPos string `json:"curPos" form:"curPos"` //当前位置
}
type ForumTopicUpdate struct {
	Id           uint64  `json:"id" form:"id" gorm:"not null;unique;primary_key;"`                                            // 话题编号
	CoverImage   string  `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:封面;size:500;"`                 //封面
	TopicGroupId *uint64 `json:"topicGroupId" form:"topicGroupId" gorm:"type:bigint(20);column:topic_group_id;comment:分组id;"` //分组id
	Hot          *int    `json:"hot" form:"hot" gorm:"column:hot;comment:是否热门：0 否 1 是;size:10;"`                              //是否热门：0 否 1 是
}
