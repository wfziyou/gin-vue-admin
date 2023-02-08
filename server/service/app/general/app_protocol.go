package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppProtocolService struct {
}

// CreateProtocol 创建Protocol记录
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) CreateProtocol(hkProtocol general.Protocol) (err error) {
	err = global.GVA_DB.Create(&hkProtocol).Error
	return err
}

// DeleteProtocol 删除Protocol记录
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) DeleteProtocol(hkProtocol general.Protocol) (err error) {
	err = global.GVA_DB.Delete(&hkProtocol).Error
	return err
}

// DeleteProtocolByIds 批量删除Protocol记录
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) DeleteProtocolByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.Protocol{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProtocol 更新Protocol记录
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) UpdateProtocol(hkProtocol general.Protocol) (err error) {
	err = global.GVA_DB.Save(&hkProtocol).Error
	return err
}

// GetProtocol 根据id获取Protocol记录
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) GetProtocol(id uint64) (hkProtocol general.Protocol, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkProtocol).Error
	return
}

// GetProtocolByName 用名字查询协议
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) GetProtocolByName(name string) (hkProtocol general.Protocol, err error) {
	err = global.GVA_DB.Where("name = ?", name).First(&hkProtocol).Error
	return
}

// GetProtocolInfoList 分页获取Protocol记录
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) GetProtocolInfoList(info generalReq.ProtocolSearch) (list []general.Protocol, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&general.Protocol{})
	var hkProtocols []general.Protocol
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if len(info.Name) > 0 {
		db = db.Where("name = ?", info.Name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkProtocols).Error
	return hkProtocols, total, err
}
