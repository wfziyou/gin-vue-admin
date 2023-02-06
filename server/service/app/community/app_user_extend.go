package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type AppUserExtendService struct {
}

// CreateHkUserExtend 创建HkUserExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserExtendService *AppUserExtendService) CreateHkUserExtend(hkUserExtend community.HkUserExtend) (err error) {
	err = global.GVA_DB.Create(&hkUserExtend).Error
	return err
}

// DeleteHkUserExtend 删除HkUserExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserExtendService *AppUserExtendService) DeleteHkUserExtend(hkUserExtend community.HkUserExtend) (err error) {
	err = global.GVA_DB.Delete(&hkUserExtend).Error
	return err
}

// DeleteHkUserExtendByIds 批量删除HkUserExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserExtendService *AppUserExtendService) DeleteHkUserExtendByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkUserExtend{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkUserExtend 更新HkUserExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserExtendService *AppUserExtendService) UpdateHkUserExtend(hkUserExtend community.HkUserExtend) (err error) {
	err = global.GVA_DB.Save(&hkUserExtend).Error
	return err
}

// GetHkUserExtend 根据id获取HkUserExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserExtendService *AppUserExtendService) GetHkUserExtend(id uint) (hkUserExtend community.HkUserExtend, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserExtend).Error
	return
}

// GetHkUserExtendInfoList 分页获取HkUserExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserExtendService *AppUserExtendService) GetHkUserExtendInfoList(info communityReq.HkUserExtendSearch) (list []community.HkUserExtend, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkUserExtend{})
	var hkUserExtends []community.HkUserExtend
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUserExtends).Error
	return hkUserExtends, total, err
}
