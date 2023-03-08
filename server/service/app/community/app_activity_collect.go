package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HkActivityCollectService struct {
}

// CreateHkActivityCollect 创建HkActivityCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityCollectService *HkActivityCollectService) CreateHkActivityCollect(hkActivityCollect community.HkActivityCollect) (err error) {
	err = global.GVA_DB.Create(&hkActivityCollect).Error
	return err
}

// DeleteHkActivityCollect 删除HkActivityCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityCollectService *HkActivityCollectService) DeleteHkActivityCollect(hkActivityCollect community.HkActivityCollect) (err error) {
	err = global.GVA_DB.Delete(&hkActivityCollect).Error
	return err
}

// DeleteHkActivityCollectByIds 批量删除HkActivityCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityCollectService *HkActivityCollectService) DeleteHkActivityCollectByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkActivityCollect{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkActivityCollect 更新HkActivityCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityCollectService *HkActivityCollectService) UpdateHkActivityCollect(hkActivityCollect community.HkActivityCollect) (err error) {
	err = global.GVA_DB.Save(&hkActivityCollect).Error
	return err
}

// GetHkActivityCollect 根据id获取HkActivityCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityCollectService *HkActivityCollectService) GetHkActivityCollect(id uint64) (hkActivityCollect community.HkActivityCollect, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkActivityCollect).Error
	return
}

// GetHkActivityCollectInfoList 分页获取HkActivityCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityCollectService *HkActivityCollectService) GetHkActivityCollectInfoList(info communityReq.HkActivityCollectSearch) (list []community.HkActivityCollect, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkActivityCollect{})
	var hkActivityCollects []community.HkActivityCollect
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkActivityCollects).Error
	return hkActivityCollects, total, err
}
