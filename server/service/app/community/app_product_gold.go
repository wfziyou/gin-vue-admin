package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProductGoldService struct {
}

// CreateProductGold 创建ProductGold记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *ProductGoldService) CreateProductGold(hkProductGold *community.ProductGold) (err error) {
	err = global.GVA_DB.Create(hkProductGold).Error
	return err
}

// DeleteProductGold 删除ProductGold记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *ProductGoldService) DeleteProductGold(hkProductGold community.ProductGold) (err error) {
	err = global.GVA_DB.Delete(&hkProductGold).Error
	return err
}

// DeleteProductGoldByIds 批量删除ProductGold记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *ProductGoldService) DeleteProductGoldByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ProductGold{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProductGold 更新ProductGold记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *ProductGoldService) UpdateProductGold(hkProductGold community.ProductGold) (err error) {
	err = global.GVA_DB.Save(&hkProductGold).Error
	return err
}

// GetProductGold 根据id获取ProductGold记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *ProductGoldService) GetProductGold(id uint) (hkProductGold community.ProductGold, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkProductGold).Error
	return
}

// GetProductGoldInfoList 分页获取ProductGold记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *ProductGoldService) GetProductGoldInfoList() (list []community.ProductGoldInfo, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.ProductGold{})
	var hkProductGolds []community.ProductGoldInfo

	err = db.Find(&hkProductGolds).Error
	return hkProductGolds, err
}
