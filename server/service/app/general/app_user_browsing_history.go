package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppUserBrowsingHistoryService struct {
}

// CreateUserBrowsingHistory 创建UserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) CreateUserBrowsingHistory(hkUserBrowsingHistory general.UserBrowsingHistory) (err error) {
	err = global.GVA_DB.Create(&hkUserBrowsingHistory).Error
	return err
}

// DeleteUserBrowsingHistory 删除UserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) DeleteUserBrowsingHistory(hkUserBrowsingHistory general.UserBrowsingHistory) (err error) {
	err = global.GVA_DB.Delete(&hkUserBrowsingHistory).Error
	return err
}

// DeleteUserBrowsingHistoryByIds 批量删除UserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) DeleteUserBrowsingHistoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.UserBrowsingHistory{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserBrowsingHistory 更新UserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) UpdateUserBrowsingHistory(hkUserBrowsingHistory general.UserBrowsingHistory) (err error) {
	err = global.GVA_DB.Save(&hkUserBrowsingHistory).Error
	return err
}

// GetUserBrowsingHistory 根据id获取UserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) GetUserBrowsingHistory(id uint) (hkUserBrowsingHistory general.UserBrowsingHistory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserBrowsingHistory).Error
	return
}

// GetUserBrowsingHistoryInfoList 分页获取UserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) GetUserBrowsingHistoryInfoList(info generalReq.UserBrowsingHistorySearch) (list []general.UserBrowsingHistory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&general.UserBrowsingHistory{})
	var hkUserBrowsingHistorys []general.UserBrowsingHistory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartUpdatedAt != nil && info.EndUpdatedAt != nil {
		db = db.Where("updated_at BETWEEN ? AND ?", info.StartUpdatedAt, info.EndUpdatedAt)
	}
	db = db.Where("user_id = ?", info.UserId)

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUserBrowsingHistorys).Error
	return hkUserBrowsingHistorys, total, err
}
