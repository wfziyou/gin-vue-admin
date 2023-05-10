// 自动生成模板UserExtend
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserCoverImage 结构体
type UserCoverImage struct {
	global.GvaModelApp
	CoverImage string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:主页封面;size:256;"` //主页封面
	Type       int    `json:"type" form:"type" gorm:"column:type;comment:类型;size:10;"`                         //类型
}

// TableName UserCoverImage 表名
func (UserCoverImage) TableName() string {
	return "hk_user_cover_image"
}
