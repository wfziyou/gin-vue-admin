// 自动生成模板HkCommentThumbsUp
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkCommentThumbsUp 结构体
type HkCommentThumbsUp struct {
	global.GVA_MODEL
	TenantId   string     `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:19;"`
	Time       *time.Time `json:"time" form:"time" gorm:"column:time;comment:点赞时间;"`
	CommentId  *int       `json:"commentId" form:"commentId" gorm:"column:comment_id;comment:评论编号;size:19;"`
	CreateUser *int       `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept *int       `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
}

// TableName HkCommentThumbsUp 表名
func (HkCommentThumbsUp) TableName() string {
	return "hk_comment_thumbs_up"
}
