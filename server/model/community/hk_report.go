// 自动生成模板HkReport
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkReport 结构体
type HkReport struct {
      global.GVA_MODEL
      TenantId  string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
      ReportUserId  *int `json:"reportUserId" form:"reportUserId" gorm:"column:report_user_id;comment:被举报用户编号;size:19;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:举报用户编号;size:19;"`
      CircleId  *int `json:"circleId" form:"circleId" gorm:"column:circle_id;comment:圈子_编号;size:19;"`
      ReasonId  *int `json:"reasonId" form:"reasonId" gorm:"column:reason_id;comment:举报原因_编号;size:19;"`
      ReportType  *int `json:"reportType" form:"reportType" gorm:"column:report_type;comment:举报类型:0用户举报、1评论举报、2内容举报-帖子、3内容举报-视频、4内容举报-动态、5内容举报-话题;size:10;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:举报内容;size:200;"`
      ContentAttachment  string `json:"contentAttachment" form:"contentAttachment" gorm:"column:content_attachment;comment:内容附件;size:400;"`
      CurStatus  *int `json:"curStatus" form:"curStatus" gorm:"column:cur_status;comment:处理状态：0 未处理、1已处理;size:10;"`
      HandleUserId  *int `json:"handleUserId" form:"handleUserId" gorm:"column:handle_user_id;comment:操作用户_编号;size:19;"`
      HandleType  *int `json:"handleType" form:"handleType" gorm:"column:handle_type;comment:处理方式:0无效处理（不予处理）、账号禁言;size:19;"`
      DurationId  *int `json:"durationId" form:"durationId" gorm:"column:duration_id;comment:禁言时长_编号;size:19;"`
      HandleScore  *int `json:"handleScore" form:"handleScore" gorm:"column:handle_score;comment:积分处理:0不扣分、1扣分;size:10;"`
      ScoreExperience  *int `json:"scoreExperience" form:"scoreExperience" gorm:"column:score_experience;comment:经验;size:10;"`
      ScoreCommunity  *int `json:"scoreCommunity" form:"scoreCommunity" gorm:"column:score_community;comment:社区积分;size:10;"`
      ScoreBuy  *int `json:"scoreBuy" form:"scoreBuy" gorm:"column:score_buy;comment:购物积分;size:10;"`
      ScoreDownload  *int `json:"scoreDownload" form:"scoreDownload" gorm:"column:score_download;comment:下载值;size:10;"`
      Unlock  *int `json:"unlock" form:"unlock" gorm:"column:unlock;comment:是否解除：0 否、是;size:10;"`
      CreateUser  *int `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
      CreateDept  *int `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
      UpdateUser  *int `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
      IsDeleted  *int `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
}


// TableName HkReport 表名
func (HkReport) TableName() string {
  return "hk_report"
}

