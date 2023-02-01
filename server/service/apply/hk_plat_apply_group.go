package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apply"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/apply/request"
)

type HkPlatApplyGroupService struct {
}

// CreateHkPlatApplyGroup 创建HkPlatApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkPlatApplyGroupService *HkPlatApplyGroupService) CreateHkPlatApplyGroup(hkPlatApplyGroup apply.HkPlatApplyGroup) (err error) {
	err = global.GVA_DB.Create(&hkPlatApplyGroup).Error
	return err
}

// DeleteHkPlatApplyGroup 删除HkPlatApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkPlatApplyGroupService *HkPlatApplyGroupService)DeleteHkPlatApplyGroup(hkPlatApplyGroup apply.HkPlatApplyGroup) (err error) {
	err = global.GVA_DB.Delete(&hkPlatApplyGroup).Error
	return err
}

// DeleteHkPlatApplyGroupByIds 批量删除HkPlatApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkPlatApplyGroupService *HkPlatApplyGroupService)DeleteHkPlatApplyGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.HkPlatApplyGroup{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkPlatApplyGroup 更新HkPlatApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkPlatApplyGroupService *HkPlatApplyGroupService)UpdateHkPlatApplyGroup(hkPlatApplyGroup apply.HkPlatApplyGroup) (err error) {
	err = global.GVA_DB.Save(&hkPlatApplyGroup).Error
	return err
}

// GetHkPlatApplyGroup 根据id获取HkPlatApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkPlatApplyGroupService *HkPlatApplyGroupService)GetHkPlatApplyGroup(id uint) (hkPlatApplyGroup apply.HkPlatApplyGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkPlatApplyGroup).Error
	return
}

// GetHkPlatApplyGroupInfoList 分页获取HkPlatApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkPlatApplyGroupService *HkPlatApplyGroupService)GetHkPlatApplyGroupInfoList(info applyReq.HkPlatApplyGroupSearch) (list []apply.HkPlatApplyGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&apply.HkPlatApplyGroup{})
    var hkPlatApplyGroups []apply.HkPlatApplyGroup
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkPlatApplyGroups).Error
	return  hkPlatApplyGroups, total, err
}
