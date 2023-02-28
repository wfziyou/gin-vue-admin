package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppMiniProgramService struct {
}

// CreateMiniProgram 创建MiniProgram记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMiniProgramService *AppMiniProgramService) CreateMiniProgram(hkMiniProgram general.MiniProgram) (err error) {
	err = global.GVA_DB.Create(&hkMiniProgram).Error
	return err
}

// DeleteMiniProgram 删除MiniProgram记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMiniProgramService *AppMiniProgramService) DeleteMiniProgram(hkMiniProgram general.MiniProgram) (err error) {
	err = global.GVA_DB.Delete(&hkMiniProgram).Error
	return err
}

// DeleteMiniProgramByIds 批量删除MiniProgram记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMiniProgramService *AppMiniProgramService) DeleteMiniProgramByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.MiniProgram{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMiniProgram 更新MiniProgram记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMiniProgramService *AppMiniProgramService) UpdateMiniProgram(hkMiniProgram general.MiniProgram) (err error) {
	err = global.GVA_DB.Save(&hkMiniProgram).Error
	return err
}

// GetMiniProgram 根据id获取MiniProgram记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMiniProgramService *AppMiniProgramService) GetMiniProgram(id uint64) (hkMiniProgram general.MiniProgramBaseInfo, err error) {
	err = global.GVA_DB.Raw("SELECT t.name,t.icon,t.company_name,t.program_id,p.version,p.code,p.packet_address,t.hidden FROM hk_mini_program t LEFT JOIN hk_mini_program_packet p ON  t.cur_packet_id = p.id WHERE t.id = ?", id).First(&hkMiniProgram).Error
	//err = global.GVA_DB.Where("id = ?", id).First(&hkMiniProgram).Error
	return
}

// GetMiniProgramInfoList 分页获取MiniProgram记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMiniProgramService *AppMiniProgramService) GetMiniProgramInfoList(info generalReq.MiniProgramSearch) (list []general.MiniProgram, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&general.MiniProgram{})
	var hkMiniPrograms []general.MiniProgram
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkMiniPrograms).Error
	return hkMiniPrograms, total, err
}
