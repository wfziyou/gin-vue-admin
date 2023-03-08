package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HkUserSignService struct {
}

// CreateHkUserSign 创建HkUserSign记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserSignService *HkUserSignService) CreateHkUserSign(hkUserSign community.HkUserSign) (err error) {
	err = global.GVA_DB.Create(&hkUserSign).Error
	return err
}

// DeleteHkUserSign 删除HkUserSign记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserSignService *HkUserSignService) DeleteHkUserSign(hkUserSign community.HkUserSign) (err error) {
	err = global.GVA_DB.Delete(&hkUserSign).Error
	return err
}

// DeleteHkUserSignByIds 批量删除HkUserSign记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserSignService *HkUserSignService) DeleteHkUserSignByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkUserSign{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkUserSign 更新HkUserSign记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserSignService *HkUserSignService) UpdateHkUserSign(hkUserSign community.HkUserSign) (err error) {
	err = global.GVA_DB.Save(&hkUserSign).Error
	return err
}

// GetHkUserSign 根据id获取HkUserSign记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserSignService *HkUserSignService) GetHkUserSign(id uint64) (hkUserSign community.HkUserSign, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserSign).Error
	return
}

// GetHkUserSignInfoList 分页获取HkUserSign记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserSignService *HkUserSignService) GetHkUserSignInfoList(info communityReq.HkUserSignSearch) (list []community.HkUserSign, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkUserSign{})
	var hkUserSigns []community.HkUserSign
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUserSigns).Error
	return hkUserSigns, total, err
}
