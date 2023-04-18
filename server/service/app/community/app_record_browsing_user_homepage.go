package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type RecordBrowsingUserHomepageService struct {
}

// CreateRecordBrowsingUserHomepage 创建RecordBrowsingUserHomepage记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkRecordBrowsingUserHomepageService *RecordBrowsingUserHomepageService) CreateRecordBrowsingUserHomepage(hkRecordBrowsingUserHomepage *community.RecordBrowsingUserHomepage) (err error) {
	err = global.GVA_DB.Create(hkRecordBrowsingUserHomepage).Error
	return err
}

// DeleteRecordBrowsingUserHomepage 删除RecordBrowsingUserHomepage记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkRecordBrowsingUserHomepageService *RecordBrowsingUserHomepageService) DeleteRecordBrowsingUserHomepage(hkRecordBrowsingUserHomepage community.RecordBrowsingUserHomepage) (err error) {
	err = global.GVA_DB.Delete(&hkRecordBrowsingUserHomepage).Error
	return err
}

// DeleteRecordBrowsingUserHomepageByIds 批量删除RecordBrowsingUserHomepage记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkRecordBrowsingUserHomepageService *RecordBrowsingUserHomepageService) DeleteRecordBrowsingUserHomepageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.RecordBrowsingUserHomepage{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRecordBrowsingUserHomepage 更新RecordBrowsingUserHomepage记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkRecordBrowsingUserHomepageService *RecordBrowsingUserHomepageService) UpdateRecordBrowsingUserHomepage(hkRecordBrowsingUserHomepage community.RecordBrowsingUserHomepage) (err error) {
	err = global.GVA_DB.Save(&hkRecordBrowsingUserHomepage).Error
	return err
}

// GetRecordBrowsingUserHomepage 根据id获取RecordBrowsingUserHomepage记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkRecordBrowsingUserHomepageService *RecordBrowsingUserHomepageService) GetRecordBrowsingUserHomepage(id uint64) (hkRecordBrowsingUserHomepage community.RecordBrowsingUserHomepage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkRecordBrowsingUserHomepage).Error
	return
}

// GetRecordBrowsingUserHomepageInfoList 分页获取RecordBrowsingUserHomepage记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkRecordBrowsingUserHomepageService *RecordBrowsingUserHomepageService) GetRecordBrowsingUserHomepageInfoList(info communityReq.RecordBrowsingUserHomepageSearch) (list []community.RecordBrowsingUserHomepage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.RecordBrowsingUserHomepage{})
	var hkRecordBrowsingUserHomepages []community.RecordBrowsingUserHomepage
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkRecordBrowsingUserHomepages).Error
	return hkRecordBrowsingUserHomepages, total, err
}

func (hkRecordBrowsingUserHomepageService *RecordBrowsingUserHomepageService) GetFrequentBrowsingUserList(userId uint64, info request.PageInfo) (list []community.UserBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.RecordBrowsingUserHomepage{}).Preload("UserBaseInfo")
	var hkRecordBrowsingUserHomepages []community.RecordBrowsingUserHomepage

	db = db.Where("user_id = ?", userId)
	//更新时间降序排列
	db = db.Order("updated_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkRecordBrowsingUserHomepages).Error

	var userBaseInfos []community.UserBaseInfo
	if err == nil {
		if len(hkRecordBrowsingUserHomepages) > 0 {
			for _, info := range hkRecordBrowsingUserHomepages {
				userBaseInfos = append(userBaseInfos, info.UserBaseInfo)
			}
		}
	}

	return userBaseInfos, total, err
}
