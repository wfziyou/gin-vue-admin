package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type HkUserCircleApplyService struct {
}

// CreateHkUserCircleApply 创建HkUserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserCircleApplyService *HkUserCircleApplyService) CreateHkUserCircleApply(hkUserCircleApply community.HkUserCircleApply) (err error) {
	err = global.GVA_DB.Create(&hkUserCircleApply).Error
	return err
}

// DeleteHkUserCircleApply 删除HkUserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserCircleApplyService *HkUserCircleApplyService) DeleteHkUserCircleApply(hkUserCircleApply community.HkUserCircleApply) (err error) {
	err = global.GVA_DB.Delete(&hkUserCircleApply).Error
	return err
}

// DeleteHkUserCircleApplyByIds 批量删除HkUserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserCircleApplyService *HkUserCircleApplyService) DeleteHkUserCircleApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkUserCircleApply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkUserCircleApply 更新HkUserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserCircleApplyService *HkUserCircleApplyService) UpdateHkUserCircleApply(hkUserCircleApply community.HkUserCircleApply) (err error) {
	err = global.GVA_DB.Save(&hkUserCircleApply).Error
	return err
}

// GetHkUserCircleApply 根据id获取HkUserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserCircleApplyService *HkUserCircleApplyService) GetHkUserCircleApply(id uint) (hkUserCircleApply community.HkUserCircleApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserCircleApply).Error
	return
}

// GetHkUserCircleApplyInfoList 分页获取HkUserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserCircleApplyService *HkUserCircleApplyService) GetHkUserCircleApplyInfoList(info communityReq.HkUserCircleApplySearch) (list []community.HkUserCircleApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkUserCircleApply{})
	var hkUserCircleApplys []community.HkUserCircleApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUserCircleApplys).Error
	return hkUserCircleApplys, total, err
}
