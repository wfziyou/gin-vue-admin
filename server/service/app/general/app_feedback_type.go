package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FeedbackTypeService struct {
}

// CreateFeedbackType 创建FeedbackType记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackTypeService *FeedbackTypeService) CreateFeedbackType(hkFeedbackType *general.FeedbackType) (err error) {
	err = global.GVA_DB.Create(hkFeedbackType).Error
	return err
}

// DeleteFeedbackType 删除FeedbackType记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackTypeService *FeedbackTypeService) DeleteFeedbackType(hkFeedbackType general.FeedbackType) (err error) {
	err = global.GVA_DB.Delete(&hkFeedbackType).Error
	return err
}

// DeleteFeedbackTypeByIds 批量删除FeedbackType记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackTypeService *FeedbackTypeService) DeleteFeedbackTypeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]general.FeedbackType{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFeedbackType 更新FeedbackType记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackTypeService *FeedbackTypeService) UpdateFeedbackType(hkFeedbackType general.FeedbackType) (err error) {
	err = global.GVA_DB.Save(&hkFeedbackType).Error
	return err
}

// GetFeedbackType 根据id获取FeedbackType记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackTypeService *FeedbackTypeService) GetFeedbackType(id uint64) (hkFeedbackType general.FeedbackType, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkFeedbackType).Error
	return
}

// GetFeedbackTypeInfoList 获取FeedbackType记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFeedbackTypeService *FeedbackTypeService) GetFeedbackTypeInfoList() (list []general.FeedbackType, err error) {
	// 创建db
	db := global.GVA_DB.Model(&general.FeedbackType{})
	var hkFeedbackTypes []general.FeedbackType
	err = db.Find(&hkFeedbackTypes).Error
	return hkFeedbackTypes, err
}
