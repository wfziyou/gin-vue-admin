package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type HkForumPostsGroupService struct {
}

// CreateHkForumPostsGroup 创建HkForumPostsGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumPostsGroupService *HkForumPostsGroupService) CreateHkForumPostsGroup(hkForumPostsGroup community.HkForumPostsGroup) (err error) {
	err = global.GVA_DB.Create(&hkForumPostsGroup).Error
	return err
}

// DeleteHkForumPostsGroup 删除HkForumPostsGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumPostsGroupService *HkForumPostsGroupService)DeleteHkForumPostsGroup(hkForumPostsGroup community.HkForumPostsGroup) (err error) {
	err = global.GVA_DB.Delete(&hkForumPostsGroup).Error
	return err
}

// DeleteHkForumPostsGroupByIds 批量删除HkForumPostsGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumPostsGroupService *HkForumPostsGroupService)DeleteHkForumPostsGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForumPostsGroup{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkForumPostsGroup 更新HkForumPostsGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumPostsGroupService *HkForumPostsGroupService)UpdateHkForumPostsGroup(hkForumPostsGroup community.HkForumPostsGroup) (err error) {
	err = global.GVA_DB.Save(&hkForumPostsGroup).Error
	return err
}

// GetHkForumPostsGroup 根据id获取HkForumPostsGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumPostsGroupService *HkForumPostsGroupService)GetHkForumPostsGroup(id uint) (hkForumPostsGroup community.HkForumPostsGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumPostsGroup).Error
	return
}

// GetHkForumPostsGroupInfoList 分页获取HkForumPostsGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumPostsGroupService *HkForumPostsGroupService)GetHkForumPostsGroupInfoList(info communityReq.HkForumPostsGroupSearch) (list []community.HkForumPostsGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&community.HkForumPostsGroup{})
    var hkForumPostsGroups []community.HkForumPostsGroup
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostsGroups).Error
	return  hkForumPostsGroups, total, err
}
