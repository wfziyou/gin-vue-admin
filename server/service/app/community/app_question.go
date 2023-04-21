package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type QuestionService struct {
}

// CreateQuestion 创建Question记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionService *QuestionService) CreateQuestion(userId uint64, hkQuestion communityReq.CreateQuestion) (err error) {
	//err = global.GVA_DB.Create(&hkQuestion).Error
	return err
}

// DeleteQuestion 删除Question记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionService *QuestionService) DeleteQuestion(hkQuestion community.ForumPosts) (err error) {
	err = global.GVA_DB.Delete(&hkQuestion).Error
	return err
}

func (questionService *QuestionService) DeleteQuestionById(id uint64) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumPosts{}, "id in =", id).Error
	return err
}

// DeleteQuestionByIds 批量删除Question记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionService *QuestionService) DeleteQuestionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.ForumPosts{}, "id in ?", ids.Ids).Error
	return err
}

//// UpdateQuestion 更新Question记录
//// Author [piexlmax](https://github.com/piexlmax)
//func (questionService *QuestionService) UpdateQuestion(hkQuestion communityReq.UpdateQuestionReq) (err error) {
//	//err = global.GVA_DB.Save(&hkQuestion).Error
//	return err
//}

// GetQuestion 根据id获取Question记录
// Author [piexlmax](https://github.com/piexlmax)
func (questionService *QuestionService) GetQuestion(id uint64) (hkQuestion community.ForumPosts, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkQuestion).Error
	return
}

//// GetQuestionInfoList 分页获取Question记录
//// Author [piexlmax](https://github.com/piexlmax)
//func (questionService *QuestionService) GetQuestionInfoList(info communityReq.QuestionSearch) (list []community.ForumPostsBaseInfo, total int64, err error) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	// 创建db
//	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{})
//	var hkQuestions []community.ForumPostsBaseInfo
//	// 如果有条件搜索 下方会自动创建搜索语句
//	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
//		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
//	}
//	err = db.Count(&total).Error
//	if err != nil {
//		return
//	}
//
//	err = db.Limit(limit).Offset(offset).Find(&hkQuestions).Error
//	return hkQuestions, total, err
//}
