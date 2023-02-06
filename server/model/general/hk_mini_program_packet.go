// 自动生成模板HkMiniProgramPacket
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkMiniProgramPacket 结构体
type HkMiniProgramPacket struct {
	global.GVA_MODEL
	MiniProgramId *int   `json:"miniProgramId" form:"miniProgramId" gorm:"column:mini_program_id;comment:小程序id;size:19;"` //小程序id
	Name          string `json:"name" form:"name" gorm:"column:name;comment:包名;size:80;"`                                 //包名
	Version       string `json:"version" form:"version" gorm:"column:version;comment:版本;size:80;"`                        //版本
	PacketAddress string `json:"packetAddress" form:"packetAddress" gorm:"column:packet_address;comment:访问地址;size:256;"`  //访问地址
}

// TableName HkMiniProgramPacket 表名
func (HkMiniProgramPacket) TableName() string {
	return "hk_mini_program_packet"
}
