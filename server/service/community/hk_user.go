package community

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type HkUserService struct {
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: userInter community.HkUser, err error

func (hkUserService *HkUserService) Register(u community.HkUser) (userInter community.HkUser, err error) {
	var user community.HkUser
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

func (hkUserService *HkUserService) Login(u *community.HkUser) (userInter *community.HkUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user community.HkUser
	err = global.GVA_DB.Where("account = ?", u.Account).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		//need to do
		//MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: userInter *model.SysUser,err error

func (hkUserService *HkUserService) ChangePassword(u *community.HkUser, newPassword string) (userInter *community.HkUser, err error) {
	var user community.HkUser
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

// CreateHkUser 创建HkUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserService *HkUserService) CreateHkUser(hkUser community.HkUser) (err error) {
	err = global.GVA_DB.Create(&hkUser).Error
	return err
}

// DeleteHkUser 删除HkUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserService *HkUserService) DeleteHkUser(hkUser community.HkUser) (err error) {
	err = global.GVA_DB.Delete(&hkUser).Error
	return err
}

// DeleteHkUserByIds 批量删除HkUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserService *HkUserService) DeleteHkUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]community.HkUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHkUser 更新HkUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserService *HkUserService) UpdateHkUser(hkUser community.HkUser) (err error) {
	err = global.GVA_DB.Save(&hkUser).Error
	return err
}

// GetHkUser 根据id获取HkUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserService *HkUserService) GetHkUser(id uint) (hkUser community.HkUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&hkUser).Error
	return
}

// GetHkUserInfoList 分页获取HkUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserService *HkUserService) GetHkUserInfoList(info communityReq.HkUserSearch) (list []community.HkUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkUser{})
	var hkUsers []community.HkUser
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

// AppGetHkUserInfoList 分页获取HkUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (hkUserService *HkUserService) AppGetHkUserInfoList(info appReq.UserSearch) (list []community.HkUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.HkUser{})
	var hkUsers []community.HkUser
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
