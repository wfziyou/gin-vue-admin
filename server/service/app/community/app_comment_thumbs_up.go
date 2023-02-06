package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type AppCommentThumbsUpService struct {
}

// CreateHkCommentThumbsUp 创建HkCommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) CreateHkCommentThumbsUp(hkCommentThumbsUp community.HkCommentThumbsUp) (err error) {
	err = global.GVA_DB.Create(&hkCommentThumbsUp).Error
	return err
}

// DeleteHkCommentThumbsUp 删除HkCommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) DeleteHkCommentThumbsUp(hkCommentThumbsUp community.HkCommentThumbsUp) (err error) {
	err = global.GVA_DB.Delete(&hkCommentThumbsUp).Error
	return err
}

// DeleteHkCommentThumbsUpByIds 批量删除HkCommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) DeleteHkCommentThumbsUpByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkCommentThumbsUp{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkCommentThumbsUp 更新HkCommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) UpdateHkCommentThumbsUp(hkCommentThumbsUp community.HkCommentThumbsUp) (err error) {
	err = global.GVA_DB.Save(&hkCommentThumbsUp).Error
	return err
}

// GetHkCommentThumbsUp 根据id获取HkCommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) GetHkCommentThumbsUp(id uint) (hkCommentThumbsUp community.HkCommentThumbsUp, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCommentThumbsUp).Error
	return
}

// GetHkCommentThumbsUpInfoList 分页获取HkCommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) GetHkCommentThumbsUpInfoList(info communityReq.HkCommentThumbsUpSearch) (list []community.HkCommentThumbsUp, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkCommentThumbsUp{})
	var hkCommentThumbsUps []community.HkCommentThumbsUp
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCommentThumbsUps).Error
	return hkCommentThumbsUps, total, err
}
