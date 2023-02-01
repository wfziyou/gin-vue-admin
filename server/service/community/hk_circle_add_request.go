package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type HkCircleAddRequestService struct {
}

// CreateHkCircleAddRequest 创建HkCircleAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleAddRequestService *HkCircleAddRequestService) CreateHkCircleAddRequest(hkCircleAddRequest community.HkCircleAddRequest) (err error) {
	err = global.GVA_DB.Create(&hkCircleAddRequest).Error
	return err
}

// DeleteHkCircleAddRequest 删除HkCircleAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleAddRequestService *HkCircleAddRequestService)DeleteHkCircleAddRequest(hkCircleAddRequest community.HkCircleAddRequest) (err error) {
	err = global.GVA_DB.Delete(&hkCircleAddRequest).Error
	return err
}

// DeleteHkCircleAddRequestByIds 批量删除HkCircleAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleAddRequestService *HkCircleAddRequestService)DeleteHkCircleAddRequestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkCircleAddRequest{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkCircleAddRequest 更新HkCircleAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleAddRequestService *HkCircleAddRequestService)UpdateHkCircleAddRequest(hkCircleAddRequest community.HkCircleAddRequest) (err error) {
	err = global.GVA_DB.Save(&hkCircleAddRequest).Error
	return err
}

// GetHkCircleAddRequest 根据id获取HkCircleAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleAddRequestService *HkCircleAddRequestService)GetHkCircleAddRequest(id uint) (hkCircleAddRequest community.HkCircleAddRequest, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleAddRequest).Error
	return
}

// GetHkCircleAddRequestInfoList 分页获取HkCircleAddRequest记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleAddRequestService *HkCircleAddRequestService)GetHkCircleAddRequestInfoList(info communityReq.HkCircleAddRequestSearch) (list []community.HkCircleAddRequest, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&community.HkCircleAddRequest{})
    var hkCircleAddRequests []community.HkCircleAddRequest
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkCircleAddRequests).Error
	return  hkCircleAddRequests, total, err
}
