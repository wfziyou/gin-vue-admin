package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HkFocusUserService struct {
}

// CreateHkFocusUser 创建HkFocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *HkFocusUserService) CreateHkFocusUser(hkFocusUser community.HkFocusUser) (err error) {
	err = global.GVA_DB.Create(&hkFocusUser).Error
	return err
}

// DeleteHkFocusUser 删除HkFocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *HkFocusUserService) DeleteHkFocusUser(hkFocusUser community.HkFocusUser) (err error) {
	err = global.GVA_DB.Delete(&hkFocusUser).Error
	return err
}

// DeleteHkFocusUserByIds 批量删除HkFocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *HkFocusUserService) DeleteHkFocusUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkFocusUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkFocusUser 更新HkFocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *HkFocusUserService) UpdateHkFocusUser(hkFocusUser community.HkFocusUser) (err error) {
	err = global.GVA_DB.Save(&hkFocusUser).Error
	return err
}

// GetHkFocusUser 根据id获取HkFocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *HkFocusUserService) GetHkFocusUser(id uint64) (hkFocusUser community.HkFocusUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkFocusUser).Error
	return
}

// GetHkFocusUserInfoList 分页获取HkFocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *HkFocusUserService) GetHkFocusUserInfoList(info communityReq.HkFocusUserSearch) (list []community.HkFocusUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkFocusUser{})
	var hkFocusUsers []community.HkFocusUser
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
