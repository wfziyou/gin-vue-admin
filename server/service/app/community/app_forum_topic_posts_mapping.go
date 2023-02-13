package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForumTopicPostsMappingService struct {
}

// CreateHkForumTopicPostsMapping 创建HkForumTopicPostsMapping记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTopicPostsMappingService *AppForumTopicPostsMappingService) CreateHkForumTopicPostsMapping(hkForumTopicPostsMapping community.ForumTopicPostsMapping) (err error) {
	err = global.GVA_DB.Create(&hkForumTopicPostsMapping).Error
	return err
}

// DeleteHkForumTopicPostsMapping 删除HkForumTopicPostsMapping记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTopicPostsMappingService *AppForumTopicPostsMappingService) DeleteHkForumTopicPostsMapping(hkForumTopicPostsMapping community.ForumTopicPostsMapping) (err error) {
	err = global.GVA_DB.Delete(&hkForumTopicPostsMapping).Error
	return err
}

// DeleteHkForumTopicPostsMappingByIds 批量删除HkForumTopicPostsMapping记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTopicPostsMappingService *AppForumTopicPostsMappingService) DeleteHkForumTopicPostsMappingByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumTopicPostsMapping{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkForumTopicPostsMapping 更新HkForumTopicPostsMapping记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTopicPostsMappingService *AppForumTopicPostsMappingService) UpdateHkForumTopicPostsMapping(hkForumTopicPostsMapping community.ForumTopicPostsMapping) (err error) {
	err = global.GVA_DB.Save(&hkForumTopicPostsMapping).Error
	return err
}

// GetHkForumTopicPostsMapping 根据id获取HkForumTopicPostsMapping记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTopicPostsMappingService *AppForumTopicPostsMappingService) GetHkForumTopicPostsMapping(id uint) (hkForumTopicPostsMapping community.ForumTopicPostsMapping, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumTopicPostsMapping).Error
	return
}

// GetHkForumTopicPostsMappingInfoList 分页获取HkForumTopicPostsMapping记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTopicPostsMappingService *AppForumTopicPostsMappingService) GetHkForumTopicPostsMappingInfoList(info communityReq.HkForumTopicPostsMappingSearch) (list []community.ForumTopicPostsMapping, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumTopicPostsMapping{})
	var hkForumTopicPostsMappings []community.ForumTopicPostsMapping
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumTopicPostsMappings).Error
	return hkForumTopicPostsMappings, total, err
}
