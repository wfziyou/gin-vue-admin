package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppUserCollectService struct {
}

// CreateUserCollect 创建UserCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCollectService *AppUserCollectService) CreateUserCollect(hkUserCollect general.UserCollect) (err error) {
	err = global.GVA_DB.Where("user_id = ? and posts_id = ?", hkUserCollect.UserId, hkUserCollect.PostsId).First(&hkUserCollect).Error
	if err == nil {
		err = global.GVA_DB.Save(&hkUserCollect).Error
		err = appUserCollectService.UpdateCollectNum(hkUserCollect.PostsId)
	} else {
		err = global.GVA_DB.Save(&hkUserCollect).Error
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

// DeleteUserCollect 删除UserCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCollectService *AppUserCollectService) DeleteUserCollect(hkUserCollect general.UserCollect) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&hkUserCollect).Error
	if err == nil {
		err = appUserCollectService.UpdateCollectNum(hkUserCollect.PostsId)
	}
	return err
}

// DeleteUserCollectByIds 批量删除UserCollect记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserCollectService *AppUserCollectService) DeleteUserCollectByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]general.UserCollect{}, "id in ?", ids.Ids).Error
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
func (appUserCollectService *AppUserCollectService) GetUserCollectInfoList(info generalReq.UserCollectSearch) (list []general.UserCollect, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&general.UserCollect{})
	var hkUserCollects []general.UserCollect
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	db = db.Where("user_id = ?", info.UserId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUserCollects).Error
	return hkUserCollects, total, err
}
