package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCircleRelationService struct {
}

// CreateCircleRelation 创建CircleRelation记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRelationService *AppCircleRelationService) CreateCircleRelation(hkCircleRelation community.CircleRelation) (err error) {
	err = global.GVA_DB.Create(&hkCircleRelation).Error
	return err
}

// DeleteCircleRelation 删除CircleRelation记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRelationService *AppCircleRelationService) DeleteCircleRelation(hkCircleRelation community.CircleRelation) (err error) {
	err = global.GVA_DB.Delete(&hkCircleRelation).Error
	return err
}
func (appCircleRelationService *AppCircleRelationService) DeleteCircleRelationByOtherCircleId(circleId uint64, otherCircleId uint64) (err error) {
	err = global.GVA_DB.Delete(&[]community.CircleRelation{}, "circle_id = ? AND other_circle_id = ?", circleId, otherCircleId).Error
	return err
}

// DeleteCircleRelationByIds 批量删除CircleRelation记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRelationService *AppCircleRelationService) DeleteCircleRelationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.CircleRelation{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCircleRelation 更新CircleRelation记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRelationService *AppCircleRelationService) UpdateCircleRelation(hkCircleRelation community.CircleRelation) (err error) {
	err = global.GVA_DB.Save(&hkCircleRelation).Error
	return err
}

// GetCircleRelation 根据id获取CircleRelation记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRelationService *AppCircleRelationService) GetCircleRelation(id uint) (hkCircleRelation community.CircleRelation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleRelation).Error
	return
}

// GetCircleRelationInfoList 分页获取CircleRelation记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleRelationService *AppCircleRelationService) GetCircleRelationInfoList(info communityReq.CircleRelationSearch) (list []community.CircleRelation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleRelation{})
	var hkCircleRelations []community.CircleRelation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircleRelations).Error
	return hkCircleRelations, total, err
}

func (appCircleRelationService *AppCircleRelationService) GetChildCircleList(circleId uint64, page request.PageInfo) (list []community.CircleBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleRelation{})
	var hkCircleRelations []community.CircleRelation

	db = db.Where("circle_id = ? and relation_type = 0", circleId)
	if len(page.Keyword) > 0 {
		db = db.Where("other_circle_name LIKE ?", "%"+page.Keyword+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	//
	err = db.Limit(limit).Offset(offset).Find(&hkCircleRelations).Error
	size := len(hkCircleRelations)
	if err == nil && size > 0 {
		var ids = make([]uint64, size)
		for index, v := range hkCircleRelations {
			ids[index] = v.ID
		}
		err = global.GVA_DB.Model(&community.CircleBaseInfo{}).Where("id in ?", ids).Find(&list).Error
	}
	return list, total, err
}
