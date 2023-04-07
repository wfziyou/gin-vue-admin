package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ActivityService struct {
}

// CreateActivity 创建Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *ActivityService) CreateActivity(hkActivity community.Activity) (err error) {
	err = global.GVA_DB.Create(&hkActivity).Error
	return err
}

// DeleteActivity 删除Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *ActivityService) DeleteActivity(hkActivity community.Activity) (err error) {
	err = global.GVA_DB.Delete(&hkActivity).Error
	return err
}

// DeleteActivityByIds 批量删除Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *ActivityService) DeleteActivityByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.Activity{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateActivity 更新Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *ActivityService) UpdateActivity(hkActivity community.Activity) (err error) {
	err = global.GVA_DB.Save(&hkActivity).Error
	return err
}

// GetActivity 根据id获取Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *ActivityService) GetActivity(id uint64) (hkActivity community.Activity, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkActivity).Error
	return
}

// GetActivityInfoList 分页获取Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *ActivityService) GetActivityInfoList(info communityReq.ActivitySearch) (list []community.Activity, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.Activity{})
	var hkActivitys []community.Activity
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
