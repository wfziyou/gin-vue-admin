package general

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"gorm.io/gorm"
)

type AppUserCollectService struct {
}

// CreateUserCollect 创建UserCollect记录
func (appUserCollectService *AppUserCollectService) CreateUserCollect(userId uint64, posts community.ForumPosts) (err error) {
	var hkUserCollect = general.UserCollect{UserId: userId, PostsId: posts.ID, Category: posts.Category}
	err = global.GVA_DB.Where("user_id = ? and posts_id = ?", userId, posts.ID).First(&hkUserCollect).Error
	if err == nil {
		err = global.GVA_DB.Save(&hkUserCollect).Error
		err = appUserCollectService.UpdateCollectNum(hkUserCollect.PostsId)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&hkUserCollect).Error
	}
	return err
}
func (appUserCollectService *AppUserCollectService) UpdateCollectNum(postsIdd uint64) (err error) {
	var total int64 = 0
	db := global.GVA_DB.Model(&general.UserCollect{}).Where("posts_id = ?", postsIdd)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(community.ForumPosts{}).Where("id = ?", postsIdd).Update("collect_num", total).Error
	return err
}

// DeleteUserCollect 删除收藏记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCollectService *AppUserCollectService) DeleteUserCollect(userId uint64, postsId uint64) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&general.UserCollect{}, "user_id = ? AND posts_id = ?", userId, postsId).Error
	if err == nil {
		err = appUserCollectService.UpdateCollectNum(postsId)
	}
	return err
}

// DeleteUserCollectByIds 批量删除收藏记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCollectService *AppUserCollectService) DeleteUserCollectByIds(userId uint64, ids []uint64) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]general.UserCollect{}, "user_id = ? AND posts_id in ?", userId, ids).Error
	if err == nil {
		for _, postsId := range ids {
			err = appUserCollectService.UpdateCollectNum(postsId)
		}
	}
	return err
}

// UpdateUserCollect 更新UserCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCollectService *AppUserCollectService) UpdateUserCollect(hkUserCollect general.UserCollect) (err error) {
	err = global.GVA_DB.Save(&hkUserCollect).Error
	return err
}

// GetUserCollect 根据id获取UserCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCollectService *AppUserCollectService) GetUserCollect(id uint) (hkUserCollect general.UserCollect, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserCollect).Error
	return
}
func (appUserCollectService *AppUserCollectService) GetUserCollectEx(userId uint64, postsIds []uint64) (hkUserCollect []general.UserCollect, num int, err error) {
	err = global.GVA_DB.Where("user_id = ? and posts_id in  ?", userId, postsIds).Find(&hkUserCollect).Error
	if err == nil {
		num = len(hkUserCollect)
	}
	return
}
func (appUserCollectService *AppUserCollectService) GetUserIsCollect(userId uint64, list []community.ForumPostsBaseInfo) (err error) {
	var size = len(list)
	if size > 0 {

		var ids = make([]uint64, size)
		for index, v := range list {
			ids[index] = v.ID
		}
		if data, num, err := appUserCollectService.GetUserCollectEx(userId, ids); err == nil && num > 0 {
			for _, v := range data {
				for i, forum := range list {
					if forum.ID == v.PostsId {
						list[i].Collect = 1
						break
					}
				}
			}
		}
	}
	return
}

// GetUserCollectInfoList 分页获取UserCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCollectService *AppUserCollectService) GetUserCollectInfoList(info generalReq.UserCollectSearch) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&general.UserCollect{})
	var hkUserCollects []general.UserCollect

	db = db.Where("user_id = ?", info.UserId)
	if info.Category > 0 {
		db = db.Where("category = ?", info.Category)
	}
	//创建时间降序排列
	db = db.Order("created_at desc")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUserCollects).Error
	if err == nil {
		var size = len(hkUserCollects)
		if size > 0 {
			var ids = make([]uint64, size)
			for index, v := range hkUserCollects {
				ids[index] = v.PostsId
			}
			//查询最新发布的帖子
			db1 := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
			var hkForumPosts []community.ForumPostsBaseInfo
			db1 = db1.Where("id in ?", ids)
			db1 = db1.Where("check_status = ?",
				community.PostsCheckStatusPass)
			err = db1.Find(&hkForumPosts).Error
			return hkForumPosts, total, err
		}
	}
	return []community.ForumPostsBaseInfo{}, total, err
}
