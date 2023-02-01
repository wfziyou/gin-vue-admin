// 自动生成模板HkForumTopic
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkForumTopic 结构体
type HkForumTopic struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"`
      CoverImage  string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:封面;size:500;"`
      TopicGroupId  *int `json:"topicGroupId" form:"topicGroupId" gorm:"column:topic_group_id;comment:分组id;size:19;"`
      Type  *int `json:"type" form:"type" gorm:"column:type;comment:话题类型：0 全局，1圈子;size:10;"`
      Hot  *int `json:"hot" form:"hot" gorm:"column:hot;comment:是否热门：0 否 1 是;size:10;"`
      DiscussNum  *int `json:"discussNum" form:"discussNum" gorm:"column:discuss_num;comment:讨论数;size:10;"`
      AttentionNum  *int `json:"attentionNum" form:"attentionNum" gorm:"column:attention_num;comment:关注数;size:10;"`
      CircleId  *int `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`
      ReviewStatus  *int `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态：0 未处理 1 通过，2驳回;size:10;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDeleted  *int `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
}


// TableName HkForumTopic 表名
func (HkForumTopic) TableName() string {
  return "hk_forum_topic"
}

