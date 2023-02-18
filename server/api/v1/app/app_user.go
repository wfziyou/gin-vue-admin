package app

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	communityRes "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service"
	"math/rand"

	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

type UserApi struct {
}

/*************************************
用户
**************************************/

// GetCaptcha 获取验证码
// @Tags      App_User
// @Summary   获取验证码
// @accept    application/json
// @Produce   application/json
// @Param data body communityReq.CaptchaReq true "获取验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router    /app/user/getCaptcha [post]
func (userApi *UserApi) GetCaptcha(c *gin.Context) {
	var obj communityReq.CaptchaReq
	_ = c.ShouldBindJSON(&obj)

	//类型：0 测试，1注册，2修改密码，3绑定电话，4忘记密码，5绑定银行
	var TemplateCode string = global.GVA_CONFIG.AliyunSms.SmsTemplate.Test
	if obj.Type == 1 {
		TemplateCode = global.GVA_CONFIG.AliyunSms.SmsTemplate.Register
	} else if obj.Type == 2 {
		TemplateCode = global.GVA_CONFIG.AliyunSms.SmsTemplate.ChangePwd
	} else if obj.Type == 3 {
		TemplateCode = global.GVA_CONFIG.AliyunSms.SmsTemplate.BindTel
	} else if obj.Type == 4 {
		TemplateCode = global.GVA_CONFIG.AliyunSms.SmsTemplate.ResetPwd
	} else if obj.Type == 5 {
		TemplateCode = global.GVA_CONFIG.AliyunSms.SmsTemplate.BindBank
	}
	code := fmt.Sprintf("{\"code\":\"%06v\"}", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	key := fmt.Sprintf("%d-%s", obj.Type, obj.Telephone)
	if cacheCaptcha, err := cacheSmsService.GetCacheSms(key); err == redis.Nil {
		if err := cacheSmsService.SetCacheSms(key, code); err != nil {
			global.GVA_LOG.Error("设置验证码到缓存失败!", zap.Error(err))
			response.FailWithMessage("设置验证码到缓存失败", c)
			return
		}

		if err := service.ServiceGroupApp.SendAliSms([]string{obj.Telephone}, TemplateCode, code); err != nil {
			global.GVA_LOG.Error("发送验证码失败!", zap.Error(err))
			response.FailWithMessage("发送验证码失败", c)
			return
		} else {
			response.OkWithData("发送成功", c)
			return
		}
	} else if err != nil {
		global.GVA_LOG.Error("设置验证码到缓存失败!", zap.Error(err))
		response.FailWithMessage("设置验证码到缓存失败", c)
	} else {
		//global.GVA_LOG.Error("aaa:" + cacheCaptcha.Overtime.String() + " : " + time.Now().String())
		if cacheCaptcha.Overtime.After(time.Now()) {
			response.FailWithMessage("发送太频繁，稍后再试", c)
			return
		}

		if err := cacheSmsService.SetCacheSms(key, code); err != nil {
			global.GVA_LOG.Error("设置验证码到缓存失败!", zap.Error(err))
			response.FailWithMessage("设置验证码到缓存失败", c)
			return
		}

		if err := service.ServiceGroupApp.SendAliSms([]string{obj.Telephone}, TemplateCode, code); err != nil {
			global.GVA_LOG.Error("发送验证码失败!", zap.Error(err))
			response.FailWithMessage("发送验证码失败", c)
			return
		} else {
			response.OkWithData("发送成功", c)
			return
		}
	}
}

// Register 注册
// @Tags     App_User
// @Summary  注册
// @Produce   application/json
// @Param    data  body communityReq.Register true  "注册"
// @Success  200   {object}  response.Response{data=communityRes.SysUserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /app/user/register [post]
func (userApi *UserApi) Register(c *gin.Context) {
	var r communityReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user := &community.User{Account: r.Account, Password: r.Password}
	userReturn, err := appUserService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(communityRes.SysUserResponse{User: userReturn}, "注册成功", c)
}

// LoginPwd 用户登录(账号密码)
// @Tags     App_User
// @Summary  用户登录(账号密码)
// @Produce   application/json
// @Param    data  body      communityReq.LoginPwd  true  "用户登录(账号密码)"
// @Success  200   {object}  response.Response{data=communityRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /app/user/loginPwd [post]
func (userApi *UserApi) LoginPwd(c *gin.Context) {
	var l communityReq.LoginPwd
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.AppLoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = openCaptcha == 0 || openCaptcha < utils.InterfaceToInt(v)

	if !oc {
		u := &community.User{Account: l.Account, Password: l.Password}
		user, err := appUserService.Login(u)
		if err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Status != 0 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		userApi.TokenNext(c, *user)
		return
	} else if openCaptcha == 0 {
		u := &community.User{Account: l.Account, Password: l.Password}
		user, err := appUserService.Login(u)
		if err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Status != 0 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		userApi.TokenNext(c, *user)
		return
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("连续登录失败次数超过最大次数，请稍后再试", c)
}

// LoginTelephone 用户登录(手机)
// @Tags     App_User
// @Summary  用户登录(手机)
// @Produce   application/json
// @Param    data  body      communityReq.LoginTelephone   true  "用户登录(手机)"
// @Success  200   {object}  response.Response{data=communityRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /app/user/loginTelephone [post]
func (userApi *UserApi) LoginTelephone(c *gin.Context) {
	//var l communityReq.LoginTelephone

}

// LoginThird 用户登录(第三方授权)
// @Tags     App_User
// @Summary  用户登录(第三方授权)
// @Produce   application/json
// @Param    data  body      communityReq.LoginThird    true  "用户登录(第三方授权)"
// @Success  200   {object}  response.Response{data=communityRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /app/user/loginThird [post]
func (userApi *UserApi) LoginThird(c *gin.Context) {
	//var l communityReq.LoginThird

}

// TokenNext 登录以后签发jwt
func (userApi *UserApi) TokenNext(c *gin.Context, user community.User) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.Uuid,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Account,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(communityRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Account); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Account); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(communityRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Account); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(communityRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// ResetPassword 重置密码
// @Tags      App_User
// @Summary   重置密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      communityReq.ResetPasswordReq    true  "重置密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router    /app/user/resetPassword [post]
func (userApi *UserApi) ResetPassword(c *gin.Context) {
	var req communityReq.ResetPasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ResetPasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	u := &community.User{GvaModelApp: global.GvaModelApp{ID: uid}, Password: req.Password}
	_, err = appUserService.ChangePassword(u, req.Password)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// BindTelephone 绑定手机
// @Tags     App_User
// @Summary  绑定手机
// @Produce   application/json
// @Param    data  body      communityReq.BindTelephoneReq   true  "绑定手机"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router   /app/user/bindTelephone [post]
func (userApi *UserApi) BindTelephone(c *gin.Context) {
	//var l communityReq.BindTelephoneReq

}

// BindEmail 绑定邮箱
// @Tags     App_User
// @Summary  绑定邮箱
// @Produce   application/json
// @Param    data  body      communityReq.BindEmailReq  true  "绑定邮箱"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router   /app/user/bindEmail [post]
func (userApi *UserApi) BindEmail(c *gin.Context) {
	//var l communityReq.BindEmailReq

}

// GetUserBaseInfo 用id查询UserBaseInfo
// @Tags App_User
// @Summary 用id查询UserBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询UserBaseInfo"
// @Success  200   {object}  response.Response{data=common.UserBaseInfo,msg=string}  "返回common.UserBaseInfo"
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
		response.OkWithData(common.UserBaseInfo{
			ID:        user.ID,
			NickName:  user.NickName,
			Phone:     user.Phone,
			Email:     user.Email,
			HeaderImg: user.HeaderImg,
			Birthday:  user.Birthday,
			Sex:       user.Sex,
		}, c)
	}
}

// SetSelfBaseInfo
// @Tags      App_User
// @Summary   设置用户基础信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      communityReq.SetSelfBaseInfoReq    true  "设置用户基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router    /app/user/setSelfBaseInfo [put]
func (userApi *UserApi) SetSelfBaseInfo(c *gin.Context) {
	var user communityReq.SetSelfBaseInfoReq
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = utils.GetUserID(c)
	err = appUserService.SetSelfBaseInfo(user)

	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// GetUserList 分页获取User列表
// @Tags App_User
// @Summary 分页获取User列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.UserSearch true "分页获取User列表"
// @Success 200 {object}  response.PageResult{List=[]community.User,msg=string} "返回common.User"
// @Router /app/user/getUserList [get]
func (userApi *UserApi) GetUserList(c *gin.Context) {
	var pageInfo communityReq.UserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := appUserService.AppGetUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
