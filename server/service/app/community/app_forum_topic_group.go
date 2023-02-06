package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type AppForumTopicGroupService struct {
}

// CreateHkForumTopicGroup 创建HkForumTopicGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) CreateHkForumTopicGroup(hkForumTopicGroup community.HkForumTopicGroup) (err error) {
	err = global.GVA_DB.Create(&hkForumTopicGroup).Error
	return err
}

// DeleteHkForumTopicGroup 删除HkForumTopicGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) DeleteHkForumTopicGroup(hkForumTopicGroup community.HkForumTopicGroup) (err error) {
	err = global.GVA_DB.Delete(&hkForumTopicGroup).Error
	return err
}

// DeleteHkForumTopicGroupByIds 批量删除HkForumTopicGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) DeleteHkForumTopicGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForumTopicGroup{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkForumTopicGroup 更新HkForumTopicGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) UpdateHkForumTopicGroup(hkForumTopicGroup community.HkForumTopicGroup) (err error) {
	err = global.GVA_DB.Save(&hkForumTopicGroup).Error
	return err
}

// GetHkForumTopicGroup 根据id获取HkForumTopicGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) GetHkForumTopicGroup(id uint) (hkForumTopicGroup community.HkForumTopicGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumTopicGroup).Error
	return
}

// GetHkForumTopicGroupInfoList 分页获取HkForumTopicGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicGroupService *AppForumTopicGroupService) GetHkForumTopicGroupInfoList(info communityReq.HkForumTopicGroupSearch) (list []community.HkForumTopicGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkForumTopicGroup{})
	var hkForumTopicGroups []community.HkForumTopicGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumTopicGroups).Error
	return hkForumTopicGroups, total, err
}
