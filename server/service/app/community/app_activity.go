package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HkActivityService struct {
}

// CreateHkActivity 创建HkActivity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *HkActivityService) CreateHkActivity(hkActivity community.HkActivity) (err error) {
	err = global.GVA_DB.Create(&hkActivity).Error
	return err
}

// DeleteHkActivity 删除HkActivity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *HkActivityService) DeleteHkActivity(hkActivity community.HkActivity) (err error) {
	err = global.GVA_DB.Delete(&hkActivity).Error
	return err
}

// DeleteHkActivityByIds 批量删除HkActivity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *HkActivityService) DeleteHkActivityByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkActivity{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkActivity 更新HkActivity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *HkActivityService) UpdateHkActivity(hkActivity community.HkActivity) (err error) {
	err = global.GVA_DB.Save(&hkActivity).Error
	return err
}

// GetHkActivity 根据id获取HkActivity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *HkActivityService) GetHkActivity(id uint64) (hkActivity community.HkActivity, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkActivity).Error
	return
}

// GetHkActivityInfoList 分页获取HkActivity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *HkActivityService) GetHkActivityInfoList(info communityReq.HkActivitySearch) (list []community.HkActivity, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkActivity{})
	var hkActivitys []community.HkActivity
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkActivitys).Error
	return hkActivitys, total, err
}
