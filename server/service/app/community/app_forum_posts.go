package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"time"
)

type AppForumPostsService struct {
}

// CreateForumPosts 创建帖子记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) CreateForumPosts(info communityReq.CreateForumPostsReq) (err error) {
	checkStatus := community.PostsCheckStatusPass
	if info.Draft == community.DraftTrue {
		checkStatus = community.PostsCheckStatusDraft
	}
	forumPosts := community.ForumPosts{
		UserId:          info.UserId,
		CircleId:        info.CircleId,
		Category:        info.Category,
		Title:           info.Title,
		CoverImage:      info.CoverImage,
		ContentType:     info.ContentType,
		ContentMarkdown: info.ContentMarkdown,
		ContentHtml:     info.ContentHtml,
		Video:           info.Video,
		Attachment:      info.Attachment,
		Anonymity:       info.Anonymity,
		CheckStatus:     checkStatus,
		IsPublic:        community.ForumPostsIsPublicTrue,
		PowerComment:    community.ForumPostsPowerCommentOpen,
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
func (appForumPostsService *AppForumPostsService) GetForumPosts(id uint64) (hkForumPosts community.ForumPosts, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumPosts).Error
	return
}

func (appForumPostsService *AppForumPostsService) FindForumPosts(selectUserId uint64, id uint64) (hkForumPosts community.ForumPosts, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("UserInfo").First(&hkForumPosts).Error
	if err == nil {
		isFocus, isFan, _ := GetIsFocusAndIsFan(selectUserId, &hkForumPosts.UserInfo)
		hkForumPosts.UserInfo.IsFocus = isFocus
		hkForumPosts.UserInfo.IsFan = isFan
	}
	return
}

// GetRecommendPostsList 分页获取推荐帖子列表
func (appForumPostsService *AppForumPostsService) GetRecommendPostsList(selectUserId uint64, channelId uint64, page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var forumPosts []community.ForumPostsBaseInfo

	db = db.Where("channel_id = ?", channelId)

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&forumPosts).Error
	if err == nil {
		SetPostsListUserIsFocus(selectUserId, forumPosts)
	}
	return forumPosts, total, err
}

// GetGlobalTopInfoList 分页获全局置顶资讯列表
func (appForumPostsService *AppForumPostsService) GetGlobalTopInfoList(selectUserId uint64) (list []community.ForumPostsBaseInfo, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var forumPosts []community.ForumPostsBaseInfo

	db = db.Where("top = 1 and is_public = 1 and check_status=? and category = ?",
		community.PostsCheckStatusPass,
		community.PostsCategoryNews)

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(utils.HomePageTopNewsNum).Find(&forumPosts).Error
	if err == nil {
		SetPostsListUserIsFocus(selectUserId, forumPosts)
	}
	return forumPosts, total, err
}

// GetGlobalRecommendInfoList 分页获全局推荐资讯列表
func (appForumPostsService *AppForumPostsService) GetGlobalRecommendInfoList(selectUserId uint64, page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var forumPosts []community.ForumPostsBaseInfo

	db = db.Where("is_public = 1 and check_status=? and category = ?",
		community.PostsCheckStatusPass,
		community.PostsCategoryNews)

	if len(page.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+page.Keyword+"%")
	}

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&forumPosts).Error
	if err == nil {
		SetPostsListUserIsFocus(selectUserId, forumPosts)
	}
	return forumPosts, total, err
}

// GetNearbyRecommendPostsList 分页获附近推荐帖子列表
func (appForumPostsService *AppForumPostsService) GetNearbyRecommendPostsList(selectUserId uint64, curPos string, page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var forumPosts []community.ForumPostsBaseInfo

	if len(page.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+page.Keyword+"%")
	}

	db = db.Where("circle_id = 0 and is_public = 1 and check_status=?",
		community.PostsCheckStatusPass)

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&forumPosts).Error
	if err == nil {
		SetPostsListUserIsFocus(selectUserId, forumPosts)
	}
	return forumPosts, total, err
}

// GetGlobalRecommendQuestionList 分页获取全局推荐问题列表
func (appForumPostsService *AppForumPostsService) GetGlobalRecommendQuestionList(selectUserId uint64, curPos string, page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var forumPosts []community.ForumPostsBaseInfo

	if len(page.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+page.Keyword+"%")
	}

	db = db.Where("circle_id = 0 and is_public = 1 AND check_status=? AND category = ?",
		community.PostsCheckStatusPass,
		community.PostsCategoryQuestion)

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&forumPosts).Error
	if err == nil {
		SetPostsListUserIsFocus(selectUserId, forumPosts)
	}
	return forumPosts, total, err
}

