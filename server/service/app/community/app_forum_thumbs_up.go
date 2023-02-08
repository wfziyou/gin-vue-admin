package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForumThumbsUpService struct {
}

// CreateForumThumbsUp 创建ForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) CreateForumThumbsUp(hkForumThumbsUp community.ForumThumbsUp) (err error) {
	err = global.GVA_DB.Create(&hkForumThumbsUp).Error
	return err
}

// DeleteForumThumbsUp 删除ForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) DeleteForumThumbsUp(hkForumThumbsUp community.ForumThumbsUp) (err error) {
	err = global.GVA_DB.Delete(&hkForumThumbsUp).Error
	return err
}

// DeleteForumThumbsUpByIds 批量删除ForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) DeleteForumThumbsUpByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumThumbsUp{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateForumThumbsUp 更新ForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) UpdateForumThumbsUp(hkForumThumbsUp community.ForumThumbsUp) (err error) {
	err = global.GVA_DB.Save(&hkForumThumbsUp).Error
	return err
}

// GetForumThumbsUp 根据id获取ForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) GetForumThumbsUp(id uint) (hkForumThumbsUp community.ForumThumbsUp, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumThumbsUp).Error
	return
}

// GetForumThumbsUpInfoList 分页获取ForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) GetForumThumbsUpInfoList(info communityReq.ForumThumbsUpSearch) (list []community.ForumThumbsUp, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumThumbsUp{})
	var hkForumThumbsUps []community.ForumThumbsUp
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
