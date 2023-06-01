// 自动生成模板UserExtend
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserHeaderImage 结构体
type UserHeaderImage struct {
	global.GvaModelApp
	HeaderImg string `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:主页封面;size:256;"` //头像
	Type      int    `json:"type" form:"type" gorm:"column:type;comment:类型;size:10;"`                    //类型
}

// TableName UserCoverImage 表名
func (UserHeaderImage) TableName() string {
	return "hk_user_header_image"
}
