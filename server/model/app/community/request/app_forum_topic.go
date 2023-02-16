package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ForumTopicSearch struct {
	Name         string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                                           //名称
	TopicGroupId uint64 `json:"topicGroupId" form:"topicGroupId" gorm:"type:bigint(20);column:topic_group_id;comment:分组id;"`       //分组id
	Type         *int   `json:"type" form:"type" gorm:"column:type;comment:话题类型：0 全局，1圈子;size:10;"`                                //话题类型：0 全局，1圈子
	Hot          *int   `json:"hot" form:"hot" gorm:"column:hot;comment:是否热门：0 否 1 是;size:10;"`                                    //是否热门：0 否 1 是
	CircleId     uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`                   //圈子_编号
	ReviewStatus *int   `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态：0 未处理 1 通过，2驳回;size:10;"` //审核状态：0 未处理 1 通过，2驳回

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
