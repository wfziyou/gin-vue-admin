package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type AppCircleClassifyService struct {
}

// CreateHkCircleClassify 创建HkCircleClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) CreateHkCircleClassify(hkCircleClassify community.HkCircleClassify) (err error) {
	err = global.GVA_DB.Create(&hkCircleClassify).Error
	return err
}

// DeleteHkCircleClassify 删除HkCircleClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) DeleteHkCircleClassify(hkCircleClassify community.HkCircleClassify) (err error) {
	err = global.GVA_DB.Delete(&hkCircleClassify).Error
	return err
}

// DeleteHkCircleClassifyByIds 批量删除HkCircleClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) DeleteHkCircleClassifyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkCircleClassify{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkCircleClassify 更新HkCircleClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) UpdateHkCircleClassify(hkCircleClassify community.HkCircleClassify) (err error) {
	err = global.GVA_DB.Save(&hkCircleClassify).Error
	return err
}

// GetHkCircleClassify 根据id获取HkCircleClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) GetHkCircleClassify(id uint) (hkCircleClassify community.HkCircleClassify, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleClassify).Error
	return
}

// GetHkCircleClassifyInfoList 分页获取HkCircleClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) GetHkCircleClassifyInfoList(info communityReq.HkCircleClassifySearch) (list []community.HkCircleClassify, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkCircleClassify{})
	var hkCircleClassifys []community.HkCircleClassify
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircleClassifys).Error
	return hkCircleClassifys, total, err
}
