package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppReportService struct {
}

// CreateReport 创建Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportService *AppReportService) CreateReport(userId uint64, info communityReq.CreateReportReq) (err error) {
	obj := community.Report{
		UserId:            userId,
		ReportType:        info.ReportType,
		ReportId:          info.ReportId,
		Content:           info.Content,
		ContentAttachment: info.ContentAttachment,
	}
	err = global.GVA_DB.Create(&obj).Error
	return err
}

// DeleteReport 删除Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportService *AppReportService) DeleteReport(hkReport community.Report) (err error) {
	err = global.GVA_DB.Delete(&hkReport).Error
	return err
}

// DeleteReportByIds 批量删除Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportService *AppReportService) DeleteReportByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.Report{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateReport 更新Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportService *AppReportService) UpdateReport(hkReport community.Report) (err error) {
	err = global.GVA_DB.Save(&hkReport).Error
	return err
}

// GetReport 根据id获取Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportService *AppReportService) GetReport(id uint64) (hkReport community.Report, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkReport).Error
	return
}

// GetReportInfoList 分页获取Report记录
// Author [piexlmax](https://github.com/piexlmax)
func (appReportService *AppReportService) GetReportInfoList(info communityReq.ReportSearch) (list []community.Report, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.Report{})
	var hkReports []community.Report
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.ReasonId != nil {
		db = db.Where("reason_id = ?", info.ReasonId)
	}
	if info.ReportType != nil {
		db = db.Where("report_type = ?", info.ReportType)
	}
	if info.CurStatus != nil {
		db = db.Where("cur_status = ?", info.CurStatus)
	}
	if info.HandleUserId != nil {
		db = db.Where("handle_user_id = ?", info.HandleUserId)
	}
	if info.HandleType != nil {
		db = db.Where("handle_type = ?", info.HandleType)
	}
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
