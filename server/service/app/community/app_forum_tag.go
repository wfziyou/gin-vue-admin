package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForumTagService struct {
}

// CreateForumTag 创建ForumTag记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagService *AppForumTagService) CreateForumTag(hkForumTag community.ForumTag) (err error) {
	err = global.GVA_DB.Create(&hkForumTag).Error
	return err
}

// DeleteForumTag 删除ForumTag记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagService *AppForumTagService) DeleteForumTag(hkForumTag community.ForumTag) (err error) {
	err = global.GVA_DB.Delete(&hkForumTag).Error
	return err
}

// DeleteForumTagByIds 批量删除ForumTag记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagService *AppForumTagService) DeleteForumTagByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumTag{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateForumTag 更新ForumTag记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagService *AppForumTagService) UpdateForumTag(hkForumTag community.ForumTag) (err error) {
	err = global.GVA_DB.Save(&hkForumTag).Error
	return err
}

// GetForumTag 根据id获取ForumTag记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagService *AppForumTagService) GetForumTag(id uint) (hkForumTag community.ForumTag, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumTag).Error
	return
}

// GetForumTagInfoList 分页获取ForumTag记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTagService *AppForumTagService) GetForumTagInfoList(info communityReq.ForumTagSearch) (list []community.ForumTag, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumTag{})
	var hkForumTags []community.ForumTag
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumTags).Error
	return hkForumTags, total, err
}
