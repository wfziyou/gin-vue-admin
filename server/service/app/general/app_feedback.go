package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	generalReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/general/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FeedbackService struct {
}

// CreateFeedback 创建Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackService *FeedbackService) CreateFeedback(userId uint64, typeId uint64, typeName string, des string, attachment string) (err error) {
	err = global.GVA_DB.Create(&general.Feedback{
		UserId:     userId,
		TypeId:     typeId,
		Type:       typeName,
		Des:        des,
		Attachment: attachment,
	}).Error
	return err
}

// DeleteFeedback 删除Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackService *FeedbackService) DeleteFeedback(hkFeedback general.Feedback) (err error) {
	err = global.GVA_DB.Delete(&hkFeedback).Error
	return err
}

// DeleteFeedbackByIds 批量删除Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackService *FeedbackService) DeleteFeedbackByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.Feedback{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFeedback 更新Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackService *FeedbackService) UpdateFeedback(info generalReq.UpdateFeedbackReq) (err error) {
	db := global.GVA_DB.Model(&general.Feedback{})
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
func (hkFeedbackService *FeedbackService) GetFeedback(id uint64) (hkFeedback general.Feedback, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkFeedback).Error
	return
}

// GetFeedbackInfoList 分页获取Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackService *FeedbackService) GetFeedbackInfoList(info generalReq.FeedbackSearch) (list []general.Feedback, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&general.Feedback{})
	var hkFeedbacks []general.Feedback
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
