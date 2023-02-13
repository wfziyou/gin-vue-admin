package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CircleRequestSearch struct {
	Type             *int   `json:"type" form:"type" gorm:"column:type;comment:类型：0官方圈子 ，1用户圈子;size:10;"`                               //类型：0官方圈子 ，1用户圈子
	Name             string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                                          //圈子名称
	CircleClassifyId uint64 `json:"circleClassifyId" form:"circleClassifyId" gorm:"column:circle_classify_id;comment:圈子分类_编号;size:19;"` //圈子分类_编号
	Slogan           string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                                    //圈子标语
	Des              string `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`                                           //圈子简介
	CheckStatus      *int   `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态：0 未处理 1 通过，2驳回;size:10;"`     //审核状态：0 未处理 1 通过，2驳回

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
