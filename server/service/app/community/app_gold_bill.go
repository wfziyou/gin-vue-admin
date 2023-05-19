package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GoldBillService struct {
}

// CreateGoldBill 创建GoldBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkGoldBillService *GoldBillService) CreateGoldBill(hkGoldBill *community.GoldBill) (err error) {
	err = global.GVA_DB.Create(hkGoldBill).Error
	return err
}

// DeleteGoldBill 删除GoldBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkGoldBillService *GoldBillService) DeleteGoldBill(hkGoldBill community.GoldBill) (err error) {
	err = global.GVA_DB.Delete(&hkGoldBill).Error
	return err
}
func (hkGoldBillService *GoldBillService) UpdateGoldBillLinkId(hkGoldBill community.GoldBill, linkId uint64) (err error) {
	db := global.GVA_DB.Model(hkGoldBill)
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	updateData["link_id"] = linkId
	err = db.Where("id = ?", hkGoldBill.ID).Updates(updateData).Error
	return err
}

// DeleteGoldBillByIds 批量删除GoldBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkGoldBillService *GoldBillService) DeleteGoldBillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.GoldBill{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateGoldBill 更新GoldBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkGoldBillService *GoldBillService) UpdateGoldBill(hkGoldBill community.GoldBill) (err error) {
	err = global.GVA_DB.Save(&hkGoldBill).Error
	return err
}

// GetGoldBill 根据id获取GoldBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkGoldBillService *GoldBillService) GetGoldBill(id uint64) (hkGoldBill community.GoldBill, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkGoldBill).Error
	return
}

// GetGoldBillInfoList 分页获取GoldBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkGoldBillService *GoldBillService) GetGoldBillInfoList(userId uint64, info communityReq.GoldBillSearch) (list []community.GoldBill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.GoldBill{})
	var hkGoldBills []community.GoldBill

	db = db.Where("user_id = ?", userId)
	if info.Pm != nil {
		db = db.Where("pm = ?", info.Pm)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
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
