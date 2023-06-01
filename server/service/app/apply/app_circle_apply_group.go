package apply

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"strconv"
)

type AppCircleApplyGroupService struct {
}

func (appCircleApplyGroupService *AppCircleApplyGroupService) SetCircleApplyGroupSort(info applyReq.ParamSetCircleApplyGroupSort) (err error) {
	if len(info.GroupIds) == 0 {
		return nil
	}
	tmp := utils.SplitToUint64List(info.GroupIds, ",")
	sql := "update hk_circle_apply_group set sort = CASE id"
	sqlIds := "("
	size := len(tmp)
	for index, id := range tmp {
		sql += fmt.Sprintf(" WHEN %d THEN %d", id, size-index)
		if index != 0 {
			sqlIds = sqlIds + "," + strconv.FormatUint(id, 10)
		} else {
			sqlIds = sqlIds + strconv.FormatUint(id, 10)
		}
	}
	sqlIds += ")"
	sql = sql + " END WHERE id IN" + sqlIds
	err = global.GVA_DB.Exec(sql).Error
	return err
}
func (appCircleApplyGroupService *AppCircleApplyGroupService) AddCircleApplyGroup(circleId uint64, name string) (err error) {
	err = global.GVA_DB.Create(&apply.CircleApplyGroup{
		CircleId: circleId,
		Name:     name,
	}).Error
	return err
}

// DeleteCircleApplyGroup 删除CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) DeleteCircleApplyGroup(hkCircleApplyGroup apply.CircleApplyGroup) (err error) {
	err = global.GVA_DB.Delete(&hkCircleApplyGroup).Error
	return err
}

// DeleteCircleApplyGroupByIds 批量删除CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) DeleteCircleApplyGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.CircleApplyGroup{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCircleApplyGroup 更新CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) UpdateCircleApplyGroup(hkCircleApplyGroup apply.CircleApplyGroup) (err error) {
	err = global.GVA_DB.Save(&hkCircleApplyGroup).Error
	return err
}

// GetCircleApplyGroup 根据id获取CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) GetCircleApplyGroup(id uint64) (hkCircleApplyGroup apply.CircleApplyGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleApplyGroup).Error
	return
}

// GetCircleApplyGroupInfoList 分页获取CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) GetCircleApplyGroupInfoList(info applyReq.CircleApplyGroupSearch) (list []apply.CircleApplyGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&apply.CircleApplyGroup{})
	var hkCircleApplyGroups []apply.CircleApplyGroup
	if info.CircleId != 0 {
		db = db.Where("circle_id = ?", info.CircleId)
	}
	if len(info.Keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+info.Keyword+"%")
	}
	if info.ParentId != nil {
		db = db.Where("parent_id = ?", info.ParentId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	db = db.Order("sort desc")
	err = db.Limit(limit).Offset(offset).Find(&hkCircleApplyGroups).Error
	return hkCircleApplyGroups, total, err
}

// GetCircleApplyGroupInfoListAll 分页获取CircleApplyGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyGroupService *AppCircleApplyGroupService) GetCircleApplyGroupInfoListAll(info applyReq.CircleApplyGroupSearchAll) (list []apply.CircleApplyGroup, err error) {
	// 创建db
	db := global.GVA_DB.Model(&apply.CircleApplyGroup{})
	var hkCircleApplyGroups []apply.CircleApplyGroup

	if info.CircleId != 0 {
		db = db.Where("circle_id = ?", info.CircleId)
	}
	if len(info.Name) > 0 {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.ParentId != nil {
		db = db.Where("parent_id = ?", info.ParentId)
	}
	db = db.Order("sort desc")
	err = db.Find(&hkCircleApplyGroups).Error
	return hkCircleApplyGroups, err
}
