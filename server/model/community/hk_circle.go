// 自动生成模板HkCircle
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// HkCircle 结构体
type HkCircle struct {
	global.GVA_MODEL
	TenantId         string     `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`                                                             //租户ID
	Type             *int       `json:"type" form:"type" gorm:"column:type;comment:类型：0官方圈子、1用户圈子、2小区;size:10;"`                                                            //类型：0官方圈子、1用户圈子、2小区
	Name             string     `json:"name" form:"name" gorm:"column:name;comment:圈子名称;size:20;"`                                                                          //圈子名称
	Logo             string     `json:"logo" form:"logo" gorm:"column:logo;comment:圈子Logo;size:500;"`                                                                       //圈子Logo
	CircleClassifyId *int       `json:"circleClassifyId" form:"circleClassifyId" gorm:"column:circle_classify_id;comment:圈子分类_编号;size:19;"`                                 //圈子分类_编号
	Slogan           string     `json:"slogan" form:"slogan" gorm:"column:slogan;comment:圈子标语;size:20;"`                                                                    //圈子标语;size:20
	Des              string     `json:"des" form:"des" gorm:"column:des;comment:圈子简介;size:1000;"`                                                                           //圈子简介;size:1000
	Protocol         string     `json:"protocol" form:"protocol" gorm:"column:protocol;comment:圈子规约;size:1000;"`                                                            //圈子规约;size:1000
	BackImage        string     `json:"backImage" form:"backImage" gorm:"column:back_image;comment:圈子背景图;size:500;"`                                                        //圈子背景图;size:500
	SupportCategory  string     `json:"supportCategory" form:"supportCategory" gorm:"column:support_category;comment:支持内容类别(json数组)：0视频、1动态、2资讯、3公告、4文章、5问答、6建议;size:500;"` //支持内容类别(json数组)：0视频、1动态、2资讯、3公告、4文章、5问答、6建议;size:500;
	Pay              *int       `json:"pay" form:"pay" gorm:"column:pay;comment:付费：0 否，1是;size:10;"`                                                                        //付费：0 否，1是;size:10
	Process          *int       `json:"process" form:"process" gorm:"column:process;comment:是否开启版块内容人工审核：0 否，1是;size:10;"`                                                  //是否开启版块内容人工审核：0 否，1是
	Property         *int       `json:"property" form:"property" gorm:"column:property;comment:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）;size:10;"`                               //:圈子属性： 0公开（自由加入），1公开（审核加入），2私密（邀请加入）
	View             *int       `json:"view" form:"view" gorm:"column:view;comment:板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到;size:10;"`                         //板块可见性： 0不在社区中显示，不能被搜索到，1不在社区中显示，可以被搜索到，2在社区中显示，可以被搜索到
	PowerAdd         *int       `json:"powerAdd" form:"powerAdd" gorm:"column:power_add;comment:圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户;size:10;"`                             //圈子加入权限：0 所有人，1指定用户组，2指定部门和成员，3仅邀请的用户
	PowerView        *int       `json:"powerView" form:"powerView" gorm:"column:power_view;comment:圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                                //圈子内浏览权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerPublish     *int       `json:"powerPublish" form:"powerPublish" gorm:"column:power_publish;comment:圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                       //圈子内发布权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerComment     *int       `json:"powerComment" form:"powerComment" gorm:"column:power_comment;comment:圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组;size:10;"`                       //圈子内评论权限：0 所有人，1版块用户，2版主，3指定用户组
	PowerAddUser     string     `json:"powerAddUser" form:"powerAddUser" gorm:"column:power_add_user;comment:圈子加入权限_指定部门和成员(json数组);size:500;"`                             //圈子加入权限_指定部门和成员(json数组);size:500
	PowerViewUser    string     `json:"powerViewUser" form:"powerViewUser" gorm:"column:power_view_user;comment:圈子内浏览权限_指定部门和用户(json数组);size:500;"`                         //圈子内浏览权限_指定部门和用户(json数组);size:500
	PowerPublishUser string     `json:"powerPublishUser" form:"powerPublishUser" gorm:"column:power_publish_user;comment:圈子内发布权限_指定部门和用户(json数组);size:500;"`                //圈子内发布权限_指定部门和用户(json数组);size:500
	PowerCommentUser string     `json:"powerCommentUser" form:"powerCommentUser" gorm:"column:power_comment_user;comment:圈子内评论权限_指定部门和用户(json数组);size:500;"`                //圈子内评论权限_指定部门和用户(json数组);size:500
	NoLimitUserGroup string     `json:"noLimitUserGroup" form:"noLimitUserGroup" gorm:"column:no_limit_user_group;comment:不受限用户组(json数组);size:500;"`                        //不受限用户组(json数组);size:500
	NewUserFocus     *int       `json:"newUserFocus" form:"newUserFocus" gorm:"column:new_user_focus;comment:新注册用户默认关注：0 否，1是;size:10;"`                                    //新注册用户默认关注：0 否，1是
	Sort             *int       `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                                                            //排序
	CreateUser       *int       `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept       *int       `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
	CreateTime       *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
	UpdateUser       *int       `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改人;size:19;"`
	UpdateTime       *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
	Status           *int       `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	IsDeleted        *int       `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;size:10;"`
}

// TableName HkCircle 表名
func (HkCircle) TableName() string {
	return "hk_circle"
}
