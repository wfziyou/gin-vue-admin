package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
)

type FocusUserService struct {
}

// CreateFocusUser 创建FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) CreateFocusUser(hkFocusUser community.FocusUser) (err error) {
	err = global.GVA_DB.Create(&hkFocusUser).Error
	return err
}

// DeleteFocusUser 删除FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) DeleteFocusUser(userId uint64, focusUserId uint64) (err error) {
	err = global.GVA_DB.Delete(&[]community.FocusUser{}, "user_id = ? and focus_user_id = ?", userId, focusUserId).Error
	return err
}

// DeleteFocusUserByIds 批量删除FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) DeleteFocusUserByIds(Ids []uint64) (err error) {
	err = global.GVA_DB.Delete(&[]community.FocusUser{}, "id in ?", Ids).Error
	return err
}

// UpdateFocusUser 更新FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) UpdateFocusUser(hkFocusUser community.FocusUser) (err error) {
	err = global.GVA_DB.Save(&hkFocusUser).Error
	return err
}

// GetFocusUser 根据id获取FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) GetFocusUser(userId uint64, focusUserId uint64) (focusUserInfo community.FocusUserInfo, err error) {
	focusUser := community.FocusUser{}
	err = global.GVA_DB.Where("userId = ?,focusUserId = ?", userId, focusUserId).First(&focusUser).Error
	if err == nil {
		userBaseInfo := community.UserBaseInfo{}
		err = global.GVA_DB.Where("id = ?", focusUserId).First(&userBaseInfo).Error
		if err != nil {
			focusUserInfo.Id = focusUser.ID
			focusUserInfo.FocusUserId = focusUser.FocusUserId
			focusUserInfo.Remark = focusUser.Remark
			focusUserInfo.Tag = focusUser.Tag
			focusUserInfo.Account = userBaseInfo.Account
			focusUserInfo.RealName = userBaseInfo.RealName
			focusUserInfo.HeaderImg = userBaseInfo.HeaderImg
			focusUserInfo.Sex = userBaseInfo.Sex
			focusUserInfo.Description = userBaseInfo.Description
		}
	}
	return
}

// GetFocusUserInfoList 分页获取FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) GetFocusUserInfoList(info communityReq.FocusUserSearch) (list []community.FocusUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.FocusUser{})
	var hkFocusUsers []community.FocusUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkFocusUsers).Error
	return hkFocusUsers, total, err
}
