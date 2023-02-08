package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppUserExtendService struct {
}

// CreateUserExtend 创建UserExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserExtendService *AppUserExtendService) CreateUserExtend(hkUserExtend community.UserExtend) (err error) {
	err = global.GVA_DB.Create(&hkUserExtend).Error
	return err
}

// DeleteUserExtend 删除UserExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserExtendService *AppUserExtendService) DeleteUserExtend(hkUserExtend community.UserExtend) (err error) {
	err = global.GVA_DB.Delete(&hkUserExtend).Error
	return err
}

// DeleteUserExtendByIds 批量删除UserExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserExtendService *AppUserExtendService) DeleteUserExtendByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.UserExtend{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserExtend 更新UserExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserExtendService *AppUserExtendService) UpdateUserExtend(hkUserExtend community.UserExtend) (err error) {
	err = global.GVA_DB.Save(&hkUserExtend).Error
	return err
}

// GetUserExtend 根据id获取UserExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserExtendService *AppUserExtendService) GetUserExtend(id uint) (hkUserExtend community.UserExtend, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserExtend).Error
	return
}

// GetUserExtendInfoList 分页获取UserExtend记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserExtendService *AppUserExtendService) GetUserExtendInfoList(info communityReq.UserExtendSearch) (list []community.UserExtend, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.UserExtend{})
	var hkUserExtends []community.UserExtend
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUserExtends).Error
	return hkUserExtends, total, err
}
