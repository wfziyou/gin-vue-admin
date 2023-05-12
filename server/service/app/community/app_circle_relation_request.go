package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCircleRelationRequestService struct {
}

// CreateCircleRelationRequest 创建CircleRelationRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleRelationRequestService *AppCircleRelationRequestService) CreateCircleRelationRequest(hkCircleRelationRequest *community.CircleRelationRequest) (err error) {
	err = global.GVA_DB.Create(hkCircleRelationRequest).Error
	return err
}

// DeleteCircleRelationRequest 删除CircleRelationRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleRelationRequestService *AppCircleRelationRequestService) DeleteCircleRelationRequest(hkCircleRelationRequest community.CircleRelationRequest) (err error) {
	err = global.GVA_DB.Delete(&hkCircleRelationRequest).Error
	return err
}

// DeleteCircleRelationRequestByIds 批量删除CircleRelationRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleRelationRequestService *AppCircleRelationRequestService) DeleteCircleRelationRequestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.CircleRelationRequest{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCircleRelationRequest 更新CircleRelationRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleRelationRequestService *AppCircleRelationRequestService) UpdateCircleRelationRequest(hkCircleRelationRequest community.CircleRelationRequest) (err error) {
	err = global.GVA_DB.Save(&hkCircleRelationRequest).Error
	return err
}

// GetCircleRelationRequest 根据id获取CircleRelationRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleRelationRequestService *AppCircleRelationRequestService) GetCircleRelationRequest(id uint) (hkCircleRelationRequest community.CircleRelationRequest, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleRelationRequest).Error
	return
}

// GetCircleRelationRequestInfoList 分页获取CircleRelationRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleRelationRequestService *AppCircleRelationRequestService) GetCircleRelationRequestInfoList(info communityReq.CircleRelationRequestSearch) (list []community.CircleRelationRequest, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleRelationRequest{})
	var hkCircleRelationRequests []community.CircleRelationRequest
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircleRelationRequests).Error
	return hkCircleRelationRequests, total, err
}
