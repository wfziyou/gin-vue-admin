package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HkUserBillService struct {
}

// CreateHkUserBill 创建HkUserBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBillService *HkUserBillService) CreateHkUserBill(hkUserBill *community.HkUserBill) (err error) {
	err = global.GVA_DB.Create(hkUserBill).Error
	return err
}

// DeleteHkUserBill 删除HkUserBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBillService *HkUserBillService) DeleteHkUserBill(hkUserBill community.HkUserBill) (err error) {
	err = global.GVA_DB.Delete(&hkUserBill).Error
	return err
}

// DeleteHkUserBillByIds 批量删除HkUserBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBillService *HkUserBillService) DeleteHkUserBillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkUserBill{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkUserBill 更新HkUserBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBillService *HkUserBillService) UpdateHkUserBill(hkUserBill community.HkUserBill) (err error) {
	err = global.GVA_DB.Save(&hkUserBill).Error
	return err
}

// GetHkUserBill 根据id获取HkUserBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBillService *HkUserBillService) GetHkUserBill(id uint64) (hkUserBill community.HkUserBill, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserBill).Error
	return
}

// GetHkUserBillInfoList 分页获取HkUserBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBillService *HkUserBillService) GetHkUserBillInfoList(info communityReq.HkUserBillSearch) (list []community.HkUserBill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkUserBill{})
	var hkUserBills []community.HkUserBill
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUserBills).Error
	return hkUserBills, total, err
}
