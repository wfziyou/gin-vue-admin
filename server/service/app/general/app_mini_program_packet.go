package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppMiniProgramPacketService struct {
}

// CreateMiniProgramPacket 创建MiniProgramPacket记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMiniProgramPacketService *AppMiniProgramPacketService) CreateMiniProgramPacket(hkMiniProgramPacket general.MiniProgramPacket) (err error) {
	err = global.GVA_DB.Create(&hkMiniProgramPacket).Error
	return err
}

// DeleteMiniProgramPacket 删除MiniProgramPacket记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMiniProgramPacketService *AppMiniProgramPacketService) DeleteMiniProgramPacket(hkMiniProgramPacket general.MiniProgramPacket) (err error) {
	err = global.GVA_DB.Delete(&hkMiniProgramPacket).Error
	return err
}

// DeleteMiniProgramPacketByIds 批量删除MiniProgramPacket记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMiniProgramPacketService *AppMiniProgramPacketService) DeleteMiniProgramPacketByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.MiniProgramPacket{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMiniProgramPacket 更新MiniProgramPacket记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMiniProgramPacketService *AppMiniProgramPacketService) UpdateMiniProgramPacket(hkMiniProgramPacket general.MiniProgramPacket) (err error) {
	err = global.GVA_DB.Save(&hkMiniProgramPacket).Error
	return err
}

// GetMiniProgramPacket 根据id获取MiniProgramPacket记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMiniProgramPacketService *AppMiniProgramPacketService) GetMiniProgramPacket(id uint) (hkMiniProgramPacket general.MiniProgramPacket, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkMiniProgramPacket).Error
	return
}

// GetMiniProgramPacketInfoList 分页获取MiniProgramPacket记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMiniProgramPacketService *AppMiniProgramPacketService) GetMiniProgramPacketInfoList(info generalReq.MiniProgramPacketSearch) (list []general.MiniProgramPacket, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&general.MiniProgramPacket{})
	var hkMiniProgramPackets []general.MiniProgramPacket
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkMiniProgramPackets).Error
	return hkMiniProgramPackets, total, err
}
