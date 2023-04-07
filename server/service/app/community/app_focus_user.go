package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
)

type FocusUserService struct {
}

// CreateFocusUser 创建FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) CreateFocusUser(hkFocusUser community.FocusUser) (err error) {
	err = global.GVA_DB.Create(&hkFocusUser).Error
	return err
}

// DeleteFocusUser 删除FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) DeleteFocusUser(hkFocusUser community.FocusUser) (err error) {
	err = global.GVA_DB.Delete(&hkFocusUser).Error
	return err
}

// DeleteFocusUserByIds 批量删除FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) DeleteFocusUserByIds(Ids []uint64) (err error) {
	err = global.GVA_DB.Delete(&[]community.FocusUser{}, "id in ?", Ids).Error
	return err
}

// UpdateFocusUser 更新FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) UpdateFocusUser(hkFocusUser community.FocusUser) (err error) {
	err = global.GVA_DB.Save(&hkFocusUser).Error
	return err
}

// GetFocusUser 根据id获取FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) GetFocusUser(id uint64) (hkFocusUser community.FocusUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkFocusUser).Error
	return
}

// GetFocusUserInfoList 分页获取FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) GetFocusUserInfoList(info communityReq.FocusUserSearch) (list []community.FocusUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.FocusUser{})
	var hkFocusUsers []community.FocusUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkFocusUsers).Error
	return hkFocusUsers, total, err
}
