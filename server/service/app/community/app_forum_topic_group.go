package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForumTopicGroupService struct {
}

// CreateForumTopicGroup 创建话题Group记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) CreateForumTopicGroup(hkForumTopicGroup community.ForumTopicGroup) (err error) {
	err = global.GVA_DB.Create(&hkForumTopicGroup).Error
	return err
}

// DeleteForumTopicGroup 删除话题Group记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) DeleteForumTopicGroup(hkForumTopicGroup community.ForumTopicGroup) (err error) {
	err = global.GVA_DB.Delete(&hkForumTopicGroup).Error
	return err
}

// DeleteForumTopicGroupByIds 批量删除话题Group记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) DeleteForumTopicGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumTopicGroup{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateForumTopicGroup 更新话题Group记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) UpdateForumTopicGroup(hkForumTopicGroup community.ForumTopicGroup) (err error) {
	err = global.GVA_DB.Save(&hkForumTopicGroup).Error
	return err
}

// GetForumTopicGroup 根据id获取ForumTopicGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) GetForumTopicGroup(id uint64) (hkForumTopicGroup community.ForumTopicGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumTopicGroup).Error
	return
}

// GetForumTopicGroupInfoList 分页获取ForumTopicGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) GetForumTopicGroupInfoList(info communityReq.ForumTopicGroupSearch) (list []community.ForumTopicGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumTopicGroup{})
	var hkForumTopicGroups []community.ForumTopicGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if len(info.Name) > 0 {
		db = db.Where("name LIKE %", "%"+info.Name+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumTopicGroups).Error
	return hkForumTopicGroups, total, err
}

// GetForumTopicGroupInfoListAll 分页获取ForumTopicGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) GetForumTopicGroupInfoListAll(info communityReq.ForumTopicGroupSearch) (list []community.ForumTopicGroup, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.ForumTopicGroup{})
	var hkForumTopicGroups []community.ForumTopicGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if len(info.Name) > 0 {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Find(&hkForumTopicGroups).Error
	return hkForumTopicGroups, total, err
}
