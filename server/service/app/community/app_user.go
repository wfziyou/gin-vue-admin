package community

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type AppUserService struct {
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: userInter community.User, err error

func (appUserService *AppUserService) Register(user community.User) (userInter community.User, err error) {
	var userTmp community.User
	if !errors.Is(global.GVA_DB.Where("account = ?", user.Account).First(&userTmp).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	user.Password = utils.BcryptHash(user.Password)
	user.Uuid = uuid.NewV4()
	err = global.GVA_DB.Create(&user).Error
	if err == nil {
		var userCovers []community.UserCoverImage
		err := global.GVA_DB.Model(&community.UserCoverImage{}).Where("account = ?", user.Account).Find(&userCovers)
		size := len(userCovers)
		var coverImage string
		if err == nil && size > 0 {
			rand.Seed(time.Now().Unix())
			coverImage = userCovers[rand.Intn(size)].CoverImage
		}
		if len(coverImage) > 0 {
			appUserService.UpdateUserCover(user.ID, coverImage)
		}
	}
	return user, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (appUserService *AppUserService) Login(u *community.User) (userInter *community.User, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user community.User
	err = global.GVA_DB.Where("account = ?", u.Account).Preload("Authorities").Preload("Authority").Preload("UserExtend").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}

		//if user.UserExtend.CircleId != 0 {
		//	var hkCircle community.Circle
		//	if global.GVA_DB.Where("id = ?", user.UserExtend.CircleId).First(&hkCircle).Error == nil {
		//		user.UserExtend.CircleName = hkCircle.Name
		//	}
		//}
	}
	return &user, err
}
func (appUserService *AppUserService) GetUserByPhone(phone string) (userInter *community.User, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user *community.User
	var users []community.User
	err = global.GVA_DB.Where("phone = ?", phone).Preload("Authorities").Preload("Authority").Preload("UserExtend").Find(&users).Error
	if err != nil {
		return nil, errors.New("数据库错误")
	}
	if len(users) > 0 {
		user = &users[0]
	}
	return user, err
}

// ChangePassword 修改用户密码
func (appUserService *AppUserService) ChangePassword(u *community.User, newPassword string) (userInter *community.User, err error) {
	var user community.User
	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&user).Error
	return &user, err
}

// ResetPassword 重置密码
func (appUserService *AppUserService) ResetPassword(user *community.User, newPassword string) (err error) {
	db := global.GVA_DB.Model(user)
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	updateData["password"] = utils.BcryptHash(newPassword)
	err = db.Where("id = ?", user.ID).Updates(updateData).Error
	return err
}
func (appUserService *AppUserService) BindTelephone(userId uint64, telephone string) (err error) {
	db := global.GVA_DB.Model(&community.User{GvaModelApp: global.GvaModelApp{ID: userId}})
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	updateData["phone"] = telephone
	err = db.Where("id = ?", userId).Updates(updateData).Error
	return err
}
func (appUserService *AppUserService) DecreaseGold(userId uint64, num uint64, title string, titleIcon string, mark string) (err error) {
	if num == 0 {
		return nil
	}
	obj := community.UserExtend{GvaModelApp: global.GvaModelApp{ID: userId}}
	db := global.GVA_DB.Model(&obj)
	err = db.Where("id = ?", userId).First(&obj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("金币不够")
	} else if err != nil {
		return err
	} else {
		if obj.CurrencyGold < num {
			return errors.New("金币不够")
		}
		beforeNumber := obj.CurrencyGold
		currency := obj.CurrencyGold - num
		var updateData map[string]interface{}
		updateData = make(map[string]interface{})
		updateData["currency_gold"] = currency
		err = db.Where("id = ?", userId).Updates(updateData).Error
		if err == nil {
			err = global.GVA_DB.Create(&community.GoldBill{
				UserId:       userId,
				Pm:           community.GoldBillTypeDecrease,
				Title:        title,
				TitleIcon:    titleIcon,
				BeforeNumber: beforeNumber,
				ChangeNumber: num,
				Balance:      currency,
				Mark:         mark,
			}).Error
		}
	}
	return err
}
func (appUserService *AppUserService) IncreaseGold(userId uint64, num uint64, title string, titleIcon string, mark string) (err error) {
	if num == 0 {
		return nil
	}
	obj := community.UserExtend{GvaModelApp: global.GvaModelApp{ID: userId}}
	db := global.GVA_DB.Model(&obj)
	err = db.Where("id = ?", userId).First(&obj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		obj.CurrencyGold = num
		err = global.GVA_DB.Create(&obj).Error
		if err == nil {
			if err == nil {
				err = global.GVA_DB.Create(&community.GoldBill{
					UserId:       userId,
					Pm:           community.GoldBillTypeIncrease,
					Title:        title,
					TitleIcon:    titleIcon,
					BeforeNumber: 0,
					ChangeNumber: num,
					Balance:      num,
					Mark:         mark,
				}).Error
			}
		}
		return err
	} else if err == nil {
		beforeNumber := obj.CurrencyGold
		currency := obj.CurrencyGold + num
		var updateData map[string]interface{}
		updateData = make(map[string]interface{})
		updateData["currency_gold"] = currency
		err = db.Where("id = ?", userId).Updates(updateData).Error
		if err == nil {
			err = global.GVA_DB.Create(&community.GoldBill{
				UserId:       userId,
				Pm:           community.GoldBillTypeIncrease,
				Title:        title,
				TitleIcon:    titleIcon,
				BeforeNumber: beforeNumber,
				ChangeNumber: num,
				Balance:      currency,
				Mark:         mark,
			}).Error
		}
	}
	return err
}

