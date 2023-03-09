package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCircleService struct {
}

// CreateCircle 创建Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) CreateCircle(hkCircle community.Circle) (err error) {
	err = global.GVA_DB.Create(&hkCircle).Error
	return err
}

// DeleteCircle 删除Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) DeleteCircle(hkCircle community.Circle) (err error) {
	err = global.GVA_DB.Delete(&hkCircle).Error
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
	if len(req.BackImage) > 0 {
		updateData["back_image"] = req.BackImage
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

// GetCircle 根据id获取Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) GetCircle(id uint64) (hkCircle community.Circle, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircle).Error
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
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.CircleClassifyId != 0 {
		db = db.Where("circle_classify_id = ?", info.CircleClassifyId)
	}
	if len(info.Name) > 0 {
		db = db.Where("name LIKE '%?%'", info.Name)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircles).Error
	return hkCircles, total, err
}

// GetSelfCircleList 分页获取Circle记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleService *AppCircleService) GetSelfCircleList(info communityReq.SelfCircleSearch) (list []community.CircleBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CircleUser{})
	var circleUser []community.CircleUser

	if info.CircleId != 0 {
		db = db.Where("circle_id = ?'", info.CircleId)
	}
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
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
