package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCircleApplyGroupService struct {
}

// CreateCircleApplyGroup 创建CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) CreateCircleApplyGroup(hkCircleApplyGroup apply.CircleApplyGroup) (err error) {
	err = global.GVA_DB.Create(&hkCircleApplyGroup).Error
	return err
}

// DeleteCircleApplyGroup 删除CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) DeleteCircleApplyGroup(hkCircleApplyGroup apply.CircleApplyGroup) (err error) {
	err = global.GVA_DB.Delete(&hkCircleApplyGroup).Error
	return err
}

// DeleteCircleApplyGroupByIds 批量删除CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) DeleteCircleApplyGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.CircleApplyGroup{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCircleApplyGroup 更新CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) UpdateCircleApplyGroup(hkCircleApplyGroup apply.CircleApplyGroup) (err error) {
	err = global.GVA_DB.Save(&hkCircleApplyGroup).Error
	return err
}

// GetCircleApplyGroup 根据id获取CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) GetCircleApplyGroup(id uint) (hkCircleApplyGroup apply.CircleApplyGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleApplyGroup).Error
	return
}

// GetCircleApplyGroupInfoList 分页获取CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) GetCircleApplyGroupInfoList(info applyReq.CircleApplyGroupSearch) (list []apply.CircleApplyGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&apply.CircleApplyGroup{})
	var hkCircleApplyGroups []apply.CircleApplyGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CircleId != 0 {
		db = db.Where("circle_id = ?", info.CircleId)
	}
	if len(info.Name) > 0 {
		db = db.Where("name LIKE '?%'", info.Name)
	}
	if info.ParentId != nil {
		db = db.Where("parent_id = ?", info.ParentId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircleApplyGroups).Error
	return hkCircleApplyGroups, total, err
}

// GetCircleApplyGroupInfoListAll 分页获取CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) GetCircleApplyGroupInfoListAll(info applyReq.CircleApplyGroupSearchAll) (list []apply.CircleApplyGroup, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&apply.CircleApplyGroup{})
	var hkCircleApplyGroups []apply.CircleApplyGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CircleId != 0 {
		db = db.Where("circle_id = ?", info.CircleId)
	}
	if len(info.Name) > 0 {
		db = db.Where("name LIKE '?%'", info.Name)
	}
	if info.ParentId != nil {
		db = db.Where("parent_id = ?", info.ParentId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Find(&hkCircleApplyGroups).Error
	return hkCircleApplyGroups, total, err
}
