package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppReportReasonService struct {
}

// CreateReportReason 创建ReportReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportReasonService *AppReportReasonService) CreateReportReason(hkReportReason community.ReportReason) (err error) {
	err = global.GVA_DB.Create(&hkReportReason).Error
	return err
}

// DeleteReportReason 删除ReportReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportReasonService *AppReportReasonService) DeleteReportReason(hkReportReason community.ReportReason) (err error) {
	err = global.GVA_DB.Delete(&hkReportReason).Error
	return err
}

// DeleteReportReasonByIds 批量删除ReportReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportReasonService *AppReportReasonService) DeleteReportReasonByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ReportReason{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateReportReason 更新ReportReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportReasonService *AppReportReasonService) UpdateReportReason(hkReportReason community.ReportReason) (err error) {
	err = global.GVA_DB.Save(&hkReportReason).Error
	return err
}

// GetReportReason 根据id获取ReportReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportReasonService *AppReportReasonService) GetReportReason(id uint64) (hkReportReason community.ReportReason, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkReportReason).Error
	return
}

// GetReportReasonInfoList 分页获取ReportReason记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportReasonService *AppReportReasonService) GetReportReasonInfoList(info communityReq.ReportReasonSearch) (list []community.ReportReason, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ReportReason{})
	var hkReportReasons []community.ReportReason
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
