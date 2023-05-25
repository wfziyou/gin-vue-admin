package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"gorm.io/gorm"
)

type AppCircleUserService struct {
}

func (appCircleUserService *AppCircleUserService) CreateCircleUser(circleUser community.CircleUser) (err error) {
	err = global.GVA_DB.Create(&circleUser).Error
	if err == nil {
		err = appCircleUserService.UpdateCircleUserCount(circleUser.CircleId)
		appCircleUserService.UpdateUserNumCircle(circleUser.UserId)
	}
	return err
}
func (appCircleUserService *AppCircleUserService) DeleteCircleUser(circleUser community.CircleUser) (err error) {
	err = global.GVA_DB.Delete(&circleUser).Error
	if err == nil {
		err = appCircleUserService.UpdateCircleUserCount(circleUser.CircleId)
		appCircleUserService.UpdateUserNumCircle(circleUser.UserId)
	}
	return err
}
func (appCircleUserService *AppCircleUserService) DeleteCircleAllUser(circleId uint64) (err error) {
	var circleUser []community.CircleUser
	err = global.GVA_DB.Model(&community.CircleUser{}).Where("circle_id = ?", circleId).Find(&circleUser).Error
	if err == nil && len(circleUser) > 0 {
		err = global.GVA_DB.Delete(&[]community.CircleUser{}, "circle_id = ?", circleId).Error
		if err == nil {
			for _, obj := range circleUser {
				appCircleUserService.UpdateUserNumCircle(obj.UserId)
			}
		}
	}
	return err
}
func (appCircleUserService *AppCircleUserService) DeleteCircleUsers(circleId uint64, userIds []uint64) (err error) {
	if len(userIds) == 0 {
		return nil
	}
	err = global.GVA_DB.Delete(&[]community.CircleUser{}, "circle_id = ? and user_id in ?", circleId, userIds).Error
	if err == nil {
		err = appCircleUserService.UpdateCircleUserCount(circleId)
		for _, userId := range userIds {
			appCircleUserService.UpdateUserNumCircle(userId)
		}
	}
	return err
}
func (appCircleUserService *AppCircleUserService) GetUserHaveCircles(userId uint64, circleIds []uint64) (circleUser []community.CircleUser, num int, err error) {
	err = global.GVA_DB.Where("user_id = ? and circle_id in  ?", userId, circleIds).Find(&circleUser).Error
	if err == nil {
		num = len(circleUser)
	}
	return
}
func (appCircleUserService *AppCircleUserService) GetUserHaveCircle(userId uint64, list []community.CircleBaseInfo) (err error) {
	var size = len(list)
	if size > 0 {
		var ids = make([]uint64, size)
		for index, v := range list {
			ids[index] = v.ID
		}
		if data, num, err := appCircleUserService.GetUserHaveCircles(userId, ids); err == nil && num > 0 {
			for _, v := range data {
				for i, obj := range list {
					if obj.ID == v.CircleId {
						list[i].HaveCircle = 1
						break
					}
				}
			}
		}
	}
	return
}
func (appCircleUserService *AppCircleUserService) GetCircleUser(circleId uint64, userId uint64) (circleUser community.CircleUser, err error) {
	err = global.GVA_DB.Where("circle_id = ? and user_id = ?", circleId, userId).First(&circleUser).Error
	return
}
func (appCircleUserService *AppCircleUserService) SelfCircleTop(circleId uint64, userId uint64) (err error) {
	db := global.GVA_DB.Model(&community.CircleUser{})
	var maxSort int
	err = db.Where("user_id = ?", userId).Select("MAX(sort)").Find(&maxSort).Error
	if err != nil {
		return err
	}
	maxSort += 1
	err = global.GVA_DB.Model(community.CircleUser{}).Where("circle_id = ? AND user_id = ?", circleId, userId).Update("sort", maxSort).Error
	return err
}
func (appCircleUserService *AppCircleUserService) SelfCircleCancelTop(circleId uint64, userId uint64) (err error) {
	err = global.GVA_DB.Model(community.CircleUser{}).Where("circle_id = ? AND user_id = ?", circleId, userId).Update("sort", 0).Error
	return err
}

