package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/general"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/general/request"
)

type HkUserCollectService struct {
}

// CreateHkUserCollect 创建HkUserCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserCollectService *HkUserCollectService) CreateHkUserCollect(hkUserCollect general.HkUserCollect) (err error) {
	err = global.GVA_DB.Create(&hkUserCollect).Error
	return err
}

// DeleteHkUserCollect 删除HkUserCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserCollectService *HkUserCollectService)DeleteHkUserCollect(hkUserCollect general.HkUserCollect) (err error) {
	err = global.GVA_DB.Delete(&hkUserCollect).Error
	return err
}

// DeleteHkUserCollectByIds 批量删除HkUserCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserCollectService *HkUserCollectService)DeleteHkUserCollectByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.HkUserCollect{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkUserCollect 更新HkUserCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserCollectService *HkUserCollectService)UpdateHkUserCollect(hkUserCollect general.HkUserCollect) (err error) {
	err = global.GVA_DB.Save(&hkUserCollect).Error
	return err
}

// GetHkUserCollect 根据id获取HkUserCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserCollectService *HkUserCollectService)GetHkUserCollect(id uint) (hkUserCollect general.HkUserCollect, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserCollect).Error
	return
}

// GetHkUserCollectInfoList 分页获取HkUserCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserCollectService *HkUserCollectService)GetHkUserCollectInfoList(info generalReq.HkUserCollectSearch) (list []general.HkUserCollect, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&general.HkUserCollect{})
    var hkUserCollects []general.HkUserCollect
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkUserCollects).Error
	return  hkUserCollects, total, err
}
