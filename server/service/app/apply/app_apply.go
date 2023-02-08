package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppApplyService struct {
}

// CreateApply 创建Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) CreateApply(hkApply apply.Apply) (err error) {
	err = global.GVA_DB.Create(&hkApply).Error
	return err
}

// DeleteApply 删除Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) DeleteApply(hkApply apply.Apply) (err error) {
	err = global.GVA_DB.Delete(&hkApply).Error
	return err
}

// DeleteApplyByIds 批量删除Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) DeleteApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.Apply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateApply 更新Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) UpdateApply(hkApply apply.Apply) (err error) {
	err = global.GVA_DB.Save(&hkApply).Error
	return err
}

// GetApply 根据id获取Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) GetApply(id uint64) (hkApply apply.Apply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkApply).Error
	return
}

// GetApplyInfoList 分页获取Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) GetApplyInfoList(info applyReq.ApplySearch) (list []apply.Apply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&apply.Apply{})
	var hkApplys []apply.Apply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkApplys).Error
	return hkApplys, total, err
}

// GetApplyInfoListAll 分页获取Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) GetApplyInfoListAll(info applyReq.ApplySearchAll) (list []apply.Apply, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&apply.Apply{})
	var hkApplys []apply.Apply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Find(&hkApplys).Error
	return hkApplys, total, err
}
