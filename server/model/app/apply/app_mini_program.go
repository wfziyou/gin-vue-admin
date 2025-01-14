// 自动生成模板MiniProgram
package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MiniProgram 结构体
type MiniProgram struct {
	global.GvaModelApp
	Name             string `json:"name" form:"name" gorm:"column:name;comment:名称;size:80;"`                                            //名称
	Icon             string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                                           //图标
	CompanyName      string `json:"companyName" form:"companyName" gorm:"column:company_name;comment:公司名称;size:256;"`                   //公司名称
	ProgramId        string `json:"programId" form:"programId" gorm:"column:program_id;comment:小程序id;size:80;"`                         //小程序id
	CurPacketId      uint64 `json:"curPacketId" form:"curPacketId" gorm:"column:cur_packet_id;comment:当前包id;size:12;"`                  //当前包id
	CurVersion       string `json:"curVersion" form:"curVersion" gorm:"column:cur_version;comment:当前版本;size:12;"`                       //当前版本
	CurPacketAddress string `json:"curPacketAddress" form:"curPacketAddress" gorm:"column:cur_packet_address;comment:当前包下载地址;size:12;"` //当前包下载地址
	Hidden           *int   `json:"hidden" form:"hidden" gorm:"column:hidden;comment:隐藏(0否 1是);size:10;"`                               //隐藏(0否 1是)
}
type MiniProgramBaseInfo struct {
	Id            uint64 `json:"id" form:"id" gorm:"primarykey"`                                                           // 主键ID
	Name          string `json:"name" form:"name" gorm:"column:name;comment:名称;size:80;"`                                  //名称
	Icon          string `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:256;"`                                 //图标
	CompanyName   string `json:"companyName" form:"companyName" gorm:"column:company_name;comment:公司名称;size:256;"`         //公司名称
	ProgramId     string `json:"programId" form:"programId" gorm:"column:program_id;comment:小程序id;size:80;"`               //小程序id
	Version       string `json:"version" form:"version" gorm:"column:version;comment:当前版本;size:12;"`                       //当前版本
	Code          uint64 `json:"code" form:"code" gorm:"column:code;comment:版本code;size:12;"`                              //版本code
	PacketAddress string `json:"packetAddress" form:"packetAddress" gorm:"column:packet_address;comment:当前包下载地址;size:12;"` //当前包下载地址
	Hidden        *int   `json:"hidden" form:"hidden" gorm:"column:hidden;comment:隐藏(0否 1是);size:10;"`                     //隐藏(0否 1是)
}

// TableName MiniProgram 表名
func (MiniProgram) TableName() string {
	return "hk_mini_program"
}
