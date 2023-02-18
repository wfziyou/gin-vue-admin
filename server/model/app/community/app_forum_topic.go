// 自动生成模板ForumTopic
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ForumTopic 结构体
type ForumTopic struct {
	global.GvaModelApp
	Name         string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                                           //名称
	CoverImage   string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:封面;size:500;"`                       //封面
	TopicGroupId uint64 `json:"topicGroupId" form:"topicGroupId" gorm:"type:bigint(20);column:topic_group_id;comment:分组id;"`       //分组id
	Type         int    `json:"type" form:"type" gorm:"column:type;comment:话题类型：0 全局，1圈子;size:10;"`                                //话题类型：0 全局，1圈子
	Hot          int    `json:"hot" form:"hot" gorm:"column:hot;comment:是否热门：0 否 1 是;size:10;"`                                    //是否热门：0 否 1 是
	DiscussNum   int    `json:"discussNum" form:"discussNum" gorm:"column:discuss_num;comment:讨论数;size:10;"`                       //讨论数
	AttentionNum int    `json:"attentionNum" form:"attentionNum" gorm:"column:attention_num;comment:关注数;size:10;"`                 //关注数
	CircleId     uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`                   //圈子_编号
	ReviewStatus int    `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态：0 未处理 1 通过，2驳回;size:10;"` //审核状态：0 未处理 1 通过，2驳回
}

// TableName ForumTopic 表名
func (ForumTopic) TableName() string {
	return "hk_forum_topic"
}

// ForumTopicBaseInfo 结构体
type ForumTopicBaseInfo struct {
	ID   uint64 `json:"id" form:"id" gorm:"primarykey"`                          // 主键ID
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"` //名称
}

// TableName ForumTopic 表名
func (ForumTopicBaseInfo) TableName() string {
	return "hk_forum_topic"
}
