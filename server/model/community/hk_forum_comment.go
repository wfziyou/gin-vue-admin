// 自动生成模板HkForumComment
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkForumComment 结构体
type HkForumComment struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      PostsId  *int `json:"postsId" form:"postsId" gorm:"column:posts_id;comment:帖子编号;size:19;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:评论者;size:19;"`
      CommentTime  *time.Time `json:"commentTime" form:"commentTime" gorm:"column:comment_time;comment:评论时间;"`
      CommentContent  string `json:"commentContent" form:"commentContent" gorm:"column:comment_content;comment:评论内容;"`
      LikeTimes  *int `json:"likeTimes" form:"likeTimes" gorm:"column:like_times;comment:点赞数;size:10;"`
      ParentId  *int `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父评论编号;size:19;"`
      CheckUser  *int `json:"checkUser" form:"checkUser" gorm:"column:check_user;comment:审核人;size:19;"`
      CheckTime  *time.Time `json:"checkTime" form:"checkTime" gorm:"column:check_time;comment:审核时间;"`
      CheckStatus  *int `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态(0未审批 1通过 2拒绝);size:10;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDeleted  *int `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
}


// TableName HkForumComment 表名
func (HkForumComment) TableName() string {
  return "hk_forum_comment"
}

