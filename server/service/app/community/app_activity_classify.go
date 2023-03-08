package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HkActivityClassifyService struct {
}

// CreateHkActivityClassify 创建HkActivityClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityClassifyService *HkActivityClassifyService) CreateHkActivityClassify(hkActivityClassify community.HkActivityClassify) (err error) {
	err = global.GVA_DB.Create(&hkActivityClassify).Error
	return err
}

// DeleteHkActivityClassify 删除HkActivityClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityClassifyService *HkActivityClassifyService) DeleteHkActivityClassify(hkActivityClassify community.HkActivityClassify) (err error) {
	err = global.GVA_DB.Delete(&hkActivityClassify).Error
	return err
}

// DeleteHkActivityClassifyByIds 批量删除HkActivityClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityClassifyService *HkActivityClassifyService) DeleteHkActivityClassifyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkActivityClassify{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkActivityClassify 更新HkActivityClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityClassifyService *HkActivityClassifyService) UpdateHkActivityClassify(hkActivityClassify community.HkActivityClassify) (err error) {
	err = global.GVA_DB.Save(&hkActivityClassify).Error
	return err
}

// GetHkActivityClassify 根据id获取HkActivityClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityClassifyService *HkActivityClassifyService) GetHkActivityClassify(id uint64) (hkActivityClassify community.HkActivityClassify, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkActivityClassify).Error
	return
}

// GetHkActivityClassifyInfoList 分页获取HkActivityClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityClassifyService *HkActivityClassifyService) GetHkActivityClassifyInfoList(info communityReq.HkActivityClassifySearch) (list []community.HkActivityClassify, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkActivityClassify{})
	var hkActivityClassifys []community.HkActivityClassify
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkActivityClassifys).Error
	return hkActivityClassifys, total, err
}
