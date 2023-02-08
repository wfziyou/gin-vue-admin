package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCircleService struct {
}

// CreateCircle 创建Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) CreateCircle(hkCircle community.Circle) (err error) {
	err = global.GVA_DB.Create(&hkCircle).Error
	return err
}

// DeleteCircle 删除Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) DeleteCircle(hkCircle community.Circle) (err error) {
	err = global.GVA_DB.Delete(&hkCircle).Error
	return err
}

// DeleteCircleByIds 批量删除Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) DeleteCircleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.Circle{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCircle 更新Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) UpdateCircle(hkCircle community.Circle) (err error) {
	err = global.GVA_DB.Save(&hkCircle).Error
	return err
}

// GetCircle 根据id获取Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) GetCircle(id uint64) (hkCircle community.Circle, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircle).Error
	return
}

// GetCircleInfoList 分页获取Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) GetCircleInfoList(info communityReq.CircleSearch) (list []community.CircleBaseInfo, total int64, err error) {
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

// GetSelfCircleList 分页获取Circle记录
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
