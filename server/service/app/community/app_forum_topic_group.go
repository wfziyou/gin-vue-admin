package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
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
func (appForumTopicGroupService *AppForumTopicGroupService) GetForumTopicGroupInfoList(info request.PageInfo) (list []community.ForumTopicGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumTopicGroup{})
	var hkForumTopicGroups []community.ForumTopicGroup
	if len(info.Keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+info.Keyword+"%")
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
func (appForumTopicGroupService *AppForumTopicGroupService) GetForumTopicGroupInfoListAll(keyword string) (list []community.ForumTopicGroup, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.ForumTopicGroup{})
	var hkForumTopicGroups []community.ForumTopicGroup

	if len(keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Find(&hkForumTopicGroups).Error
	return hkForumTopicGroups, total, err
}
