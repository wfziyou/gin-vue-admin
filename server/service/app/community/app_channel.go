package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
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

// GetChannelInfoList 分页获取HkChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkChannelService *HkChannelService) GetChannelInfoList() (list []community.ChannelInfo, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.ChannelInfo{})
	var hkChannels []community.ChannelInfo

	err = db.Find(&hkChannels).Error
	return hkChannels, total, err
}

// GetChannelInfoListById 分页获取HkChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkChannelService *HkChannelService) GetChannelInfoListById(ids string) (list []community.ChannelInfo, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.ChannelInfo{})
	var hkChannels []community.ChannelInfo

	err = db.Where("id in(?)", ids).Find(&hkChannels).Error
	return hkChannels, total, err
}
