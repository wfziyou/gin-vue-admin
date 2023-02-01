package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type HkForbiddenSpeakService struct {
}

// CreateHkForbiddenSpeak 创建HkForbiddenSpeak记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForbiddenSpeakService *HkForbiddenSpeakService) CreateHkForbiddenSpeak(hkForbiddenSpeak community.HkForbiddenSpeak) (err error) {
	err = global.GVA_DB.Create(&hkForbiddenSpeak).Error
	return err
}

// DeleteHkForbiddenSpeak 删除HkForbiddenSpeak记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForbiddenSpeakService *HkForbiddenSpeakService)DeleteHkForbiddenSpeak(hkForbiddenSpeak community.HkForbiddenSpeak) (err error) {
	err = global.GVA_DB.Delete(&hkForbiddenSpeak).Error
	return err
}

// DeleteHkForbiddenSpeakByIds 批量删除HkForbiddenSpeak记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForbiddenSpeakService *HkForbiddenSpeakService)DeleteHkForbiddenSpeakByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForbiddenSpeak{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkForbiddenSpeak 更新HkForbiddenSpeak记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForbiddenSpeakService *HkForbiddenSpeakService)UpdateHkForbiddenSpeak(hkForbiddenSpeak community.HkForbiddenSpeak) (err error) {
	err = global.GVA_DB.Save(&hkForbiddenSpeak).Error
	return err
}

// GetHkForbiddenSpeak 根据id获取HkForbiddenSpeak记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForbiddenSpeakService *HkForbiddenSpeakService)GetHkForbiddenSpeak(id uint) (hkForbiddenSpeak community.HkForbiddenSpeak, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForbiddenSpeak).Error
	return
}

// GetHkForbiddenSpeakInfoList 分页获取HkForbiddenSpeak记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForbiddenSpeakService *HkForbiddenSpeakService)GetHkForbiddenSpeakInfoList(info communityReq.HkForbiddenSpeakSearch) (list []community.HkForbiddenSpeak, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&community.HkForbiddenSpeak{})
    var hkForbiddenSpeaks []community.HkForbiddenSpeak
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkForbiddenSpeaks).Error
	return  hkForbiddenSpeaks, total, err
}
