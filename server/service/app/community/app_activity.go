package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type ActivityService struct {
}

// CreateActivity 创建Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *ActivityService) CreateActivity(userId uint64, info communityReq.CreateActivityReq) (activity community.ForumPosts, err error) {
	checkStatus := community.PostsCheckStatusPass
	if info.Draft == community.DraftTrue {
		checkStatus = community.PostsCheckStatusDraft
	}
	obj := community.ForumPosts{
		CircleId:           info.CircleId,
		UserId:             userId,
		Title:              info.Title,
		CoverImage:         info.CoverImage,
		ChannelId:          community.ChannelCodeActivity,
		ContentType:        community.ContentTypeHtml,
		Category:           community.PostsCategoryActivity,
		ContentHtml:        info.Content,
		ActivityStartAt:    info.ActivityStartAt,
		ActivityEndAt:      info.ActivityEndAt,
		ActivityAddress:    info.ActivityAddress,
		ActivityUserNum:    info.ActivityUserNum,
		ActivityCurUserNum: 1,
		PayCurrency:        utils.CurrencyGold,
		PayNum:             info.PayNum,
		CheckStatus:        checkStatus,
		IsPublic:           community.ForumPostsIsPublicTrue,
		PowerComment:       community.ForumPostsPowerCommentOpen,
	}
	err = global.GVA_DB.Create(&obj).Error
	if err == nil {
		err = global.GVA_DB.Create(&community.ActivityUser{
			ActivityId: obj.ActivityId,
			UserId:     userId,
		}).Error
		if len(info.TopicId) > 0 {
			tmp := utils.SplitToUint64List(info.TopicId, ",")
			for _, topicId := range tmp {
				err = global.GVA_DB.Create(&community.ForumTopicPostsMapping{
					TopicId: topicId,
					PostsId: obj.ID,
				}).Error
			}
		}
	}
	return obj, err
}

// DeleteActivity 删除Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *ActivityService) DeleteActivity(hkActivity community.ForumPosts) (err error) {
	err = global.GVA_DB.Delete(&hkActivity).Error
	return err
}

func (hkActivityService *ActivityService) DeleteActivityById(id uint64) (err error) {
	err = global.GVA_DB.Where("id = ?", id).Delete(&community.ForumPosts{}).Error
	return err
}

// DeleteActivityByIds 批量删除Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *ActivityService) DeleteActivityByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumPosts{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateActivity 更新Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityService *ActivityService) UpdateActivity(info communityReq.UpdateActivityReq) (err error) {
	//err = global.GVA_DB.Save(&hkActivity).Error
	db := global.GVA_DB.Model(&community.ForumPosts{})
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})

	if len(info.Title) > 0 {
		updateData["title"] = info.Title
	}
	if len(info.Content) > 0 {
		updateData["content_html"] = info.Content
	}
	if info.ActivityStartAt != nil {
		updateData["activity_start_at"] = info.ActivityStartAt
	}
	if info.ActivityEndAt != nil {
		updateData["activity_end_at"] = info.ActivityEndAt
	}
	if len(info.ActivityAddress) > 0 {
		updateData["activity_address"] = info.ActivityAddress
	}
	if info.ActivityUserNum != nil {
		updateData["activity_user_num"] = info.ActivityUserNum
	}
	if info.PayNum != nil {
		updateData["pay_num"] = info.PayNum
	}
	err = db.Where("id = ?", info.Id).Updates(updateData).Error
	return err
}

// GetActivity 根据id获取Activity记录
func (hkActivityService *ActivityService) GetActivity(id uint64) (hkActivity community.ForumPosts, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkActivity).Error
	return
}
func (hkActivityService *ActivityService) GetActivityByUser(selectUserId uint64, id uint64) (hkActivity community.ForumPosts, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("UserInfo").First(&hkActivity).Error
	if err == nil {
		isFocus, isFan, _ := GetIsFocusAndIsFan(selectUserId, &hkActivity.UserInfo)
		hkActivity.UserInfo.IsFocus = isFocus
		hkActivity.UserInfo.IsFan = isFan
	}
	return
}

// GetActivityInfoList 分页获取Activity记录
func (hkActivityService *ActivityService) GetActivityInfoList(selectUserId uint64, info communityReq.ActivitySearch) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var hkActivitys []community.ForumPostsBaseInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkActivitys).Error
	if err == nil {
		SetPostsListUserIsFocus(selectUserId, hkActivitys)
	}
	return hkActivitys, total, err
}

func (hkActivityService *ActivityService) CreateActivityDynamic(userId uint64, activity community.ForumPosts, content string, attachment string) (err error) {
	err = global.GVA_DB.Create(&community.ForumPosts{
		CircleId:   activity.CircleId,
		ActivityId: activity.ID,
		Category:   community.PostsCategoryDynamic,
		UserId:     userId,
		Title:      content,
		Attachment: attachment,
	}).Error
	return err
}
func (hkActivityService *ActivityService) GetActivityDynamic(id uint64) (dynamic community.ForumPosts, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dynamic).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("动态不存在")
		return
	} else if err != nil {
		return
	}
	if dynamic.ActivityId == 0 {
		err = errors.New("不是活动动态")
		return
	}
	return
}
func (hkActivityService *ActivityService) DeleteActivityDynamic(id uint64) (err error) {
	err = global.GVA_DB.Where("id = ?", id).Delete(&community.ForumPosts{}).Error
	return err
}
func (hkActivityService *ActivityService) DeleteActivityDynamicByIds(activityId uint64, ids []uint64) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumPosts{}, "activity_id = ? AND id = ?", activityId, ids).Error
	return err
}

func (hkActivityService *ActivityService) GetActivityDynamicList(selectUserId uint64, activityId uint64, page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var activityDynamicList []community.ForumPostsBaseInfo

	db = db.Where("activity_id = ?", activityId)

	if len(page.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+page.Keyword+"%")
	}

	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&activityDynamicList).Error
	if err == nil {
		SetPostsListUserIsFocus(selectUserId, activityDynamicList)
	}
	return activityDynamicList, total, err
}
