package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/general"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/general/request"
)

type HkUserBrowsingHistoryService struct {
}

// CreateHkUserBrowsingHistory 创建HkUserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBrowsingHistoryService *HkUserBrowsingHistoryService) CreateHkUserBrowsingHistory(hkUserBrowsingHistory general.HkUserBrowsingHistory) (err error) {
	err = global.GVA_DB.Create(&hkUserBrowsingHistory).Error
	return err
}

// DeleteHkUserBrowsingHistory 删除HkUserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBrowsingHistoryService *HkUserBrowsingHistoryService)DeleteHkUserBrowsingHistory(hkUserBrowsingHistory general.HkUserBrowsingHistory) (err error) {
	err = global.GVA_DB.Delete(&hkUserBrowsingHistory).Error
	return err
}

// DeleteHkUserBrowsingHistoryByIds 批量删除HkUserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBrowsingHistoryService *HkUserBrowsingHistoryService)DeleteHkUserBrowsingHistoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.HkUserBrowsingHistory{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkUserBrowsingHistory 更新HkUserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBrowsingHistoryService *HkUserBrowsingHistoryService)UpdateHkUserBrowsingHistory(hkUserBrowsingHistory general.HkUserBrowsingHistory) (err error) {
	err = global.GVA_DB.Save(&hkUserBrowsingHistory).Error
	return err
}

// GetHkUserBrowsingHistory 根据id获取HkUserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBrowsingHistoryService *HkUserBrowsingHistoryService)GetHkUserBrowsingHistory(id uint) (hkUserBrowsingHistory general.HkUserBrowsingHistory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserBrowsingHistory).Error
	return
}

// GetHkUserBrowsingHistoryInfoList 分页获取HkUserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBrowsingHistoryService *HkUserBrowsingHistoryService)GetHkUserBrowsingHistoryInfoList(info generalReq.HkUserBrowsingHistorySearch) (list []general.HkUserBrowsingHistory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&general.HkUserBrowsingHistory{})
    var hkUserBrowsingHistorys []general.HkUserBrowsingHistory
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkUserBrowsingHistorys).Error
	return  hkUserBrowsingHistorys, total, err
}