func (appCircleUserService *AppCircleUserService) GetCircleUserInfo(selectUserId uint64, circleId uint64, userId uint64) (circleUserInfo community.CircleUserInfo, err error) {
	var circleUser community.CircleUser
	err = global.GVA_DB.Where("circle_id = ? and user_id = ?", circleId, userId).Preload("UserInfo").First(&circleUser).Error
	if err == nil {
		isFocus, isFan, _ := GetIsFocusAndIsFan(selectUserId, &circleUser.UserInfo)
		circleUser.UserInfo.IsFocus = isFocus
		circleUser.UserInfo.IsFan = isFan

		circleUserInfo.Remark = circleUser.Remark
		circleUserInfo.Tag = circleUser.Tag
		circleUserInfo.Power = circleUser.Power
		circleUserInfo.UserBaseInfo = circleUser.UserInfo
		return circleUserInfo, err
	}
	return
}
func (appCircleUserService *AppCircleUserService) GetCircleUserInfoList(selectUserId uint64, info communityReq.CircleUserSearch) (list []community.CircleUserInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleUser{}).Preload("UserInfo")
	var hkCircleUsers []community.CircleUser

	db = db.Where("circle_id = ?", info.CircleId)

	if info.Power != nil && *info.Power == community.CircleUserPowerManager {
		db = db.Where("power = ?", info.Power)
	}
	if len(info.Keyword) > 0 {
		db = db.Where("remark = ?", "%"+info.Keyword+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	var tmp = make([]community.CircleUserInfo, 0, len(hkCircleUsers))
	err = db.Limit(limit).Offset(offset).Find(&hkCircleUsers).Error
	if err == nil {
		SetCircleUserUserIsFocus(selectUserId, hkCircleUsers)
		for _, obj := range hkCircleUsers {
			tmp = append(tmp, community.CircleUserInfo{
				Remark:       obj.Remark,
				Tag:          obj.Tag,
				Power:        obj.Power,
				UserBaseInfo: obj.UserInfo,
			})
		}
	}
	return tmp, total, err
}
func (appCircleUserService *AppCircleUserService) GetCircleUserCount(circleId uint64) (total int64, err error) {
	db := global.GVA_DB.Model(&community.CircleUser{}).Where("circle_id = ?", circleId)
	err = db.Where("circle_id = ?", circleId).Count(&total).Error
	return total, err
}
func (appCircleUserService *AppCircleUserService) UpdateCircleUser(info communityReq.UpdateCircleUserReq) (err error) {
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	if len(info.Remark) > 0 {
		updateData["remark"] = info.Remark
	}
	if len(info.Tag) > 0 {
		updateData["tag"] = info.Tag
	}

	if info.Sort != nil {
		updateData["sort"] = info.Sort
	}
	err = global.GVA_DB.Model(&community.CircleUser{}).Where("circle_id = ? AND user_id = ?", info.CircleId, info.UserId).Updates(updateData).Error
	return err
}
func (appCircleUserService *AppCircleUserService) SetCircleUserPower(info communityReq.SetCircleUserPowerReq) (err error) {
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	updateData["power"] = info.Power
	err = global.GVA_DB.Model(&community.CircleUser{}).Where("circle_id = ? AND user_id = ?", info.CircleId, info.UserId).Updates(updateData).Error
	return err
}

func (appCircleUserService *AppCircleUserService) UpdateCircleUserCount(circleId uint64) (err error) {
	num, err := appCircleUserService.GetCircleUserCount(circleId)
	if err == nil {
		err = global.GVA_DB.Model(community.Circle{}).Where("id = ?", circleId).Update("user_num", num).Error
	}
	return err
}
func (appCircleUserService *AppCircleUserService) UpdateUserNumCircle(userId uint64) (err error) {
	numCircle := int64(0)
	err = global.GVA_DB.Model(&community.CircleUser{}).Where("user_id = ?", userId).Count(&numCircle).Error
	if err != nil {
		return
	}

	obj := community.UserExtend{GvaModelApp: global.GvaModelApp{ID: userId}, NumCircle: numCircle}
	db := global.GVA_DB.Model(&obj)
	if errors.Is(db.Where("id = ?", userId).First(&obj).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&obj).Error
	} else {
		var updateData map[string]interface{}
		updateData = make(map[string]interface{})
		updateData["num_circle"] = numCircle
		err = db.Where("id = ?", userId).Updates(updateData).Error
	}
	return err
}
