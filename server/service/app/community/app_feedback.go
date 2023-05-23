package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FeedbackService struct {
}

// CreateFeedback 创建Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackService *FeedbackService) CreateFeedback(userId uint64, des string, attachment string) (err error) {
	err = global.GVA_DB.Create(&community.Feedback{
		UserId:     userId,
		Des:        des,
		Attachment: attachment,
	}).Error
	return err
}

// DeleteFeedback 删除Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackService *FeedbackService) DeleteFeedback(hkFeedback community.Feedback) (err error) {
	err = global.GVA_DB.Delete(&hkFeedback).Error
	return err
}

// DeleteFeedbackByIds 批量删除Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackService *FeedbackService) DeleteFeedbackByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.Feedback{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFeedback 更新Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackService *FeedbackService) UpdateFeedback(info communityReq.UpdateFeedbackReq) (err error) {
	db := global.GVA_DB.Model(&community.Feedback{})
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	if info.CheckStatus != nil {
		updateData["check_status"] = info.CheckStatus
	}
	if len(info.Process) > 0 {
		updateData["process"] = info.Process
	}
	err = db.Where("id = ?", info.Id).Updates(updateData).Error
	return err
}

// GetFeedback 根据id获取Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackService *FeedbackService) GetFeedback(id uint64) (hkFeedback community.Feedback, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkFeedback).Error
	return
}

// GetFeedbackInfoList 分页获取Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackService *FeedbackService) GetFeedbackInfoList(info communityReq.FeedbackSearch) (list []community.Feedback, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.Feedback{})
	var hkFeedbacks []community.Feedback
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.Keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+info.Keyword+"%")
		db = db.Where("type LIKE ?", "%"+info.Keyword+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkFeedbacks).Error
	return hkFeedbacks, total, err
}
