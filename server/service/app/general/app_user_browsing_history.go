package general

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type AppUserBrowsingHistoryService struct {
}

// CreateUserBrowsingHistory 创建UserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) CreateUserBrowsingHistory(userId uint64, postsId uint64, category int) (err error) {
	var obj = general.UserBrowsingHistory{UserId: userId, PostsId: postsId, Category: category}
	err = global.GVA_DB.Where("user_id = ? and posts_id = ?", userId, postsId).First(&obj).Error
	if err == nil {
		err = global.GVA_DB.Save(&obj).Error
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&obj).Error
	}
	return err
}

// DeleteUserBrowsingHistory 删除浏览历史记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) DeleteUserBrowsingHistory(hkUserBrowsingHistory general.UserBrowsingHistory) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&hkUserBrowsingHistory).Error
	return err
}
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) DeleteUserBrowsingHistoryById(userId uint64, id uint64) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&general.UserBrowsingHistory{}, "user_id = ? AND id in =", userId, id).Error
	return err
}

// DeleteUserBrowsingHistoryByIds 批量删除浏览历史记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) DeleteUserBrowsingHistoryByIds(userId uint64, ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]general.UserBrowsingHistory{}, "user_id = ? AND id in ?", userId, ids.Ids).Error
	return err
}
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) DeletAlleUserBrowsingHistory(userId uint64, category int) (err error) {
	if category > 0 {
		err = global.GVA_DB.Unscoped().Delete(&[]general.UserBrowsingHistory{}, "user_id = ? AND category = ?", userId, category).Error
	} else {
		err = global.GVA_DB.Unscoped().Delete(&[]general.UserBrowsingHistory{}, "user_id = ?", userId).Error

	}
	return err
}

// UpdateUserBrowsingHistory 更新UserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) UpdateUserBrowsingHistory(hkUserBrowsingHistory general.UserBrowsingHistory) (err error) {
	err = global.GVA_DB.Save(&hkUserBrowsingHistory).Error
	return err
}

// GetUserBrowsingHistory 根据id获取UserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) GetUserBrowsingHistory(id uint) (hkUserBrowsingHistory general.UserBrowsingHistory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserBrowsingHistory).Error
	return
}

// GetUserBrowsingHistoryInfoList 分页获取UserBrowsingHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserBrowsingHistoryService *AppUserBrowsingHistoryService) GetUserBrowsingHistoryInfoList(userId uint64, info generalReq.UserBrowsingHistorySearch) (list []general.UserBrowsingHistoryInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&general.UserBrowsingHistoryInfo{})
	var hkUserBrowsingHistorys []general.UserBrowsingHistoryInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartUpdatedAt != nil && info.EndUpdatedAt != nil {
		db = db.Where("updated_at BETWEEN ? AND ?", info.StartUpdatedAt, info.EndUpdatedAt)
	}
	if userId > 0 {
		db = db.Where("user_id = ?", userId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUserBrowsingHistorys).Error
	if err == nil {
		var size = len(hkUserBrowsingHistorys)
		if size > 0 {
			var ids = make([]uint64, size)
			for index, v := range hkUserBrowsingHistorys {
				ids[index] = v.PostsId
			}
			//查询最新发布的帖子
			db1 := global.GVA_DB.Model(&community.ForumPostsBaseInfo{})
			var hkForumPosts []community.ForumPostsBaseInfo
			db1 = db1.Where("id in ?", ids)
			err = db1.Find(&hkForumPosts).Error
			if err == nil {
				for x, obj := range hkUserBrowsingHistorys {
					for _, posts := range hkForumPosts {
						if obj.PostsId == posts.ID {
							hkUserBrowsingHistorys[x].ID = posts.ID
							hkUserBrowsingHistorys[x].Title = posts.Title
							hkUserBrowsingHistorys[x].Category = posts.Category
							hkUserBrowsingHistorys[x].Attachment = posts.Attachment
							hkUserBrowsingHistorys[x].CoverImage = posts.CoverImage
							break
						}
					}
				}
			}
		}
	}
	return hkUserBrowsingHistorys, total, err
}
