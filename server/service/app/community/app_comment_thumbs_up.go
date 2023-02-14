package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCommentThumbsUpService struct {
}

// CreateCommentThumbsUp 创建CommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) CreateCommentThumbsUp(hkCommentThumbsUp community.CommentThumbsUp) (err error) {
	err = global.GVA_DB.Create(&hkCommentThumbsUp).Error
	return err
}

// DeleteCommentThumbsUp 删除CommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) DeleteCommentThumbsUp(info communityReq.DeleteCommentThumbsUp) (err error) {
	err = global.GVA_DB.Where("user_id = ? and comment_id = ?", info.UserId, info.CommentId).Delete(&community.CommentThumbsUp{}).Error
	return err
}

// DeleteCommentThumbsUpByIds 批量删除CommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) DeleteCommentThumbsUpByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.CommentThumbsUp{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCommentThumbsUp 更新CommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) UpdateCommentThumbsUp(hkCommentThumbsUp community.CommentThumbsUp) (err error) {
	err = global.GVA_DB.Save(&hkCommentThumbsUp).Error
	return err
}

// GetCommentThumbsUp 根据id获取CommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) GetCommentThumbsUp(id uint) (hkCommentThumbsUp community.CommentThumbsUp, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCommentThumbsUp).Error
	return
}

// GetCommentThumbsUpInfoList 分页获取CommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) GetCommentThumbsUpInfoList(info communityReq.CommentThumbsUpSearch) (list []community.CommentThumbsUp, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CommentThumbsUp{})
	var hkCommentThumbsUps []community.CommentThumbsUp
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCommentThumbsUps).Error
	return hkCommentThumbsUps, total, err
}
