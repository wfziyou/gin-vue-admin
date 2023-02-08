package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForumTagGroupService struct {
}

// CreateForumTagGroup 创建ForumTagGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagGroupService *AppForumTagGroupService) CreateForumTagGroup(hkForumTagGroup community.ForumTagGroup) (err error) {
	err = global.GVA_DB.Create(&hkForumTagGroup).Error
	return err
}

// DeleteForumTagGroup 删除ForumTagGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagGroupService *AppForumTagGroupService) DeleteForumTagGroup(hkForumTagGroup community.ForumTagGroup) (err error) {
	err = global.GVA_DB.Delete(&hkForumTagGroup).Error
	return err
}

// DeleteForumTagGroupByIds 批量删除ForumTagGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagGroupService *AppForumTagGroupService) DeleteForumTagGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumTagGroup{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateForumTagGroup 更新ForumTagGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagGroupService *AppForumTagGroupService) UpdateForumTagGroup(hkForumTagGroup community.ForumTagGroup) (err error) {
	err = global.GVA_DB.Save(&hkForumTagGroup).Error
	return err
}

// GetForumTagGroup 根据id获取ForumTagGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagGroupService *AppForumTagGroupService) GetForumTagGroup(id uint) (hkForumTagGroup community.ForumTagGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumTagGroup).Error
	return
}

// GetForumTagGroupInfoList 分页获取ForumTagGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagGroupService *AppForumTagGroupService) GetForumTagGroupInfoList(info communityReq.ForumTagGroupSearch) (list []community.ForumTagGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumTagGroup{})
	var hkForumTagGroups []community.ForumTagGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumTagGroups).Error
	return hkForumTagGroups, total, err
}
