package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForbiddenSpeakReasonService struct {
}

// CreateForbiddenSpeakReason 创建ForbiddenSpeakReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakReasonService *AppForbiddenSpeakReasonService) CreateForbiddenSpeakReason(hkForbiddenSpeakReason community.ForbiddenSpeakReason) (err error) {
	err = global.GVA_DB.Create(&hkForbiddenSpeakReason).Error
	return err
}

// DeleteForbiddenSpeakReason 删除ForbiddenSpeakReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakReasonService *AppForbiddenSpeakReasonService) DeleteForbiddenSpeakReason(hkForbiddenSpeakReason community.ForbiddenSpeakReason) (err error) {
	err = global.GVA_DB.Delete(&hkForbiddenSpeakReason).Error
	return err
}

// DeleteForbiddenSpeakReasonByIds 批量删除ForbiddenSpeakReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakReasonService *AppForbiddenSpeakReasonService) DeleteForbiddenSpeakReasonByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForbiddenSpeakReason{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateForbiddenSpeakReason 更新ForbiddenSpeakReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakReasonService *AppForbiddenSpeakReasonService) UpdateForbiddenSpeakReason(hkForbiddenSpeakReason community.ForbiddenSpeakReason) (err error) {
	err = global.GVA_DB.Save(&hkForbiddenSpeakReason).Error
	return err
}

// GetForbiddenSpeakReason 根据id获取ForbiddenSpeakReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakReasonService *AppForbiddenSpeakReasonService) GetForbiddenSpeakReason(id uint) (hkForbiddenSpeakReason community.ForbiddenSpeakReason, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForbiddenSpeakReason).Error
	return
}

// GetForbiddenSpeakReasonInfoList 分页获取ForbiddenSpeakReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakReasonService *AppForbiddenSpeakReasonService) GetForbiddenSpeakReasonInfoList(info communityReq.ForbiddenSpeakReasonSearch) (list []community.ForbiddenSpeakReason, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForbiddenSpeakReason{})
	var hkForbiddenSpeakReasons []community.ForbiddenSpeakReason
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
