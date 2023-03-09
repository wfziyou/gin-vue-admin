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
)

type AppUserService struct {
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: userInter community.User, err error

func (appUserService *AppUserService) Register(u community.User) (userInter community.User, err error) {
	var user community.User
	if !errors.Is(global.GVA_DB.Where("account = ?", u.Account).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.Uuid = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return u, err
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
		//need to do
		//MenuServiceApp.UserAuthorityDefaultRouter(&user)
		if user.UserExtend.CircleId != 0 {
			var hkCircle community.Circle
			if global.GVA_DB.Where("id = ?", user.UserExtend.CircleId).First(&hkCircle).Error == nil {
				user.UserExtend.CircleName = hkCircle.Name
			}
		}
	} else {

	}
	return &user, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: userInter *model.SysUser,err error

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

func (appUserService *AppUserService) SetSelfBaseInfo(req communityReq.SetSelfBaseInfoReq) (err error) {
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
	err = db.Where("id = ?", req.ID).Updates(updateData).Error
	return err
}

// GetUser 根据id获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetUser(id uint64) (hkUser community.User, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUser).Error
	return
}

// GetUserInfoList 分页获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetUserInfoList(info communityReq.UserSearch) (list []community.User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.User{})
	var hkUsers []community.User
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&hkUsers).Error
	return hkUsers, total, err
}

// AppGetUserInfoList 分页获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) AppGetUserInfoList(info communityReq.UserSearch) (list []community.User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.User{})
	var hkUsers []community.User
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if len(info.Account) > 0 {
		db = db.Where("title account '?%'", info.Account)
	}

	if len(info.NickName) > 0 {
		db = db.Where("nick_name LIKE '?%'", info.NickName)
	}
	if info.Sex != nil {
		db = db.Where("sex = ?", info.Sex)
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
	channel = data.ChannelId
	return
}

// UpdateUserChannel 更新用户频道
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) UpdateUserChannel(userId uint64, channel string) (err error) {
	var updateData map[string]interface{}
	updateData = make(map[string]interface{})
	updateData["channel_id"] = channel

	db := global.GVA_DB.Model(&community.UserExtend{})
	err = db.Where("id = ?", userId).Updates(updateData).Error
	return err
}
