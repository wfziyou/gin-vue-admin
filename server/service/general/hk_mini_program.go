package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/general"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/general/request"
)

type HkMiniProgramService struct {
}

// CreateHkMiniProgram 创建HkMiniProgram记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkMiniProgramService *HkMiniProgramService) CreateHkMiniProgram(hkMiniProgram general.HkMiniProgram) (err error) {
	err = global.GVA_DB.Create(&hkMiniProgram).Error
	return err
}

// DeleteHkMiniProgram 删除HkMiniProgram记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkMiniProgramService *HkMiniProgramService)DeleteHkMiniProgram(hkMiniProgram general.HkMiniProgram) (err error) {
	err = global.GVA_DB.Delete(&hkMiniProgram).Error
	return err
}

// DeleteHkMiniProgramByIds 批量删除HkMiniProgram记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkMiniProgramService *HkMiniProgramService)DeleteHkMiniProgramByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.HkMiniProgram{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkMiniProgram 更新HkMiniProgram记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkMiniProgramService *HkMiniProgramService)UpdateHkMiniProgram(hkMiniProgram general.HkMiniProgram) (err error) {
	err = global.GVA_DB.Save(&hkMiniProgram).Error
	return err
}

// GetHkMiniProgram 根据id获取HkMiniProgram记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkMiniProgramService *HkMiniProgramService)GetHkMiniProgram(id uint) (hkMiniProgram general.HkMiniProgram, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkMiniProgram).Error
	return
}

// GetHkMiniProgramInfoList 分页获取HkMiniProgram记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkMiniProgramService *HkMiniProgramService)GetHkMiniProgramInfoList(info generalReq.HkMiniProgramSearch) (list []general.HkMiniProgram, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&general.HkMiniProgram{})
    var hkMiniPrograms []general.HkMiniProgram
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkMiniPrograms).Error
	return  hkMiniPrograms, total, err
}
