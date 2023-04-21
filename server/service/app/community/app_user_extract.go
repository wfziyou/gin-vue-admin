package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HkUserExtractService struct {
}

// CreateHkUserExtract 创建HkUserExtract记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserExtractService *HkUserExtractService) CreateHkUserExtract(hkUserExtract *community.HkUserExtract) (err error) {
	err = global.GVA_DB.Create(hkUserExtract).Error
	return err
}

// DeleteHkUserExtract 删除HkUserExtract记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserExtractService *HkUserExtractService) DeleteHkUserExtract(hkUserExtract community.HkUserExtract) (err error) {
	err = global.GVA_DB.Delete(&hkUserExtract).Error
	return err
}

// DeleteHkUserExtractByIds 批量删除HkUserExtract记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserExtractService *HkUserExtractService) DeleteHkUserExtractByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkUserExtract{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkUserExtract 更新HkUserExtract记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserExtractService *HkUserExtractService) UpdateHkUserExtract(hkUserExtract community.HkUserExtract) (err error) {
	err = global.GVA_DB.Save(&hkUserExtract).Error
	return err
}

// GetHkUserExtract 根据id获取HkUserExtract记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserExtractService *HkUserExtractService) GetHkUserExtract(id uint) (hkUserExtract community.HkUserExtract, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserExtract).Error
	return
}

// GetHkUserExtractInfoList 分页获取HkUserExtract记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserExtractService *HkUserExtractService) GetHkUserExtractInfoList(info communityReq.HkUserExtractSearch) (list []community.HkUserExtract, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkUserExtract{})
	var hkUserExtracts []community.HkUserExtract
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUserExtracts).Error
	return hkUserExtracts, total, err
}
