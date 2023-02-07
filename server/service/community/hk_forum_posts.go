package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type HkForumPostsService struct {
}

// CreateHkForumPosts 创建HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumPostsService *HkForumPostsService) CreateHkForumPosts(hkForumPosts community.HkForumPosts) (err error) {
	err = global.GVA_DB.Create(&hkForumPosts).Error
	return err
}

// DeleteHkForumPosts 删除HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumPostsService *HkForumPostsService) DeleteHkForumPosts(hkForumPosts community.HkForumPosts) (err error) {
	err = global.GVA_DB.Delete(&hkForumPosts).Error
	return err
}

// DeleteHkForumPostsByIds 批量删除HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumPostsService *HkForumPostsService) DeleteHkForumPostsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForumPosts{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkForumPosts 更新HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumPostsService *HkForumPostsService) UpdateHkForumPosts(hkForumPosts community.HkForumPosts) (err error) {
	err = global.GVA_DB.Save(&hkForumPosts).Error
	return err
}

// GetHkForumPosts 根据id获取HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumPostsService *HkForumPostsService) GetHkForumPosts(id uint) (hkForumPosts community.HkForumPosts, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumPosts).Error
	return
}

// GetHkForumPostsInfoList 分页获取HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumPostsService *HkForumPostsService) GetHkForumPostsInfoList(info communityReq.HkForumPostsSearch) (list []community.HkForumPosts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkForumPosts{})
	var hkForumPostss []community.HkForumPosts
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}
