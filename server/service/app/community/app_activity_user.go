package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HkActivityUserService struct {
}

// CreateHkActivityUser 创建HkActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *HkActivityUserService) CreateHkActivityUser(hkActivityUser community.HkActivityUser) (err error) {
	err = global.GVA_DB.Create(&hkActivityUser).Error
	return err
}

// DeleteHkActivityUser 删除HkActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *HkActivityUserService) DeleteHkActivityUser(hkActivityUser community.HkActivityUser) (err error) {
	err = global.GVA_DB.Delete(&hkActivityUser).Error
	return err
}

// DeleteHkActivityUserByIds 批量删除HkActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *HkActivityUserService) DeleteHkActivityUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkActivityUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkActivityUser 更新HkActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *HkActivityUserService) UpdateHkActivityUser(hkActivityUser community.HkActivityUser) (err error) {
	err = global.GVA_DB.Save(&hkActivityUser).Error
	return err
}

// GetHkActivityUser 根据id获取HkActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *HkActivityUserService) GetHkActivityUser(id uint64) (hkActivityUser community.HkActivityUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkActivityUser).Error
	return
}

// GetHkActivityUserInfoList 分页获取HkActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *HkActivityUserService) GetHkActivityUserInfoList(info communityReq.HkActivityUserSearch) (list []community.HkActivityUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkActivityUser{})
	var hkActivityUsers []community.HkActivityUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkActivityUsers).Error
	return hkActivityUsers, total, err
}
