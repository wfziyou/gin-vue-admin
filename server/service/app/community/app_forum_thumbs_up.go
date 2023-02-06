package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type AppForumThumbsUpService struct {
}

// CreateHkForumThumbsUp 创建HkForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) CreateHkForumThumbsUp(hkForumThumbsUp community.HkForumThumbsUp) (err error) {
	err = global.GVA_DB.Create(&hkForumThumbsUp).Error
	return err
}

// DeleteHkForumThumbsUp 删除HkForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) DeleteHkForumThumbsUp(hkForumThumbsUp community.HkForumThumbsUp) (err error) {
	err = global.GVA_DB.Delete(&hkForumThumbsUp).Error
	return err
}

// DeleteHkForumThumbsUpByIds 批量删除HkForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) DeleteHkForumThumbsUpByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForumThumbsUp{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkForumThumbsUp 更新HkForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) UpdateHkForumThumbsUp(hkForumThumbsUp community.HkForumThumbsUp) (err error) {
	err = global.GVA_DB.Save(&hkForumThumbsUp).Error
	return err
}

// GetHkForumThumbsUp 根据id获取HkForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) GetHkForumThumbsUp(id uint) (hkForumThumbsUp community.HkForumThumbsUp, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumThumbsUp).Error
	return
}

// GetHkForumThumbsUpInfoList 分页获取HkForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) GetHkForumThumbsUpInfoList(info communityReq.HkForumThumbsUpSearch) (list []community.HkForumThumbsUp, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkForumThumbsUp{})
	var hkForumThumbsUps []community.HkForumThumbsUp
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumThumbsUps).Error
	return hkForumThumbsUps, total, err
}
