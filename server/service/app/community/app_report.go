package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type AppReportService struct {
}

// CreateHkReport 创建HkReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportService *AppReportService) CreateHkReport(hkReport community.HkReport) (err error) {
	err = global.GVA_DB.Create(&hkReport).Error
	return err
}

// DeleteHkReport 删除HkReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportService *AppReportService) DeleteHkReport(hkReport community.HkReport) (err error) {
	err = global.GVA_DB.Delete(&hkReport).Error
	return err
}

// DeleteHkReportByIds 批量删除HkReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportService *AppReportService) DeleteHkReportByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkReport{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkReport 更新HkReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportService *AppReportService) UpdateHkReport(hkReport community.HkReport) (err error) {
	err = global.GVA_DB.Save(&hkReport).Error
	return err
}

// GetHkReport 根据id获取HkReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportService *AppReportService) GetHkReport(id uint) (hkReport community.HkReport, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkReport).Error
	return
}

// GetHkReportInfoList 分页获取HkReport记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportService *AppReportService) GetHkReportInfoList(info communityReq.HkReportSearch) (list []community.HkReport, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkReport{})
	var hkReports []community.HkReport
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkReports).Error
	return hkReports, total, err
}
