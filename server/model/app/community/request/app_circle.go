package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CircleForumPostsSearch struct {
	CircleId  uint64 `json:"circleId" form:"circleId" "`   //圈子_编号
	ChannelId uint64 `json:"channelId" form:"channelId" "` //频道_编号
	request.PageInfo
}

type UserCircleForumPostsSearch struct {
	Category *int   `json:"category" form:"category" "`                      //类别：1视频、2动态、5文章、6问答、7活动
	GroupId  *int   `json:"groupId" form:"groupId" "`                        //帖子分类编号
	Title    string `json:"title" form:"title" gorm:"column:title;size:80;"` //标题
	Top      *int   `json:"top" form:"top" `                                 //置顶：0否、1是
	Marrow   *int   `json:"marrow" form:"marrow" `                           //精华：0否、1是

	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"` //创建时间（开始）
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`     //创建时间（结束）
	request.PageInfo
}

type SelfCircleSearch struct {
	Type int `json:"type" form:"type"` //类型：0全部、2管理、3创建
	request.PageInfo
}
type SelfCircleRequestSearch struct {
	request.PageInfo
}

type EnterCircleReq struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   uint64 `json:"-"`                         //用户ID
}
type ExitCircleReq struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   uint64 `json:"-"`                         //用户ID
}

type ParamSelfCircleTop struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
}
type ParamSelfCircleCancelTop struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
}

type FindCircleUserReq struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   uint64 `json:"userId"`                    //用户ID
}
type ApplyEnterCircleReq struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   uint64 `json:"-"`                         //用户ID
	Reason   string `json:"reason" form:"reason"`      //申请理由
}
type ApproveEnterCircleReq struct {
	CircleId    uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	CheckStatus int32  `json:"checkStatus"`               //审核状态：1通过、2驳回
	ID          uint64 `json:"id" `                       //申请编号
}
type DeleteCircleForumPostsReq struct {
	CircleId uint64   `json:"circleId" form:"circleId" ` //圈子_编号
	Ids      []uint64 `json:"ids" form:"ids" `           //帖子Ids
}
type DeleteCircleUserReq struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   uint64 `json:"userId" form:"userId" `     //用户ID
}
type DeleteCircleUsersReq struct {
	CircleId uint64   `json:"circleId" form:"circleId" ` //圈子_编号
	UserIds  []uint64 `json:"userIds" form:"userIds" `   //用户Ids
}
type UpdateCircleUserReq struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   uint64 `json:"userId" form:"userId" `     //用户ID
	Remark   string `json:"remark" form:"remark"`      //备注
	Tag      string `json:"tag" form:"tag"`            //标签
	Sort     *int   `json:"sort" form:"sort"`          //用户的圈子排序
}
type SetCircleUserPowerReq struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	UserId   uint64 `json:"userId" form:"userId" `     //用户ID
	Power    int    `json:"power" form:"power"`        //权限：0普通 1圈主
}

// CreateCircleRequestReq 创建圈子申请参数
type CreateCircleRequestReq struct {
	Name             string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                                                              //圈子名称
	Logo             string `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`                                                             //圈子Logo
	CircleClassifyId uint64 `json:"circleClassifyId" form:"circleClassifyId" gorm:"type:bigint(20);column:circle_classify_id;comment:圈子分类_编号;"`           //圈子分类_编号
	Slogan           string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                                                        //圈子标语
	Des              string `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`                                                               //圈子简介
	Protocol         string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:圈子规约;size:1000;"`                                                //圈子规约
	CoverImage       string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:圈子背景图;size:500;"`                                        //圈子背景图
	Property         int    `json:"property" form:"property" gorm:"column:property;comment:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）;size:10;"` //:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）
}
type ParamDestroyCircle struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
}

