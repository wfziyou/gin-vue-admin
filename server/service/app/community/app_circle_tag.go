package community

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
	"strconv"
)

type AppCircleTagService struct {
}

func (appCircleTagService *AppCircleTagService) AddCircleTag(circleId uint64, name string) (tag community.CircleTag, err error) {
	var obj = community.CircleTag{CircleId: circleId, Name: name}
	err = global.GVA_DB.Where("circle_id = ? AND name = ?", circleId, name).First(&obj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&obj).Error
	}
	return obj, err
}
func (appCircleTagService *AppCircleTagService) DeleteCircleTags(circleId uint64, names []string) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]community.CircleTag{}, "circle_id = ? and name in ?", circleId, names).Error
	return err
}
func (appCircleTagService *AppCircleTagService) GetCircleTagList(circleId uint64) (list []community.CircleTag, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.CircleTag{})
	var hkCircleTags []community.CircleTag
	db = db.Where("circle_id = ?", circleId)
	db = db.Order("sort desc")
	err = db.Find(&hkCircleTags).Error
	return hkCircleTags, err
}
func (appCircleTagService *AppCircleTagService) SetCircleTagSort(info communityReq.ParamSetCircleTagSort) (err error) {
	if len(info.TagIds) == 0 {
		return nil
	}
	tmp := utils.SplitToUint64List(info.TagIds, ",")
	sql := "update hk_circle_tag set sort = CASE id"
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
