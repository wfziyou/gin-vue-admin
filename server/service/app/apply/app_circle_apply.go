package apply

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/apply"
	applyReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/apply/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/general"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AppCircleApplyService struct {
}

// CreateCircleApply 创建CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) CreateCircleApply(circleId uint64, info applyReq.ParamAddCircleApply) (err error) {
	var tmpApply = apply.Apply{}
	tmpApply.OwerType = utils.ApplyOwnerTypeCircle
	tmpApply.CircleId = circleId
	tmpApply.Name = info.Name
	tmpApply.Icon = info.Icon
	tmpApply.Type = info.Type
	tmpApply.ApplyParameters = info.ApplyParameters

	if info.Type == utils.ApplyTypeH5 {
		tmpApply.ApplyAddress = info.ApplyAddress
	} else {
		tmpApply.Type = utils.ApplyTypeMimiProgram
		tmpApply.MiniProgramId = info.MiniProgramId
	}

	err = global.GVA_DB.Create(&tmpApply).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Create(&apply.CircleApply{
		CircleId:     circleId,
		ApplyGroupId: info.ApplyGroupId,
		ApplyId:      tmpApply.ID,
		Power:        info.Power,
	}).Error
	if err != nil {
		global.GVA_DB.Delete(&tmpApply)
	}
	return err
}

// DeleteCircleApply 删除CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) DeleteCircleApply(hkCircleApply apply.CircleApply) (err error) {
	err = global.GVA_DB.Delete(&hkCircleApply).Error
	return err
}
func (appCircleApplyService *AppCircleApplyService) DeleteCircleApplyByApplyId(circleId uint64, applyId uint64) (err error) {
	err = global.GVA_DB.Delete(&[]apply.CircleApply{}, "circle_id = ? AND apply_id = ?", circleId, applyId).Error
	return err
}

// DeleteCircleApplyByIds 批量删除CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) DeleteCircleApplyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]apply.CircleApply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCircleApply 更新CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) UpdateCircleApply(info applyReq.ParamUpdateCircleApply) (err error) {
	//var updateData map[string]interface{}
	//updateData = make(map[string]interface{})
	//if len(info.Name) > 0 {
	//	updateData["name"] = info.Name
	//}
	//
	//db := global.GVA_DB.Model(&community.Circle{})
	//err = db.Where("id = ?", info.ID).Updates(updateData).Error
	return err
}

// GetCircleApply 根据id获取CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) GetCircleApply(id uint64) (hkCircleApply apply.CircleApply, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkCircleApply).Error
	return
}
func (appCircleApplyService *AppCircleApplyService) GetCircleApplyCountByGroupId(groupId uint64) (count int64, err error) {
	db := global.GVA_DB.Model(&apply.CircleApply{})
	db = db.Where("apply_group_id = ?", groupId)
	err = db.Count(&count).Error
	return count, err
}

// GetCircleApplyInfoList 分页获取CircleApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCircleApplyService *AppCircleApplyService) GetCircleApplyInfoList(info applyReq.CircleApplySearch, isMember bool) (list []apply.CircleApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&apply.CircleApply{}).Preload("Apply")
	var hkCircleApplys []apply.CircleApply
	db = db.Where("circle_id = ?", info.CircleId)

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkCircleApplys).Error
	if err == nil && len(hkCircleApplys) > 0 {
		var ids = make([]uint64, 0, len(hkCircleApplys))

		for _, obj := range hkCircleApplys {
			if obj.Apply.Type == utils.ApplyTypeMimiProgram {
				ids = append(ids, obj.Apply.MiniProgramId)
			}
		}
		if len(ids) > 0 {
			var miniProgram []general.MiniProgram
			err1 := global.GVA_DB.Where("id in ?", ids).Find(&miniProgram).Error
			if err1 == nil && len(miniProgram) > 0 {
				for index, obj := range hkCircleApplys {
					if obj.Apply.Type == utils.ApplyTypeMimiProgram {
						for _, program := range miniProgram {
							if obj.Apply.MiniProgramId == program.ID {
								hkCircleApplys[index].Apply.ProgramId = program.ProgramId
								break
							}
						}
					}
				}
			}
		}
	}
	return hkCircleApplys, total, err
}

