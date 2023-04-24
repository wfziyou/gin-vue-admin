package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCircleUserService struct {
}

// CreateCircleUser 创建CircleUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleUserService *AppCircleUserService) CreateCircleUser(hkCircleUser community.CircleUser) (err error) {
	err = global.GVA_DB.Create(&hkCircleUser).Error
	return err
}

// DeleteCircleUser 删除圈子用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleUserService *AppCircleUserService) DeleteCircleUser(hkCircleUser community.CircleUser) (err error) {
	err = global.GVA_DB.Delete(&hkCircleUser).Error
	return err
}
func (appCircleUserService *AppCircleUserService) DeleteCircleUserInfo(hkCircleUser community.CircleUserInfo) (err error) {
	err = global.GVA_DB.Delete(&hkCircleUser).Error
	return err
}
func (appCircleUserService *AppCircleUserService) DeleteCircleUsers(circleId uint64, userIds []uint64) (err error) {
	err = global.GVA_DB.Delete(&[]community.CircleUser{}, "circle_id = ? and user_id in ?", circleId, userIds).Error
	return err
}

// DeleteCircleUserByIds 批量删除圈子用户记录
func (appCircleUserService *AppCircleUserService) DeleteCircleUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.CircleUser{}, "id in ?", ids.Ids).Error
	return err
}
func (appCircleUserService *AppCircleUserService) GetUserHaveCircles(userId uint64, circleIds []uint64) (hkCircleUser []community.CircleUser, num int, err error) {
	err = global.GVA_DB.Where("user_id = ? and circle_id in  ?", userId, circleIds).Find(&hkCircleUser).Error
	if err == nil {
		num = len(hkCircleUser)
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

// UpdateCircleUser 更新圈子用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleUserService *AppCircleUserService) UpdateCircleUser(info communityReq.UpdateCircleUserReq) (err error) {
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	if len(info.Remark) > 0 {
		updateData["remark"] = info.Remark
	}
	if len(info.Tag) > 0 {
		updateData["tag"] = info.Tag
	}
	if info.Power != nil {
		updateData["power"] = info.Power
	}
	if info.Sort != nil {
		updateData["sort"] = info.Sort
	}
	err = global.GVA_DB.Model(&community.CircleUser{}).Where("circle_id = ? AND user_id = ?", info.CircleId, info.UserId).Updates(updateData).Error
	return err
}

// GetCircleUser 根据id获取CircleUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleUserService *AppCircleUserService) GetCircleUser(id uint64) (hkCircleUser community.CircleUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleUser).Error
	return
}

func (appCircleUserService *AppCircleUserService) GetCircleUserEx(circleId uint64, userId uint64) (hkCircleUser community.CircleUser, err error) {
	err = global.GVA_DB.Where("circle_id = ? and user_id = ?", circleId, userId).First(&hkCircleUser).Error
	return
}

func (appCircleUserService *AppCircleUserService) GetCircleUserInfo(circleId uint64, userId uint64) (hkCircleUser community.CircleUserInfo, err error) {
	err = global.GVA_DB.Where("circle_id = ? and user_id = ?", circleId, userId).Preload("UserInfo").First(&hkCircleUser).Error
	return
}

// SetUserCurCircle 根据id获取CircleUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleUserService *AppCircleUserService) SetUserCurCircle(userId uint64, circleId uint64) (err error) {
	var userExtend community.UserExtend
	userExtend.ID = userId
	err = global.GVA_DB.Where("id = ?", userId).First(&userExtend).Error
	if err == nil {
		db := global.GVA_DB.Model(&userExtend)
		err = db.Where("id = ?", userId).Update("circle_id", circleId).Error
	} else {
		err = global.GVA_DB.Save(&userExtend).Error
		return err
	}
	return err
}

// GetCircleUserInfoList 分页获取CircleUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleUserService *AppCircleUserService) GetCircleUserInfoList(info communityReq.CircleUserSearch) (list []community.CircleUserInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleUserInfo{}).Preload("UserInfo")
	var hkCircleUsers []community.CircleUserInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	db = db.Where("circle_id = ? AND user_id = ?", info.CircleId, info.UserId)
	if info.CircleId != 0 {

	}
	if len(info.Keyword) > 0 {
		db = db.Where("remark = ?", "%"+info.Keyword+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircleUsers).Error
	return hkCircleUsers, total, err
}

func (appCircleUserService *AppCircleUserService) GetCircleUserInfoListCount(info communityReq.CircleUserSearch) (total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleUser{})
	err = db.Where(&community.CircleUser{CircleId: info.CircleId, UserId: info.UserId}).Limit(limit).Offset(offset).Count(&total).Error

	return total, err
}
