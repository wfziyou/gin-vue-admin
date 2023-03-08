package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HkChannelService struct {
}

// CreateHkChannel 创建HkChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkChannelService *HkChannelService) CreateHkChannel(hkChannel community.HkChannel) (err error) {
	err = global.GVA_DB.Create(&hkChannel).Error
	return err
}

// DeleteHkChannel 删除HkChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkChannelService *HkChannelService) DeleteHkChannel(hkChannel community.HkChannel) (err error) {
	err = global.GVA_DB.Delete(&hkChannel).Error
	return err
}

// DeleteHkChannelByIds 批量删除HkChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkChannelService *HkChannelService) DeleteHkChannelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkChannel{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkChannel 更新HkChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkChannelService *HkChannelService) UpdateHkChannel(hkChannel community.HkChannel) (err error) {
	err = global.GVA_DB.Save(&hkChannel).Error
	return err
}

// GetHkChannel 根据id获取HkChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkChannelService *HkChannelService) GetHkChannel(id uint64) (hkChannel community.HkChannel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkChannel).Error
	return
}

// GetHkChannelInfoList 分页获取HkChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkChannelService *HkChannelService) GetHkChannelInfoList(info communityReq.HkChannelSearch) (list []community.HkChannel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkChannel{})
	var hkChannels []community.HkChannel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkChannels).Error
	return hkChannels, total, err
}
