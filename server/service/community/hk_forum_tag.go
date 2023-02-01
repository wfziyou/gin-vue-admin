package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type HkForumTagService struct {
}

// CreateHkForumTag 创建HkForumTag记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTagService *HkForumTagService) CreateHkForumTag(hkForumTag community.HkForumTag) (err error) {
	err = global.GVA_DB.Create(&hkForumTag).Error
	return err
}

// DeleteHkForumTag 删除HkForumTag记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTagService *HkForumTagService)DeleteHkForumTag(hkForumTag community.HkForumTag) (err error) {
	err = global.GVA_DB.Delete(&hkForumTag).Error
	return err
}

// DeleteHkForumTagByIds 批量删除HkForumTag记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTagService *HkForumTagService)DeleteHkForumTagByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForumTag{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkForumTag 更新HkForumTag记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTagService *HkForumTagService)UpdateHkForumTag(hkForumTag community.HkForumTag) (err error) {
	err = global.GVA_DB.Save(&hkForumTag).Error
	return err
}

// GetHkForumTag 根据id获取HkForumTag记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTagService *HkForumTagService)GetHkForumTag(id uint) (hkForumTag community.HkForumTag, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumTag).Error
	return
}

// GetHkForumTagInfoList 分页获取HkForumTag记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTagService *HkForumTagService)GetHkForumTagInfoList(info communityReq.HkForumTagSearch) (list []community.HkForumTag, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&community.HkForumTag{})
    var hkForumTags []community.HkForumTag
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkForumTags).Error
	return  hkForumTags, total, err
}
