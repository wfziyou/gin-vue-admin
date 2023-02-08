package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppUserCircleApplyService struct {
}

// CreateUserCircleApply 创建UserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCircleApplyService *AppUserCircleApplyService) CreateUserCircleApply(hkUserCircleApply community.UserCircleApply) (err error) {
	err = global.GVA_DB.Create(&hkUserCircleApply).Error
	return err
}

// DeleteUserCircleApply 删除UserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCircleApplyService *AppUserCircleApplyService) DeleteUserCircleApply(hkUserCircleApply community.UserCircleApply) (err error) {
	err = global.GVA_DB.Delete(&hkUserCircleApply).Error
	return err
}

// DeleteUserCircleApplyByIds 批量删除UserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCircleApplyService *AppUserCircleApplyService) DeleteUserCircleApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.UserCircleApply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserCircleApply 更新UserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCircleApplyService *AppUserCircleApplyService) UpdateUserCircleApply(hkUserCircleApply community.UserCircleApply) (err error) {
	err = global.GVA_DB.Save(&hkUserCircleApply).Error
	return err
}

// GetUserCircleApply 根据id获取UserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCircleApplyService *AppUserCircleApplyService) GetUserCircleApply(id uint) (hkUserCircleApply community.UserCircleApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserCircleApply).Error
	return
}

// GetUserCircleApplyInfoList 分页获取UserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCircleApplyService *AppUserCircleApplyService) GetUserCircleApplyInfoList(info communityReq.UserCircleApplySearch) (list []community.UserCircleApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.UserCircleApply{})
	var hkUserCircleApplys []community.UserCircleApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUserCircleApplys).Error
	return hkUserCircleApplys, total, err
}

// GetUserCircleApplyInfoListALL 获取UserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCircleApplyService *AppUserCircleApplyService) GetUserCircleApplyInfoListALL(info communityReq.UserCircleApplySearch) (list []community.UserCircleApply, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.UserCircleApply{})
	var hkUserCircleApplys []community.UserCircleApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Find(&hkUserCircleApplys).Error
	return hkUserCircleApplys, total, err
}
