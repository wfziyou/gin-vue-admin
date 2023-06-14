// 自动生成模板Circle
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

//圈子类型：0官方圈子、1用户圈子、2小区
const (
	//CircleTypeOfficial 0官方圈子
	CircleTypeOfficial = 0
	//CircleTypeUser 1用户圈子
	CircleTypeUser = 1
	//CircleTypeCommunity 2小区
	CircleTypeCommunity = 2
)

//圈子频道 1全部 2资讯 3动态 4问答 5活动 6文章
const (
	//CircleChannelAll 1全部
	CircleChannelAll = 1
	//CircleChannelNews 2资讯
	CircleChannelNews = 2
	//CircleChannelDynamic 3动态
	CircleChannelDynamic = 3
	//CircleChannelQuestion 4问答
	CircleChannelQuestion = 4
	//CircleChannelActivity 5活动
	CircleChannelActivity = 5
	//CircleChannelArticle 6文章
	CircleChannelArticle = 6
)

// Circle 结构体
type Circle struct {
	global.GvaModelApp
	Type                    int         `json:"type" form:"type" gorm:"column:type;comment:类型：0官方圈子、1用户圈子、2小区;size:10;"`                                                                           //类型：0官方圈子、1用户圈子、2小区
	Name                    string      `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                                                                                         //圈子名称
	Logo                    string      `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`                                                                                      //圈子Logo
	CircleClassifyId        uint64      `json:"circleClassifyId" form:"circleClassifyId" gorm:"type:bigint(20);column:circle_classify_id;comment:圈子分类_编号"`                                         //圈子分类_编号
	Slogan                  string      `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                                                                                   //圈子标语;size:20
	Des                     string      `json:"des" form:"des" gorm:"column:des;comment:圈子简介;"`                                                                                                    //圈子简介;size:1000
	Protocol                string      `json:"protocol" form:"protocol" gorm:"column:protocol;comment:圈子规约;size:1000;"`                                                                           //圈子规约;size:1000
	CoverImage              string      `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:圈子背景图;size:500;"`                                                                    //圈子背景图;size:500
	Pay                     int         `json:"pay" form:"pay" gorm:"column:pay;comment:付费：0否、1是;size:10;"`                                                                                        //付费：0否、1是;size:10
	Process                 int         `json:"process" form:"process" gorm:"column:process;comment:是否开启版块内容人工审核：0 否，1是;size:10;"`                                                                 //是否开启版块内容人工审核：0 否，1是
	Property                int         `json:"property" form:"property" gorm:"column:property;comment:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）;size:10;"`                                              //:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）
	View                    int         `json:"view" form:"view" gorm:"column:view;comment:板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到;size:10;"`                                        //板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到
	OutUserShowIntro        int         `json:"outUserShowIntro" form:"outUserShowIntro" gorm:"column:out_user_show_intro;default:1;comment:圈外成员是否显示介绍：0不显示、1显示;size:10;"`                         //圈外成员是否显示介绍：0不显示、1显示
	OutUserShowApply        int         `json:"outUserShowApply" form:"outUserShowApply" gorm:"column:out_user_show_apply;default:1;comment:圈外成员是否显示应用：0不显示、1能显示;size:10;"`                        //圈外成员是否显示应用：0不显示、1能显示
	OutUserLookSubCircle    int         `json:"outUserLookSubCircle" form:"outUserLookSubCircle" gorm:"column:out_user_look_sub_circle;default:1;comment:圈外成员是否查看子圈子：0不能、1能;size:10;"`             //圈外成员是否查看子圈子：0不能、1能
	OutUserShowNews         int         `json:"outUserShowNews" form:"outUserShowNews" gorm:"column:out_user_show_news;default:1;comment:圈外成员是否显示资讯：0不显示、1显示;size:10;"`                            //圈外成员是否显示资讯：0不显示、1显示
	PowerLookCircleUserList int         `json:"powerLookCircleUserList" form:"powerLookCircleUserList" gorm:"column:powerLookCircleUserList;default:2;comment:用户列表查看权限：0不能查看、1所有人、2圈内成员;size:10;"` //用户列表查看权限：0不能查看、1所有人、2圈内成员
	PowerLookCircleUserInfo int         `json:"powerLookCircleUserInfo" form:"powerLookCircleUserInfo" gorm:"column:powerLookCircleUserInfo;default:0;comment:用户信息查看权限：0不能查看、1所有人、2圈内成员;size:10;"` //用户信息查看权限：0不能查看、1所有人、2圈内成员
	PowerSearch             int         `json:"powerSearch" form:"powerSearch" gorm:"column:power_search;comment:圈子帖子搜索权限：0不能搜索、1所有人、2圈内成员;size:10;"`                                              //圈子帖子搜索权限：0不能搜索、1所有人、2圈内成员
	PowerAdd                int         `json:"powerAdd" form:"powerAdd" gorm:"column:power_add;comment:圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户;size:10;"`                                            //圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户
	PowerView               int         `json:"powerView" form:"powerView" gorm:"column:power_view;comment:圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                                               //圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerPublish            int         `json:"powerPublish" form:"powerPublish" gorm:"column:power_publish;comment:圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                                      //圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerComment            int         `json:"powerComment" form:"powerComment" gorm:"column:power_comment;comment:圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                                      //圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerAddUser            string      `json:"powerAddUser" form:"powerAddUser" gorm:"column:power_add_user;comment:圈子加入权限_指定部门和成员(json数组);"`                                                     //圈子加入权限_指定部门和成员(json数组)
	PowerViewUser           string      `json:"powerViewUser" form:"powerViewUser" gorm:"column:power_view_user;comment:圈子内浏览权限_指定部门和用户(json数组);"`                                                 //圈子内浏览权限_指定部门和用户(json数组)
	PowerPublishUser        string      `json:"powerPublishUser" form:"powerPublishUser" gorm:"column:power_publish_user;comment:圈子内发布权限_指定部门和用户(json数组);"`                                        //圈子内发布权限_指定部门和用户(json数组)
	PowerCommentUser        string      `json:"powerCommentUser" form:"powerCommentUser" gorm:"column:power_comment_user;comment:圈子内评论权限_指定部门和用户(json数组);"`                                        //圈子内评论权限_指定部门和用户(json数组)
	NoLimitUserGroup        string      `json:"noLimitUserGroup" form:"noLimitUserGroup" gorm:"column:no_limit_user_group;comment:不受限用户组(json数组);"`                                                //不受限用户组(json数组)
	NewUserFocus            int         `json:"newUserFocus" form:"newUserFocus" gorm:"column:new_user_focus;comment:新注册用户默认关注：0 否，1是;size:10;"`                                                   //新注册用户默认关注：0 否，1是
	ChannelId               string      `json:"-"gorm:"column:channel_id;comment:频道_编号;"`                                                                                                          //频道_编号
	ApplyId                 string      `json:"-" gorm:"column:apply_id;comment:应用_编号;"`                                                                                                           //应用_编号
	UserNum                 int         `json:"userNum" form:"userNum" gorm:"column:user_num;comment:用户数;size:10;"`                                                                                //用户数
	UpdateForumPostsTime    *time.Time  `json:"-" gorm:"column:update_forum_posts_time;comment:发布帖子时间;"`                                                                                           //发布帖子时间
	Sort                    int         `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                                                                           //排序
	HaveCircle              int         `json:"haveCircle" gorm:"-"`                                                                                                                               //是否在圈子里：0否、1是
	UserPower               int         `json:"userPower" gorm:"-"`                                                                                                                                //用户权限：0普通、1管理、2圈主
	Tags                    []CircleTag `json:"tags"`
}

// TableName Circle 表名
func (Circle) TableName() string {
	return "hk_circle"
}

// Circle 结构体
type CircleBaseInfo struct {
	global.GvaModelApp
	Type             int    `json:"type" form:"type" gorm:"column:type;comment:类型：0官方圈子、1用户圈子、2小区;size:10;"`                                   //类型：0官方圈子、1用户圈子、2小区
	Name             string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                                                 //圈子名称
	Logo             string `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`                                              //圈子Logo
	CircleClassifyId uint64 `json:"circleClassifyId" form:"circleClassifyId" gorm:"type:bigint(20);column:circle_classify_id;comment:圈子分类_编号"` //圈子分类_编号
	Slogan           string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                                           //圈子标语;size:20
	Pay              int    `json:"pay" form:"pay" gorm:"column:pay;comment:付费：0否、1是;size:10;"`                                                //付费：0否、1是;size:10
	Process          int    `json:"process" form:"process" gorm:"column:process;comment:是否开启版块内容人工审核：0 否，1是;size:10;"`                         //是否开启版块内容人工审核：0 否，1是
	Property         int    `json:"property" form:"property" gorm:"column:property;comment:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）;size:10;"`      //:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）
	UserNum          int    `json:"userNum" form:"userNum" gorm:"column:user_num;comment:用户数;size:10;"`                                        //用户数
	Sort             int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                                   //排序
	HaveCircle       int    `json:"haveCircle" gorm:"-"`                                                                                       //是在圈子里：0否、1是
}
type CircleInfo struct {
	global.GvaModelAppEx
	Name       string `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                      //圈子名称
	Logo       string `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`                   //圈子Logo
	Slogan     string `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                //圈子标语;size:20
	Des        string `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`                       //圈子简介;size:1000
	Protocol   string `json:"protocol" form:"protocol" gorm:"column:protocol;comment:圈子规约;size:1000;"`        //圈子规约;size:1000
	CoverImage string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:圈子背景图;size:500;"` //圈子背景图;size:500
}
type CirclePower struct {
	global.GvaModelAppEx
	Process                 int    `json:"process" form:"process" gorm:"column:process;comment:是否开启版块内容人工审核：0 否，1是;size:10;"`                                                                 //是否开启版块内容人工审核：0 否，1是
	Property                int    `json:"property" form:"property" gorm:"column:property;comment:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）;size:10;"`                                              //:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）
	View                    int    `json:"view" form:"view" gorm:"column:view;comment:板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到;size:10;"`                                        //板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到
	OutUserShowIntro        int    `json:"outUserShowIntro" form:"outUserShowIntro" gorm:"column:out_user_show_intro;default:1;comment:圈外成员是否显示介绍：0不显示、1显示;size:10;"`                         //圈外成员是否显示介绍：0不显示、1显示
	OutUserShowApply        int    `json:"outUserShowApply" form:"outUserShowApply" gorm:"column:out_user_show_apply;default:1;comment:圈外成员是否显示应用：0不显示、1能显示;size:10;"`                        //圈外成员是否显示应用：0不显示、1能显示
	OutUserLookSubCircle    int    `json:"outUserLookSubCircle" form:"outUserLookSubCircle" gorm:"column:out_user_look_sub_circle;default:1;comment:圈外成员是否查看子圈子：0不能、1能;size:10;"`             //圈外成员是否查看子圈子：0不能、1能
	OutUserShowNews         int    `json:"outUserShowNews" form:"outUserShowNews" gorm:"column:out_user_show_news;default:1;comment:圈外成员是否显示资讯：0不显示、1显示;size:10;"`                            //圈外成员是否显示资讯：0不显示、1显示
	PowerLookCircleUserList int    `json:"powerLookCircleUserList" form:"powerLookCircleUserList" gorm:"column:powerLookCircleUserList;default:2;comment:用户列表查看权限：0不能查看、1所有人、2圈内成员;size:10;"` //用户列表查看权限：0不能查看、1所有人、2圈内成员
	PowerLookCircleUserInfo int    `json:"powerLookCircleUserInfo" form:"powerLookCircleUserInfo" gorm:"column:powerLookCircleUserInfo;default:0;comment:用户信息查看权限：0不能查看、1所有人、2圈内成员;size:10;"` //用户信息查看权限：0不能查看、1所有人、2圈内成员
	PowerSearch             int    `json:"powerSearch" form:"powerSearch" gorm:"column:power_search;comment:圈子帖子搜索权限：0不能搜索、1所有人、2圈内成员;size:10;"`                                              //圈子帖子搜索权限：0不能搜索、1所有人、2圈内成员
	PowerAdd                int    `json:"powerAdd" form:"powerAdd" gorm:"column:power_add;comment:圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户;size:10;"`                                            //圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户
	PowerView               int    `json:"powerView" form:"powerView" gorm:"column:power_view;comment:圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                                               //圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerPublish            int    `json:"powerPublish" form:"powerPublish" gorm:"column:power_publish;comment:圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                                      //圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerComment            int    `json:"powerComment" form:"powerComment" gorm:"column:power_comment;comment:圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                                      //圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerAddUser            string `json:"powerAddUser" form:"powerAddUser" gorm:"column:power_add_user;comment:圈子加入权限_指定部门和成员(json数组);size:500;"`                                            //圈子加入权限_指定部门和成员(json数组);size:500
	PowerViewUser           string `json:"powerViewUser" form:"powerViewUser" gorm:"column:power_view_user;comment:圈子内浏览权限_指定部门和用户(json数组);size:500;"`                                        //圈子内浏览权限_指定部门和用户(json数组);size:500
	PowerPublishUser        string `json:"powerPublishUser" form:"powerPublishUser" gorm:"column:power_publish_user;comment:圈子内发布权限_指定部门和用户(json数组);size:500;"`                               //圈子内发布权限_指定部门和用户(json数组);size:500
	PowerCommentUser        string `json:"powerCommentUser" form:"powerCommentUser" gorm:"column:power_comment_user;comment:圈子内评论权限_指定部门和用户(json数组);size:500;"`                               //圈子内评论权限_指定部门和用户(json数组);size:500
	NoLimitUserGroup        string `json:"noLimitUserGroup" form:"noLimitUserGroup" gorm:"column:no_limit_user_group;comment:不受限用户组(json数组);size:500;"`                                       //不受限用户组(json数组);size:500
}

// TableName Circle 表名
func (CircleBaseInfo) TableName() string {
	return "hk_circle"
}