// CreateUser 创建User记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) CreateUser(hkUser community.User) (err error) {
	err = global.GVA_DB.Create(&hkUser).Error
	return err
}

// DeleteUser 删除User记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) DeleteUser(hkUser community.User) (err error) {
	err = global.GVA_DB.Delete(&hkUser).Error
	return err
}

// DeleteUserByIds 批量删除User记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) DeleteUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.User{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUser 更新User记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) UpdateUser(hkUser community.User) (err error) {
	err = global.GVA_DB.Save(&hkUser).Error
	return err
}

func (appUserService *AppUserService) SetSelfBaseInfo(userId uint64, req communityReq.SetSelfBaseInfoReq) (err error) {
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	if len(req.NickName) > 0 {
		updateData["nick_name"] = req.NickName
	}
	if len(req.HeaderImg) > 0 {
		updateData["header_img"] = req.HeaderImg
	}

	if len(req.Birthday) > 0 {
		updateData["birthday"] = req.Birthday
	}
	if req.Sex != nil {
		updateData["sex"] = req.Sex
	}
	if len(req.Description) > 0 {
		updateData["description"] = req.Description
	}

	db := global.GVA_DB.Model(&community.User{})
	err = db.Where("id = ?", userId).Updates(updateData).Error
	return err
}

// GetUser 根据id获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetUser(id uint64) (hkUser community.User, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("UserExtend").First(&hkUser).Error
	return
}
func (appUserService *AppUserService) GetUserInfo(selectUserId uint64, id uint64) (userInfo community.UserInfo, err error) {
	user := community.User{}
	err = global.GVA_DB.Where("id = ?", id).Preload("UserExtend").First(&user).Error
	if err == nil {
		userInfo.ID = user.ID
		userInfo.UserType = user.UserType
		userInfo.Account = user.Account
		userInfo.NickName = user.NickName
		userInfo.RealName = user.RealName
		userInfo.HeaderImg = user.HeaderImg
		userInfo.Birthday = user.Birthday
		userInfo.Sex = user.Sex
		userInfo.Description = user.Description
		userInfo.CoverImage = user.UserExtend.CoverImage
		userInfo.NumCircle = user.UserExtend.NumCircle
		userInfo.NumFocus = user.UserExtend.NumFocus
		userInfo.NumFan = user.UserExtend.NumFan

		isFocus, isFan, _ := GetIsFocusAndIsFan(selectUserId, &userInfo.UserBaseInfo)
		userInfo.IsFocus = isFocus
		userInfo.IsFan = isFan
	}
	return
}
func SetPostsListUserIsFocus(selectUserId uint64, list []community.ForumPostsBaseInfo) {
	size := len(list)
	if size == 0 {
		return
	}
	var ids = make([]uint64, 0, size)
	num := 0
	for _, obj := range list {
		if obj.UserInfo.ID != selectUserId {
			ids = append(ids, obj.UserInfo.ID)
			num++
		}
	}
	if num == 0 {
		return
	}
	var focusUserList []community.FocusUser
	err := global.GVA_DB.Model(&community.FocusUser{}).Where("user_id = ? AND focus_user_id in ?", selectUserId, ids).Find(&focusUserList).Error
	if err != nil {
		return
	}

	for index, obj := range list {
		if obj.UserInfo.ID != selectUserId {
			for _, focus := range focusUserList {
				if obj.UserId == focus.FocusUserId {
					list[index].UserInfo.IsFocus = 1
					break
				}
			}
		}
	}
}
func SetForumCommentUserIsFocus(selectUserId uint64, list []community.ForumComment) {
	size := len(list)
	if size == 0 {
		return
	}
	var ids = make([]uint64, 0, size)
	num := 0
	for _, obj := range list {
		if obj.UserInfo.ID != selectUserId {
			ids = append(ids, obj.UserInfo.ID)
			num++
		}
	}
	if num == 0 {
		return
	}
	var focusUserList []community.FocusUser
	err := global.GVA_DB.Model(&community.FocusUser{}).Where("user_id = ? AND focus_user_id in ?", selectUserId, ids).Find(&focusUserList).Error
	if err != nil {
		return
	}

	for index, obj := range list {
		if obj.UserInfo.ID != selectUserId {
			for _, focus := range focusUserList {
				if obj.UserId == focus.FocusUserId {
					list[index].UserInfo.IsFocus = 1
					break
				}
			}
		}
	}
}
func SetCircleUserUserIsFocus(selectUserId uint64, list []community.CircleUser) {
	size := len(list)
	if size == 0 {
		return
	}
	var ids = make([]uint64, 0, size)
	num := 0
	for _, obj := range list {
		if obj.UserInfo.ID != selectUserId {
			ids = append(ids, obj.UserInfo.ID)
			num++
		}
	}
	if num == 0 {
		return
	}
	var focusUserList []community.FocusUser
	err := global.GVA_DB.Model(&community.FocusUser{}).Where("user_id = ? AND focus_user_id in ?", selectUserId, ids).Find(&focusUserList).Error
	if err != nil {
		return
	}

	for index, obj := range list {
		if obj.UserInfo.ID != selectUserId {
			for _, focus := range focusUserList {
				if obj.UserId == focus.FocusUserId {
					list[index].UserInfo.IsFocus = 1
					break
				}
			}
		}
	}
}

