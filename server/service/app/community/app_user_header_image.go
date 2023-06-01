package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserHeaderImageService struct {
}

// CreateUserHeaderImage 创建UserHeaderImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *UserHeaderImageService) CreateUserHeaderImage(hkUserHeaderImage *community.UserHeaderImage) (err error) {
	err = global.GVA_DB.Create(hkUserHeaderImage).Error
	return err
}

// DeleteUserHeaderImage 删除UserHeaderImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *UserHeaderImageService) DeleteUserHeaderImage(hkUserHeaderImage community.UserHeaderImage) (err error) {
	err = global.GVA_DB.Delete(&hkUserHeaderImage).Error
	return err
}

// DeleteUserHeaderImageByIds 批量删除UserHeaderImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *UserHeaderImageService) DeleteUserHeaderImageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.UserHeaderImage{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserHeaderImage 更新UserHeaderImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *UserHeaderImageService) UpdateUserHeaderImage(hkUserHeaderImage community.UserHeaderImage) (err error) {
	err = global.GVA_DB.Save(&hkUserHeaderImage).Error
	return err
}

// GetUserHeaderImage 根据id获取UserHeaderImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *UserHeaderImageService) GetUserHeaderImage(id uint64) (hkUserHeaderImage community.UserHeaderImage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserHeaderImage).Error
	return
}

// GetUserHeaderImageList 分页获取UserHeaderImage记录
func (service *UserHeaderImageService) GetUserHeaderImageList(info communityReq.UserHeaderImageSearch) (list []community.UserHeaderImage, err error) {
	// 创建db
	db := global.GVA_DB.Model(&community.UserHeaderImage{})
	var hkUserHeaderImages []community.UserHeaderImage

	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
	}
	err = db.Find(&hkUserHeaderImages).Error
	return hkUserHeaderImages, err
}
