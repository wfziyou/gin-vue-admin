package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppVersionService struct {
}

// CreateAppVersion 创建AppVersion记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *AppVersionService) CreateAppVersion(appVersion general.AppVersion) (err error) {
	err = global.GVA_DB.Create(&appVersion).Error
	return err
}

// DeleteAppVersion 删除AppVersion记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *AppVersionService) DeleteAppVersion(appVersion general.AppVersion) (err error) {
	err = global.GVA_DB.Delete(&appVersion).Error
	return err
}

// DeleteAppVersionByIds 批量删除AppVersion记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *AppVersionService) DeleteAppVersionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.AppVersion{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppVersion 更新AppVersion记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *AppVersionService) UpdateAppVersion(appVersion general.AppVersion) (err error) {
	err = global.GVA_DB.Save(&appVersion).Error
	return err
}

// GetAppVersion 根据id获取AppVersion记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *AppVersionService) GetAppVersion(id uint64) (appVersion general.AppVersion, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appVersion).Error
	return
}

func (service *AppVersionService) CheckAppUpdate(info generalReq.ParamCheckAppUpdate) (obj general.AppVersion, err error) {
	var number int64
	err = global.GVA_DB.Model(&general.AppVersion{}).Where("platform = ? AND version > ? AND force_update = 1", info.Platform, info.CurVersion).Count(&number).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(&general.AppVersion{}).Where("platform = ?", info.Platform).Order("created_at desc").Limit(1).First(&obj).Error
	if err != nil {
		return
	}
	if number > 0 {
		obj.ForceUpdate = 1
	}
	return obj, err
}
