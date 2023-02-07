package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type HkCircleUserService struct {
}

// CreateHkCircleUser 创建HkCircleUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleUserService *HkCircleUserService) CreateHkCircleUser(hkCircleUser community.HkCircleUser) (err error) {
	err = global.GVA_DB.Create(&hkCircleUser).Error
	return err
}

// DeleteHkCircleUser 删除HkCircleUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleUserService *HkCircleUserService) DeleteHkCircleUser(hkCircleUser community.HkCircleUser) (err error) {
	err = global.GVA_DB.Delete(&hkCircleUser).Error
	return err
}

// DeleteHkCircleUserByIds 批量删除HkCircleUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleUserService *HkCircleUserService) DeleteHkCircleUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkCircleUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkCircleUser 更新HkCircleUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleUserService *HkCircleUserService) UpdateHkCircleUser(hkCircleUser community.HkCircleUser) (err error) {
	err = global.GVA_DB.Save(&hkCircleUser).Error
	return err
}

// GetHkCircleUser 根据id获取HkCircleUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleUserService *HkCircleUserService) GetHkCircleUser(id uint) (hkCircleUser community.HkCircleUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleUser).Error
	return
}

// GetHkCircleUserInfoList 分页获取HkCircleUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleUserService *HkCircleUserService) GetHkCircleUserInfoList(info communityReq.HkCircleUserSearch) (list []community.HkCircleUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkCircleUser{})
	var hkCircleUsers []community.HkCircleUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircleUsers).Error
	return hkCircleUsers, total, err
}
