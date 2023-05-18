package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PayTypeService struct {
}

// CreatePayType 创建PayType记录
func (service *PayTypeService) CreatePayType(hkPayType *community.PayType) (err error) {
	err = global.GVA_DB.Create(hkPayType).Error
	return err
}

// DeletePayType 删除PayType记录
func (service *PayTypeService) DeletePayType(hkPayType community.PayType) (err error) {
	err = global.GVA_DB.Delete(&hkPayType).Error
	return err
}

// DeletePayTypeByIds 批量删除PayType记录
func (service *PayTypeService) DeletePayTypeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.PayType{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePayType 更新PayType记录
func (service *PayTypeService) UpdatePayType(hkPayType community.PayType) (err error) {
	err = global.GVA_DB.Save(&hkPayType).Error
	return err
}

// GetPayType 根据id获取PayType记录
func (service *PayTypeService) GetPayType(id uint) (hkPayType community.PayType, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkPayType).Error
	return
}

// GetPayTypeInfoList 分页获取PayType记录
func (service *PayTypeService) GetPayTypeInfoList() (list []community.PayTypeInfo, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.PayType{})
	var hkPayTypes []community.PayTypeInfo

	err = db.Find(&hkPayTypes).Error
	return hkPayTypes, err
}
