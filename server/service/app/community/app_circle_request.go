package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type AppCircleRequestService struct {
}

// CreateHkCircleRequest 创建HkCircleRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRequestService *AppCircleRequestService) CreateHkCircleRequest(hkCircleRequest community.HkCircleRequest) (err error) {
	err = global.GVA_DB.Create(&hkCircleRequest).Error
	return err
}

// DeleteHkCircleRequest 删除HkCircleRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRequestService *AppCircleRequestService) DeleteHkCircleRequest(hkCircleRequest community.HkCircleRequest) (err error) {
	err = global.GVA_DB.Delete(&hkCircleRequest).Error
	return err
}

// DeleteHkCircleRequestByIds 批量删除HkCircleRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRequestService *AppCircleRequestService) DeleteHkCircleRequestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkCircleRequest{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkCircleRequest 更新HkCircleRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRequestService *AppCircleRequestService) UpdateHkCircleRequest(hkCircleRequest community.HkCircleRequest) (err error) {
	err = global.GVA_DB.Save(&hkCircleRequest).Error
	return err
}

// GetHkCircleRequest 根据id获取HkCircleRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRequestService *AppCircleRequestService) GetHkCircleRequest(id uint) (hkCircleRequest community.HkCircleRequest, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleRequest).Error
	return
}

// GetHkCircleRequestInfoList 分页获取HkCircleRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRequestService *AppCircleRequestService) GetHkCircleRequestInfoList(info communityReq.HkCircleRequestSearch) (list []community.HkCircleRequest, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkCircleRequest{})
	var hkCircleRequests []community.HkCircleRequest
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircleRequests).Error
	return hkCircleRequests, total, err
}
