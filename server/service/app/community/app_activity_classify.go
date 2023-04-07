package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ActivityClassifyService struct {
}

// CreateActivityClassify 创建活动分类记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityClassifyService *ActivityClassifyService) CreateActivityClassify(hkActivityClassify community.ActivityClassify) (err error) {
	err = global.GVA_DB.Create(&hkActivityClassify).Error
	return err
}

// DeleteActivityClassify 删除ActivityClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityClassifyService *ActivityClassifyService) DeleteActivityClassify(hkActivityClassify community.ActivityClassify) (err error) {
	err = global.GVA_DB.Delete(&hkActivityClassify).Error
	return err
}

// DeleteActivityClassifyByIds 批量删除ActivityClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityClassifyService *ActivityClassifyService) DeleteActivityClassifyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ActivityClassify{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateActivityClassify 更新ActivityClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityClassifyService *ActivityClassifyService) UpdateActivityClassify(hkActivityClassify community.ActivityClassify) (err error) {
	err = global.GVA_DB.Save(&hkActivityClassify).Error
	return err
}

// GetActivityClassify 根据id获取ActivityClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityClassifyService *ActivityClassifyService) GetActivityClassify(id uint64) (hkActivityClassify community.ActivityClassify, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkActivityClassify).Error
	return
}

// GetActivityClassifyInfoList 分页获取ActivityClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityClassifyService *ActivityClassifyService) GetActivityClassifyInfoList(info communityReq.ActivityClassifySearch) (list []community.ActivityClassify, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ActivityClassify{})
	var hkActivityClassifys []community.ActivityClassify
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkActivityClassifys).Error
	return hkActivityClassifys, total, err
}
