package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCircleApplyService struct {
}

// CreateHkCircleApply 创建HkCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) CreateHkCircleApply(hkCircleApply apply.HkCircleApply) (err error) {
	err = global.GVA_DB.Create(&hkCircleApply).Error
	return err
}

// DeleteHkCircleApply 删除HkCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) DeleteHkCircleApply(hkCircleApply apply.HkCircleApply) (err error) {
	err = global.GVA_DB.Delete(&hkCircleApply).Error
	return err
}

// DeleteHkCircleApplyByIds 批量删除HkCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) DeleteHkCircleApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.HkCircleApply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkCircleApply 更新HkCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) UpdateHkCircleApply(hkCircleApply apply.HkCircleApply) (err error) {
	err = global.GVA_DB.Save(&hkCircleApply).Error
	return err
}

// GetHkCircleApply 根据id获取HkCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) GetHkCircleApply(id uint) (hkCircleApply apply.HkCircleApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleApply).Error
	return
}

// GetHkCircleApplyInfoList 分页获取HkCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) GetHkCircleApplyInfoList(info applyReq.HkCircleApplySearch) (list []apply.HkCircleApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&apply.HkCircleApply{})
	var hkCircleApplys []apply.HkCircleApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircleApplys).Error
	return hkCircleApplys, total, err
}

// GetHkCircleApplyInfoListAll 分页获取HkCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) GetHkCircleApplyInfoListAll(info applyReq.HkCircleApplySearchAll) (list []apply.HkCircleApply, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&apply.HkCircleApply{})
	var hkCircleApplys []apply.HkCircleApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Find(&hkCircleApplys).Error
	return hkCircleApplys, total, err
}
