package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForbiddenSpeakDurationService struct {
}

// CreateForbiddenSpeakDuration 创建ForbiddenSpeakDuration记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakDurationService *AppForbiddenSpeakDurationService) CreateForbiddenSpeakDuration(hkForbiddenSpeakDuration community.ForbiddenSpeakDuration) (err error) {
	err = global.GVA_DB.Create(&hkForbiddenSpeakDuration).Error
	return err
}

// DeleteForbiddenSpeakDuration 删除ForbiddenSpeakDuration记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakDurationService *AppForbiddenSpeakDurationService) DeleteForbiddenSpeakDuration(hkForbiddenSpeakDuration community.ForbiddenSpeakDuration) (err error) {
	err = global.GVA_DB.Delete(&hkForbiddenSpeakDuration).Error
	return err
}

// DeleteForbiddenSpeakDurationByIds 批量删除ForbiddenSpeakDuration记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakDurationService *AppForbiddenSpeakDurationService) DeleteForbiddenSpeakDurationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForbiddenSpeakDuration{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateForbiddenSpeakDuration 更新ForbiddenSpeakDuration记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakDurationService *AppForbiddenSpeakDurationService) UpdateForbiddenSpeakDuration(hkForbiddenSpeakDuration community.ForbiddenSpeakDuration) (err error) {
	err = global.GVA_DB.Save(&hkForbiddenSpeakDuration).Error
	return err
}

// GetForbiddenSpeakDuration 根据id获取ForbiddenSpeakDuration记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakDurationService *AppForbiddenSpeakDurationService) GetForbiddenSpeakDuration(id uint) (hkForbiddenSpeakDuration community.ForbiddenSpeakDuration, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForbiddenSpeakDuration).Error
	return
}

// GetForbiddenSpeakDurationInfoList 分页获取ForbiddenSpeakDuration记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakDurationService *AppForbiddenSpeakDurationService) GetForbiddenSpeakDurationInfoList(info communityReq.ForbiddenSpeakDurationSearch) (list []community.ForbiddenSpeakDuration, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForbiddenSpeakDuration{})
	var hkForbiddenSpeakDurations []community.ForbiddenSpeakDuration
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
