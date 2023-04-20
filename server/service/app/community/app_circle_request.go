package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCircleRequestService struct {
}

// CreateCircleRequest 创建圈子请求记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRequestService *AppCircleRequestService) CreateCircleRequest(hkCircleRequest community.CircleRequest) (err error) {
	err = global.GVA_DB.Create(&hkCircleRequest).Error
	return err
}

// DeleteCircleRequest 删除CircleRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRequestService *AppCircleRequestService) DeleteCircleRequest(hkCircleRequest community.CircleRequest) (err error) {
	err = global.GVA_DB.Delete(&hkCircleRequest).Error
	return err
}

// DeleteCircleRequestByIds 批量删除CircleRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRequestService *AppCircleRequestService) DeleteCircleRequestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.CircleRequest{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCircleRequest 更新CircleRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRequestService *AppCircleRequestService) UpdateCircleRequest(hkCircleRequest community.CircleRequest) (err error) {
	err = global.GVA_DB.Save(&hkCircleRequest).Error
	return err
}

// GetCircleRequest 根据id获取CircleRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRequestService *AppCircleRequestService) GetCircleRequest(id uint64) (hkCircleRequest community.CircleRequest, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleRequest).Error
	return
}

// GetCircleRequestInfoList 分页获取CircleRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRequestService *AppCircleRequestService) GetCircleRequestInfoList(info communityReq.CircleRequestSearch) (list []community.CircleRequest, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleRequest{})
	var hkCircleRequests []community.CircleRequest
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if len(info.Name) > 0 {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.CircleClassifyId != 0 {
		db = db.Where("circle_classify_id = ?", info.CircleClassifyId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircleRequests).Error
	return hkCircleRequests, total, err
}