// GetGlobalRecommendActivityList 分页获全局推荐活动列表
func (appForumPostsService *AppForumPostsService) GetGlobalRecommendActivityList(selectUserId uint64, page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var forumPosts []community.ForumPostsBaseInfo

	if len(page.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+page.Keyword+"%")
	}

	db = db.Where("circle_id = 0 and is_public = 1 and check_status=? and category = ?",
		community.PostsCheckStatusPass,
		community.PostsCategoryActivity)

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&forumPosts).Error
	if err == nil {
		SetPostsListUserIsFocus(selectUserId, forumPosts)
	}
	return forumPosts, total, err
}

// GetForumPostsInfoList 分页获取ForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) GetForumPostsInfoList(selectUserId uint64, info communityReq.ForumPostsSearch) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var forumPosts []community.ForumPostsBaseInfo
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

	err = db.Limit(limit).Offset(offset).Find(&forumPosts).Error
	if err == nil {
		SetPostsListUserIsFocus(selectUserId, forumPosts)
	}
	return forumPosts, total, err
}

func (appForumPostsService *AppForumPostsService) GetUserForumPostsList(selectUserId uint64, userId uint64, info communityReq.UserForumPostsSearch) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var forumPosts []community.ForumPostsBaseInfo

	db = db.Where("user_id = ?", userId)
	if selectUserId != userId {
		db = db.Where("is_public = ? AND check_status = ?", community.ForumPostsIsPublicTrue, community.PostsCheckStatusPass)
	} else {
		db = db.Where("check_status != ?", community.PostsCheckStatusDraft)
	}

	//类别：1视频、2动态、3资讯、4公告、5文章、6问答、7活动
	if info.Category != 0 {
		db = db.Where("category = ?", info.Category)
	}

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&forumPosts).Error
	if err == nil {
		size := len(list)
		if size > 0 {
			if selectUserId != list[0].UserInfo.ID {
				isFocus, isFan, _ := GetIsFocusAndIsFan(selectUserId, &list[0].UserInfo)
				for i := 0; i < size; i++ {
					list[i].UserInfo.IsFocus = isFocus
					list[i].UserInfo.IsFan = isFan
				}
			}
		}
	}
	if err == nil {
		SetPostsListUserIsFocus(selectUserId, forumPosts)
	}
	return forumPosts, total, err
}

// GetCircleForumPostsList 分页获取ForumPosts记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) GetCircleForumPostsList(selectUserId uint64, info communityReq.CircleForumPostsSearch) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var forumPosts []community.ForumPostsBaseInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	//圈子_编号
	db = db.Where("circle_id = ?", info.CircleId)

	//帖子分类编号
	if info.ChannelId > community.CircleChannelAll {
		if info.ChannelId == community.CircleChannelQuestion {
			db = db.Where("channel_id = 0 AND category = ?", community.PostsCategoryQuestion)
		} else if info.ChannelId == community.CircleChannelDynamic {
			db = db.Where("channel_id = 0 AND category = ?", community.PostsCategoryDynamic)
		} else if info.ChannelId == community.CircleChannelNews {
			db = db.Where("channel_id = 0 AND category = ?", community.PostsCategoryNews)
		} else if info.ChannelId == community.CircleChannelActivity {
			db = db.Where("channel_id = 0 AND category = ?", community.PostsCategoryActivity)
		} else if info.ChannelId == community.CircleChannelArticle {
			db = db.Where("channel_id = 0 AND category = ?", community.PostsCategoryArticle)
		} else {
			db = db.Where("channel_id = ?", info.ChannelId)
		}
	}
	if len(info.Keyword) != 0 {
		db = db.Where("title LIKE ?", "%"+info.Keyword+"%")
	}
	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&forumPosts).Error
	if err == nil {
		SetPostsListUserIsFocus(selectUserId, forumPosts)
	}
	return forumPosts, total, err
}

