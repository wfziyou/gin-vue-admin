package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type AppForumCommentService struct {
}

// CreateHkForumComment 创建HkForumComment记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumCommentService *AppForumCommentService) CreateHkForumComment(hkForumComment community.HkForumComment) (err error) {
	err = global.GVA_DB.Create(&hkForumComment).Error
	return err
}

// DeleteHkForumComment 删除HkForumComment记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumCommentService *AppForumCommentService) DeleteHkForumComment(hkForumComment community.HkForumComment) (err error) {
	err = global.GVA_DB.Delete(&hkForumComment).Error
	return err
}

// DeleteHkForumCommentByIds 批量删除HkForumComment记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumCommentService *AppForumCommentService) DeleteHkForumCommentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForumComment{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkForumComment 更新HkForumComment记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumCommentService *AppForumCommentService) UpdateHkForumComment(hkForumComment community.HkForumComment) (err error) {
	err = global.GVA_DB.Save(&hkForumComment).Error
	return err
}

// GetHkForumComment 根据id获取HkForumComment记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumCommentService *AppForumCommentService) GetHkForumComment(id uint) (hkForumComment community.HkForumComment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumComment).Error
	return
}

// GetHkForumCommentInfoList 分页获取HkForumComment记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumCommentService *AppForumCommentService) GetHkForumCommentInfoList(info communityReq.HkForumCommentSearch) (list []community.HkForumComment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkForumComment{})
	var hkForumComments []community.HkForumComment
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumComments).Error
	return hkForumComments, total, err
}
