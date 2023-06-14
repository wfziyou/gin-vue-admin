package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
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
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.RecordBrowsingUserHomepage{}).Preload("UserBaseInfo")
	var hkRecordBrowsingUserHomepages []community.RecordBrowsingUserHomepage

	db = db.Where("browser = ?", userId)
	//更新时间降序排列
	db = db.Order("browse_time desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(utils.BrowsingUserShowNum).Offset(0).Find(&hkRecordBrowsingUserHomepages).Error

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

func (hkRecordBrowsingUserHomepageService *RecordBrowsingUserHomepageService) BrowsingUser(browser uint64, userId uint64) (err error) {
	if browser == userId {
		return
	}
	obj := community.RecordBrowsingUserHomepage{}
	err = global.GVA_DB.Where("browser= ? AND user_id = ?", browser, userId).First(&obj).Error
	curTime := time.Now()
	curTimeStr := strconv.FormatInt(curTime.Unix(), 10)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		obj.Browser = browser
		obj.UserId = userId
		obj.BrowseTime = curTime
		obj.BrowseNum = 1
		obj.BrowseData = curTimeStr
		err = global.GVA_DB.Create(&obj).Error
		return
	} else if err != nil {
		return
	}

	browseTime := curTime
	strList := strings.Split(obj.BrowseData, ",")
	len := len(strList)
	if len >= utils.BrowsingUserNum {
		browseData := ""
		strList = append(strList, curTimeStr)
		for index, item := range strList {
			if index > 1 && index < utils.BrowsingUserNum {
				browseData = browseData + "," + item
			} else if index == 1 {
				browseData = item
				if tmp, err := strconv.ParseInt(item, 10, 64); err == nil {
					browseTime = time.Unix(tmp, 0)
				}
			} else if index > utils.BrowsingUserNum {
				obj.BrowseNum = utils.BrowsingUserNum
				break
			}
		}
		browseData = browseData + "," + curTimeStr
		obj.BrowseTime = browseTime
		obj.BrowseData = browseData
	} else if len == 0 {
		obj.BrowseNum = 1
		obj.BrowseData = curTimeStr
	} else {
		obj.BrowseNum = len + 1
		obj.BrowseData = obj.BrowseData + "," + curTimeStr
	}
	err = global.GVA_DB.Updates(obj).Error
	return
}
