package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ThirdPlatformService struct {
}

// CreateHkThirdPlatform 创建HkThirdPlatform记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkThirdPlatformService *ThirdPlatformService) CreateHkThirdPlatform(hkThirdPlatform *community.HkThirdPlatform) (err error) {
	err = global.GVA_DB.Create(hkThirdPlatform).Error
	return err
}

// DeleteHkThirdPlatform 删除HkThirdPlatform记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkThirdPlatformService *ThirdPlatformService) DeleteHkThirdPlatform(hkThirdPlatform community.HkThirdPlatform) (err error) {
	err = global.GVA_DB.Delete(&hkThirdPlatform).Error
	return err
}

// DeleteHkThirdPlatformByIds 批量删除HkThirdPlatform记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkThirdPlatformService *ThirdPlatformService) DeleteHkThirdPlatformByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkThirdPlatform{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkThirdPlatform 更新HkThirdPlatform记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkThirdPlatformService *ThirdPlatformService) UpdateHkThirdPlatform(hkThirdPlatform community.HkThirdPlatform) (err error) {
	err = global.GVA_DB.Save(&hkThirdPlatform).Error
	return err
}

// GetHkThirdPlatform 根据id获取HkThirdPlatform记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkThirdPlatformService *ThirdPlatformService) GetHkThirdPlatform(id uint) (hkThirdPlatform community.HkThirdPlatform, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkThirdPlatform).Error
	return
}

// GetThirdPlatformInfoList 分页获取HkThirdPlatform记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkThirdPlatformService *ThirdPlatformService) GetThirdPlatformInfoList() (list []community.HkThirdPlatform, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.HkThirdPlatform{})
	var hkThirdPlatforms []community.HkThirdPlatform

	err = db.Where("status = 1").Find(&hkThirdPlatforms).Error
	return hkThirdPlatforms, err
}
