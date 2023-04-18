package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type RecordBrowsingCircleHomepageService struct {
}

// CreateRecordBrowsingCircleHomepage 创建RecordBrowsingCircleHomepage记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkRecordBrowsingCircleHomepageService *RecordBrowsingCircleHomepageService) CreateRecordBrowsingCircleHomepage(hkRecordBrowsingCircleHomepage *community.RecordBrowsingCircleHomepage) (err error) {
	err = global.GVA_DB.Create(hkRecordBrowsingCircleHomepage).Error
	return err
}

// DeleteRecordBrowsingCircleHomepage 删除RecordBrowsingCircleHomepage记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkRecordBrowsingCircleHomepageService *RecordBrowsingCircleHomepageService) DeleteRecordBrowsingCircleHomepage(hkRecordBrowsingCircleHomepage community.RecordBrowsingCircleHomepage) (err error) {
	err = global.GVA_DB.Delete(&hkRecordBrowsingCircleHomepage).Error
	return err
}

// DeleteRecordBrowsingCircleHomepageByIds 批量删除RecordBrowsingCircleHomepage记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkRecordBrowsingCircleHomepageService *RecordBrowsingCircleHomepageService) DeleteRecordBrowsingCircleHomepageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.RecordBrowsingCircleHomepage{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRecordBrowsingCircleHomepage 更新RecordBrowsingCircleHomepage记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkRecordBrowsingCircleHomepageService *RecordBrowsingCircleHomepageService) UpdateRecordBrowsingCircleHomepage(hkRecordBrowsingCircleHomepage community.RecordBrowsingCircleHomepage) (err error) {
	err = global.GVA_DB.Save(&hkRecordBrowsingCircleHomepage).Error
	return err
}

// GetRecordBrowsingCircleHomepage 根据id获取RecordBrowsingCircleHomepage记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkRecordBrowsingCircleHomepageService *RecordBrowsingCircleHomepageService) GetRecordBrowsingCircleHomepage(id uint64) (hkRecordBrowsingCircleHomepage community.RecordBrowsingCircleHomepage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkRecordBrowsingCircleHomepage).Error
	return
}

// GetRecordBrowsingCircleHomepageInfoList 分页获取RecordBrowsingCircleHomepage记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkRecordBrowsingCircleHomepageService *RecordBrowsingCircleHomepageService) GetRecordBrowsingCircleHomepageInfoList(info communityReq.RecordBrowsingCircleHomepageSearch) (list []community.RecordBrowsingCircleHomepage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.RecordBrowsingCircleHomepage{})
	var hkRecordBrowsingCircleHomepages []community.RecordBrowsingCircleHomepage
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkRecordBrowsingCircleHomepages).Error
	return hkRecordBrowsingCircleHomepages, total, err
}
