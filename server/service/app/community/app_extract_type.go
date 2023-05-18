package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ExtractTypeService struct {
}

// CreateExtractType 创建ExtractType记录
func (service *ExtractTypeService) CreateExtractType(hkExtractType *community.ExtractType) (err error) {
	err = global.GVA_DB.Create(hkExtractType).Error
	return err
}

// DeleteExtractType 删除ExtractType记录
func (service *ExtractTypeService) DeleteExtractType(hkExtractType community.ExtractType) (err error) {
	err = global.GVA_DB.Delete(&hkExtractType).Error
	return err
}

// DeleteExtractTypeByIds 批量删除ExtractType记录
func (service *ExtractTypeService) DeleteExtractTypeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ExtractType{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateExtractType 更新ExtractType记录
func (service *ExtractTypeService) UpdateExtractType(hkExtractType community.ExtractType) (err error) {
	err = global.GVA_DB.Save(&hkExtractType).Error
	return err
}

// GetExtractType 根据id获取ExtractType记录
func (service *ExtractTypeService) GetExtractType(id uint) (hkExtractType community.ExtractType, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkExtractType).Error
	return
}

// GetExtractTypeInfoList 分页获取ExtractType记录
func (service *ExtractTypeService) GetExtractTypeInfoList() (list []community.ExtractType, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.ExtractType{})
	var hkExtractTypes []community.ExtractType

	err = db.Find(&hkExtractTypes).Error
	return hkExtractTypes, err
}