func (appForumPostsService *AppForumPostsService) GetFocusUserPostsList(userId uint64, page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)

	//查询关注用户
	type tmp struct {
		ID         uint64
		UserExtend community.UserExtend `gorm:"foreignKey:ID;references:ID;"` //用户扩展
	}
	var focusUsers []tmp
	focusUserDb := global.GVA_DB.Model(&community.FocusUser{}).Select("focus_user_id as id").Where("user_id = ?", userId)
	err = focusUserDb.Find(&focusUsers).Error
	if err != nil {
		return
	}
	var focusUsersSize = len(focusUsers)
	var hkForumPosts []community.ForumPostsBaseInfo
	if focusUsersSize == 0 {
		return hkForumPosts, total, err
	}
	var ids = make([]uint64, focusUsersSize)
	for index, v := range focusUsers {
		ids[index] = v.ID
	}

	//查询最近发布帖子的关注用户
	type tmpUpdate struct {
		ID uint64
		Tm *time.Time
	}
	var focusUsersTm []tmpUpdate
	db1 := global.GVA_DB.Model(&community.UserExtend{}).Select("id,update_forum_posts_time as tm").Where("id in ?", ids)
	db1 = db1.Order("update_forum_posts_time desc")
	err = db1.Find(&focusUsersTm).Error
	if err != nil {
		return
	}
	var size = len(focusUsersTm)
	if size == 0 {
		return hkForumPosts, total, err
	}
	var idEx = make([]uint64, size)
	for index, v := range focusUsersTm {
		idEx[index] = v.ID
	}
	//查询最新发布的帖子
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	db = db.Where("user_id in ?", idEx)
	db = db.Where("is_public = 1 and check_status=?",
		community.PostsCheckStatusPass)

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPosts).Error
	if err == nil {
		//因为是关注用户的帖子，所以帖子的用户都是关注状态
		for index, _ := range hkForumPosts {
			hkForumPosts[index].UserInfo.IsFocus = 1
		}
		//SetPostsListUserIsFocus(userId, hkForumPosts)
	}

	return hkForumPosts, total, err
}

