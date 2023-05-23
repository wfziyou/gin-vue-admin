package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HelpTypeService struct {
}

// CreateHelpType 创建HelpType记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkHelpTypeService *HelpTypeService) CreateHelpType(hkHelpType *community.HelpType) (err error) {
	err = global.GVA_DB.Create(hkHelpType).Error
	return err
}

// DeleteHelpType 删除HelpType记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkHelpTypeService *HelpTypeService) DeleteHelpType(hkHelpType community.HelpType) (err error) {
	err = global.GVA_DB.Delete(&hkHelpType).Error
	return err
}

// DeleteHelpTypeByIds 批量删除HelpType记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkHelpTypeService *HelpTypeService) DeleteHelpTypeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HelpType{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHelpType 更新HelpType记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkHelpTypeService *HelpTypeService) UpdateHelpType(hkHelpType community.HelpType) (err error) {
	err = global.GVA_DB.Save(&hkHelpType).Error
	return err
}

// GetHelpType 根据id获取HelpType记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkHelpTypeService *HelpTypeService) GetHelpType(id uint64) (hkHelpType community.HelpType, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkHelpType).Error
	return
}

// GetHelpTypeInfoList 获取HelpType记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkHelpTypeService *HelpTypeService) GetHelpTypeInfoList() (list []community.HelpType, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.HelpType{})
	var hkHelpTypes []community.HelpType
	err = db.Find(&hkHelpTypes).Error
	return hkHelpTypes, err
}
