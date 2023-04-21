package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HkOrderService struct {
}

// CreateHkOrder 创建HkOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkOrderService *HkOrderService) CreateHkOrder(hkOrder *community.HkOrder) (err error) {
	err = global.GVA_DB.Create(hkOrder).Error
	return err
}

// DeleteHkOrder 删除HkOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkOrderService *HkOrderService) DeleteHkOrder(hkOrder community.HkOrder) (err error) {
	err = global.GVA_DB.Delete(&hkOrder).Error
	return err
}

// DeleteHkOrderByIds 批量删除HkOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkOrderService *HkOrderService) DeleteHkOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkOrder{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkOrder 更新HkOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkOrderService *HkOrderService) UpdateHkOrder(hkOrder community.HkOrder) (err error) {
	err = global.GVA_DB.Save(&hkOrder).Error
	return err
}

// GetHkOrder 根据id获取HkOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkOrderService *HkOrderService) GetHkOrder(id uint64) (hkOrder community.HkOrder, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkOrder).Error
	return
}

// GetHkOrderInfoList 分页获取HkOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkOrderService *HkOrderService) GetHkOrderInfoList(info communityReq.HkOrderSearch) (list []community.HkOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkOrder{})
	var hkOrders []community.HkOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkOrders).Error
	return hkOrders, total, err
}
