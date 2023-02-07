package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/general/request"
)

type HkBugReportService struct {
}

// CreateHkBugReport 创建HkBugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkBugReportService *HkBugReportService) CreateHkBugReport(hkBugReport general.HkBugReport) (err error) {
	err = global.GVA_DB.Create(&hkBugReport).Error
	return err
}

// DeleteHkBugReport 删除HkBugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkBugReportService *HkBugReportService) DeleteHkBugReport(hkBugReport general.HkBugReport) (err error) {
	err = global.GVA_DB.Delete(&hkBugReport).Error
	return err
}

// DeleteHkBugReportByIds 批量删除HkBugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkBugReportService *HkBugReportService) DeleteHkBugReportByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.HkBugReport{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkBugReport 更新HkBugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkBugReportService *HkBugReportService) UpdateHkBugReport(hkBugReport general.HkBugReport) (err error) {
	err = global.GVA_DB.Save(&hkBugReport).Error
	return err
}

// GetHkBugReport 根据id获取HkBugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkBugReportService *HkBugReportService) GetHkBugReport(id uint) (hkBugReport general.HkBugReport, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkBugReport).Error
	return
}

// GetHkBugReportInfoList 分页获取HkBugReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkBugReportService *HkBugReportService) GetHkBugReportInfoList(info generalReq.HkBugReportSearch) (list []general.HkBugReport, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&general.HkBugReport{})
	var hkBugReports []general.HkBugReport
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
