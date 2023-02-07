package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCircleService struct {
}

// CreateHkCircle 创建HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) CreateHkCircle(hkCircle community.HkCircle) (err error) {
	err = global.GVA_DB.Create(&hkCircle).Error
	return err
}

// DeleteHkCircle 删除HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) DeleteHkCircle(hkCircle community.HkCircle) (err error) {
	err = global.GVA_DB.Delete(&hkCircle).Error
	return err
}

// DeleteHkCircleByIds 批量删除HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) DeleteHkCircleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkCircle{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkCircle 更新HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) UpdateHkCircle(hkCircle community.HkCircle) (err error) {
	err = global.GVA_DB.Save(&hkCircle).Error
	return err
}

// GetHkCircle 根据id获取HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) GetHkCircle(id uint) (hkCircle community.HkCircle, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircle).Error
	return
}

// GetHkCircleInfoList 分页获取HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) GetHkCircleInfoList(info communityReq.CircleSearch) (list []community.CircleBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleBaseInfo{})
	var hkCircles []community.CircleBaseInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircles).Error
	return hkCircles, total, err
}

// GetSelfCircleList 分页获取HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) GetSelfCircleList(info communityReq.SelfCircleSearch) (list []community.CircleBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleBaseInfo{})
	var hkCircles []community.CircleBaseInfo

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircles).Error
	return hkCircles, total, err
}
