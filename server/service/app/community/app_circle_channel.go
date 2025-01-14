package community

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
)

type AppCircleChannelService struct {
}

// CreateCircleChannel 创建CircleChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (circleChannelService *AppCircleChannelService) CreateCircleChannel(hkChannel community.CircleChannel) (err error) {
	err = global.GVA_DB.Create(&hkChannel).Error
	return err
}

// DeleteCircleChannelByIds 批量删除CircleChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (circleChannelService *AppCircleChannelService) DeleteCircleChannelByIds(circleId uint64, ids []uint64) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]community.CircleChannel{}, "circle_id = ? AND id in ?", circleId, ids).Error
	return err
}

// UpdateCircleChannel 更新CircleChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (circleChannelService *AppCircleChannelService) UpdateCircleChannel(hkChannel community.CircleChannel) (err error) {
	err = global.GVA_DB.Save(&hkChannel).Error
	return err
}

// GetCircleChannel 根据id获取CircleChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (circleChannelService *AppCircleChannelService) GetCircleChannel(id uint64) (hkChannel community.CircleChannelInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkChannel).Error
	return
}

// GetChannelInfoList 分页获取CircleChannel记录
func (circleChannelService *AppCircleChannelService) GetChannelInfoList() (list []community.CircleChannelInfo, total int64, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.CircleChannelInfo{})
	var hkChannels []community.CircleChannelInfo

	err = db.Find(&hkChannels).Error
	return hkChannels, total, err
}

// GetChannelInfoListById 分页获取CircleChannel记录
func (circleChannelService *AppCircleChannelService) GetChannelInfoListById(ids string) (list []community.CircleChannelInfo, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.CircleChannelInfo{})
	var hkChannels []community.CircleChannelInfo
	sql := fmt.Sprintf("id in(%s)", ids)
	err = db.Where(sql).Find(&hkChannels).Error
	return hkChannels, err
}

//GetDefaultChannelInfoList 获取默认频道列表
func (circleChannelService *AppCircleChannelService) GetDefaultChannelInfoList() (list []community.CircleChannelInfo, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.CircleChannel{})
	var hkChannels []community.CircleChannelInfo
	err = db.Where("circle_id = 0").Order("id").Find(&hkChannels).Error
	return hkChannels, err
}

func (circleChannelService *AppCircleChannelService) GetCircleChannelListAll(circleId uint64) (list []community.CircleChannelInfoEx, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.CircleChannel{})
	var hkChannels []community.CircleChannelInfoEx
	err = db.Where("circle_id = 0 OR circle_id = ?", circleId).Order("id").Find(&hkChannels).Error
	return hkChannels, err
}