type UpdateCircleReq struct {
	ID               uint64 `json:"id" `                                                                                                                                                   //圈子编号
	Name             string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                                                                                         //圈子名称
	Logo             string `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`                                                                                        //圈子Logo
	Slogan           string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                                                                                   //圈子标语;size:20
	Des              string `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`                                                                                          //圈子简介;size:1000
	Protocol         string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:圈子规约;size:1000;"`                                                                           //圈子规约;size:1000
	CoverImage       string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:圈子背景图;size:500;"`                                                                   //圈子背景图;size:500
	Process          *int   `json:"process" form:"process" gorm:"column:process;comment:是否开启版块内容人工审核：0 否，1是;size:10;"`                                                       //是否开启版块内容人工审核：0 否，1是
	Property         *int   `json:"property" form:"property" gorm:"column:property;comment:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）;size:10;"`                            //:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）
	View             *int   `json:"view" form:"view" gorm:"column:view;comment:板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到;size:10;"` //板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到
	PowerAdd         *int   `json:"powerAdd" form:"powerAdd" gorm:"column:power_add;comment:圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户;size:10;"`                     //圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户
	PowerView        *int   `json:"powerView" form:"powerView" gorm:"column:power_view;comment:圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                              //圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerPublish     *int   `json:"powerPublish" form:"powerPublish" gorm:"column:power_publish;comment:圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                     //圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerComment     *int   `json:"powerComment" form:"powerComment" gorm:"column:power_comment;comment:圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                     //圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerAddUser     string `json:"powerAddUser" form:"powerAddUser" gorm:"column:power_add_user;comment:圈子加入权限_指定部门和成员(json数组);size:500;"`                                 //圈子加入权限_指定部门和成员(json数组);size:500
	PowerViewUser    string `json:"powerViewUser" form:"powerViewUser" gorm:"column:power_view_user;comment:圈子内浏览权限_指定部门和用户(json数组);size:500;"`                            //圈子内浏览权限_指定部门和用户(json数组);size:500
	PowerPublishUser string `json:"powerPublishUser" form:"powerPublishUser" gorm:"column:power_publish_user;comment:圈子内发布权限_指定部门和用户(json数组);size:500;"`                   //圈子内发布权限_指定部门和用户(json数组);size:500
	PowerCommentUser string `json:"powerCommentUser" form:"powerCommentUser" gorm:"column:power_comment_user;comment:圈子内评论权限_指定部门和用户(json数组);size:500;"`                   //圈子内评论权限_指定部门和用户(json数组);size:500
	NoLimitUserGroup string `json:"noLimitUserGroup" form:"noLimitUserGroup" gorm:"column:no_limit_user_group;comment:不受限用户组(json数组);size:500;"`                                   //不受限用户组(json数组);size:500
	NewUserFocus     *int   `json:"newUserFocus" form:"newUserFocus" gorm:"column:new_user_focus;comment:新注册用户默认关注：0 否，1是;size:10;"`                                            //新注册用户默认关注：0 否，1是
}
type UpdateCircleBaseInfoReq struct {
	ID         uint64 `json:"id" `                                                                                 //圈子编号
	Name       string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                       //圈子名称
	Logo       string `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`                      //圈子Logo
	Slogan     string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                 //圈子标语;size:20
	Des        string `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`                        //圈子简介;size:1000
	Protocol   string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:圈子规约;size:1000;"`         //圈子规约;size:1000
	CoverImage string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:圈子背景图;size:500;"` //圈子背景图;size:500
}
type UpdateCirclePowerReq struct {
	ID               uint64 `json:"id" `                                                                                                                                                   //圈子编号
	Process          *int   `json:"process" form:"process" gorm:"column:process;comment:是否开启版块内容人工审核：0 否，1是;size:10;"`                                                       //是否开启版块内容人工审核：0 否，1是
	Property         *int   `json:"property" form:"property" gorm:"column:property;comment:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）;size:10;"`                            //:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）
	View             *int   `json:"view" form:"view" gorm:"column:view;comment:板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到;size:10;"` //板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到
	PowerAdd         *int   `json:"powerAdd" form:"powerAdd" gorm:"column:power_add;comment:圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户;size:10;"`                     //圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户
	PowerView        *int   `json:"powerView" form:"powerView" gorm:"column:power_view;comment:圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                              //圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerPublish     *int   `json:"powerPublish" form:"powerPublish" gorm:"column:power_publish;comment:圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                     //圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerComment     *int   `json:"powerComment" form:"powerComment" gorm:"column:power_comment;comment:圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                     //圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerAddUser     string `json:"powerAddUser" form:"powerAddUser" gorm:"column:power_add_user;comment:圈子加入权限_指定部门和成员(json数组);size:500;"`                                 //圈子加入权限_指定部门和成员(json数组);size:500
	PowerViewUser    string `json:"powerViewUser" form:"powerViewUser" gorm:"column:power_view_user;comment:圈子内浏览权限_指定部门和用户(json数组);size:500;"`                            //圈子内浏览权限_指定部门和用户(json数组);size:500
	PowerPublishUser string `json:"powerPublishUser" form:"powerPublishUser" gorm:"column:power_publish_user;comment:圈子内发布权限_指定部门和用户(json数组);size:500;"`                   //圈子内发布权限_指定部门和用户(json数组);size:500
	PowerCommentUser string `json:"powerCommentUser" form:"powerCommentUser" gorm:"column:power_comment_user;comment:圈子内评论权限_指定部门和用户(json数组);size:500;"`                   //圈子内评论权限_指定部门和用户(json数组);size:500
	NoLimitUserGroup string `json:"noLimitUserGroup" form:"noLimitUserGroup" gorm:"column:no_limit_user_group;comment:不受限用户组(json数组);size:500;"`                                   //不受限用户组(json数组);size:500
	NewUserFocus     *int   `json:"newUserFocus" form:"newUserFocus" gorm:"column:new_user_focus;comment:新注册用户默认关注：0 否，1是;size:10;"`                                            //新注册用户默认关注：0 否，1是
}

