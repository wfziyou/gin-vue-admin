package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppBugReportService struct {
}

// CreateBugReport 创建BugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBugReportService *AppBugReportService) CreateBugReport(hkBugReport general.BugReport) (err error) {
	err = global.GVA_DB.Create(&hkBugReport).Error
	return err
}

// DeleteBugReport 删除BugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBugReportService *AppBugReportService) DeleteBugReport(hkBugReport general.BugReport) (err error) {
	err = global.GVA_DB.Delete(&hkBugReport).Error
	return err
}

// DeleteBugReportByIds 批量删除BugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBugReportService *AppBugReportService) DeleteBugReportByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.BugReport{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBugReport 更新BugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBugReportService *AppBugReportService) UpdateBugReport(hkBugReport general.BugReport) (err error) {
	err = global.GVA_DB.Save(&hkBugReport).Error
	return err
}

// GetBugReport 根据id获取BugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBugReportService *AppBugReportService) GetBugReport(id uint64) (hkBugReport general.BugReport, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkBugReport).Error
	return
}

// GetBugReportInfoList 分页获取BugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBugReportService *AppBugReportService) GetBugReportInfoList(info generalReq.BugReportSearch) (list []general.BugReport, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&general.BugReport{})
	var hkBugReports []general.BugReport
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkBugReports).Error
	return hkBugReports, total, err
}

// AppGetBugReportInfoList 分页获取BugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appBugReportService *AppBugReportService) AppGetBugReportInfoList(info generalReq.BugReportSearch) (list []general.BugReport, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&general.BugReport{})
	var hkBugReports []general.BugReport
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkBugReports).Error
	return hkBugReports, total, err
}
