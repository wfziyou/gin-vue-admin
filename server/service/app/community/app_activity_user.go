package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ActivityUserService struct {
}

// CreateActivityUser 创建ActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *ActivityUserService) CreateActivityUser(hkActivityUser community.ActivityUser) (err error) {
	err = global.GVA_DB.Create(&hkActivityUser).Error
	return err
}

// DeleteActivityUser 删除ActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *ActivityUserService) DeleteActivityUser(hkActivityUser community.ActivityUser) (err error) {
	err = global.GVA_DB.Delete(&hkActivityUser).Error
	return err
}

// DeleteActivityUserByIds 批量删除ActivityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkActivityUserService *ActivityUserService) DeleteActivityUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ActivityUser{}, "id in ?", ids.Ids).Error
	return err
}

func (hkActivityUserService *ActivityUserService) DeleteActivityUserByUserIds(activityId uint64, userIds []uint64) (err error) {
	err = global.GVA_DB.Delete(&[]community.ActivityUser{}, "activity_id=? and user_id in ?", activityId, userIds).Error
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
func (hkActivityUserService *ActivityUserService) GetActivityUser(id uint64) (hkActivityUser community.ActivityUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkActivityUser).Error
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

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkActivityUsers).Error
	return hkActivityUsers, total, err
}
