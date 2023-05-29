package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type AppCircleService struct {
}

// CreateCircle 创建Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) CreateCircle(hkCircle community.Circle) (circle community.Circle, err error) {
	err = global.GVA_DB.Create(&hkCircle).Error
	return circle, err
}

// DeleteCircle 删除Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) DeleteCircle(hkCircle community.Circle) (err error) {
	err = global.GVA_DB.Delete(&hkCircle).Error
	return err
}
func (appCircleService *AppCircleService) UnscopedDeleteCircle(hkCircle community.Circle) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&hkCircle).Error
	return err
}

// DeleteCircleByIds 批量删除Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) DeleteCircleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.Circle{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCircle 更新Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) UpdateCircle(req communityReq.UpdateCircleReq) (err error) {
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	if len(req.Name) > 0 {
		updateData["name"] = req.Name
	}
	if len(req.Logo) > 0 {
		updateData["logo"] = req.Logo
	}
	if len(req.Slogan) > 0 {
		updateData["slogan"] = req.Slogan
	}
	if len(req.Des) > 0 {
		updateData["des"] = req.Des
	}
	if len(req.Protocol) > 0 {
		updateData["protocol"] = req.Protocol
	}
	if len(req.CoverImage) > 0 {
		updateData["cover_image"] = req.CoverImage
	}
	if req.Process != nil {
		updateData["process"] = req.Process
	}
	if req.Property != nil {
		updateData["property"] = req.Property
	}
	if req.View != nil {
		updateData["view"] = req.View
	}
	if req.PowerAdd != nil {
		updateData["power_add"] = req.PowerAdd
	}
	if req.PowerView != nil {
		updateData["power_view"] = req.PowerView
	}
	if req.PowerPublish != nil {
		updateData["power_publish"] = req.PowerPublish
	}
	if req.PowerComment != nil {
		updateData["power_comment"] = req.PowerComment
	}
	if len(req.PowerAddUser) > 0 {
		updateData["power_add_user"] = req.PowerAddUser
	}
	if len(req.PowerViewUser) > 0 {
		updateData["power_view_user"] = req.PowerViewUser
	}
	if len(req.PowerPublishUser) > 0 {
		updateData["power_publish_user"] = req.PowerPublishUser
	}
	if len(req.PowerCommentUser) > 0 {
		updateData["power_comment_user"] = req.PowerCommentUser
	}
	if len(req.NoLimitUserGroup) > 0 {
		updateData["no_limit_user_group"] = req.NoLimitUserGroup
	}
	if req.NewUserFocus != nil {
		updateData["new_user_focus"] = req.NewUserFocus
	}

	db := global.GVA_DB.Model(&community.CircleBaseInfo{})
	err = db.Where("id = ?", req.ID).Updates(updateData).Error
	return err
}
func (appCircleService *AppCircleService) UpdateCircleBaseInfo(req communityReq.UpdateCircleBaseInfoReq) (err error) {
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	if len(req.Name) > 0 {
		updateData["name"] = req.Name
	}
	if len(req.Logo) > 0 {
		updateData["logo"] = req.Logo
	}
	if len(req.Slogan) > 0 {
		updateData["slogan"] = req.Slogan
	}
	if len(req.Des) > 0 {
		updateData["des"] = req.Des
	}
	if len(req.Protocol) > 0 {
		updateData["protocol"] = req.Protocol
	}
	if len(req.CoverImage) > 0 {
		updateData["cover_image"] = req.CoverImage
	}

	db := global.GVA_DB.Model(&community.Circle{})
	err = db.Where("id = ?", req.ID).Updates(updateData).Error
	return err
}
func (appCircleService *AppCircleService) UpdateCirclePower(req communityReq.UpdateCirclePowerReq) (err error) {
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	if req.Process != nil {
		updateData["process"] = req.Process
	}
	if req.Property != nil {
		updateData["property"] = req.Property
	}
	if req.View != nil {
		updateData["view"] = req.View
	}
	if req.PowerAdd != nil {
		updateData["power_add"] = req.PowerAdd
	}
	if req.PowerView != nil {
		updateData["power_view"] = req.PowerView
	}
	if req.PowerPublish != nil {
		updateData["power_publish"] = req.PowerPublish
	}
	if req.PowerComment != nil {
		updateData["power_comment"] = req.PowerComment
	}
	if len(req.PowerAddUser) > 0 {
		updateData["power_add_user"] = req.PowerAddUser
	}
	if len(req.PowerViewUser) > 0 {
		updateData["power_view_user"] = req.PowerViewUser
	}
	if len(req.PowerPublishUser) > 0 {
		updateData["power_publish_user"] = req.PowerPublishUser
	}
	if len(req.PowerCommentUser) > 0 {
		updateData["power_comment_user"] = req.PowerCommentUser
	}
	if len(req.NoLimitUserGroup) > 0 {
		updateData["no_limit_user_group"] = req.NoLimitUserGroup
	}
	if req.NewUserFocus != nil {
		updateData["new_user_focus"] = req.NewUserFocus
	}

	db := global.GVA_DB.Model(&community.Circle{})
	err = db.Where("id = ?", req.ID).Updates(updateData).Error
	return err
}

