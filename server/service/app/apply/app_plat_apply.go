package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppPlatApplyService struct {
}

// CreateHkPlatApply 创建HkPlatApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyService *AppPlatApplyService) CreateHkPlatApply(hkPlatApply apply.HkPlatApply) (err error) {
	err = global.GVA_DB.Create(&hkPlatApply).Error
	return err
}

// DeleteHkPlatApply 删除HkPlatApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyService *AppPlatApplyService) DeleteHkPlatApply(hkPlatApply apply.HkPlatApply) (err error) {
	err = global.GVA_DB.Delete(&hkPlatApply).Error
	return err
}

// DeleteHkPlatApplyByIds 批量删除HkPlatApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyService *AppPlatApplyService) DeleteHkPlatApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.HkPlatApply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkPlatApply 更新HkPlatApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyService *AppPlatApplyService) UpdateHkPlatApply(hkPlatApply apply.HkPlatApply) (err error) {
	err = global.GVA_DB.Save(&hkPlatApply).Error
	return err
}

// GetHkPlatApply 根据id获取HkPlatApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyService *AppPlatApplyService) GetHkPlatApply(id uint) (hkPlatApply apply.HkPlatApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkPlatApply).Error
	return
}

// GetHkPlatApplyInfoList 分页获取HkPlatApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyService *AppPlatApplyService) GetHkPlatApplyInfoList(info applyReq.HkPlatApplySearch) (list []apply.HkPlatApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&apply.HkPlatApply{})
	var hkPlatApplys []apply.HkPlatApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkPlatApplys).Error
	return hkPlatApplys, total, err
}
