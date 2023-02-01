package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/apply"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/apply/request"
)

type HkCircleApplyGroupService struct {
}

// CreateHkCircleApplyGroup 创建HkCircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleApplyGroupService *HkCircleApplyGroupService) CreateHkCircleApplyGroup(hkCircleApplyGroup apply.HkCircleApplyGroup) (err error) {
	err = global.GVA_DB.Create(&hkCircleApplyGroup).Error
	return err
}

// DeleteHkCircleApplyGroup 删除HkCircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleApplyGroupService *HkCircleApplyGroupService)DeleteHkCircleApplyGroup(hkCircleApplyGroup apply.HkCircleApplyGroup) (err error) {
	err = global.GVA_DB.Delete(&hkCircleApplyGroup).Error
	return err
}

// DeleteHkCircleApplyGroupByIds 批量删除HkCircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleApplyGroupService *HkCircleApplyGroupService)DeleteHkCircleApplyGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.HkCircleApplyGroup{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkCircleApplyGroup 更新HkCircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleApplyGroupService *HkCircleApplyGroupService)UpdateHkCircleApplyGroup(hkCircleApplyGroup apply.HkCircleApplyGroup) (err error) {
	err = global.GVA_DB.Save(&hkCircleApplyGroup).Error
	return err
}

// GetHkCircleApplyGroup 根据id获取HkCircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleApplyGroupService *HkCircleApplyGroupService)GetHkCircleApplyGroup(id uint) (hkCircleApplyGroup apply.HkCircleApplyGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleApplyGroup).Error
	return
}

// GetHkCircleApplyGroupInfoList 分页获取HkCircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleApplyGroupService *HkCircleApplyGroupService)GetHkCircleApplyGroupInfoList(info applyReq.HkCircleApplyGroupSearch) (list []apply.HkCircleApplyGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&apply.HkCircleApplyGroup{})
    var hkCircleApplyGroups []apply.HkCircleApplyGroup
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkCircleApplyGroups).Error
	return  hkCircleApplyGroups, total, err
}