// GetCircleApplyInfoListAll 分页获取CircleApply记录
func (appCircleApplyService *AppCircleApplyService) GetCircleApplyInfoListAll(info applyReq.CircleApplySearchAll, isMember bool) (list []apply.CircleApply, err error) {
	db := global.GVA_DB.Model(&apply.CircleApply{}).Preload("Apply")
	var hkCircleApplys []apply.CircleApply

	db = db.Where("circle_id = ?", info.CircleId)

	err = db.Find(&hkCircleApplys).Error
	if err == nil && len(hkCircleApplys) > 0 {
		var ids = make([]uint64, 0, len(hkCircleApplys))

		for _, obj := range hkCircleApplys {
			if obj.Apply.Type == utils.ApplyTypeMimiProgram {
				ids = append(ids, obj.Apply.MiniProgramId)
			}
		}
		if len(ids) > 0 {
			var miniProgram []general.MiniProgram
			err1 := global.GVA_DB.Where("id in ?", ids).Find(&miniProgram).Error
			if err1 == nil && len(miniProgram) > 0 {
				for index, obj := range hkCircleApplys {
					if obj.Apply.Type == utils.ApplyTypeMimiProgram {
						for _, program := range miniProgram {
							if obj.Apply.MiniProgramId == program.ID {
								hkCircleApplys[index].Apply.ProgramId = program.ProgramId
								break
							}
						}
					}
				}
			}
		}
	}
	return hkCircleApplys, err
}
func (appCircleApplyService *AppCircleApplyService) GetCircleHotApplyList(circleId uint64, applyIds string, isMember bool) (list []apply.CircleApply, total int64, err error) {
	if len(applyIds) == 0 {
		return
	}
	db := global.GVA_DB.Model(&apply.CircleApply{}).Preload("Apply")

	power := 0
	if isMember == true {
		power = 1
	}
	sql := fmt.Sprintf("id in(%s) AND circle_id = %d AND power <= %d", applyIds, circleId, power)
	db = db.Where(sql)

	err = global.GVA_DB.Model(&apply.CircleApply{}).Where("circle_id = ? AND power <= ?", circleId, power).Count(&total).Error
	if err != nil {
		return
	}

	var hkCircleApplys []apply.CircleApply
	err = db.Find(&hkCircleApplys).Error
	size := len(hkCircleApplys)
	if err == nil && size > 0 {
		newCircleApplys := make([]apply.CircleApply, 0, size)
		tmp := utils.SplitToUint64List(applyIds, ",")
		for _, id := range tmp {
			for _, obj := range hkCircleApplys {
				if obj.ID == id {
					newCircleApplys = append(newCircleApplys, obj)
					break
				}
			}
		}

		var ids = make([]uint64, 0, size)
		for _, obj := range newCircleApplys {
			if obj.Apply.Type == utils.ApplyTypeMimiProgram {
				ids = append(ids, obj.Apply.MiniProgramId)
			}
		}
		if len(ids) > 0 {
			var miniProgram []general.MiniProgram
			err1 := global.GVA_DB.Where("id in ?", ids).Find(&miniProgram).Error
			if err1 == nil && len(miniProgram) > 0 {
				for index, obj := range newCircleApplys {
					if obj.Apply.Type == utils.ApplyTypeMimiProgram {
						for _, program := range miniProgram {
							if obj.Apply.MiniProgramId == program.ID {
								newCircleApplys[index].Apply.ProgramId = program.ProgramId
								break
							}
						}
					}
				}
			}
		}
		return newCircleApplys, total, err
	}
	return nil, total, err
}
