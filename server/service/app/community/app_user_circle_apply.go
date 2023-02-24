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

// GetUserCircleApplyInfoListALL 获取UserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCircleApplyService *AppUserCircleApplyService) GetUserCircleApplyInfoListALL(info communityReq.UserCircleApplySearch) (list []community.UserCircleApply, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.UserCircleApply{}).Preload("Apply")
	var hkUserCircleApplys []community.UserCircleApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CircleId != 0 {
		db = db.Where("circle_id = ?", info.CircleId)
	}
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Find(&hkUserCircleApplys).Error
	return hkUserCircleApplys, total, err
}

// GetUserCircleApplyInfoListALL 获取UserCircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCircleApplyService *AppUserCircleApplyService) SetUserCircleApply(info communityReq.UserCircleApplyUpdate) (err error) {
	// 创建db
	db := global.GVA_DB
	err = db.Delete(&community.UserCircleApply{}, "user_id = ? and circle_id = ?", info.UserId, info.CircleId).Error
	if len(info.Apply) > 0 {
		for index, _ := range info.Apply {
			info.Apply[index].CircleId = info.CircleId
			info.Apply[index].UserId = info.UserId
		}
		err = db.Create(&info.Apply).Error
	}
	return err
}
