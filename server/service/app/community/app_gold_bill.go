package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HkGoldBillService struct {
}

// CreateHkGoldBill 创建HkGoldBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkGoldBillService *HkGoldBillService) CreateHkGoldBill(hkGoldBill *community.HkGoldBill) (err error) {
	err = global.GVA_DB.Create(hkGoldBill).Error
	return err
}

// DeleteHkGoldBill 删除HkGoldBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkGoldBillService *HkGoldBillService) DeleteHkGoldBill(hkGoldBill community.HkGoldBill) (err error) {
	err = global.GVA_DB.Delete(&hkGoldBill).Error
	return err
}

// DeleteHkGoldBillByIds 批量删除HkGoldBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkGoldBillService *HkGoldBillService) DeleteHkGoldBillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkGoldBill{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkGoldBill 更新HkGoldBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkGoldBillService *HkGoldBillService) UpdateHkGoldBill(hkGoldBill community.HkGoldBill) (err error) {
	err = global.GVA_DB.Save(&hkGoldBill).Error
	return err
}

// GetHkGoldBill 根据id获取HkGoldBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkGoldBillService *HkGoldBillService) GetHkGoldBill(id uint64) (hkGoldBill community.HkGoldBill, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkGoldBill).Error
	return
}

// GetHkGoldBillInfoList 分页获取HkGoldBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkGoldBillService *HkGoldBillService) GetHkGoldBillInfoList(info communityReq.HkGoldBillSearch) (list []community.HkGoldBill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkGoldBill{})
	var hkGoldBills []community.HkGoldBill
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkGoldBills).Error
	return hkGoldBills, total, err
}
