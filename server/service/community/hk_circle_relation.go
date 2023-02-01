package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type HkCircleRelationService struct {
}

// CreateHkCircleRelation 创建HkCircleRelation记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleRelationService *HkCircleRelationService) CreateHkCircleRelation(hkCircleRelation community.HkCircleRelation) (err error) {
	err = global.GVA_DB.Create(&hkCircleRelation).Error
	return err
}

// DeleteHkCircleRelation 删除HkCircleRelation记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleRelationService *HkCircleRelationService)DeleteHkCircleRelation(hkCircleRelation community.HkCircleRelation) (err error) {
	err = global.GVA_DB.Delete(&hkCircleRelation).Error
	return err
}

// DeleteHkCircleRelationByIds 批量删除HkCircleRelation记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleRelationService *HkCircleRelationService)DeleteHkCircleRelationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkCircleRelation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkCircleRelation 更新HkCircleRelation记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleRelationService *HkCircleRelationService)UpdateHkCircleRelation(hkCircleRelation community.HkCircleRelation) (err error) {
	err = global.GVA_DB.Save(&hkCircleRelation).Error
	return err
}

// GetHkCircleRelation 根据id获取HkCircleRelation记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleRelationService *HkCircleRelationService)GetHkCircleRelation(id uint) (hkCircleRelation community.HkCircleRelation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleRelation).Error
	return
}

// GetHkCircleRelationInfoList 分页获取HkCircleRelation记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleRelationService *HkCircleRelationService)GetHkCircleRelationInfoList(info communityReq.HkCircleRelationSearch) (list []community.HkCircleRelation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&community.HkCircleRelation{})
    var hkCircleRelations []community.HkCircleRelation
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkCircleRelations).Error
	return  hkCircleRelations, total, err
}
