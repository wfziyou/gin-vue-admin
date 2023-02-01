package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
)

type HkForumThumbsUpService struct {
}

// CreateHkForumThumbsUp 创建HkForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumThumbsUpService *HkForumThumbsUpService) CreateHkForumThumbsUp(hkForumThumbsUp community.HkForumThumbsUp) (err error) {
	err = global.GVA_DB.Create(&hkForumThumbsUp).Error
	return err
}

// DeleteHkForumThumbsUp 删除HkForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumThumbsUpService *HkForumThumbsUpService)DeleteHkForumThumbsUp(hkForumThumbsUp community.HkForumThumbsUp) (err error) {
	err = global.GVA_DB.Delete(&hkForumThumbsUp).Error
	return err
}

// DeleteHkForumThumbsUpByIds 批量删除HkForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumThumbsUpService *HkForumThumbsUpService)DeleteHkForumThumbsUpByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkForumThumbsUp{},"id in ?",ids.Ids).Error
	return err
}

// UpdateHkForumThumbsUp 更新HkForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumThumbsUpService *HkForumThumbsUpService)UpdateHkForumThumbsUp(hkForumThumbsUp community.HkForumThumbsUp) (err error) {
	err = global.GVA_DB.Save(&hkForumThumbsUp).Error
	return err
}

// GetHkForumThumbsUp 根据id获取HkForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumThumbsUpService *HkForumThumbsUpService)GetHkForumThumbsUp(id uint) (hkForumThumbsUp community.HkForumThumbsUp, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkForumThumbsUp).Error
	return
}

// GetHkForumThumbsUpInfoList 分页获取HkForumThumbsUp记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkForumThumbsUpService *HkForumThumbsUpService)GetHkForumThumbsUpInfoList(info communityReq.HkForumThumbsUpSearch) (list []community.HkForumThumbsUp, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&community.HkForumThumbsUp{})
    var hkForumThumbsUps []community.HkForumThumbsUp
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&hkForumThumbsUps).Error
	return  hkForumThumbsUps, total, err
}