func (appCircleService *AppCircleService) FindCircle(id uint64) (hkCircle community.Circle, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Tags").First(&hkCircle).Error
	if err == nil {
		var hkCircleTags []community.CircleTag
		if err := global.GVA_DB.Model(&community.CircleTag{}).Where("circle_id = ?", id).Find(&hkCircleTags).Error; err == nil {
			hkCircle.Tags = hkCircleTags
		}
	}
	return
}

// GetCircle 根据id获取Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) GetCircle(id uint64) (hkCircle community.Circle, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Tags").First(&hkCircle).Error
	return
}
func (appCircleService *AppCircleService) GetCircleInfo(id uint64) (hkCircle community.CircleInfo, err error) {
	err = global.GVA_DB.Model(&community.Circle{}).Where("id = ?", id).First(&hkCircle).Error
	return
}
func (appCircleService *AppCircleService) GetCirclePower(id uint64) (hkCircle community.CirclePower, err error) {
	err = global.GVA_DB.Model(&community.Circle{}).Where("id = ?", id).First(&hkCircle).Error
	return
}

// GetCircleInfoList 分页获取Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) GetCircleInfoList(info communityReq.CircleSearch) (list []community.CircleBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleBaseInfo{})
	var hkCircles []community.CircleBaseInfo
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.CircleClassifyId != 0 {
		db = db.Where("circle_classify_id = ?", info.CircleClassifyId)
	}
	if len(info.Keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+info.Keyword+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircles).Error
	return hkCircles, total, err
}

func (appCircleService *AppCircleService) GetSelfCircleList(userId uint64, info communityReq.SelfCircleSearch) (list []community.CircleBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleUser{})
	var circleUser []community.CircleUser

	db = db.Where("user_id = ?", userId)
	if info.Power != nil {
		db = db.Where("power = ?", info.Power)
	}
	if len(info.Keyword) > 0 {
		db = db.Where("circle_name LIKE ?", "%"+info.Keyword+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if info.Power == nil {
		db = db.Order("sort desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&circleUser).Error
	var circleIds = make([]uint64, len(circleUser))
	if err != nil {
		return
	}

	for index, item := range circleUser {
		circleIds[index] = item.CircleId
	}
	var hkCircles []community.CircleBaseInfo
	err = global.GVA_DB.Where("id in ?", circleIds).Find(&hkCircles).Error

	return hkCircles, total, err
}

// UpdateCircleChannel 更新用户频道
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) UpdateCircleChannel(circleId uint64, channel string) (err error) {
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	updateData["channel_id"] = channel
	err = global.GVA_DB.Model(&community.Circle{}).Where("id = ?", circleId).Updates(updateData).Error
	return err
}
func (appCircleService *AppCircleService) AddCircleChannel(circleId uint64, name string) (channel community.CircleChannel, err error) {
	err = global.GVA_DB.Where("circle_id = ? AND name = ?", circleId, name).First(&channel).Error
	if err == nil {
		return channel, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	var tmp = community.CircleChannel{
		CircleId: circleId,
		Name:     name,
	}

	err = global.GVA_DB.Create(&tmp).Error
	return tmp, err
}
