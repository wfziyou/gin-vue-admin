package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCircleClassifyService struct {
}

// CreateCircleClassify 创建CircleClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) CreateCircleClassify(hkCircleClassify community.CircleClassify) (err error) {
	err = global.GVA_DB.Create(&hkCircleClassify).Error
	return err
}

// DeleteCircleClassify 删除CircleClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) DeleteCircleClassify(hkCircleClassify community.CircleClassify) (err error) {
	err = global.GVA_DB.Delete(&hkCircleClassify).Error
	return err
}

// DeleteCircleClassifyByIds 批量删除CircleClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) DeleteCircleClassifyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.CircleClassify{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCircleClassify 更新CircleClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) UpdateCircleClassify(hkCircleClassify community.CircleClassify) (err error) {
	err = global.GVA_DB.Save(&hkCircleClassify).Error
	return err
}

// GetCircleClassify 根据id获取圈子分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) GetCircleClassify(id uint64) (hkCircleClassify community.CircleClassify, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleClassify).Error
	return
}

// GetCircleClassifyInfoList 分页获取圈子分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) GetCircleClassifyInfoList(info communityReq.CircleClassifySearch) (list []community.CircleClassify, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleClassify{})
	var hkCircleClassifys []community.CircleClassify
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

// GetCircleClassifyInfoList 分页获取圈子分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleClassifyService *AppCircleClassifyService) GetCircleClassifyInfoListAll(info communityReq.CircleClassifySearch) (list []community.CircleClassify, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.CircleClassify{})
	var hkCircleClassifys []community.CircleClassify
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Find(&hkCircleClassifys).Error
	return hkCircleClassifys, total, err
}
