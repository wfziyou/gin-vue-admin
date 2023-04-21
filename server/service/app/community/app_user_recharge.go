package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HkUserRechargeService struct {
}

// CreateHkUserRecharge 创建HkUserRecharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserRechargeService *HkUserRechargeService) CreateHkUserRecharge(hkUserRecharge *community.HkUserRecharge) (err error) {
	err = global.GVA_DB.Create(hkUserRecharge).Error
	return err
}

// DeleteHkUserRecharge 删除HkUserRecharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserRechargeService *HkUserRechargeService) DeleteHkUserRecharge(hkUserRecharge community.HkUserRecharge) (err error) {
	err = global.GVA_DB.Delete(&hkUserRecharge).Error
	return err
}

// DeleteHkUserRechargeByIds 批量删除HkUserRecharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserRechargeService *HkUserRechargeService) DeleteHkUserRechargeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkUserRecharge{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkUserRecharge 更新HkUserRecharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserRechargeService *HkUserRechargeService) UpdateHkUserRecharge(hkUserRecharge community.HkUserRecharge) (err error) {
	err = global.GVA_DB.Save(&hkUserRecharge).Error
	return err
}

// GetHkUserRecharge 根据id获取HkUserRecharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserRechargeService *HkUserRechargeService) GetHkUserRecharge(id uint64) (hkUserRecharge community.HkUserRecharge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserRecharge).Error
	return
}

// GetHkUserRechargeInfoList 分页获取HkUserRecharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserRechargeService *HkUserRechargeService) GetHkUserRechargeInfoList(info communityReq.HkUserRechargeSearch) (list []community.HkUserRecharge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkUserRecharge{})
	var hkUserRecharges []community.HkUserRecharge
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUserRecharges).Error
	return hkUserRecharges, total, err
}
