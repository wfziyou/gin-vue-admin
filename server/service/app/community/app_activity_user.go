package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"gorm.io/gorm"
)

type ActivityUserService struct {
}

// CreateActivityUser 创建ActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *ActivityUserService) CreateActivityUser(hkActivityUser community.ActivityUser) (err error) {
	err = global.GVA_DB.Create(&hkActivityUser).Error
	if err == nil {
		err = UpdateActivityUserNum(hkActivityUser.ActivityId)
	}
	return err
}
func (hkActivityUserService *ActivityUserService) AddActivityUser(activityUser community.ActivityUser) (err error) {
	var obj = community.ActivityUser{}
	err = global.GVA_DB.Where("user_id = ? and activity_id = ?", obj.UserId, obj.ActivityId).First(&obj).Error
	if err == nil {
		//		err = global.GVA_DB.Save(&obj).Error
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&activityUser).Error
		if err == nil {
			var activityUserNum int64
			err = global.GVA_DB.Model(&community.ActivityUser{}).Count(&activityUserNum).Error
			if err == nil {
				err = UpdateActivityUserNum(obj.ActivityId)
			}
		}
	}
	return err
}

// DeleteActivityUser 删除ActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *ActivityUserService) DeleteActivityUser(hkActivityUser community.ActivityUser) (err error) {
	err = global.GVA_DB.Delete(&hkActivityUser).Error
	if err == nil {
		var activityUserNum int64
		err = global.GVA_DB.Model(&community.ActivityUser{}).Count(&activityUserNum).Error
		if err == nil {
			err = UpdateActivityUserNum(hkActivityUser.ActivityId)
		}
	}
	return err
}

func (hkActivityUserService *ActivityUserService) DeleteActivityUserByUserIds(activityId uint64, userIds []uint64) (err error) {
	err = global.GVA_DB.Delete(&[]community.ActivityUser{}, "activity_id = ? AND user_id in ?", activityId, userIds).Error
	if err == nil {
		var activityUserNum int64
		err = global.GVA_DB.Model(&community.ActivityUser{}).Count(&activityUserNum).Error
		if err == nil {
			err = UpdateActivityUserNum(activityId)
		}
	}
	return err
}

// UpdateActivityUser 更新ActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *ActivityUserService) UpdateActivityUser(hkActivityUser community.ActivityUser) (err error) {
	err = global.GVA_DB.Save(&hkActivityUser).Error
	return err
}

// GetActivityUser 根据id获取ActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *ActivityUserService) GetActivityUser(activityId uint64, userId uint64) (hkActivityUser community.ActivityUser, err error) {
	err = global.GVA_DB.Where("activity_id = ? AND user_id = ?", activityId, userId).First(&hkActivityUser).Error
	return
}

// GetActivityUserInfoList 分页获取ActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *ActivityUserService) GetActivityUserInfoList(info communityReq.ActivityUserSearch) (list []community.ActivityUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ActivityUser{})
	var hkActivityUsers []community.ActivityUser

	if len(info.Keyword) > 0 {
		db = db.Where("remark LIKE ?", "%"+info.Keyword+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkActivityUsers).Error
	return hkActivityUsers, total, err
}
func UpdateActivityUserNum(activityId uint64) (err error) {
	var activityUserNum int64
	err = global.GVA_DB.Model(&community.ActivityUser{}).Count(&activityUserNum).Error
	if err == nil {
		db := global.GVA_DB.Model(community.ForumPosts{})
		var updateData map[string]interface{}
		updateData = make(map[string]interface{})
		updateData["activity_cur_user_num"] = activityUserNum
		err = db.Where("activity_id = ?", activityId).Updates(updateData).Error
	}
	return err
}
