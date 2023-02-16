package request

// CreateForumTopic 结构体
type CreateForumTopicReq struct {
	Name         string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`                                     //名称
	CoverImage   string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:封面;size:500;"`                 //封面
	TopicGroupId uint64 `json:"topicGroupId" form:"topicGroupId" gorm:"type:bigint(20);column:topic_group_id;comment:分组id;"` //分组id
	Type         int    `json:"type" form:"type" gorm:"column:type;comment:话题类型：0 全局，1圈子;size:10;"`                          //话题类型：0 全局，1圈子
	CircleId     uint64 `json:"circleId" form:"circleId" gorm:"type:bigint(20);column:circle_id;comment:圈子_编号;"`             //圈子_编号
}
