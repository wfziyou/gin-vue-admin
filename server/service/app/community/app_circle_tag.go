package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"gorm.io/gorm"
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
	err = db.Find(&hkCircleTags).Error
	return hkCircleTags, err
}
