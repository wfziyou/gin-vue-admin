package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
)

type AppCircleTagService struct {
}

func (appCircleTagService *AppCircleTagService) CreateCircleTag(circleId uint64, name string) (err error) {
	//var tag = community.CircleTag{}
	//err = global.GVA_DB.Where("name = ?", name).First(&tag).Error
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	err = global.GVA_DB.Create(&tag).Error
	//}
	return err
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
