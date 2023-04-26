package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FocusUserService struct {
}

// CreateFocusUser 创建FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) CreateFocusUser(userId uint64, userNickName string, focusUser community.User, mutual int) (err error) {
	hkFocusUser := community.FocusUser{UserId: userId, NickName: userNickName, FocusUserId: focusUser.ID, FocusNickName: focusUser.NickName, Mutual: mutual}
	err = global.GVA_DB.Where("user_id = ? and focus_user_id = ?", userId, focusUser.ID).First(&hkFocusUser).Error
	err = global.GVA_DB.Save(&hkFocusUser).Error
	return err
}

// DeleteFocusUser 删除FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) DeleteFocusUser(userId uint64, focusUserId uint64) (err error) {
	err = global.GVA_DB.Delete(&[]community.FocusUser{}, "user_id = ? and focus_user_id = ?", userId, focusUserId).Error
	return err
}
func (hkFocusUserService *FocusUserService) DeleteFocusUserObj(info community.FocusUser) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&info).Error
	return err
}

// DeleteFocusUserByIds 批量删除FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) DeleteFocusUserByIds(Ids []uint64) (err error) {
	err = global.GVA_DB.Delete(&[]community.FocusUser{}, "id in ?", Ids).Error
	return err
}

// UpdateFocusUser 更新FocusUser记录
func (hkFocusUserService *FocusUserService) UpdateFocusUser(hkFocusUser community.FocusUser) (err error) {
	err = global.GVA_DB.Save(&hkFocusUser).Error
	return err
}

func (hkFocusUserService *FocusUserService) UpdateFocusUserEx(userId uint64, info communityReq.UpdateFocusUserReq) (err error) {
	db := global.GVA_DB.Model(&community.FocusUser{})
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	if len(info.Remark) > 0 {
		updateData["remark"] = info.Remark
	}
	if len(info.Remark) > 0 {
		updateData["tag"] = info.Tag
	}
	err = db.Where("user_id = ? AND focus_user_id = ?", userId, info.UserId).Updates(updateData).Error
	return err
}

// GetFocusUser 根据id获取FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) GetFocusUser(userId uint64, focusUserId uint64) (focusUserInfo community.FocusUserInfo, err error) {
	focusUser := community.FocusUser{}
	err = global.GVA_DB.Where("user_id = ? AND focus_user_id = ?", userId, focusUserId).First(&focusUser).Error
	if err == nil {
		user := community.User{}
		err = global.GVA_DB.Where("id = ?", focusUserId).Preload("UserExtend").First(&user).Error
		focusUserInfo.ID = focusUserId
		focusUserInfo.Mutual = focusUser.Mutual
		if err == nil {
			focusUserInfo.UserType = user.UserType
			focusUserInfo.Account = user.Account
			focusUserInfo.NickName = user.NickName
			focusUserInfo.RealName = user.RealName
			focusUserInfo.HeaderImg = user.HeaderImg
			focusUserInfo.Sex = user.Sex
			focusUserInfo.Description = user.Description
			focusUserInfo.NumCircle = user.UserExtend.NumCircle
			focusUserInfo.NumFocus = user.UserExtend.NumFocus
			focusUserInfo.NumFan = user.UserExtend.NumFan
		}
	}
	return
}
func (hkFocusUserService *FocusUserService) GetFocusUserObj(userId uint64, focusUserId uint64) (focusUser community.FocusUser, err error) {
	err = global.GVA_DB.Where("user_id = ? AND focus_user_id = ?", userId, focusUserId).First(&focusUser).Error
	return
}

// GetFocusUserInfoList 分页获取FocusUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkFocusUserService *FocusUserService) GetFocusUserInfoList(userId uint64, info communityReq.FocusUserSearch) (list []community.FocusUserInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.FocusUser{}).Select("focus_user_id as id,mutual")
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("user_id = ?", userId)
	if len(info.Keyword) > 0 {
		db = db.Where("focus_nick_name LIKE ?", "%"+info.Keyword+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	var hkFocusUsers []community.FocusUserInfo
	err = db.Limit(limit).Offset(offset).Find(&hkFocusUsers).Error
	if err == nil {
		var size = len(hkFocusUsers)
		if size > 0 {
			var ids = make([]uint64, size)
			for index, v := range hkFocusUsers {
				ids[index] = v.ID
			}

			var users []community.User
			err = global.GVA_DB.Where("id in ?", ids).Preload("UserExtend").Find(&users).Error

			if err == nil && len(users) > 0 {
				for _, user := range users {
					for i, obj := range hkFocusUsers {
						if obj.ID == user.ID {
							hkFocusUsers[i].UserType = user.UserType
							hkFocusUsers[i].Account = user.Account
							hkFocusUsers[i].NickName = user.NickName
							hkFocusUsers[i].RealName = user.RealName
							hkFocusUsers[i].HeaderImg = user.HeaderImg
							hkFocusUsers[i].Sex = user.Sex
							hkFocusUsers[i].Description = user.Description
							hkFocusUsers[i].NumCircle = user.UserExtend.NumCircle
							hkFocusUsers[i].NumFocus = user.UserExtend.NumFocus
							hkFocusUsers[i].NumFan = user.UserExtend.NumFan
							break
						}
					}
				}
			}
		}

	}
	return hkFocusUsers, total, err
}
func (hkFocusUserService *FocusUserService) GetFansList(userId uint64, page request.PageInfo) (list []community.FocusUserInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.FocusUser{}).Select("user_id as id,mutual")

	db = db.Where("focus_user_id = ?", userId)
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(page.Keyword) > 0 {
		db = db.Where("nick_name LIKE ?", "%"+page.Keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	var hkFocusUsers []community.FocusUserInfo
	err = db.Limit(limit).Offset(offset).Find(&hkFocusUsers).Error
	if err == nil {
		var size = len(hkFocusUsers)
		if size > 0 {
			var ids = make([]uint64, size)
			for index, v := range hkFocusUsers {
				ids[index] = v.ID
			}

			var users []community.User
			err = global.GVA_DB.Where("id in ?", ids).Preload("UserExtend").Find(&users).Error

			if err == nil && len(users) > 0 {
				for _, user := range users {
					for i, obj := range hkFocusUsers {
						if obj.ID == user.ID {
							hkFocusUsers[i].UserType = user.UserType
							hkFocusUsers[i].Account = user.Account
							hkFocusUsers[i].NickName = user.NickName
							hkFocusUsers[i].RealName = user.RealName
							hkFocusUsers[i].HeaderImg = user.HeaderImg
							hkFocusUsers[i].Sex = user.Sex
							hkFocusUsers[i].Description = user.Description
							hkFocusUsers[i].NumCircle = user.UserExtend.NumCircle
							hkFocusUsers[i].NumFocus = user.UserExtend.NumFocus
							hkFocusUsers[i].NumFan = user.UserExtend.NumFan
							break
						}
					}
				}
			}
		}

	}
	return hkFocusUsers, total, err
}
