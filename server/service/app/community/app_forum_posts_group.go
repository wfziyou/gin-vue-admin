package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForumPostsGroupService struct {
}

// CreateForumPostsGroup 创建ForumPostsGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsGroupService *AppForumPostsGroupService) CreateForumPostsGroup(hkForumPostsGroup community.ForumPostsGroup) (err error) {
	err = global.GVA_DB.Create(&hkForumPostsGroup).Error
	return err
}

// DeleteForumPostsGroup 删除ForumPostsGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsGroupService *AppForumPostsGroupService) DeleteForumPostsGroup(hkForumPostsGroup community.ForumPostsGroup) (err error) {
	err = global.GVA_DB.Delete(&hkForumPostsGroup).Error
	return err
}

// DeleteForumPostsGroupByIds 批量删除ForumPostsGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsGroupService *AppForumPostsGroupService) DeleteForumPostsGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumPostsGroup{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateForumPostsGroup 更新ForumPostsGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsGroupService *AppForumPostsGroupService) UpdateForumPostsGroup(hkForumPostsGroup community.ForumPostsGroup) (err error) {
	err = global.GVA_DB.Save(&hkForumPostsGroup).Error
	return err
}

// GetForumPostsGroup 根据id获取ForumPostsGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsGroupService *AppForumPostsGroupService) GetForumPostsGroup(id uint) (hkForumPostsGroup community.ForumPostsGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumPostsGroup).Error
	return
}

// GetForumPostsGroupInfoList 分页获取ForumPostsGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsGroupService *AppForumPostsGroupService) GetForumPostsGroupInfoList(info communityReq.ForumPostsGroupSearch) (list []community.ForumPostsGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsGroup{})
	var hkForumPostsGroups []community.ForumPostsGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostsGroups).Error
	return hkForumPostsGroups, total, err
}
