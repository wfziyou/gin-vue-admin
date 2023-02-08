package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCircleApplyService struct {
}

// CreateCircleApply 创建CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) CreateCircleApply(hkCircleApply apply.CircleApply) (err error) {
	err = global.GVA_DB.Create(&hkCircleApply).Error
	return err
}

// DeleteCircleApply 删除CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) DeleteCircleApply(hkCircleApply apply.CircleApply) (err error) {
	err = global.GVA_DB.Delete(&hkCircleApply).Error
	return err
}

// DeleteCircleApplyByIds 批量删除CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) DeleteCircleApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.CircleApply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCircleApply 更新CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) UpdateCircleApply(hkCircleApply apply.CircleApply) (err error) {
	err = global.GVA_DB.Save(&hkCircleApply).Error
	return err
}

// GetCircleApply 根据id获取CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) GetCircleApply(id uint64) (hkCircleApply apply.CircleApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleApply).Error
	return
}

// GetCircleApplyInfoList 分页获取CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) GetCircleApplyInfoList(info applyReq.CircleApplySearch) (list []apply.CircleApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&apply.CircleApply{})
	var hkCircleApplys []apply.CircleApply
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

// GetCircleApplyInfoListAll 分页获取CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) GetCircleApplyInfoListAll(info applyReq.CircleApplySearchAll) (list []apply.CircleApply, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&apply.CircleApply{})
	var hkCircleApplys []apply.CircleApply
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
