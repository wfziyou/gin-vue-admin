package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserCoverImageService struct {
}

// CreateUserCoverImage 创建UserCoverImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *UserCoverImageService) CreateUserCoverImage(hkUserCoverImage *community.UserCoverImage) (err error) {
	err = global.GVA_DB.Create(hkUserCoverImage).Error
	return err
}

// DeleteUserCoverImage 删除UserCoverImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *UserCoverImageService) DeleteUserCoverImage(hkUserCoverImage community.UserCoverImage) (err error) {
	err = global.GVA_DB.Delete(&hkUserCoverImage).Error
	return err
}

// DeleteUserCoverImageByIds 批量删除UserCoverImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *UserCoverImageService) DeleteUserCoverImageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.UserCoverImage{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserCoverImage 更新UserCoverImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *UserCoverImageService) UpdateUserCoverImage(hkUserCoverImage community.UserCoverImage) (err error) {
	err = global.GVA_DB.Save(&hkUserCoverImage).Error
	return err
}

// GetUserCoverImage 根据id获取UserCoverImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (service *UserCoverImageService) GetUserCoverImage(id uint64) (hkUserCoverImage community.UserCoverImage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUserCoverImage).Error
	return
}

// GetUserCoverImageList 分页获取UserCoverImage记录
func (service *UserCoverImageService) GetUserCoverImageList(info communityReq.UserCoverImageSearch) (list []community.UserCoverImage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.UserCoverImage{})
	var hkUserCoverImages []community.UserCoverImage
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
	}
	err = db.Limit(limit).Offset(offset).Find(&hkUserCoverImages).Error
	return hkUserCoverImages, total, err
}