func GetIsFocusAndIsFan(selectUserId uint64, userInfo *community.UserBaseInfo) (isFocus int, isFan int, err error) {
	if selectUserId != userInfo.ID {
		var num1, num2 int64
		err1 := global.GVA_DB.Model(&community.FocusUser{}).Where("user_id = ? AND focus_user_id = ?", selectUserId, userInfo.ID).Count(&num1).Error
		err2 := global.GVA_DB.Model(&community.FocusUser{}).Where("user_id = ? AND focus_user_id = ?", userInfo.ID, selectUserId).Count(&num2).Error
		if err1 == nil && err2 == nil {
			if num1 > 0 {
				isFocus = 1
			}
			if num2 > 0 {
				isFan = 1
			}
		}
	}
	return
}
func (appUserService *AppUserService) GetIsFocusAndIsFan(selectUserId uint64, userInfo *community.UserBaseInfo) (isFocus int, isFan int, err error) {
	return GetIsFocusAndIsFan(selectUserId, userInfo)
}

// GetUserInfoList 分页获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetUserInfoList(info communityReq.UserSearch) (list []community.UserBaseInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.UserBaseInfo{})
	var hkUsers []community.UserBaseInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.Keyword) > 0 {
		account := "%" + info.Keyword + "%"
		nickName := "%" + info.Keyword + "%"
		phone := "%" + info.Keyword + "%"
		db = db.Where("account LIKE ? OR nick_name LIKE ? OR phone = ?", account, nickName, phone)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUsers).Error
	return hkUsers, total, err
}

// GetUserChannel 获取用户频道
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetUserChannel(userId uint64) (channel string, err error) {
	var data = community.UserExtend{}
	err = global.GVA_DB.Where("id = ?", userId).Select("channel_id").First(&data).Error
	if err == nil {
		channel = data.ChannelId
	}
	return
}

