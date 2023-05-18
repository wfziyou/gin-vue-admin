package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserBillService struct {
}

func (hkUserBillService *UserBillService) GetPayTypeList() (list []community.UserBill, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.UserBill{})
	var hkUserBills []community.UserBill
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Find(&hkUserBills).Error
	return hkUserBills, total, err
}

// CreateUserBill 创建UserBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBillService *UserBillService) CreateUserBill(hkUserBill *community.UserBill) (err error) {
	err = global.GVA_DB.Create(hkUserBill).Error
	return err
}

// DeleteUserBill 删除UserBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBillService *UserBillService) DeleteUserBill(hkUserBill community.UserBill) (err error) {
	err = global.GVA_DB.Delete(&hkUserBill).Error
	return err
}

// DeleteUserBillByIds 批量删除UserBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBillService *UserBillService) DeleteUserBillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.UserBill{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserBill 更新UserBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBillService *UserBillService) UpdateUserBill(hkUserBill community.UserBill) (err error) {
	err = global.GVA_DB.Save(&hkUserBill).Error
	return err
}

// GetUserBill 根据id获取UserBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBillService *UserBillService) GetUserBill(id uint64) (hkUserBill community.UserBill, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserBill).Error
	return
}

// GetUserBillInfoList 分页获取UserBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserBillService *UserBillService) GetUserBillInfoList(userId uint64, info communityReq.UserBillSearch) (list []community.UserBill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.UserBill{})
	var hkUserBills []community.UserBill

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

	err = db.Limit(limit).Offset(offset).Find(&hkUserBills).Error
	return hkUserBills, total, err
}
