package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppForumTopicService struct {
}

// CreateForumTopic 创建话题记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicService *AppForumTopicService) CreateForumTopic(hkForumTopic community.ForumTopic) (err error) {
	err = global.GVA_DB.Create(&hkForumTopic).Error
	return err
}

// DeleteForumTopic 删除话题记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicService *AppForumTopicService) DeleteForumTopic(hkForumTopic community.ForumTopic) (err error) {
	err = global.GVA_DB.Delete(&hkForumTopic).Error
	return err
}

// DeleteForumTopicByIds 批量删除话题记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicService *AppForumTopicService) DeleteForumTopicByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumTopic{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateForumTopic 更新话题记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicService *AppForumTopicService) UpdateForumTopic(hkForumTopic community.ForumTopic) (err error) {
	err = global.GVA_DB.Save(&hkForumTopic).Error
	return err
}

// GetForumTopic 根据id获取ForumTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicService *AppForumTopicService) GetForumTopic(id uint64) (hkForumTopic community.ForumTopic, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumTopic).Error
	return
}

// GetForumTopicInfoList 分页获取ForumTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (appForumTopicService *AppForumTopicService) GetForumTopicInfoList(info communityReq.ForumTopicSearch) (list []community.ForumTopic, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumTopic{})
	var hkForumTopics []community.ForumTopic
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if len(info.Name) > 0 {
		db = db.Where("name LIKE '?%'", info.Name)
	}
	if info.TopicGroupId != 0 {
		db = db.Where("topic_group_id = ?", info.TopicGroupId)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.Hot != nil {
		db = db.Where("hot = ?", info.Hot)
	}
	if info.CircleId != 0 {
		db = db.Where("circle_id = ?", info.CircleId)
	}
	if info.ReviewStatus != nil {
		db = db.Where("review_status = ?", info.ReviewStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumTopics).Error
	return hkForumTopics, total, err
}
