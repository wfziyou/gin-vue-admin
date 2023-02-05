// 自动生成模板HkMiniProgramPacket
package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HkMiniProgramPacket 结构体
type HkMiniProgramPacket struct {
	global.GVA_MODEL
	TenantId      string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;size:12;"`
	MiniProgramId *int   `json:"miniProgramId" form:"miniProgramId" gorm:"column:mini_program_id;comment:小程序id;size:19;"`
	Name          string `json:"name" form:"name" gorm:"column:name;comment:包名;size:80;"`
	Version       string `json:"version" form:"version" gorm:"column:version;comment:版本;size:80;"`
	PacketAddress string `json:"packetAddress" form:"packetAddress" gorm:"column:packet_address;comment:访问地址;size:256;"`
	CreateUser    *int   `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;size:19;"`
	CreateDept    *int   `json:"createDept" form:"createDept" gorm:"column:create_dept;comment:创建部门;size:19;"`
}

// TableName HkMiniProgramPacket 表名
func (HkMiniProgramPacket) TableName() string {
	return "hk_mini_program_packet"
}
