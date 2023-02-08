package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForbiddenSpeakService struct {
}

// CreateForbiddenSpeak 创建ForbiddenSpeak记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakService *AppForbiddenSpeakService) CreateForbiddenSpeak(hkForbiddenSpeak community.ForbiddenSpeak) (err error) {
	err = global.GVA_DB.Create(&hkForbiddenSpeak).Error
	return err
}

// DeleteForbiddenSpeak 删除ForbiddenSpeak记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakService *AppForbiddenSpeakService) DeleteForbiddenSpeak(hkForbiddenSpeak community.ForbiddenSpeak) (err error) {
	err = global.GVA_DB.Delete(&hkForbiddenSpeak).Error
	return err
}

// DeleteForbiddenSpeakByIds 批量删除ForbiddenSpeak记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakService *AppForbiddenSpeakService) DeleteForbiddenSpeakByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForbiddenSpeak{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateForbiddenSpeak 更新ForbiddenSpeak记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakService *AppForbiddenSpeakService) UpdateForbiddenSpeak(hkForbiddenSpeak community.ForbiddenSpeak) (err error) {
	err = global.GVA_DB.Save(&hkForbiddenSpeak).Error
	return err
}

// GetForbiddenSpeak 根据id获取ForbiddenSpeak记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakService *AppForbiddenSpeakService) GetForbiddenSpeak(id uint) (hkForbiddenSpeak community.ForbiddenSpeak, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForbiddenSpeak).Error
	return
}

// GetForbiddenSpeakInfoList 分页获取ForbiddenSpeak记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForbiddenSpeakService *AppForbiddenSpeakService) GetForbiddenSpeakInfoList(info communityReq.ForbiddenSpeakSearch) (list []community.ForbiddenSpeak, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForbiddenSpeak{})
	var hkForbiddenSpeaks []community.ForbiddenSpeak
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForbiddenSpeaks).Error
	return hkForbiddenSpeaks, total, err
}
