package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
)

type AppCommentThumbsUpService struct {
}

// CreateCommentThumbsUp 创建CommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) CreateCommentThumbsUp(hkCommentThumbsUp community.CommentThumbsUp) (err error) {
	err = global.GVA_DB.Create(&hkCommentThumbsUp).Error
	if err == nil {
		err = appCommentThumbsUpService.UpdateCommentLikeNum(hkCommentThumbsUp.CommentId)
	}
	return err
}
func (appCommentThumbsUpService *AppCommentThumbsUpService) UpdateCommentLikeNum(commentId uint64) (err error) {
	var total int64 = 0
	db := global.GVA_DB.Model(&community.CommentThumbsUp{}).Where("comment_id = ?", commentId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(community.ForumComment{}).Where("id = ?", commentId).Update("like_num", total).Error
	return err
}

// DeleteCommentThumbsUp 删除CommentThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCommentThumbsUpService *AppCommentThumbsUpService) DeleteCommentThumbsUp(info communityReq.DeleteCommentThumbsUp) (err error) {
	err = global.GVA_DB.Where("user_id = ? and comment_id = ?", info.UserId, info.CommentId).Delete(&community.CommentThumbsUp{}).Error
	if err == nil {
		err = appCommentThumbsUpService.UpdateCommentLikeNum(info.CommentId)
	}
	return err
}

func (appCommentThumbsUpService *AppCommentThumbsUpService) GetCommentThumbsUpEx(userId uint64, postsIds []uint64) (hkForumThumbsUp []community.CommentThumbsUp, num int, err error) {
	err = global.GVA_DB.Where("user_id = ? and comment_id in  ?", userId, postsIds).Find(&hkForumThumbsUp).Error
	if err == nil {
		num = len(hkForumThumbsUp)
	}
	return
}
func (appCommentThumbsUpService *AppCommentThumbsUpService) GetUserCommentThumbsUp(userId uint64, list []community.ForumComment) (err error) {
	var size = len(list)
	if size > 0 {
		var ids = make([]uint64, size)
		for index, v := range list {
			ids[index] = v.ID
		}
		if data, num, err := appCommentThumbsUpService.GetCommentThumbsUpEx(userId, ids); err == nil && num > 0 {
			for _, v := range data {
				for i, comment := range list {
					if comment.ID == v.CommentId {
						list[i].ThumbsUp = 1
						break
					}
				}
			}
		}
	}
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