// UpdateUserChannel 更新用户频道
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) UpdateUserChannel(userId uint64, channel string) (err error) {
	obj := community.UserExtend{GvaModelApp: global.GvaModelApp{ID: userId}, ChannelId: channel}
	db := global.GVA_DB.Model(&obj)
	if errors.Is(db.Where("id = ?", userId).First(&obj).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&obj).Error
	} else {
		var updateData map[string]interface{}
		updateData = make(map[string]interface{})
		updateData["channel_id"] = channel
		err = db.Where("id = ?", userId).Updates(updateData).Error
	}
	return err
}
func (appUserService *AppUserService) UpdateUserCoverImage(userId uint64, coverImage string) (err error) {
	obj := community.UserExtend{GvaModelApp: global.GvaModelApp{ID: userId}, CoverImage: coverImage}
	db := global.GVA_DB.Model(&obj)
	if errors.Is(db.Where("id = ?", userId).First(&obj).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&obj).Error
	} else {
		var updateData map[string]interface{}
		updateData = make(map[string]interface{})
		updateData["cover_image"] = coverImage
		err = db.Where("id = ?", userId).Updates(updateData).Error
	}
	return err
}
func (appUserService *AppUserService) UpdatePostsTime(circleId uint64, userId uint64) (err error) {
	obj := community.UserExtend{GvaModelApp: global.GvaModelApp{ID: userId}}
	db := global.GVA_DB.Model(&obj)
	curTime := time.Now()
	if errors.Is(db.Where("id = ?", userId).First(&obj).Error, gorm.ErrRecordNotFound) {
		now := curTime
		obj.UpdateForumPostsTime = &now
		err = global.GVA_DB.Create(&obj).Error
	} else {
		var updateData map[string]interface{}
		updateData = make(map[string]interface{})
		updateData["update_forum_posts_time"] = curTime
		err = db.Where("id = ?", userId).Updates(updateData).Error
	}
	if circleId > 0 {
		var circle community.Circle
		err = global.GVA_DB.Where("id = ?", circleId).First(&circle).Error
		if err == nil {
			var updateData map[string]interface{}
			updateData = make(map[string]interface{})
			updateData["update_forum_posts_time"] = curTime
			err = global.GVA_DB.Model(&community.Circle{}).Where("id = ?", circleId).Updates(updateData).Error
		}
	}
	return err
}
func (appUserService *AppUserService) UpdateUserCover(userId uint64, coverImage string) (err error) {
	obj := community.UserExtend{GvaModelApp: global.GvaModelApp{ID: userId}, CoverImage: coverImage}
	db := global.GVA_DB.Model(&obj)
	if errors.Is(db.Where("id = ?", userId).First(&obj).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&obj).Error
	} else {
		var updateData map[string]interface{}
		updateData = make(map[string]interface{})
		updateData["cover_image"] = coverImage
		err = db.Where("id = ?", userId).Updates(updateData).Error
	}
	return err
}
func (appUserService *AppUserService) UpdateUserNumFocus(userId uint64) (err error) {
	focusCount := int64(0)
	err = global.GVA_DB.Model(&community.FocusUser{}).Where("user_id = ?", userId).Count(&focusCount).Error
	if err != nil {
		return
	}

	obj := community.UserExtend{GvaModelApp: global.GvaModelApp{ID: userId}, NumFocus: focusCount}
	db := global.GVA_DB.Model(&obj)
	if errors.Is(db.Where("id = ?", userId).First(&obj).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&obj).Error
	} else {
		var updateData map[string]interface{}
		updateData = make(map[string]interface{})
		updateData["num_focus"] = focusCount
		err = db.Where("id = ?", userId).Updates(updateData).Error
	}
	return err
}
func (appUserService *AppUserService) UpdateUserNumFan(userId uint64) (err error) {
	fanCount := int64(0)
	err = global.GVA_DB.Model(&community.FocusUser{}).Where("focus_user_id = ?", userId).Count(&fanCount).Error
	if err != nil {
		return
	}

	obj := community.UserExtend{GvaModelApp: global.GvaModelApp{ID: userId}, NumFan: fanCount}
	db := global.GVA_DB.Model(&obj)
	if errors.Is(db.Where("id = ?", userId).First(&obj).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&obj).Error
	} else {
		var updateData map[string]interface{}
		updateData = make(map[string]interface{})
		updateData["num_fan"] = fanCount
		err = db.Where("id = ?", userId).Updates(updateData).Error
	}
	return err
}
func (appUserService *AppUserService) UpdateUserFocus(userId uint64) (err error) {
	focusCount := int64(0)
	fanCount := int64(0)
	err = global.GVA_DB.Model(&community.FocusUser{}).Where("user_id = ?", userId).Count(&focusCount).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(&community.FocusUser{}).Where("focus_user_id = ?", userId).Count(&fanCount).Error
	if err != nil {
		return
	}

	obj := community.UserExtend{GvaModelApp: global.GvaModelApp{ID: userId}, NumFocus: focusCount, NumFan: fanCount}
	db := global.GVA_DB.Model(&obj)
	if errors.Is(db.Where("id = ?", userId).First(&obj).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Create(&obj).Error
	} else {
		var updateData map[string]interface{}
		updateData = make(map[string]interface{})
		updateData["num_focus"] = focusCount
		updateData["num_fan"] = fanCount
		err = db.Where("id = ?", userId).Updates(updateData).Error
	}
	return err
}
