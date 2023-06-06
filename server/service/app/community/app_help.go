package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HelpService struct {
}

// CreateHelp 创建Help记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *HelpService) CreateHelp(hkHelp *community.Help) (err error) {
	err = global.GVA_DB.Create(hkHelp).Error
	return err
}

// DeleteHelp 删除Help记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *HelpService) DeleteHelp(hkHelp community.Help) (err error) {
	err = global.GVA_DB.Delete(&hkHelp).Error
	return err
}

// DeleteHelpByIds 批量删除Help记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *HelpService) DeleteHelpByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.Help{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHelp 更新Help记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *HelpService) UpdateHelp(hkHelp community.Help) (err error) {
	err = global.GVA_DB.Save(&hkHelp).Error
	return err
}

// GetHelp 根据id获取Help记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *HelpService) GetHelp(id uint64) (hkHelp community.HelpBaseInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkHelp).Error
	return
}
func (service *HelpService) GetHelpInfo(id uint64) (hkHelp community.HelpInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkHelp).Error
	return
}

// GetHelpInfoList 获取Help记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *HelpService) GetHelpList(info communityReq.HelpSearch) (list []community.HelpBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.Help{})
	var hkFeedbacks []community.HelpBaseInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+info.Keyword+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkFeedbacks).Error
	return hkFeedbacks, total, err
}
func (service *HelpService) GetMainHelpList() (mainHelp community.MainHelp, err error) {
	err = global.GVA_DB.Model(&community.HelpType{}).Find(&mainHelp.HelpType).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(&community.Help{}).Where("top > 0").Find(&mainHelp.Help).Error
	return
}
func (service *HelpService) GetHelpThumbsUpList(userId uint64, helpId uint64) (list []community.HelpThumbsUp, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.HelpThumbsUp{})
	db = db.Where("user_id = ? AND help_id = ?", userId, helpId)
	var helpThumbsUp []community.HelpThumbsUp

	err = db.Find(&helpThumbsUp).Error
	return helpThumbsUp, err
}

func (service *HelpService) SetHelpThumbsUp(userId uint64, info communityReq.HelpThumbsUpReq) (err error) {
	var helpList []community.HelpThumbsUp
	err = global.GVA_DB.Model(&community.HelpThumbsUp{}).Where("user_id = ? and help_id = ?", userId, info.HelpId).Limit(1).Find(&helpList).Error
	if err != nil {
		return
	}
	if len(helpList) > 0 {
		helpList[0].ThumbsUp = info.ThumbsUp
		if info.ThumbsUp > 0 {
			global.GVA_DB.Model(&community.HelpThumbsUp{}).Where("help_id = ? and user_id = ?", info.HelpId, userId).Update("thumbs_up", info.ThumbsUp)
		} else {
			if helpList[0].ThumbsUp == 0 {
				err = global.GVA_DB.Unscoped().Where("help_id = ? and user_id = ?", info.HelpId, userId).Delete(&community.HelpThumbsUp{}).Error
			} else {
				global.GVA_DB.Model(&community.HelpThumbsUp{}).Save(&helpList[0])
			}
		}
	} else {
		if info.ThumbsUp > 0 {
			help := community.HelpThumbsUp{
				ThumbsUp: info.ThumbsUp,
				UserId:   userId,
				HelpId:   info.HelpId,
			}
			err = global.GVA_DB.Model(&community.HelpThumbsUp{}).Create(&help).Error
			if err == nil {
				service.UpdateHelpGoodNum(info.HelpId)
			}
		}
	}
	return err
}

func (service *HelpService) UpdateHelpGoodNum(helpId uint64) (err error) {
	var total int64 = 0
	db := global.GVA_DB.Model(&community.HelpThumbsUp{}).Where("help_id = ? and thumbs_up = 1", helpId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(&community.Help{}).Where("id = ?", helpId).Update("good_num", total).Error
	return err
}
func (service *HelpService) UpdateHelpBadNum(helpId uint64) (err error) {
	var total int64 = 0
	db := global.GVA_DB.Model(&community.HelpThumbsUp{}).Where("help_id = ? and thumbs_up = 2", helpId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(&community.Help{}).Where("id = ?", helpId).Update("bad_num", total).Error
	return err
}
