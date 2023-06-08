// 自动生成模板CircleRequest
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CircleRequest 结构体
type CircleRequest struct {
	global.GvaModelApp
	UserId           uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户_编号;"`                                  //用户_编号
	Type             int    `json:"type" form:"type" gorm:"column:type;comment:类型：0官方圈子 ，1用户圈子;size:10;"`                                       //类型：0官方圈子 ，1用户圈子
	Name             string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                                                  //圈子名称
	Logo             string `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`                                               //圈子Logo
	CircleClassifyId uint64 `json:"circleClassifyId" form:"circleClassifyId" gorm:"type:bigint(20);column:circle_classify_id;comment:圈子分类_编号;"` //圈子分类_编号
	Slogan           string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                                            //圈子标语
	Des              string `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`                                                   //圈子简介
	Protocol         string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:圈子规约;size:1000;"`                                    //圈子规约
	CoverImage       string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:圈子背景图;size:500;"`                             //圈子背景图
	Property         int    `json:"property" form:"property" gorm:"column:property;comment:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）;size:10;"`       //:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）
	CheckStatus      *int   `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态：0 未处理 1 通过，2驳回;size:10;"`             //审核状态：0 未处理 1 通过，2驳回
}

// TableName CircleRequest 表名
func (CircleRequest) TableName() string {
	return "hk_circle_request"
}
