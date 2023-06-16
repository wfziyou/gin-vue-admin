package general

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type OpenScreenAdvertising struct {
	global.GvaModelAppEx
	Content  string `json:"content" form:"content" gorm:"column:content;comment:内容;size:256;"`           //内容
	Link     string `json:"link" form:"link" gorm:"column:link;comment:链接;size:256;"`                    //链接
	ShowTime int    `json:"showTime" form:"showTime" gorm:"column:show_time;comment:显示时间，单位毫秒;size:10;"` //显示时间，单位毫秒
	Type     int    `json:"type" form:"type" gorm:"column:type;comment:类型:0图片、1视频;size:10;"`             //类型:0图片、1视频
	OwnerId  uint64 `json:"ownerId" form:"ownerId" gorm:"column:owner_id;comment:拥有者id;size:19;"`        //拥有者id
}

// TableName OpenScreenAdvertising 表名
func (OpenScreenAdvertising) TableName() string {
	return "hk_open_screen_advertising"
}