type ParamSetCircleChannel struct {
	CircleId   uint64 `json:"circleId" form:"circleId" `                  //圈子_编号
	ChannelIds string `json:"channelIds" form:"channelIds" example:"1,2"` //频道编号，通过逗号分割
}

type ParamSetCircleTagSort struct {
	CircleId uint64 `json:"circleId" form:"circleId" `          //圈子_编号
	TagIds   string `json:"tagIds" form:"tagIds" example:"1,2"` //标签编号，通过逗号分割
}
type ParamAddCircleChannel struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	Name     string `json:"name" form:"name" `         //名称
}
type ParamDeleteCircleChannel struct {
	CircleId uint64   `json:"circleId" form:"circleId" ` //圈子_编号
	Ids      []uint64 `json:"ids" form:"ids"`
}

type ParamGetCircleChannel struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
}

type ParamCreateCircleTag struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
	Name     string `json:"name" form:"name"`          //标签名称
}

type ParamDeleteCircleTags struct {
	CircleId uint64   `json:"circleId" form:"circleId" ` //圈子_编号
	Ids      []uint64 `json:"ids" form:"ids"`
}
type ParamGetCircleTags struct {
	CircleId uint64 `json:"circleId" form:"circleId" ` //圈子_编号
}

type ParamRequestBecomeChildCircle struct {
	ParentCircleId uint64 `json:"parentCircleId" form:"parentCircleId" ` //父圈子编号
	CircleId       uint64 `json:"circleId" form:"circleId" `             //圈子_编号
	Des            string `json:"des" form:"des"`                        //描述
}
type ParamApproveChildCircleRequest struct {
	ID          uint64 `json:"id" `            //申请编号
	CheckStatus int    `json:"checkStatus" `   //审批状态: 1同意，2拒绝
	Des         string `json:"des" form:"des"` //描述
}
type ParamCreateChildCircle struct {
	CircleId         uint64 `json:"circleId" form:"circleId" `                                                                                        //圈子_编号
	Name             string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                                                    //圈子名称
	Logo             string `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`                                                   //圈子Logo
	CircleClassifyId uint64 `json:"circleClassifyId" form:"circleClassifyId" gorm:"type:bigint(20);column:circle_classify_id;comment:圈子分类_编号;"` //圈子分类_编号
	Slogan           string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                                              //圈子标语
	Des              string `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`                                                     //圈子简介
	Protocol         string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:圈子规约;size:1000;"`                                      //圈子规约
	CoverImage       string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:圈子背景图;size:500;"`                              //圈子背景图
}
type ParamDeleteChildCircle struct {
	CircleId      uint64 `json:"circleId" form:"circleId" `           //圈子_编号
	ChildCircleId uint64 `json:"childCircleId" form:"childCircleId" ` //子圈子编号
}
