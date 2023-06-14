package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppApplyService struct {
}

// CreateApply 创建Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) CreateApply(hkApply apply.Apply) (err error) {
	err = global.GVA_DB.Create(&hkApply).Error
	return err
}

// DeleteApply 删除Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) DeleteApply(hkApply apply.Apply) (err error) {
	err = global.GVA_DB.Delete(&hkApply).Error
	return err
}

// DeleteApplyByIds 批量删除Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) DeleteApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.Apply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateApply 更新Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) UpdateApply(hkApply apply.Apply) (err error) {
	err = global.GVA_DB.Save(&hkApply).Error
	return err
}

// GetApply 根据id获取Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) GetApply(id uint64) (hkApply apply.Apply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkApply).Error
	return
}

// GetApplyInfoList 分页获取Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) GetApplyInfoList(info applyReq.ApplySearchReq) (list []apply.Apply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&apply.Apply{})
	var hkApplys []apply.Apply
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("owner_type = ?", info.OwnerType)
	//拥有者：0平台、1圈子、2个人
	if info.OwnerType == 1 {
		if info.CircleId > 0 {
			db = db.Where("circle_id = ?", info.CircleId)
		}
	} else if info.OwnerType == 2 {
		if info.UserId > 0 {
			db = db.Where("user_id = ?", info.UserId)
		}
	}
	if len(info.Name) > 0 {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkApplys).Error
	return hkApplys, total, err
}

// GetApplyInfoListAll 分页获取Apply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appApplyService *AppApplyService) GetApplyInfoListAll(info applyReq.ApplyAllSearchReq) (list []apply.Apply, err error) {
	// 创建db
	db := global.GVA_DB.Model(&apply.Apply{})
	var hkApplys []apply.Apply
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("owner_type = ?", info.OwnerType)
	//拥有者：0平台、1圈子、2个人
	if info.OwnerType == 1 {
		if info.CircleId > 0 {
			db = db.Where("circle_id = ?", info.CircleId)
		}
	} else if info.OwnerType == 2 {
		if info.UserId > 0 {
			db = db.Where("user_id = ?", info.UserId)
		}
	}
	if len(info.Keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+info.Keyword+"%")
	}

	err = db.Find(&hkApplys).Error
	return hkApplys, err
}
