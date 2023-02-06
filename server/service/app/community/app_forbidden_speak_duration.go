package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type AppForbiddenSpeakDurationService struct {
}

// CreateHkForbiddenSpeakDuration 创建HkForbiddenSpeakDuration记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakDurationService *AppForbiddenSpeakDurationService) CreateHkForbiddenSpeakDuration(hkForbiddenSpeakDuration community.HkForbiddenSpeakDuration) (err error) {
	err = global.GVA_DB.Create(&hkForbiddenSpeakDuration).Error
	return err
}

// DeleteHkForbiddenSpeakDuration 删除HkForbiddenSpeakDuration记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakDurationService *AppForbiddenSpeakDurationService) DeleteHkForbiddenSpeakDuration(hkForbiddenSpeakDuration community.HkForbiddenSpeakDuration) (err error) {
	err = global.GVA_DB.Delete(&hkForbiddenSpeakDuration).Error
	return err
}

// DeleteHkForbiddenSpeakDurationByIds 批量删除HkForbiddenSpeakDuration记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakDurationService *AppForbiddenSpeakDurationService) DeleteHkForbiddenSpeakDurationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForbiddenSpeakDuration{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkForbiddenSpeakDuration 更新HkForbiddenSpeakDuration记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakDurationService *AppForbiddenSpeakDurationService) UpdateHkForbiddenSpeakDuration(hkForbiddenSpeakDuration community.HkForbiddenSpeakDuration) (err error) {
	err = global.GVA_DB.Save(&hkForbiddenSpeakDuration).Error
	return err
}

// GetHkForbiddenSpeakDuration 根据id获取HkForbiddenSpeakDuration记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakDurationService *AppForbiddenSpeakDurationService) GetHkForbiddenSpeakDuration(id uint) (hkForbiddenSpeakDuration community.HkForbiddenSpeakDuration, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForbiddenSpeakDuration).Error
	return
}

// GetHkForbiddenSpeakDurationInfoList 分页获取HkForbiddenSpeakDuration记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakDurationService *AppForbiddenSpeakDurationService) GetHkForbiddenSpeakDurationInfoList(info communityReq.HkForbiddenSpeakDurationSearch) (list []community.HkForbiddenSpeakDuration, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkForbiddenSpeakDuration{})
	var hkForbiddenSpeakDurations []community.HkForbiddenSpeakDuration
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForbiddenSpeakDurations).Error
	return hkForbiddenSpeakDurations, total, err
}
