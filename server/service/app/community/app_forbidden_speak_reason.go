package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type AppForbiddenSpeakReasonService struct {
}

// CreateHkForbiddenSpeakReason 创建HkForbiddenSpeakReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakReasonService *AppForbiddenSpeakReasonService) CreateHkForbiddenSpeakReason(hkForbiddenSpeakReason community.HkForbiddenSpeakReason) (err error) {
	err = global.GVA_DB.Create(&hkForbiddenSpeakReason).Error
	return err
}

// DeleteHkForbiddenSpeakReason 删除HkForbiddenSpeakReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakReasonService *AppForbiddenSpeakReasonService) DeleteHkForbiddenSpeakReason(hkForbiddenSpeakReason community.HkForbiddenSpeakReason) (err error) {
	err = global.GVA_DB.Delete(&hkForbiddenSpeakReason).Error
	return err
}

// DeleteHkForbiddenSpeakReasonByIds 批量删除HkForbiddenSpeakReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakReasonService *AppForbiddenSpeakReasonService) DeleteHkForbiddenSpeakReasonByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForbiddenSpeakReason{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkForbiddenSpeakReason 更新HkForbiddenSpeakReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakReasonService *AppForbiddenSpeakReasonService) UpdateHkForbiddenSpeakReason(hkForbiddenSpeakReason community.HkForbiddenSpeakReason) (err error) {
	err = global.GVA_DB.Save(&hkForbiddenSpeakReason).Error
	return err
}

// GetHkForbiddenSpeakReason 根据id获取HkForbiddenSpeakReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakReasonService *AppForbiddenSpeakReasonService) GetHkForbiddenSpeakReason(id uint) (hkForbiddenSpeakReason community.HkForbiddenSpeakReason, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForbiddenSpeakReason).Error
	return
}

// GetHkForbiddenSpeakReasonInfoList 分页获取HkForbiddenSpeakReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakReasonService *AppForbiddenSpeakReasonService) GetHkForbiddenSpeakReasonInfoList(info communityReq.HkForbiddenSpeakReasonSearch) (list []community.HkForbiddenSpeakReason, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkForbiddenSpeakReason{})
	var hkForbiddenSpeakReasons []community.HkForbiddenSpeakReason
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForbiddenSpeakReasons).Error
	return hkForbiddenSpeakReasons, total, err
}
