// 自动生成模板User
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	uuid "github.com/satori/go.uuid"
)

// User 结构体
type User struct {
	global.GvaModelApp
	Uuid        uuid.UUID             `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户编号;size:50;"`                                //用户编号
	UserType    int                   `json:"userType" form:"userType" gorm:"column:user_type;comment:用户平台: 1web、2app、3other;size:10;"` //用户平台: 1web、2app、3other
	Account     string                `json:"account" form:"account" gorm:"column:account;comment:账号;size:45;"`                         //账号
	Password    string                `json:"password" form:"password" gorm:"column:password;comment:密码;size:80;"`                      //密码
	NickName    string                `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:昵称;size:20;"`                     //昵称
	RealName    string                `json:"realName" form:"realName" gorm:"column:real_name;comment:真名;size:10;"`                     //真名
	HeaderImg   string                `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:头像;size:500;"`                 //头像
	Email       string                `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:45;"`                               //邮箱
	Phone       string                `json:"phone" form:"phone" gorm:"column:phone;comment:手机;size:45;"`                               //手机
	Birthday    string                `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日;size:20;"`
	Sex         int                   `json:"sex" form:"sex" gorm:"column:sex;comment:性别： 0未知、1男、2女;size:10;"`                //性别： 0未知、1男、2女
	Description string                `json:"description" form:"description" gorm:"column:description;comment:描述;size:45;"`   //描述
	ImToken     string                `json:"imToken" form:"imToken" gorm:"-"`                                                //Im回话
	AuthorityId uint64                `json:"roleId" form:"roleId" gorm:"column:role_id;default:888;comment:用户角色ID;size:20;"` // 用户角色ID
	Authority   system.SysAuthority   `json:"authority" gorm:"foreignKey:Id;references:AuthorityId;comment:用户角色"`             //用户角色
	Authorities []system.SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	UserExtend  UserExtend            `json:"userExtend" gorm:"foreignKey:ID;references:ID;comment:用户扩展"` //用户扩展
}

// TableName User 表名
func (User) TableName() string {
	return "hk_user"
}

type UserBaseInfo struct {
	ID          uint64 `json:"id" form:"id" gorm:"primarykey"`                                                           // 主键ID
	UserType    int    `json:"userType" form:"userType" gorm:"column:user_type;comment:用户平台: 1web、2app、3other;size:10;"` //用户平台: 1web、2app、3other
	Account     string `json:"account" form:"account" gorm:"column:account;comment:账号;size:45;"`                         //账号
	NickName    string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:昵称;size:20;"`                     //昵称
	RealName    string `json:"realName" form:"realName" gorm:"column:real_name;comment:真名;size:10;"`                     //真名
	HeaderImg   string `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:头像;size:500;"`                 //头像
	Birthday    string `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日;size:20;"`
	Sex         int    `json:"sex" form:"sex" gorm:"column:sex;comment:性别： 0未知、1男、2女;size:10;"`              //性别： 0未知、1男、2女
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述;size:45;"` //描述
	IsFocus     int    `json:"isFocus" form:"isFocus" gorm:"-"`                                              //是否关注:0 否、1是
	IsFan       int    `json:"isFan" form:"isFan" gorm:"-"`                                                  //是否是粉丝:0 否、1是
}
type UserInfo struct {
	UserBaseInfo
	Email         string `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:45;"`                                     //邮箱
	Phone         string `json:"phone" form:"phone" gorm:"column:phone;comment:手机;size:45;"`                                     //手机
	NumCircle     int64  `json:"numCircle" form:"numCircle" gorm:"type:bigint(20);column:num_circle;comment:圈子数;"`               //圈子数
	NumFocus      int64  `json:"numFocus" form:"numFocus" gorm:"type:bigint(20);column:num_focus;comment:关注数;"`                  //关注数
	NumFan        int64  `json:"numFan" form:"numFan" gorm:"type:bigint(20);column:num_fan;comment:粉丝数;"`                        //粉丝数
	CoverImage    string `json:"coverImage" form:"coverImage" gorm:"column:cover_image;comment:主页封面;size:256;"`                  //主页封面
	CurrencyMoney uint64 `json:"currencyMoney" form:"currencyMoney" gorm:"type:bigint(20);column:currency_money;comment:货币_零钱;"` //货币_零钱
	CurrencyGold  uint64 `json:"currencyGold" form:"currencyGold" gorm:"type:bigint(20);column:currency_gold;comment:货币_金币;"`    //货币_金币
}

// TableName User 表名
func (UserBaseInfo) TableName() string {
	return "hk_user"
}
