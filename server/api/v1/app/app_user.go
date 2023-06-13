package app

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	imReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open/model/request"
	openImService "github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
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
	if global.GVA_CONFIG.Param.UseSmsCheckCode == true {
		key := fmt.Sprintf("%d-%s", utils.VerificationBindTel, req.Telephone)
		cacheCaptcha, err := cacheSmsService.GetCacheSms(key)
		if err == redis.Nil {
			response.FailWithMessage("验证码失效", c)
			return
		} else if err != nil {
			global.GVA_LOG.Error("获取验证码失败", zap.Error(err))
			response.FailWithMessage("获取验证码失败", c)
			return
		}
		if cacheCaptcha.Overtime.Before(time.Now()) {
			response.FailWithMessage("验证码失效", c)
			return
		}
		if cacheCaptcha.Code != req.Captcha {
			response.FailWithMessage("验证码错误", c)
			return
		}
	} else {
		if global.GVA_CONFIG.Param.DefaultSmsCheckCode != req.Captcha {
			response.FailWithMessage("验证码错误", c)
			return
		}
	}

	userInfo, err := appUserService.GetUserByPhone(req.Telephone)
	if err == nil && userInfo != nil {
		response.FailWithMessage("电话号码被使用", c)
		return
	} else if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	user, err := appUserService.GetUserBaseInfo(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = appUserService.BindTelephone(userId, req.Telephone)
	if err != nil {
		global.GVA_LOG.Error("绑定电话失败!", zap.Error(err))
		response.FailWithMessage("绑定电话失败", c)
		return
	} else {
		var updateReq imReq.UpdateSelfUserInfoReq
		updateReq.UserID = strconv.FormatUint(user.ID, 10)
		updateReq.Nickname = user.NickName
		updateReq.FaceURL = user.HeaderImg
		updateReq.Gender = user.Sex
		updateReq.PhoneNumber = req.Telephone
		updateReq.BirthStr = user.Birthday
		updateReq.Email = user.Email
		updateReq.CreateTime = user.CreatedAt.Unix()
		//need to do 更新失败，需要单独处理
		if rsp, err := openImService.ServiceGroupApp.UpdateUserInfo(updateReq); err != nil {
			global.GVA_LOG.Error("调用IM失败：UpdateUserInfo."+err.Error(), zap.Error(err))
		} else if rsp.Code != 0 {
			global.GVA_LOG.Error("调用IM失败：UpdateUserInfo."+err.Error(), zap.Error(err))
		}
	}
	response.OkWithMessage("成功", c)
	return
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
	var req communityReq.BindEmailReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	key := fmt.Sprintf("email-%d-%s", utils.EmailCodeBindEmail, req.Email)
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
		} else if userInfo, err := appUserService.GetUserByEmail(req.Email); userInfo != nil {
			response.FailWithMessage("邮箱被使用", c)
			return
		} else if err == nil {
			userId := utils.GetUserID(c)
			err = appUserService.BindEmail(userId, req.Email)
			if err != nil {
				global.GVA_LOG.Error("绑定邮箱失败!", zap.Error(err))
				response.FailWithMessage("绑定邮箱失败", c)
				return
			} else {
				user, err := appUserService.GetUserBaseInfo(userId)
				if err != nil {
					global.GVA_LOG.Error("GetUserBaseInfo!", zap.Error(err))
					response.FailWithMessage(err.Error(), c)
					return
				}
				var updateReq imReq.UpdateSelfUserInfoReq
				updateReq.UserID = strconv.FormatUint(user.ID, 10)
				updateReq.Nickname = user.NickName
				updateReq.FaceURL = user.HeaderImg
				updateReq.Gender = user.Sex
				updateReq.PhoneNumber = user.Phone
				updateReq.BirthStr = user.Birthday
				updateReq.Email = req.Email
				updateReq.CreateTime = user.CreatedAt.Unix()
				//need to do 更新失败，需要单独处理
				if rsp, err := openImService.ServiceGroupApp.UpdateUserInfo(updateReq); err != nil {
					global.GVA_LOG.Error("调用IM失败：UpdateUserInfo."+err.Error(), zap.Error(err))
				} else if rsp.Code != 0 {
					global.GVA_LOG.Error("调用IM失败：UpdateUserInfo."+err.Error(), zap.Error(err))
				}
			}
			response.OkWithMessage("成功", c)
			return
		} else {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
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
	user, err := appUserService.SetSelfBaseInfo(userId, req)

	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		var req imReq.UpdateSelfUserInfoReq
		req.UserID = strconv.FormatUint(user.ID, 10)
		req.Nickname = user.NickName
		req.FaceURL = user.HeaderImg
		req.Gender = user.Sex
		req.PhoneNumber = user.Phone
		req.BirthStr = user.Birthday
		req.Email = user.Email
		req.CreateTime = user.CreatedAt.Unix()
		//need to do 更新失败，需要单独处理
		if rsp, err := openImService.ServiceGroupApp.UpdateUserInfo(req); err != nil {
			global.GVA_LOG.Error("调用IM失败：UpdateUserInfo."+err.Error(), zap.Error(err))
		} else if rsp.Code != 0 {
			global.GVA_LOG.Error("调用IM失败：UpdateUserInfo."+err.Error(), zap.Error(err))
		}
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

// SetUserCoverImage 设置用户主页封面
// @Tags      用户
// @Summary   设置用户主页封面
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      communityReq.SetUserCoverImageReq    true  "设置用户主页封面"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router    /app/user/setUserCoverImage [post]
func (userApi *UserApi) SetUserCoverImage(c *gin.Context) {
	var req communityReq.SetUserCoverImageReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	err = appUserService.UpdateUserCoverImage(userId, req.CoverImage)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}
