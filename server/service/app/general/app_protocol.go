package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/general/request"
)

type AppProtocolService struct {
}

// CreateHkProtocol 创建HkProtocol记录
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) CreateHkProtocol(hkProtocol general.HkProtocol) (err error) {
	err = global.GVA_DB.Create(&hkProtocol).Error
	return err
}

// DeleteHkProtocol 删除HkProtocol记录
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) DeleteHkProtocol(hkProtocol general.HkProtocol) (err error) {
	err = global.GVA_DB.Delete(&hkProtocol).Error
	return err
}

// DeleteHkProtocolByIds 批量删除HkProtocol记录
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) DeleteHkProtocolByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.HkProtocol{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkProtocol 更新HkProtocol记录
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) UpdateHkProtocol(hkProtocol general.HkProtocol) (err error) {
	err = global.GVA_DB.Save(&hkProtocol).Error
	return err
}

// GetHkProtocol 根据id获取HkProtocol记录
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) GetHkProtocol(id uint) (hkProtocol general.HkProtocol, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkProtocol).Error
	return
}

// GetHkProtocolInfoList 分页获取HkProtocol记录
// Author [piexlmax](https://github.com/piexlmax)
func (appProtocolService *AppProtocolService) GetHkProtocolInfoList(info generalReq.HkProtocolSearch) (list []general.HkProtocol, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&general.HkProtocol{})
	var hkProtocols []general.HkProtocol
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkProtocols).Error
	return hkProtocols, total, err
}
