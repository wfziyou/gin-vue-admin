package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForumPostsService struct {
}

// CreateHkForumPosts 创建HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) CreateHkForumPosts(hkForumPosts community.HkForumPosts) (err error) {
	err = global.GVA_DB.Create(&hkForumPosts).Error
	return err
}

// DeleteHkForumPosts 删除HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) DeleteHkForumPosts(hkForumPosts community.HkForumPosts) (err error) {
	err = global.GVA_DB.Delete(&hkForumPosts).Error
	return err
}

// DeleteHkForumPostsByIds 批量删除HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) DeleteHkForumPostsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForumPosts{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkForumPosts 更新HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) UpdateHkForumPosts(hkForumPosts community.HkForumPosts) (err error) {
	err = global.GVA_DB.Save(&hkForumPosts).Error
	return err
}

// GetHkForumPosts 根据id获取HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) GetHkForumPosts(id uint) (hkForumPosts community.HkForumPosts, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumPosts).Error
	return
}

// GetHkForumPostsInfoList 分页获取HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) GetHkForumPostsInfoList(info communityReq.HkForumPostsSearch) (list []community.HkForumPosts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkForumPosts{})
	var hkForumPostss []community.HkForumPosts
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}

// GetForumPostsInfoList 分页获取HkForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) GetForumPostsInfoList(info communityReq.CircleForumPostsSearch) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{})
	var hkForumPostss []community.ForumPostsBaseInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	//圈子_编号
	if info.CircleId != 0 {
		db = db.Where("circleId = ?", info.CircleId)
	}

	//类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）
	if info.Category != nil {
		db = db.Where("category = ?", info.Category)
	}

	//帖子分类编号
	if info.GroupId != nil {
		db = db.Where("groupId = ?", info.GroupId)
	}
	//标题
	if len(info.Title) != 0 {
		db = db.Where("title LIKE '%?%'", info.Title)
	}
	//置顶(0否 1是)
	if info.Top != nil {
		db = db.Where("top = ?", info.Top)
	}
	//精华(0否 1是)
	if info.Marrow != nil {
		db = db.Where("marrow = ?", info.Marrow)
	}
	//创建时间降序排列
	db = db.Order("create_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}

// GetUserForumPostsInfoList 分页获取用户加入圈子的所有动态列表
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) GetUserForumPostsInfoList(info communityReq.UserCircleForumPostsSearch) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{})
	var hkForumPostss []community.ForumPostsBaseInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	//用户所在圈子
	db.Where("circleId IN ?", info.ID)

	//类别（0视频、1动态、2资讯、3公告、4文章、5问答、6建议）
	if info.Category != nil {
		db = db.Where("category = ?", info.Category)
	}

	//帖子分类编号
	if info.GroupId != nil {
		db = db.Where("groupId = ?", info.GroupId)
	}
	//标题
	if len(info.Title) != 0 {
		db = db.Where("title LIKE '%?%'", info.Title)
	}
	//置顶(0否 1是)
	if info.Top != nil {
		db = db.Where("top = ?", info.Top)
	}
	//精华(0否 1是)
	if info.Marrow != nil {
		db = db.Where("marrow = ?", info.Marrow)
	}
	//创建时间降序排列
	db = db.Order("create_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}
