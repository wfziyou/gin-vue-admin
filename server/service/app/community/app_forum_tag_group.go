package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type AppForumTagGroupService struct {
}

// CreateHkForumTagGroup 创建HkForumTagGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagGroupService *AppForumTagGroupService) CreateHkForumTagGroup(hkForumTagGroup community.HkForumTagGroup) (err error) {
	err = global.GVA_DB.Create(&hkForumTagGroup).Error
	return err
}

// DeleteHkForumTagGroup 删除HkForumTagGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagGroupService *AppForumTagGroupService) DeleteHkForumTagGroup(hkForumTagGroup community.HkForumTagGroup) (err error) {
	err = global.GVA_DB.Delete(&hkForumTagGroup).Error
	return err
}

// DeleteHkForumTagGroupByIds 批量删除HkForumTagGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagGroupService *AppForumTagGroupService) DeleteHkForumTagGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForumTagGroup{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkForumTagGroup 更新HkForumTagGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagGroupService *AppForumTagGroupService) UpdateHkForumTagGroup(hkForumTagGroup community.HkForumTagGroup) (err error) {
	err = global.GVA_DB.Save(&hkForumTagGroup).Error
	return err
}

// GetHkForumTagGroup 根据id获取HkForumTagGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagGroupService *AppForumTagGroupService) GetHkForumTagGroup(id uint) (hkForumTagGroup community.HkForumTagGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumTagGroup).Error
	return
}

// GetHkForumTagGroupInfoList 分页获取HkForumTagGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagGroupService *AppForumTagGroupService) GetHkForumTagGroupInfoList(info communityReq.HkForumTagGroupSearch) (list []community.HkForumTagGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkForumTagGroup{})
	var hkForumTagGroups []community.HkForumTagGroup
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
