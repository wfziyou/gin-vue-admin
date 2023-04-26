package community

import (
	"database/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AppForumPostsService struct {
}

// CreateForumPosts 创建帖子记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) CreateForumPosts(info communityReq.CreateForumPostsReq) (err error) {
	forumPosts := community.ForumPosts{
		UserId:          info.UserId,
		CircleId:        uint64(info.CircleId),
		Category:        info.Category,
		Title:           info.Title,
		CoverImage:      info.CoverImage,
		ContentType:     info.ContentType,
		ContentMarkdown: info.ContentMarkdown,
		ContentHtml:     info.ContentHtml,
		Video:           info.Video,
		Attachment:      info.Attachment,
		Anonymity:       info.Anonymity,
	}
	err = global.GVA_DB.Create(&forumPosts).Error
	if err == nil && len(info.TopicId) > 0 {
		tmp := utils.SplitToUint64List(info.TopicId, ",")
		for _, topicId := range tmp {
			err = global.GVA_DB.Create(&community.ForumTopicPostsMapping{
				TopicId: topicId,
				PostsId: forumPosts.ID,
			}).Error
		}
	}
	return err
}

// DeleteForumPosts 删除帖子记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) DeleteForumPosts(info community.ForumPosts) (err error) {
	err = global.GVA_DB.Delete(&info).Error
	return err
}

// DeleteForumPostsByIds 批量删除帖子记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) DeleteForumPostsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumPosts{}, "id in ?", ids.Ids).Error
	return err
}
func (appForumPostsService *AppForumPostsService) DeleteCircleForumPostsByIds(circleId uint64, ids []uint64) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumPosts{}, "circle_id = ? AND id in ?", circleId, ids).Error
	return err
}

// UpdateForumPosts 更新ForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) UpdateForumPosts(hkForumPosts community.ForumPosts) (err error) {
	err = global.GVA_DB.Save(&hkForumPosts).Error
	return err
}

// GetForumPosts 根据id获取ForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) GetForumPosts(id uint64) (hkForumPosts community.ForumPosts, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("UserInfo").First(&hkForumPosts).Error
	return
}

// GetRecommendPostsList 分页获取推荐帖子列表
func (appForumPostsService *AppForumPostsService) GetRecommendPostsList(channelId uint64, page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var hkForumPostss []community.ForumPostsBaseInfo

	db = db.Where("channel_id = ?", channelId)

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}

// GetGlobalTopInfoList 分页获全局置顶资讯列表
func (appForumPostsService *AppForumPostsService) GetGlobalTopInfoList() (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := 3
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var hkForumPostss []community.ForumPostsBaseInfo

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}

// GetGlobalRecommendInfoList 分页获全局推荐资讯列表
func (appForumPostsService *AppForumPostsService) GetGlobalRecommendInfoList(page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var hkForumPostss []community.ForumPostsBaseInfo

	if len(page.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+page.Keyword+"%")
	}

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}

// GetNearbyRecommendPostsList 分页获附近推荐帖子列表
func (appForumPostsService *AppForumPostsService) GetNearbyRecommendPostsList(curPos string, page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var hkForumPostss []community.ForumPostsBaseInfo

	if len(page.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+page.Keyword+"%")
	}

	db = db.Where("channel_id = 0 and is_public = 1 and check_status=?",
		community.PostsCheckStatusPass)

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}

// GetGlobalRecommendQuestionList 分页获取全局推荐问题列表
func (appForumPostsService *AppForumPostsService) GetGlobalRecommendQuestionList(curPos string, page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var hkForumPostss []community.ForumPostsBaseInfo

	if len(page.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+page.Keyword+"%")
	}

	db = db.Where("channel_id = 0 and is_public = 1 and check_status=? and category = ?",
		community.PostsCheckStatusPass,
		community.PostsCategoryQuestion)

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}

// GetGlobalRecommendActivityList 分页获全局推荐活动列表
func (appForumPostsService *AppForumPostsService) GetGlobalRecommendActivityList(page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var hkForumPostss []community.ForumPostsBaseInfo

	if len(page.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+page.Keyword+"%")
	}

	//db = db.Where("channel_id = 0 and is_public = 1 and check_status=? and category = ?",
	//	community.PostsCheckStatusPass,
	//	community.PostsCategoryActivity)
	db = db.Where("channel_id = 0 and is_public = 1 and category = ?",
		community.PostsCategoryActivity)
	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}

// GetForumPostsInfoList 分页获取ForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) GetForumPostsInfoList(info communityReq.ForumPostsSearch) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var hkForumPostss []community.ForumPostsBaseInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	//圈子_编号
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}

	//圈子_编号
	if info.CircleId != 0 {
		db = db.Where("circle_id = ?", info.CircleId)
	}

	//类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动
	if info.Category != nil {
		db = db.Where("category = ?", info.Category)
	}

	//帖子分类编号
	if info.GroupId != nil {
		db = db.Where("group_id = ?", info.GroupId)
	}
	//标题
	if len(info.Title) != 0 {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	//置顶：0否、1是
	if info.Top != nil {
		db = db.Where("top = ?", info.Top)
	}
	//精华：0否、1是
	if info.Marrow != nil {
		db = db.Where("marrow = ?", info.Marrow)
	}

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}

// GetCircleForumPostsList 分页获取ForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) GetCircleForumPostsList(info communityReq.CircleForumPostsSearch) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var hkForumPostss []community.ForumPostsBaseInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	//圈子_编号
	db = db.Where("circle_id = ?", info.CircleId)

	//类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动
	if info.Category != nil {
		db = db.Where("category = ?", info.Category)
	}

	//帖子分类编号
	if info.ChannelId != nil {
		db = db.Where("channel_id = ?", info.ChannelId)
	}
	//标题
	if len(info.Title) != 0 {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	//置顶：0否、1是
	if info.Top != nil {
		db = db.Where("top = ?", info.Top)
	}
	//精华：0否、1是
	if info.Marrow != nil {
		db = db.Where("marrow = ?", info.Marrow)
	}
	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

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
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var hkForumPostss []community.ForumPostsBaseInfo
	db.Joins(",`hk_circle_user`")
	db = db.Where("hk_circle_user.circle_id = hk_forum_posts.circle_id and hk_circle_user.user_id =@userId", sql.Named("userId", info.UserId))

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	//类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动
	if info.Category != nil {
		db = db.Where("category = ?", info.Category)
	}

	//帖子分类编号
	if info.GroupId != nil {
		db = db.Where("group_id = ?", info.GroupId)
	}
	//标题
	if len(info.Title) != 0 {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	//置顶：0否、1是
	if info.Top != nil {
		db = db.Where("top = ?", info.Top)
	}
	//精华：0否、1是
	if info.Marrow != nil {
		db = db.Where("marrow = ?", info.Marrow)
	}
	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}
