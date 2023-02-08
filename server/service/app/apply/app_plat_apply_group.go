package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppPlatApplyGroupService struct {
}

// CreatePlatApplyGroup 创建PlatApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyGroupService *AppPlatApplyGroupService) CreatePlatApplyGroup(hkPlatApplyGroup apply.PlatApplyGroup) (err error) {
	err = global.GVA_DB.Create(&hkPlatApplyGroup).Error
	return err
}

// DeletePlatApplyGroup 删除PlatApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyGroupService *AppPlatApplyGroupService) DeletePlatApplyGroup(hkPlatApplyGroup apply.PlatApplyGroup) (err error) {
	err = global.GVA_DB.Delete(&hkPlatApplyGroup).Error
	return err
}

// DeletePlatApplyGroupByIds 批量删除PlatApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyGroupService *AppPlatApplyGroupService) DeletePlatApplyGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.PlatApplyGroup{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePlatApplyGroup 更新PlatApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyGroupService *AppPlatApplyGroupService) UpdatePlatApplyGroup(hkPlatApplyGroup apply.PlatApplyGroup) (err error) {
	err = global.GVA_DB.Save(&hkPlatApplyGroup).Error
	return err
}

// GetPlatApplyGroup 根据id获取PlatApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyGroupService *AppPlatApplyGroupService) GetPlatApplyGroup(id uint) (hkPlatApplyGroup apply.PlatApplyGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkPlatApplyGroup).Error
	return
}

// GetPlatApplyGroupInfoList 分页获取PlatApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPlatApplyGroupService *AppPlatApplyGroupService) GetPlatApplyGroupInfoList(info applyReq.PlatApplyGroupSearch) (list []apply.PlatApplyGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&apply.PlatApplyGroup{})
	var hkPlatApplyGroups []apply.PlatApplyGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkPlatApplyGroups).Error
	return hkPlatApplyGroups, total, err
}
