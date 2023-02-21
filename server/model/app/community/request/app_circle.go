package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CircleForumPostsSearch struct {
	CircleId uint64 `json:"circleId" form:"circleId" "`                      //圈子_编号
	Category *int   `json:"category" form:"category" "`                      //类别：1视频、2动态、3资讯、4公告、5文章、6问答、7建议
	GroupId  *int   `json:"groupId" form:"groupId" "`                        //帖子分类编号
	Title    string `json:"title" form:"title" gorm:"column:title;size:80;"` //标题
	Top      *int   `json:"top" form:"top" `                                 //置顶：0否、1是
	Marrow   *int   `json:"marrow" form:"marrow" `                           //精华：0否、1是

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type UserCircleForumPostsSearch struct {
	UserId   uint64 `json:"userId" form:"userId" `                           //用户ID
	Category *int   `json:"category" form:"category" "`                      //类别：1视频、2动态、3资讯、4公告、5文章、6问答、7建议
	GroupId  *int   `json:"groupId" form:"groupId" "`                        //帖子分类编号
	Title    string `json:"title" form:"title" gorm:"column:title;size:80;"` //标题
	Top      *int   `json:"top" form:"top" `                                 //置顶：0否、1是
	Marrow   *int   `json:"marrow" form:"marrow" `                           //精华：0否、1是

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type SelfCircleSearch struct {
	UserId   uint64 `json:"-"`                 //用户ID
	CircleId uint64 `json:"-"`                 //圈子ID
	Type     *int   `json:"type" form:"type" ` //类型：0官方圈子、1用户圈子、2小区
	Name     string `json:"name" form:"name" ` //搜索名字：圈子名称或圈子简介

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

// SetUserCurCircleReq 设置用户当前圈子
type SetUserCurCircleReq struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   uint64 `json:"-"`                         //用户ID
}
type EnterCircleReq struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   uint64 `json:"-"`                         //用户ID
}
type ApplyEnterCircleReq struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   uint64 `json:"-"`                         //用户ID
}
type ApproveEnterCircleReq struct {
	ID          uint64 `json:"id" `         //申请编号
	UserId      uint64 `json:"-"`           //用户ID
	CheckStatus int32  `json:"checkStatus"` //审核状态：1通过、2驳回
}

type DeleteCircleUserReq struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   uint64 `json:"userId" form:"userId" `     //用户ID
}
type UpdateCircleUserReq struct {
	CircleId *int `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   *int `json:"userId" form:"userId" `     //用户ID
}

// CreateCircleRequestReq 创建圈子申请参数
type CreateCircleRequestReq struct {
	Name             string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                                                  //圈子名称
	Logo             string `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`                                               //圈子Logo
	CircleClassifyId uint64 `json:"circleClassifyId" form:"circleClassifyId" gorm:"type:bigint(20);column:circle_classify_id;comment:圈子分类_编号;"` //圈子分类_编号
	Slogan           string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                                            //圈子标语
	Des              string `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`                                                   //圈子简介
	Protocol         string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:圈子规约;size:1000;"`                                    //圈子规约
	BackImage        string `json:"backImage" form:"backImage" gorm:"column:back_image;comment:圈子背景图;size:500;"`                                //圈子背景图
}

type UpdateCircleReq struct {
	ID               uint64 `json:"id" `                                                                                                                 //圈子编号
	Name             string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                                                           //圈子名称
	Logo             string `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`                                                        //圈子Logo
	Slogan           string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                                                     //圈子标语;size:20
	Des              string `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`                                                            //圈子简介;size:1000
	Protocol         string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:圈子规约;size:1000;"`                                             //圈子规约;size:1000
	BackImage        string `json:"backImage" form:"backImage" gorm:"column:back_image;comment:圈子背景图;size:500;"`                                         //圈子背景图;size:500
	Process          *int   `json:"process" form:"process" gorm:"column:process;comment:是否开启版块内容人工审核：0 否，1是;size:10;"`                                   //是否开启版块内容人工审核：0 否，1是
	Property         *int   `json:"property" form:"property" gorm:"column:property;comment:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）;size:10;"`                //:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）
	View             *int   `json:"view" form:"view" gorm:"column:view;comment:板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到;size:10;"`          //板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到
	PowerAdd         *int   `json:"powerAdd" form:"powerAdd" gorm:"column:power_add;comment:圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户;size:10;"`              //圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户
	PowerView        *int   `json:"powerView" form:"powerView" gorm:"column:power_view;comment:圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                 //圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerPublish     *int   `json:"powerPublish" form:"powerPublish" gorm:"column:power_publish;comment:圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`        //圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerComment     *int   `json:"powerComment" form:"powerComment" gorm:"column:power_comment;comment:圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`        //圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerAddUser     string `json:"powerAddUser" form:"powerAddUser" gorm:"column:power_add_user;comment:圈子加入权限_指定部门和成员(json数组);size:500;"`              //圈子加入权限_指定部门和成员(json数组);size:500
	PowerViewUser    string `json:"powerViewUser" form:"powerViewUser" gorm:"column:power_view_user;comment:圈子内浏览权限_指定部门和用户(json数组);size:500;"`          //圈子内浏览权限_指定部门和用户(json数组);size:500
	PowerPublishUser string `json:"powerPublishUser" form:"powerPublishUser" gorm:"column:power_publish_user;comment:圈子内发布权限_指定部门和用户(json数组);size:500;"` //圈子内发布权限_指定部门和用户(json数组);size:500
	PowerCommentUser string `json:"powerCommentUser" form:"powerCommentUser" gorm:"column:power_comment_user;comment:圈子内评论权限_指定部门和用户(json数组);size:500;"` //圈子内评论权限_指定部门和用户(json数组);size:500
	NoLimitUserGroup string `json:"noLimitUserGroup" form:"noLimitUserGroup" gorm:"column:no_limit_user_group;comment:不受限用户组(json数组);size:500;"`         //不受限用户组(json数组);size:500
	NewUserFocus     *int   `json:"newUserFocus" form:"newUserFocus" gorm:"column:new_user_focus;comment:新注册用户默认关注：0 否，1是;size:10;"`                     //新注册用户默认关注：0 否，1是
}
