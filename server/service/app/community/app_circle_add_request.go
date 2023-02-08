package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCircleAddRequestService struct {
}

// CreateCircleAddRequest 创建CircleAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleAddRequestService *AppCircleAddRequestService) CreateCircleAddRequest(hkCircleAddRequest community.CircleAddRequest) (err error) {
	err = global.GVA_DB.Create(&hkCircleAddRequest).Error
	return err
}

// DeleteCircleAddRequest 删除CircleAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleAddRequestService *AppCircleAddRequestService) DeleteCircleAddRequest(hkCircleAddRequest community.CircleAddRequest) (err error) {
	err = global.GVA_DB.Delete(&hkCircleAddRequest).Error
	return err
}

// DeleteCircleAddRequestByIds 批量删除CircleAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleAddRequestService *AppCircleAddRequestService) DeleteCircleAddRequestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.CircleAddRequest{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCircleAddRequest 更新CircleAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleAddRequestService *AppCircleAddRequestService) UpdateCircleAddRequest(hkCircleAddRequest community.CircleAddRequest) (err error) {
	err = global.GVA_DB.Save(&hkCircleAddRequest).Error
	return err
}

// GetCircleAddRequest 根据id获取CircleAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleAddRequestService *AppCircleAddRequestService) GetCircleAddRequest(id uint) (hkCircleAddRequest community.CircleAddRequest, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleAddRequest).Error
	return
}

// GetCircleAddRequestInfoList 分页获取CircleAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleAddRequestService *AppCircleAddRequestService) GetCircleAddRequestInfoList(info communityReq.CircleAddRequestSearch) (list []community.CircleAddRequest, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleAddRequest{})
	var hkCircleAddRequests []community.CircleAddRequest
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircleAddRequests).Error
	return hkCircleAddRequests, total, err
}
