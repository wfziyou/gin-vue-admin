package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppPlatApplyService struct {
}

// CreatePlatApply 创建PlatApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyService *AppPlatApplyService) CreatePlatApply(hkPlatApply apply.PlatApply) (err error) {
	err = global.GVA_DB.Create(&hkPlatApply).Error
	return err
}

// DeletePlatApply 删除PlatApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyService *AppPlatApplyService) DeletePlatApply(hkPlatApply apply.PlatApply) (err error) {
	err = global.GVA_DB.Delete(&hkPlatApply).Error
	return err
}

// DeletePlatApplyByIds 批量删除PlatApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyService *AppPlatApplyService) DeletePlatApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.PlatApply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePlatApply 更新PlatApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyService *AppPlatApplyService) UpdatePlatApply(hkPlatApply apply.PlatApply) (err error) {
	err = global.GVA_DB.Save(&hkPlatApply).Error
	return err
}

// GetPlatApply 根据id获取PlatApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyService *AppPlatApplyService) GetPlatApply(id uint) (hkPlatApply apply.PlatApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkPlatApply).Error
	return
}

// GetPlatApplyInfoList 分页获取PlatApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyService *AppPlatApplyService) GetPlatApplyInfoList(info applyReq.PlatApplySearch) (list []apply.PlatApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&apply.PlatApply{})
	var hkPlatApplys []apply.PlatApply
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
