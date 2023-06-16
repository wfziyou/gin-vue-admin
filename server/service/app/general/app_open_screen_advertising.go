package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppOpenScreenAdvertisingService struct {
}

// CreateOpenScreenAdvertising 创建OpenScreenAdvertising记录
// Author [piexlmax](https://github.com/piexlmax)
func (appOpenScreenAdvertisingService *AppOpenScreenAdvertisingService) CreateOpenScreenAdvertising(hkOpenScreenAdvertising general.OpenScreenAdvertising) (err error) {
	err = global.GVA_DB.Create(&hkOpenScreenAdvertising).Error
	return err
}

// DeleteOpenScreenAdvertising 删除OpenScreenAdvertising记录
// Author [piexlmax](https://github.com/piexlmax)
func (appOpenScreenAdvertisingService *AppOpenScreenAdvertisingService) DeleteOpenScreenAdvertising(hkOpenScreenAdvertising general.OpenScreenAdvertising) (err error) {
	err = global.GVA_DB.Delete(&hkOpenScreenAdvertising).Error
	return err
}

// DeleteOpenScreenAdvertisingByIds 批量删除OpenScreenAdvertising记录
// Author [piexlmax](https://github.com/piexlmax)
func (appOpenScreenAdvertisingService *AppOpenScreenAdvertisingService) DeleteOpenScreenAdvertisingByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.OpenScreenAdvertising{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOpenScreenAdvertising 更新OpenScreenAdvertising记录
// Author [piexlmax](https://github.com/piexlmax)
func (appOpenScreenAdvertisingService *AppOpenScreenAdvertisingService) UpdateOpenScreenAdvertising(hkOpenScreenAdvertising general.OpenScreenAdvertising) (err error) {
	err = global.GVA_DB.Save(&hkOpenScreenAdvertising).Error
	return err
}

// GetOpenScreenAdvertising 根据id获取OpenScreenAdvertising记录
// Author [piexlmax](https://github.com/piexlmax)
func (appOpenScreenAdvertisingService *AppOpenScreenAdvertisingService) GetOpenScreenAdvertising(id uint64) (hkOpenScreenAdvertising general.OpenScreenAdvertising, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkOpenScreenAdvertising).Error
	return
}

// GetOpenScreenAdvertisingByName 用名字查询协议
// Author [piexlmax](https://github.com/piexlmax)
func (appOpenScreenAdvertisingService *AppOpenScreenAdvertisingService) GetOpenScreenAdvertisingByName(name string) (hkOpenScreenAdvertising general.OpenScreenAdvertising, err error) {
	err = global.GVA_DB.Where("name = ?", name).First(&hkOpenScreenAdvertising).Error
	return
}

// GetOpenScreenAdvertisingInfo 分页获取OpenScreenAdvertising记录
// Author [piexlmax](https://github.com/piexlmax)
func (appOpenScreenAdvertisingService *AppOpenScreenAdvertisingService) GetOpenScreenAdvertisingInfo(info generalReq.GetOpenScreenAdvertisingReq) (hkOpenScreenAdvertising general.OpenScreenAdvertising, err error) {
	err = global.GVA_DB.Model(&general.OpenScreenAdvertising{}).Limit(1).First(&hkOpenScreenAdvertising).Error
	return
}
