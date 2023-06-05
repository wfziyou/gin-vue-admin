package community

// HelpThumbsUp 结构体
type HelpThumbsUp struct {
	UserId   uint64 `json:"userId" form:"userId" gorm:"type:bigint(20);column:user_id;comment:用户id;"`            //用户id
	HelpId   uint64 `json:"helpId" form:"helpId" gorm:"type:bigint(20);column:help_id;comment:帖子编号;"`            //帮助编号
	ThumbsUp int    `json:"thumbsUp" form:"thumbsUp" gorm:"column:thumbs_up;comment:是否好评：0未点赞 1有用 2无用;size:10;"` //是否好评：0未点赞 1有用 2无用
}

// TableName ForumThumbsUp 表名
func (HelpThumbsUp) TableName() string {
	return "hk_help_thumbs_up"
}
