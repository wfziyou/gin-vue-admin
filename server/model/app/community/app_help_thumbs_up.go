package community

// HelpThumbsUp 结构体
type HelpThumbsUp struct {
	UserId uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户id;"` //用户id
	HelpId uint64 `json:"helpId" form:"helpId" gorm:"type:bigint(20);column:help_id;comment:帖子编号;"` //帮助编号
	Good   int    `json:"good" form:"good" gorm:"column:good;comment:是否好评：0否 1是;size:10;"`          //是否好评：0否 1是
	Bad    int    `json:"bad" form:"bad" gorm:"column:bad;comment:是否差评：0否 1是;size:10;"`             //是否好评：0否 1是
}

// TableName ForumThumbsUp 表名
func (HelpThumbsUp) TableName() string {
	return "hk_help_thumbs_up"
}
