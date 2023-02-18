package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForumCommentService struct {
}

// CreateForumComment 创建ForumComment记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumCommentService *AppForumCommentService) CreateForumComment(info communityReq.CreateForumComment) (err error) {

	err = global.GVA_DB.Create(&community.ForumComment{
		UserId:  info.UserId,
		PostsId: info.PostsId,
	}).Error
	return err
}

// DeleteForumComment 删除ForumComment记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumCommentService *AppForumCommentService) DeleteForumComment(hkForumComment community.ForumComment) (err error) {
	err = global.GVA_DB.Delete(&hkForumComment).Error
	return err
}

// DeleteForumCommentByIds 批量删除ForumComment记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumCommentService *AppForumCommentService) DeleteForumCommentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumComment{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateForumComment 更新ForumComment记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumCommentService *AppForumCommentService) UpdateForumComment(hkForumComment community.ForumComment) (err error) {
	err = global.GVA_DB.Save(&hkForumComment).Error
	return err
}

// GetForumComment 根据id获取ForumComment记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumCommentService *AppForumCommentService) GetForumComment(id uint64) (hkForumComment community.ForumComment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumComment).Error
	return hkForumComment, err
}

// GetForumCommentInfoList 分页获取ForumComment记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumCommentService *AppForumCommentService) GetForumCommentInfoList(info communityReq.ForumCommentSearch) (list []community.ForumComment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumComment{})
	var hkForumComments []community.ForumComment
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	db = db.Where("posts_id = ?", info.PostsId)
	if info.ParentId > 0 {
		db = db.Where("parent_id = ?", info.ParentId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumComments).Error
	return hkForumComments, total, err
}
