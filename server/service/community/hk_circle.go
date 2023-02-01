package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type HkCircleService struct {
}

// CreateHkCircle 创建HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleService *HkCircleService) CreateHkCircle(hkCircle community.HkCircle) (err error) {
	err = global.GVA_DB.Create(&hkCircle).Error
	return err
}

// DeleteHkCircle 删除HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleService *HkCircleService)DeleteHkCircle(hkCircle community.HkCircle) (err error) {
	err = global.GVA_DB.Delete(&hkCircle).Error
	return err
}

// DeleteHkCircleByIds 批量删除HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleService *HkCircleService)DeleteHkCircleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkCircle{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkCircle 更新HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleService *HkCircleService)UpdateHkCircle(hkCircle community.HkCircle) (err error) {
	err = global.GVA_DB.Save(&hkCircle).Error
	return err
}

// GetHkCircle 根据id获取HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleService *HkCircleService)GetHkCircle(id uint) (hkCircle community.HkCircle, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircle).Error
	return
}

// GetHkCircleInfoList 分页获取HkCircle记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkCircleService *HkCircleService)GetHkCircleInfoList(info communityReq.HkCircleSearch) (list []community.HkCircle, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&community.HkCircle{})
    var hkCircles []community.HkCircle
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkCircles).Error
	return  hkCircles, total, err
}
