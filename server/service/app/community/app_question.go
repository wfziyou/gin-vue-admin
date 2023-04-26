package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type QuestionService struct {
}

// CreateQuestion 创建Question记录
func (questionService *QuestionService) CreateQuestion(userId uint64, info communityReq.CreateQuestion) (err error) {
	forumPosts := community.ForumPosts{
		UserId:       userId,
		CircleId:     0,
		Category:     community.PostsCategoryQuestion,
		Title:        info.Title,
		ContentType:  community.ContentTypeHtml,
		ContentHtml:  info.Content,
		Attachment:   info.Attachment,
		PayCurrency:  utils.CurrencyGold,
		PayNum:       info.PayNum,
		CheckStatus:  community.PostsCheckStatusPass,
		IsPublic:     community.ForumPostsIsPublicTrue,
		PowerComment: community.ForumPostsPowerCommentOpen,
	}
	err = global.GVA_DB.Create(&forumPosts).Error
	return err
}

// GetQuestion 根据id获取Question记录
func (questionService *QuestionService) GetQuestion(id uint64) (hkQuestion community.ForumPosts, err error) {
	err = global.GVA_DB.Where("id = ? AND category = ?", id, community.PostsCategoryQuestion).First(&hkQuestion).Error
	return
}
func (questionService *QuestionService) SetBestAnswer(question community.ForumPosts, answerId uint64) (answer community.ForumComment, err error) {

	if question.Status == utils.QuestionStatusClose {
		err = errors.New("问题已关闭")
		return
	}
	answer = community.ForumComment{GvaModelApp: global.GvaModelApp{ID: answerId}}
	err = global.GVA_DB.Model(&community.ForumComment{}).Select("posts_id,user_id").Where("id = ?", answerId).First(&answer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("答案不存在")
		return
	} else if err != nil {
		return answer, err
	}

	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	updateData["comment_id"] = answerId
	updateData["status"] = utils.QuestionStatusClose
	err = global.GVA_DB.Model(&question).Where("id = ?", question.ID).Updates(updateData).Error
	return answer, err
}

// CloseQuestion 关闭问题
func (questionService *QuestionService) CloseQuestion(userId uint64, questionId uint64) (err error) {
	obj := community.ForumPosts{}
	err = global.GVA_DB.Model(&community.ForumPosts{}).Select("user_id,status").Where("id = ?", questionId).First(&obj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("问题不存在")
	} else if err != nil {
		return err
	}

	if obj.UserId != userId {
		return errors.New("只能关闭自己的问题")
	}
	if obj.Status == utils.QuestionStatusClose {
		return nil
	}
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	updateData["status"] = utils.QuestionStatusClose
	err = global.GVA_DB.Model(&obj).Where("id = ?", questionId).Updates(updateData).Error
	return err
}

func (questionService *QuestionService) AnswerQuestion(userId uint64, questionId uint64, content string) (err error) {
	obj := community.ForumPosts{}
	err = global.GVA_DB.Model(&community.ForumPosts{}).Select("user_id,status").Where("id = ?", questionId).First(&obj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("问题不存在")
	} else if err != nil {
		return err
	}

	if obj.Status == utils.QuestionStatusClose {
		return errors.New("问题已关闭")
	}

	err = global.GVA_DB.Create(&community.ForumComment{
		PostsId:        questionId,
		UserId:         userId,
		CommentContent: content,
	}).Error
	return err
}

func (questionService *QuestionService) DelSelfAnswer(userId uint64, answerId uint64) (err error) {
	obj := community.ForumPosts{}
	err = global.GVA_DB.Model(&community.ForumComment{}).Select("user_id").Where("id = ?", answerId).First(&obj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("答案不存在")
	} else if err != nil {
		return err
	}

	if obj.UserId != userId {
		return errors.New("只能删除自己的答案")
	}

	err = global.GVA_DB.Delete(&community.ForumComment{
		GvaModelApp: global.GvaModelApp{ID: answerId},
	}).Error
	return err
}

// DeleteQuestion 删除Question记录
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

// GetForumCommentInfoList 分页获取评论记录
func (questionService *QuestionService) GetAnswerList(questionId uint64, info request.PageInfo) (list []community.ForumComment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumComment{}).Preload("UserInfo")
	var hkForumComments []community.ForumComment
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("posts_id = ?", questionId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumComments).Error
	return hkForumComments, total, err
}

func (questionService *QuestionService) GetCircleQuestionList(circleId uint64, page request.PageInfo) (list []community.ForumPostsBaseInfo, total int64, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.ForumPostsBaseInfo{}).Preload("TopicInfo").Preload("UserInfo").Preload("CircleInfo")
	var hkForumPostss []community.ForumPostsBaseInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	//圈子_编号
	db = db.Where("category = ? and circle_id = ?", community.PostsCategoryQuestion, circleId)

	//标题
	if len(page.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+page.Keyword+"%")
	}
	//创建时间降序排列
	db = db.Order("hk_forum_posts.created_at desc")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkForumPostss).Error
	return hkForumPostss, total, err
}
