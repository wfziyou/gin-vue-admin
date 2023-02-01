package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/general"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/general/request"
)

type HkMiniProgramPacketService struct {
}

// CreateHkMiniProgramPacket 创建HkMiniProgramPacket记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkMiniProgramPacketService *HkMiniProgramPacketService) CreateHkMiniProgramPacket(hkMiniProgramPacket general.HkMiniProgramPacket) (err error) {
	err = global.GVA_DB.Create(&hkMiniProgramPacket).Error
	return err
}

// DeleteHkMiniProgramPacket 删除HkMiniProgramPacket记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkMiniProgramPacketService *HkMiniProgramPacketService)DeleteHkMiniProgramPacket(hkMiniProgramPacket general.HkMiniProgramPacket) (err error) {
	err = global.GVA_DB.Delete(&hkMiniProgramPacket).Error
	return err
}

// DeleteHkMiniProgramPacketByIds 批量删除HkMiniProgramPacket记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkMiniProgramPacketService *HkMiniProgramPacketService)DeleteHkMiniProgramPacketByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.HkMiniProgramPacket{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkMiniProgramPacket 更新HkMiniProgramPacket记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkMiniProgramPacketService *HkMiniProgramPacketService)UpdateHkMiniProgramPacket(hkMiniProgramPacket general.HkMiniProgramPacket) (err error) {
	err = global.GVA_DB.Save(&hkMiniProgramPacket).Error
	return err
}

// GetHkMiniProgramPacket 根据id获取HkMiniProgramPacket记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkMiniProgramPacketService *HkMiniProgramPacketService)GetHkMiniProgramPacket(id uint) (hkMiniProgramPacket general.HkMiniProgramPacket, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkMiniProgramPacket).Error
	return
}

// GetHkMiniProgramPacketInfoList 分页获取HkMiniProgramPacket记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkMiniProgramPacketService *HkMiniProgramPacketService)GetHkMiniProgramPacketInfoList(info generalReq.HkMiniProgramPacketSearch) (list []general.HkMiniProgramPacket, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&general.HkMiniProgramPacket{})
    var hkMiniProgramPackets []general.HkMiniProgramPacket
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkMiniProgramPackets).Error
	return  hkMiniProgramPackets, total, err
}
