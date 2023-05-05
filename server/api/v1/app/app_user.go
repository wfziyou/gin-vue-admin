package app

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type UserApi struct {
}

/*************************************
用户
**************************************/

// BindTelephone 绑定手机
// @Tags 用户
// @Summary  绑定手机
// @Produce   application/json
// @Param    data  body      communityReq.BindTelephoneReq   true  "绑定手机"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router   /app/user/bindTelephone [post]
func (userApi *UserApi) BindTelephone(c *gin.Context) {
	var req communityReq.BindTelephoneReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	key := fmt.Sprintf("%d-%s", utils.VerificationBindTel, req.Telephone)
	if cacheCaptcha, err := cacheSmsService.GetCacheSms(key); err == redis.Nil {
		response.FailWithMessage("验证码失效", c)
		return
	} else if err != nil {
		global.GVA_LOG.Error("获取验证码失败", zap.Error(err))
		response.FailWithMessage("获取验证码失败", c)
		return
	} else {
		if cacheCaptcha.Overtime.Before(time.Now()) {
			response.FailWithMessage("验证码失效", c)
			return
		}
		if cacheCaptcha.Code != req.Captcha {
			response.FailWithMessage("验证码错误", c)
			return
		} else if _, err := appUserService.GetUserByPhone(req.Telephone); err == nil {
			response.FailWithMessage("电话号码被使用", c)
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			userId := utils.GetUserID(c)
			err = appUserService.BindTelephone(userId, req.Telephone)
			if err != nil {
				global.GVA_LOG.Error("绑定电话失败!", zap.Error(err))
				response.FailWithMessage("绑定电话失败", c)
				return
			}
			response.OkWithMessage("成功", c)
			return
		} else {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
}

// BindEmail 绑定邮箱
// @Tags 用户
// @Summary  绑定邮箱
// @Produce   application/json
// @Param    data  body      communityReq.BindEmailReq  true  "绑定邮箱"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router   /app/user/bindEmail [post]
func (userApi *UserApi) BindEmail(c *gin.Context) {
	//var l communityReq.BindEmailReq

}

// GetUserBaseInfo 用id查询用户基本信息
// @Tags 用户
// @Summary 用id查询用户基本信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询用户基本信息"
// @Success  200   {object}  response.Response{data=community.UserBaseInfo,msg=string}  "返回common.UserBaseInfo"
// @Router /app/user/getUserBaseInfo [get]
func (userApi *UserApi) GetUserBaseInfo(c *gin.Context) {
	var idSearch request.IdSearch

	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if user, err := appUserService.GetUser(idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(community.UserBaseInfo{
			ID:          user.ID,
			UserType:    user.UserType,
			Account:     user.Account,
			NickName:    user.NickName,
			RealName:    user.RealName,
			HeaderImg:   user.HeaderImg,
			Birthday:    user.Birthday,
			Sex:         user.Sex,
			Description: user.Description,
		}, c)
	}
}

// GetUserInfo 用id查询用户信息
// @Tags 用户
// @Summary 用id查询用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询用户信息"
// @Success  200   {object}  response.Response{data=community.UserInfo,msg=string}  "返回community.UserInfo"
// @Router /app/user/getUserInfo [get]
func (userApi *UserApi) GetUserInfo(c *gin.Context) {
	var idSearch request.IdSearch

	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId := utils.GetUserID(c)
	if userInfo, err := appUserService.GetUserInfo(userId, idSearch.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		hkRecordBrowsingUserHomepageService.BrowsingUser(userId, userInfo.ID)
		response.OkWithData(userInfo, c)
	}
}

// SetSelfBaseInfo
// @Tags      用户
// @Summary   设置用户基础信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      communityReq.SetSelfBaseInfoReq    true  "设置用户基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router    /app/user/setSelfBaseInfo [put]
func (userApi *UserApi) SetSelfBaseInfo(c *gin.Context) {
	var req communityReq.SetSelfBaseInfoReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	err = appUserService.SetSelfBaseInfo(userId, req)

	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// GetUserList 分页获取用户列表
// @Tags 用户
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.UserSearch true "分页获取用户列表"
// @Success 200 {object}  response.PageResult{List=[]community.User,msg=string} "返回common.User"
// @Router /app/user/getUserList [get]
func (userApi *UserApi) GetUserList(c *gin.Context) {
	var req communityReq.UserSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appUserService.GetUserInfoList(req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}
