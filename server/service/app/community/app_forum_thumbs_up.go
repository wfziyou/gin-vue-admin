package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForumThumbsUpService struct {
}

// CreateForumThumbsUp 创建ForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) CreateForumThumbsUp(hkForumThumbsUp community.ForumThumbsUp) (err error) {
	err = global.GVA_DB.Create(&hkForumThumbsUp).Error
	if err == nil {
		err = appForumThumbsUpService.UpdateForumPostsLikeNum(hkForumThumbsUp.PostsId)
	}
	return err
}
func (appForumThumbsUpService *AppForumThumbsUpService) UpdateForumPostsLikeNum(postsId uint64) (err error) {
	var total int64 = 0
	db := global.GVA_DB.Model(&community.ForumThumbsUp{}).Where("posts_id = ?", postsId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(&community.ForumPosts{}).Where("id = ?", postsId).Update("like_num", total).Error
	return err
}

// DeleteForumThumbsUp 删除ForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) DeleteForumThumbsUp(info communityReq.DeleteForumThumbsUp) (err error) {
	err = global.GVA_DB.Unscoped().Where("user_id = ? and posts_id = ?", info.UserId, info.PostsId).Delete(&community.ForumThumbsUp{}).Error
	if err == nil {
		err = appForumThumbsUpService.UpdateForumPostsLikeNum(info.PostsId)
	}
	return err
}

// DeleteForumThumbsUpByIds 批量删除ForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) DeleteForumThumbsUpByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]community.ForumThumbsUp{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateForumThumbsUp 更新ForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) UpdateForumThumbsUp(hkForumThumbsUp community.ForumThumbsUp) (err error) {
	err = global.GVA_DB.Save(&hkForumThumbsUp).Error
	return err
}

// GetForumThumbsUp 根据id获取ForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) GetForumThumbsUp(id uint) (hkForumThumbsUp community.ForumThumbsUp, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumThumbsUp).Error
	return
}

func (appForumThumbsUpService *AppForumThumbsUpService) GetForumThumbsUpEx(userId uint64, postsIds []uint64) (hkForumThumbsUp []community.ForumThumbsUp, num int, err error) {
	err = global.GVA_DB.Where("user_id = ? and posts_id in  ?", userId, postsIds).Find(&hkForumThumbsUp).Error
	if err == nil {
		num = len(hkForumThumbsUp)
	}
	return
}
func (appForumThumbsUpService *AppForumThumbsUpService) GetUserForumThumbsUp(userId uint64, list []community.ForumPostsBaseInfo) (err error) {
	var size = len(list)
	if size > 0 {

		var ids = make([]uint64, size)
		for index, v := range list {
			ids[index] = v.ID
		}
		if data, num, err := appForumThumbsUpService.GetForumThumbsUpEx(userId, ids); err == nil && num > 0 {
			for _, v := range data {
				for i, forum := range list {
					if forum.ID == v.PostsId {
						list[i].ThumbsUp = 1
						break
					}
				}
			}
		}
	}
	return
}

// GetForumThumbsUpInfoList 分页获取ForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumThumbsUpService *AppForumThumbsUpService) GetForumThumbsUpInfoList(info communityReq.ForumThumbsUpSearch) (list []community.ForumThumbsUp, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumThumbsUp{})
	var hkForumThumbsUps []community.ForumThumbsUp
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumThumbsUps).Error
	return hkForumThumbsUps, total, err
}
