package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserRechargeService struct {
}

// CreateUserRecharge 创建UserRecharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserRechargeService *UserRechargeService) CreateUserRecharge(hkUserRecharge *community.UserRecharge) (err error) {
	err = global.GVA_DB.Create(hkUserRecharge).Error
	return err
}

// DeleteUserRecharge 删除UserRecharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserRechargeService *UserRechargeService) DeleteUserRecharge(hkUserRecharge community.UserRecharge) (err error) {
	err = global.GVA_DB.Delete(&hkUserRecharge).Error
	return err
}

// DeleteUserRechargeByIds 批量删除UserRecharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserRechargeService *UserRechargeService) DeleteUserRechargeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.UserRecharge{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserRecharge 更新UserRecharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserRechargeService *UserRechargeService) UpdateUserRecharge(hkUserRecharge community.UserRecharge) (err error) {
	err = global.GVA_DB.Save(&hkUserRecharge).Error
	return err
}

// GetUserRecharge 根据id获取UserRecharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserRechargeService *UserRechargeService) GetUserRecharge(id uint64) (hkUserRecharge community.UserRecharge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserRecharge).Error
	return
}

// GetUserRechargeInfoList 分页获取UserRecharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserRechargeService *UserRechargeService) GetUserRechargeInfoList(info communityReq.UserRechargeSearch) (list []community.UserRecharge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.UserRecharge{})
	var hkUserRecharges []community.UserRecharge
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
