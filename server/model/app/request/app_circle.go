package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"time"
)

type ForumPostsSearch struct {
	community.HkForumPosts
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
type CircleForumPostsSearch struct {
	CircleId uint   `json:"circleId" form:"circleId" "`                      //圈子_编号
	Category *int   `json:"category" form:"category" "`                      //类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）
	GroupId  *int   `json:"groupId" form:"groupId" "`                        //帖子分类编号
	Title    string `json:"title" form:"title" gorm:"column:title;size:80;"` //标题
	Top      *int   `json:"top" form:"top" `                                 //置顶(0否 1是)
	Marrow   *int   `json:"marrow" form:"marrow" `                           //精华(0否 1是)

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type UserCircleForumPostsSearch struct {
	ID       []uint `json:"-"`
	Category *int   `json:"category" form:"category" "`                      //类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）
	GroupId  *int   `json:"groupId" form:"groupId" "`                        //帖子分类编号
	Title    string `json:"title" form:"title" gorm:"column:title;size:80;"` //标题
	Top      *int   `json:"top" form:"top" `                                 //置顶(0否 1是)
	Marrow   *int   `json:"marrow" form:"marrow" `                           //精华(0否 1是)

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type SelfCircleSearch struct {
	ID   uint   `json:"-"`
	Type *int   `json:"type" form:"type" ` //类型：0官方圈子、1用户圈子、2小区
	Name string `json:"name" form:"name" ` //搜索名字：圈子名称或圈子简介

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type CircleSearch struct {
	Type *int   `json:"type" form:"type" ` //类型：0官方圈子、1用户圈子、2小区
	Name string `json:"name" form:"name" ` //搜索名字：圈子名称或圈子简介

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

// SetUserCurCircleReq 设置用户当前圈子
type SetUserCurCircleReq struct {
	ID uint `json:"id"  gorm:"comment:编号"` // 圈子ID
}

type DeleteCircleUserReq struct {
	CircleId *int `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   *int `json:"userId" form:"userId" `     //用户ID
}
type UpdateCircleUserReq struct {
	CircleId *int `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   *int `json:"userId" form:"userId" `     //用户ID
}
type CircleUserSearch struct {
	CircleId *int `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   *int `json:"userId" form:"userId" `     //用户ID

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

// CreateCircleRequestReq 创建圈子申请参数
type CreateCircleRequestReq struct {
	Name             string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                                            //圈子名称
	Logo             string `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`                                           //圈子Logo
	CircleClassifyId *int   `json:"circleClassifyId" form:"circleClassifyId" gorm:"column:circle_classify_id;comment:圈子分类_编号;size:19;"` //圈子分类_编号
	Slogan           string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                                      //圈子标语
	Des              string `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`                                             //圈子简介
	Protocol         string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:圈子规约;size:1000;"`                              //圈子规约
	BackImage        string `json:"backImage" form:"backImage" gorm:"column:back_image;comment:圈子背景图;size:500;"`                         //圈子背景图
}
type CircleRequestSearch struct {
	Type             *int   `json:"type" form:"type" gorm:"column:type;comment:类型：0官方圈子 ，1用户圈子;size:10;"`                            //类型：0官方圈子 ，1用户圈子
	Name             string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                                             //圈子名称
	CircleClassifyId *int   `json:"circleClassifyId" form:"circleClassifyId" gorm:"column:circle_classify_id;comment:圈子分类_编号;size:19;"`  //圈子分类_编号
	Slogan           string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                                       //圈子标语
	Des              string `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`                                              //圈子简介
	CheckStatus      *int   `json:"checkStatus" form:"checkStatus" gorm:"column:check_status;comment:审核状态：0 未处理 1 通过，2驳回;size:10;"` //审核状态：0 未处理 1 通过，2驳回

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
type CircleClassifySearch struct {
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:20;"` //名称
	Des  string `json:"des" form:"des" gorm:"column:des;comment:描述;size:20;"`    //描述

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}
