package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ActivityAddRequestService struct {
}

// CreateActivityAddRequest 创建ActivityAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityAddRequestService *ActivityAddRequestService) CreateActivityAddRequest(activityAddRequest *community.ActivityAddRequest) (err error) {
	err = global.GVA_DB.Create(activityAddRequest).Error
	return err
}

// DeleteActivityAddRequest 删除ActivityAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityAddRequestService *ActivityAddRequestService) DeleteActivityAddRequest(activityAddRequest community.ActivityAddRequest) (err error) {
	err = global.GVA_DB.Delete(&activityAddRequest).Error
	return err
}

// DeleteActivityAddRequestByIds 批量删除ActivityAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityAddRequestService *ActivityAddRequestService) DeleteActivityAddRequestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ActivityAddRequest{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateActivityAddRequest 更新ActivityAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityAddRequestService *ActivityAddRequestService) UpdateActivityAddRequest(activityAddRequest community.ActivityAddRequest) (err error) {
	err = global.GVA_DB.Save(&activityAddRequest).Error
	return err
}

// GetActivityAddRequest 根据id获取ActivityAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityAddRequestService *ActivityAddRequestService) GetActivityAddRequest(id uint64) (activityAddRequest community.ActivityAddRequest, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&activityAddRequest).Error
	return
}

// GetActivityAddRequestInfoList 分页获取ActivityAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityAddRequestService *ActivityAddRequestService) GetActivityAddRequestInfoList(info communityReq.ActivityAddRequestSearch) (list []community.ActivityAddRequest, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ActivityAddRequest{})
	var hkActivityAddRequests []community.ActivityAddRequest
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkActivityAddRequests).Error
	return hkActivityAddRequests, total, err
}
