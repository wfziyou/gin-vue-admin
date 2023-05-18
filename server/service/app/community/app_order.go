package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	uuid "github.com/satori/go.uuid"
)

type OrderService struct {
}

// CreateOrder 创建Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkOrderService *OrderService) CreateOrder(userId uint64, info communityReq.ParamCreateOrder) (err error) {
	var product community.ProductGold
	err = global.GVA_DB.Where("id = ?", info.ProductId).First(&product).Error
	if err != nil {
		return errors.New("产品不存在")
	}
	if info.CartNum < 1 {
		info.CartNum = 1
	}
	totalPrice := product.Price * uint64(info.CartNum)
	totalCost := product.Cost * uint64(info.CartNum)
	order := community.Order{
		OrderId:    uuid.NewV4().String(),
		UserId:     userId,
		ProductId:  info.ProductId,
		CartNum:    info.CartNum,
		TotalPrice: totalPrice,
		Paid:       community.OrderPaidFalse,
		PayType:    info.PayType,
		Cost:       totalCost,
	}
	err = global.GVA_DB.Create(&order).Error
	return err
}

// DeleteOrder 删除Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkOrderService *OrderService) DeleteOrder(hkOrder community.Order) (err error) {
	err = global.GVA_DB.Delete(&hkOrder).Error
	return err
}

// DeleteOrderByIds 批量删除Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkOrderService *OrderService) DeleteOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.Order{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrder 更新Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkOrderService *OrderService) UpdateOrder(hkOrder community.Order) (err error) {
	err = global.GVA_DB.Save(&hkOrder).Error
	return err
}

// GetOrder 根据id获取Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkOrderService *OrderService) GetOrder(id uint64) (hkOrder community.Order, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkOrder).Error
	return
}

// GetOrderInfoList 分页获取Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkOrderService *OrderService) GetOrderInfoList(info communityReq.OrderSearch) (list []community.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.Order{})
	var hkOrders []community.Order
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
