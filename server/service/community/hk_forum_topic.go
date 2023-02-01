package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type HkForumTopicService struct {
}

// CreateHkForumTopic 创建HkForumTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTopicService *HkForumTopicService) CreateHkForumTopic(hkForumTopic community.HkForumTopic) (err error) {
	err = global.GVA_DB.Create(&hkForumTopic).Error
	return err
}

// DeleteHkForumTopic 删除HkForumTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTopicService *HkForumTopicService)DeleteHkForumTopic(hkForumTopic community.HkForumTopic) (err error) {
	err = global.GVA_DB.Delete(&hkForumTopic).Error
	return err
}

// DeleteHkForumTopicByIds 批量删除HkForumTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTopicService *HkForumTopicService)DeleteHkForumTopicByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForumTopic{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkForumTopic 更新HkForumTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTopicService *HkForumTopicService)UpdateHkForumTopic(hkForumTopic community.HkForumTopic) (err error) {
	err = global.GVA_DB.Save(&hkForumTopic).Error
	return err
}

// GetHkForumTopic 根据id获取HkForumTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTopicService *HkForumTopicService)GetHkForumTopic(id uint) (hkForumTopic community.HkForumTopic, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumTopic).Error
	return
}

// GetHkForumTopicInfoList 分页获取HkForumTopic记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumTopicService *HkForumTopicService)GetHkForumTopicInfoList(info communityReq.HkForumTopicSearch) (list []community.HkForumTopic, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&community.HkForumTopic{})
    var hkForumTopics []community.HkForumTopic
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkForumTopics).Error
	return  hkForumTopics, total, err
}