// GetUserForumPostsInfoList 分页获取用户加入圈子的所有动态列表
// Author [piexlmax](https://github.com/piexlmax)
func (appForumPostsService *AppForumPostsService) GetUserForumPostsInfoList(userId uint64, info communityReq.UserCircleForumPostsSearch) (list []community.ForumPostsBaseInfo, total int64, err error) {
	type tmp struct {
		ID uint64
	}
	var circles1 []tmp
	var hkForumPosts []community.ForumPostsBaseInfo
	dbCircleUser := global.GVA_DB.Model(&community.CircleUser{}).Select("circle_id as id").Joins("left join hk_circle on hk_circle.id = hk_circle_user.circle_id").Where("hk_circle_user.user_id = ?", userId)
	err = dbCircleUser.Order("hk_circle.update_forum_posts_time desc").Limit(utils.PageSizeMax).Offset(0).Find(&circles1).Error
	if err != nil {
		return hkForumPosts, total, err
	}
	var circleCount = len(circles1)

	if circleCount == 0 {
		return make([]community.ForumPostsBaseInfo, 0), total, err
	}
	var ids = make([]uint64, circleCount)
	for index, v := range circles1 {
		ids[index] = v.ID
	}

	//优化，need to do
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	db = db.Where("circle_id in ?", ids)

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

	err = db.Limit(limit).Offset(offset).Find(&hkForumPosts).Error
	if err == nil {
		SetPostsListUserIsFocus(userId, hkForumPosts)
	}
	return hkForumPosts, total, err
}
func (appForumPostsService *AppForumPostsService) GetDraft(id uint64) (hkForumPosts community.ForumPosts, err error) {
	err = global.GVA_DB.Where("id = ? AND check_status = ?", id, community.PostsCheckStatusDraft).First(&hkForumPosts).Error
	return
}
func (appForumPostsService *AppForumPostsService) GetDraftList(userId uint64, info generalReq.DraftSearch) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("CircleInfo")
	var forumPosts []community.ForumPostsBaseInfo

	db = db.Where("check_status = ? AND user_id = ?", community.PostsCheckStatusDraft, userId)
	if info.Category > 0 {
		db = db.Where("category = ? ", info.Category)
	}

	//创建时间降序排列
	db = db.Order("updated_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&forumPosts).Error
	return forumPosts, total, err
}
func (appForumPostsService *AppForumPostsService) DeleteAllDraft(userId uint64, category int) (err error) {
	if category > 0 {
		err = global.GVA_DB.Unscoped().Delete(&[]community.ForumPostsBaseInfo{}, "user_id = ? AND check_status = ? AND category = ?", userId, community.PostsCheckStatusDraft, category).Error
	} else {
		err = global.GVA_DB.Unscoped().Delete(&[]community.ForumPostsBaseInfo{}, "user_id = ? AND check_status = ?", userId, community.PostsCheckStatusDraft).Error
	}
	return err
}
func (appForumPostsService *AppForumPostsService) DeleteDraft(userId uint64, id uint64) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&community.ForumPostsBaseInfo{}, "id = ? AND user_id = ?", id, userId).Error
	return err
}
func (appForumPostsService *AppForumPostsService) DeleteDraftByIds(userId uint64, ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]community.ForumPostsBaseInfo{}, "id in ? AND user_id = ?", ids.Ids, userId).Error
	return err
}
func (appForumPostsService *AppForumPostsService) UpdateDraft(userId uint64, info generalReq.UpdateDraftReq) (err error) {
	var obj community.ForumPosts
	err = global.GVA_DB.Where("id = ? AND user_id = ? AND check_status = ?", info.Id, userId, community.PostsCheckStatusDraft).First(&obj).Error
	if err != nil {
		return
	}
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})

	if obj.Category != *info.Category {
		//TopicId: need to do
		updateData["circle_id"] = 0
		updateData["title"] = ""
		updateData["cover_image"] = ""
		updateData["attachment"] = ""
		updateData["pay_content"] = 0
		updateData["pay_content_look"] = 0
		updateData["pay_attachment"] = 0
		updateData["pay_num"] = 0
		updateData["content_html"] = ""
		updateData["video"] = ""
		updateData["tag"] = ""
		updateData["anonymity"] = 0
		updateData["power_comment"] = 1
		updateData["power_comment_anonymity"] = 0
		updateData["is_public"] = 1
		if obj.Category == community.PostsCategoryActivity {
			updateData["activity_start_at"] = ""
			updateData["activity_end_at"] = ""
			updateData["activity_address"] = ""
			updateData["activity_user_num"] = ""
			updateData["activity_add_approve"] = ""
		}
	}

	if info.CircleId != nil {
		updateData["circle_id"] = info.CircleId
	}
	if info.Category != nil {
		updateData["category"] = info.Category
	}
	if info.ChannelId != nil {
		updateData["channel_id"] = info.ChannelId
	}
	if len(info.Title) > 0 {
		updateData["title"] = info.Title
	}
	if len(info.CoverImage) > 0 {
		updateData["cover_image"] = info.CoverImage
	}
	if len(info.ContentHtml) > 0 {
		updateData["content_html"] = info.ContentHtml
	}
	if len(info.Video) > 0 {
		updateData["video"] = info.Video
	}
	if len(info.Attachment) > 0 {
		updateData["attachment"] = info.Attachment
	}
	if len(info.Tag) > 0 {
		updateData["tag"] = info.Tag
	}
	if info.Anonymity != nil {
		updateData["anonymity"] = info.Anonymity
	}
	if info.PowerComment != nil {
		updateData["power_comment"] = info.PowerComment
	}
	if info.PowerCommentAnonymity != nil {
		updateData["power_comment_anonymity"] = info.PowerCommentAnonymity
	}
	if info.PayContent != nil {
		updateData["pay_content"] = info.PayContent
	}
	if info.PayContentLook != nil {
		updateData["pay_content_look"] = info.PayContentLook
	}
	if info.PayAttachment != nil {
		updateData["pay_attachment"] = info.PayAttachment
	}
	if info.PayNum != nil {
		updateData["pay_num"] = info.PayNum
	}
	if len(info.ActivityStartAt) > 0 {
		updateData["activity_start_at"] = info.ActivityStartAt
	}
	if len(info.ActivityEndAt) > 0 {
		updateData["activity_end_at"] = info.ActivityEndAt
	}
	if len(info.ActivityAddress) > 0 {
		updateData["activity_address"] = info.ActivityAddress
	}
	if info.ActivityUserNum != nil {
		updateData["activity_user_num"] = info.ActivityUserNum
	}
	if info.ActivityAddApprove != nil {
		updateData["activity_add_approve"] = info.ActivityAddApprove
	}
	if info.IsPublic != nil {
		updateData["is_public"] = info.IsPublic
	}

	if info.Draft != nil && *info.Draft == 0 {
		updateData["check_status"] = community.PostsCheckStatusPass
	}
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{})
	err = db.Where("id = ? AND user_id = ? AND check_status = ?", info.Id, userId, community.PostsCheckStatusDraft).Updates(updateData).Error
	return err
}
