package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppApplyService struct {
}

// CreateHkApply 创建HkApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) CreateHkApply(hkApply apply.HkApply) (err error) {
	err = global.GVA_DB.Create(&hkApply).Error
	return err
}

// DeleteHkApply 删除HkApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) DeleteHkApply(hkApply apply.HkApply) (err error) {
	err = global.GVA_DB.Delete(&hkApply).Error
	return err
}

// DeleteHkApplyByIds 批量删除HkApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) DeleteHkApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.HkApply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkApply 更新HkApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) UpdateHkApply(hkApply apply.HkApply) (err error) {
	err = global.GVA_DB.Save(&hkApply).Error
	return err
}

// GetHkApply 根据id获取HkApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) GetHkApply(id uint) (hkApply apply.HkApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkApply).Error
	return
}

// GetHkApplyInfoList 分页获取HkApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) GetHkApplyInfoList(info applyReq.HkApplySearch) (list []apply.HkApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&apply.HkApply{})
	var hkApplys []apply.HkApply
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

// GetHkApplyInfoListAll 分页获取HkApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) GetHkApplyInfoListAll(info applyReq.HkApplySearchAll) (list []apply.HkApply, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&apply.HkApply{})
	var hkApplys []apply.HkApply
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
