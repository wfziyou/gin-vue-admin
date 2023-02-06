package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type AppReportReasonService struct {
}

// CreateHkReportReason 创建HkReportReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportReasonService *AppReportReasonService) CreateHkReportReason(hkReportReason community.HkReportReason) (err error) {
	err = global.GVA_DB.Create(&hkReportReason).Error
	return err
}

// DeleteHkReportReason 删除HkReportReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportReasonService *AppReportReasonService) DeleteHkReportReason(hkReportReason community.HkReportReason) (err error) {
	err = global.GVA_DB.Delete(&hkReportReason).Error
	return err
}

// DeleteHkReportReasonByIds 批量删除HkReportReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportReasonService *AppReportReasonService) DeleteHkReportReasonByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkReportReason{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkReportReason 更新HkReportReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportReasonService *AppReportReasonService) UpdateHkReportReason(hkReportReason community.HkReportReason) (err error) {
	err = global.GVA_DB.Save(&hkReportReason).Error
	return err
}

// GetHkReportReason 根据id获取HkReportReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportReasonService *AppReportReasonService) GetHkReportReason(id uint) (hkReportReason community.HkReportReason, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkReportReason).Error
	return
}

// GetHkReportReasonInfoList 分页获取HkReportReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportReasonService *AppReportReasonService) GetHkReportReasonInfoList(info communityReq.HkReportReasonSearch) (list []community.HkReportReason, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkReportReason{})
	var hkReportReasons []community.HkReportReason
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkReportReasons).Error
	return hkReportReasons, total, err
}
