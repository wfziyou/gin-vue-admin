package app

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	authReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/auth/request"
	authRes "github.com/flipped-aurora/gin-vue-admin/server/model/app/auth/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	emailService "github.com/flipped-aurora/gin-vue-admin/server/plugin/email/service"
	openImReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open/model/request"
	openImService "github.com/flipped-aurora/gin-vue-admin/server/plugin/im-open/service"
	yunXinImReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/im-yunxin/model/request"
	yunXinImService "github.com/flipped-aurora/gin-vue-admin/server/plugin/im-yunxin/service"
	oneLoginService "github.com/flipped-aurora/gin-vue-admin/server/plugin/oneLogin/service"
	smsService "github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"time"
)

type AuthApi struct {
}

// GetThirdPlat 获取第三方平台信息
// @Tags 鉴权认证
// @Summary  获取第三方平台信息
// @Produce   application/json
// @Success  200   {object}  response.Response{data=[]authRes.ThirdPlatInfo,msg=string}  "返回[]authRes.ThirdPlatInfo"
// @Router   /app/auth/getThirdPlat [post]
func (authApi *AuthApi) GetThirdPlat(c *gin.Context) {
	if list, err := hkThirdPlatformService.GetThirdPlatformInfoList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// LoginPwd 用户登录(账号密码)
// @Tags 鉴权认证
// @Summary  用户登录(账号密码)
// @Produce   application/json
// @Param    data  body      authReq.LoginPwd  true  "用户登录(账号密码)"
// @Success  200   {object}  response.Response{data=authRes.LoginResponse,msg=string}  "返回LoginResponse"
// @Router   /app/auth/loginPwd [post]
func (authApi *AuthApi) LoginPwd(c *gin.Context) {
	var req authReq.LoginPwd
	err := c.ShouldBindJSON(&req)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.AppLoginVerify)
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
		u := &community.User{Account: req.Account, Password: req.Password}
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
		TokenNext(c, user, req.Platform)
		return
	} else if openCaptcha == 0 {
		u := &community.User{Account: req.Account, Password: req.Password}
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
		TokenNext(c, user, req.Platform)
		return
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("连续登录失败次数超过最大次数，请稍后再试", c)
}

// LoginTelephone 用户登录(手机)
// @Tags 鉴权认证
// @Summary  用户登录(手机)
// @Produce   application/json
// @Param    data  body      authReq.LoginTelephone   true  "用户登录(手机)"
// @Success  200   {object}  response.Response{data=authRes.LoginResponse,msg=string}  "返回LoginResponse"
// @Router   /app/auth/loginTelephone [post]
func (authApi *AuthApi) LoginTelephone(c *gin.Context) {
	var req authReq.LoginTelephone
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	key := fmt.Sprintf("%d-%s", utils.VerificationLogin, req.Telephone)
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
		} else if user, err := appUserService.GetUserByPhone(req.Telephone); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else if user == nil {
			userInfo := &community.User{Phone: req.Telephone}
			userObj, err := appUserService.Register(*userInfo)
			if err != nil {
				global.GVA_LOG.Error("注册失败!", zap.Error(err))
				response.FailWithMessage(err.Error(), c)
				return
			}
			user = &userObj
			err = ImRegiser(user, req.Platform)
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}
			TokenNext(c, user, req.Platform)
			return
		} else {
			if user.Status != 0 {
				global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
				response.FailWithMessage("用户被禁止登录", c)
				return
			}
			TokenNext(c, user, req.Platform)
			return
		}
	}
}

// LoginThird 用户登录(第三方授权)
// @Tags 鉴权认证
// @Summary  用户登录(第三方授权)
// @Produce   application/json
// @Param data body authReq.LoginThird true "用户登录(第三方授权)"
// @Success  200   {object}  response.Response{data=authRes.LoginResponse,msg=string}  "返回LoginResponse"
// @Router   /app/auth/loginThird [post]
func (authApi *AuthApi) LoginThird(c *gin.Context) {
	var req authReq.LoginThird
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

}

// LoginOneClick 一键登录
// @Tags 鉴权认证
// @Summary 一键登录
// @Produce application/json
// @Param data body authReq.LoginOneClick true "一键登录"
// @Success  200   {object}  response.Response{data=authRes.LoginResponse,msg=string}  "返回LoginResponse"
// @Router /app/auth/loginOneClick [post]
func (authApi *AuthApi) LoginOneClick(c *gin.Context) {
	var req authReq.LoginOneClick
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if rsp, err := oneLoginService.ServiceGroupApp.LoginTokenValidate(req.Token); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else if user, err := appUserService.GetUserByPhone(rsp.Telephone); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else if user == nil {
		userInfo := &community.User{Phone: rsp.Telephone}
		userObj, err := appUserService.Register(*userInfo)
		if err != nil {
			global.GVA_LOG.Error("注册失败!", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
			return
		}
		user = &userObj
		err = ImRegiser(user, req.Platform)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		TokenNext(c, user, req.Platform)
		return
	} else {
		if user.Status != 0 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		TokenNext(c, user, req.Platform)
		return
	}
}

// GetLocalTelephone 获取本机手机号码
// @Tags 鉴权认证
// @Summary 获取本机手机号码
// @Produce application/json
// @Param data body authReq.LoginOneClick true "获取本机手机号码"
// @Success  200   {object}  response.Response{data=authRes.LocalTelephone,msg=string}  "返回LocalTelephone"
// @Router /app/auth/getLocalTelephone [post]
func (authApi *AuthApi) GetLocalTelephone(c *gin.Context) {
	var req authReq.ParamGetLocalTelephone
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if rsp, err := oneLoginService.ServiceGroupApp.LoginTokenValidate(req.Token); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithDetailed(authRes.LocalTelephone{Phone: rsp.Telephone}, "成功", c)
		return
	}
}

// Register 注册
// @Tags 鉴权认证
// @Summary  注册
// @Produce   application/json
// @Param    data  body authReq.Register true  "注册"
// @Success  200   {object}  response.Response{data=authRes.RegisterResponse,msg=string}  "返回RegisterResponse"
// @Router   /app/auth/register [post]
func (authApi *AuthApi) Register(c *gin.Context) {
	var req authReq.Register
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user := &community.User{Account: req.Account, Password: req.Password}
	userInfo, err := appUserService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		err = ImRegiser(&userInfo, req.Platform)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
	response.OkWithDetailed(authRes.RegisterResponse{User: userInfo}, "注册成功", c)
}

// GetSmsVerification 获取短信验证码
// @Tags 鉴权认证
// @Summary   获取短信验证码
// @accept    application/json
// @Produce   application/json
// @Param data body authReq.CaptchaReq true "获取短信验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router    /app/auth/getSmsVerification [post]
func (authApi *AuthApi) GetSmsVerification(c *gin.Context) {
	var obj authReq.CaptchaReq
	_ = c.ShouldBindJSON(&obj)

	//类型：0 测试，1注册，2修改密码，3绑定电话，4忘记密码，5绑定银行
	var TemplateCode string = global.GVA_CONFIG.AliyunSms.SmsTemplate.Test
	if obj.Type == utils.VerificationRegister {
		TemplateCode = global.GVA_CONFIG.AliyunSms.SmsTemplate.Register
	} else if obj.Type == utils.VerificationChangePwd {
		TemplateCode = global.GVA_CONFIG.AliyunSms.SmsTemplate.ChangePwd
	} else if obj.Type == utils.VerificationBindTel {
		TemplateCode = global.GVA_CONFIG.AliyunSms.SmsTemplate.BindTel
	} else if obj.Type == utils.VerificationResetPwd {
		TemplateCode = global.GVA_CONFIG.AliyunSms.SmsTemplate.ResetPwd
	} else if obj.Type == utils.VerificationBindBank {
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

		if err := smsService.ServiceGroupApp.SendAliSms([]string{obj.Telephone}, TemplateCode, code); err != nil {
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

		if err := smsService.ServiceGroupApp.SendAliSms([]string{obj.Telephone}, TemplateCode, code); err != nil {
			global.GVA_LOG.Error("发送验证码失败!", zap.Error(err))
			response.FailWithMessage("发送验证码失败", c)
			return
		} else {
			response.OkWithData("发送成功", c)
			return
		}
	}
}

// SendEmailVerification 发送Email验证码
// @Tags 鉴权认证
// @Summary   发送Email验证码
// @accept    application/json
// @Produce   application/json
// @Param data body authReq.EmailVerificationReq true "发送Email验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router    /app/auth/sendEmailVerification [post]
func (authApi *AuthApi) SendEmailVerification(c *gin.Context) {
	var obj authReq.EmailVerificationReq
	_ = c.ShouldBindJSON(&obj)

	//类型：0绑定邮箱
	var subject string = global.GVA_CONFIG.Email.EmailTemplate.BindEmailSubject
	var code string = fmt.Sprintf(global.GVA_CONFIG.Email.EmailTemplate.BindEmailBody, rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	//if obj.Type == utils.EmailCodeBindEmail {
	//	subject = global.GVA_CONFIG.Email.EmailTemplate.BindEmailSubject
	//	code = fmt.Sprintf(global.GVA_CONFIG.Email.EmailTemplate.BindEmailBody, rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	//}
	key := fmt.Sprintf("email-%d-%s", obj.Type, obj.Email)
	if cacheCaptcha, err := cacheSmsService.GetCacheSms(key); err == redis.Nil {
		if err := cacheSmsService.SetCacheSms(key, code); err != nil {
			global.GVA_LOG.Error("设置验证码到缓存失败!", zap.Error(err))
			response.FailWithMessage("设置验证码到缓存失败", c)
			return
		}

		if err := emailService.ServiceGroupApp.SendEmail(obj.Email, subject, code); err != nil {
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
		if cacheCaptcha.Overtime.After(time.Now()) {
			response.FailWithMessage("发送太频繁，稍后再试", c)
			return
		}

		if err := cacheSmsService.SetCacheSms(key, code); err != nil {
			global.GVA_LOG.Error("设置验证码到缓存失败!", zap.Error(err))
			response.FailWithMessage("设置验证码到缓存失败", c)
			return
		}

		if err := emailService.ServiceGroupApp.SendEmail(obj.Email, subject, code); err != nil {
			global.GVA_LOG.Error("发送验证码失败!", zap.Error(err))
			response.FailWithMessage("发送验证码失败", c)
			return
		} else {
			response.OkWithData("发送成功", c)
			return
		}
	}
}

// ResetPassword 重置密码
// @Tags 鉴权认证
// @Summary   重置密码
// @Produce  application/json
// @Param     data  body      authReq.ResetPasswordReq    true  "重置密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router    /app/auth/resetPassword [post]
func (authApi *AuthApi) ResetPassword(c *gin.Context) {
	var req authReq.ResetPasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	key := fmt.Sprintf("%d-%s", utils.VerificationResetPwd, req.Telephone)
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
		} else if user, err := appUserService.GetUserByPhone(req.Telephone); err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		} else if user == nil {
			response.FailWithMessage("账号不存在", c)
			return
		} else {
			if user.Status != 0 {
				global.GVA_LOG.Error("账号被禁止")
				response.FailWithMessage("账号被禁止", c)
				return
			}
			err = appUserService.ResetPassword(user, req.Password)
			if err != nil {
				global.GVA_LOG.Error("修改密码失败!", zap.Error(err))
				response.FailWithMessage("修改密码失败", c)
				return
			}
			response.OkWithMessage("修改成功", c)
			return
		}
	}
}

// TokenNext 登录以后签发jwt
func TokenNext(c *gin.Context, user *community.User, platform int) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.Uuid,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Account,
		AuthorityId: user.AuthorityId,
	})
	clientIp := c.ClientIP()
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if err := ImLogin(user, platform, clientIp); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//var req imReq.UpdateActionReq
	//req.Accid = user.Account
	//req.Token = token
	//if rsp, err := imService.ServiceGroupApp.UpdateAction(req); err != nil {
	//	response.FailWithMessage("调研IM失败", c)
	//	return
	//} else if rsp.Code != 200 {
	//	response.FailWithMessage("更新IM会话失败", c)
	//	return
	//}

	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(authRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Account); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Account); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(authRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
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
		response.OkWithDetailed(authRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

func ImLogin(user *community.User, platform int, clientIp string) error {
	if global.GVA_CONFIG.System.ImType == "open-im" {
		userToken := openImReq.UserTokenReq{
			Platform: platform,
			UserID:   strconv.FormatUint(user.ID, 10),
			LoginIp:  clientIp,
		}
		if rsp, err := openImService.ServiceGroupApp.UserToken(userToken); err != nil {
			global.GVA_LOG.Debug("open-im：GetUserInfo."+err.Error(), zap.Error(err))
			return err
		} else if rsp.Code == 802 {
			err = ImRegiser(user, platform)
			if err != nil {
				return err
			} else {
				_, err = openImService.ServiceGroupApp.UserToken(userToken)
			}
		} else if rsp.Code != 0 {
			global.GVA_LOG.Debug("open-im：UserToken code:" + strconv.Itoa(rsp.Code))
			return errors.New("open-im：UserToken code:" + strconv.Itoa(rsp.Code))
		} else {
			user.ImToken = rsp.UserToken.Token
		}
	} else if global.GVA_CONFIG.System.ImType == "yunxin-im" {
		var req yunXinImReq.UserGetUinfosActionReq
		req.Accids = []string{utils.UuidTo32String(user.Uuid)}
		if rsp, err := yunXinImService.ServiceGroupApp.UserGetUinfosAction(req); err != nil {
			global.GVA_LOG.Debug("调用IM失败：UserGetUinfosAction."+err.Error(), zap.Error(err))
			return err
		} else if rsp.Code == 414 {
			err = ImRegiser(user, platform)
			if err != nil {
				return err
			}
		} else {
			global.GVA_LOG.Debug("调用IM：UserGetUinfosAction code:" + strconv.Itoa(rsp.Code))
		}
	}
	return nil
}

// LoginPwdTest 用户登录(账号密码)
// @Tags 鉴权认证
// @Summary  用户登录(账号密码)
// @Produce   application/json
// @Param    data  body      authReq.LoginPwd  true  "用户登录(账号密码)"
// @Success  200   {object}  response.Response{data=authRes.LoginResponse,msg=string}  "返回LoginResponse"
// @Router   /app/auth/loginPwdTest [post]
func (authApi *AuthApi) LoginPwdTest(c *gin.Context) {
	var req authReq.LoginPwd
	err := c.ShouldBindJSON(&req)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.AppLoginVerify)
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
		u := &community.User{Account: req.Account, Password: req.Password}
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
		TokenNextTest(c, user, req.Platform)
		return
	} else if openCaptcha == 0 {
		u := &community.User{Account: req.Account, Password: req.Password}
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
		TokenNextTest(c, user, req.Platform)
		return
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("连续登录失败次数超过最大次数，请稍后再试", c)
}

func TokenNextTest(c *gin.Context, user *community.User, platform int) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.Uuid,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Account,
		AuthorityId: user.AuthorityId,
	})
	clientIp := c.ClientIP()
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if err := ImLoginTest(user, platform, clientIp); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//var req imReq.UpdateActionReq
	//req.Accid = user.Account
	//req.Token = token
	//if rsp, err := imService.ServiceGroupApp.UpdateAction(req); err != nil {
	//	response.FailWithMessage("调研IM失败", c)
	//	return
	//} else if rsp.Code != 200 {
	//	response.FailWithMessage("更新IM会话失败", c)
	//	return
	//}

	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(authRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Account); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Account); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(authRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
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
		response.OkWithDetailed(authRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}
func ImLoginTest(user *community.User, platform int, clientIp string) error {
	userToken := openImReq.UserTokenReq{
		Platform: platform,
		UserID:   strconv.FormatUint(user.ID, 10),
		LoginIp:  clientIp,
	}
	if rsp, err := openImService.ServiceGroupApp.UserToken(userToken); err != nil {
		global.GVA_LOG.Debug("open-im：GetUserInfo."+err.Error(), zap.Error(err))
		return err
	} else if rsp.Code == 802 {
		err = ImRegiserTest(user, platform)
		if err != nil {
			return err
		} else {
			_, err = openImService.ServiceGroupApp.UserToken(userToken)
		}
	} else if rsp.Code != 0 {
		global.GVA_LOG.Debug("open-im：UserToken code:" + strconv.Itoa(rsp.Code))
		return errors.New("open-im：UserToken code:" + strconv.Itoa(rsp.Code))
	} else {
		user.ImToken = rsp.UserToken.Token
	}
	return nil
}
func ImRegiserTest(user *community.User, platform int) error {
	var req openImReq.RegisterReq
	req.Platform = platform
	req.UserID = strconv.FormatUint(user.ID, 10)
	req.Nickname = user.NickName
	req.FaceURL = user.HeaderImg
	req.Gender = user.Sex
	req.PhoneNumber = user.Phone
	req.BirthStr = user.Birthday
	req.Email = user.Email
	req.CreateTime = user.CreatedAt.Unix()

	if rsp, err := openImService.ServiceGroupApp.UserRegister(req); err != nil {
		global.GVA_LOG.Error("调用IM失败：UserCreateAction."+err.Error(), zap.Error(err))
		return err
	} else if rsp.Code != 0 {
		global.GVA_LOG.Error("调用IM失败：UserCreateAction."+err.Error(), zap.Error(err))
		return errors.New(fmt.Sprintf("Code:%d ErrMsg:%s", rsp.Code, rsp.ErrMsg))
	} else {
		user.ImToken = rsp.UserToken.Token
	}
	return nil
}

// ImRegiser 调用IM注册
func ImRegiser(user *community.User, platform int) error {
	if global.GVA_CONFIG.System.ImType == "open-im" {
		var req openImReq.RegisterReq
		req.Platform = platform
		req.UserID = strconv.FormatUint(user.ID, 10)
		req.Nickname = user.NickName
		req.FaceURL = user.HeaderImg
		req.Gender = user.Sex
		req.PhoneNumber = user.Phone
		req.BirthStr = user.Birthday
		req.Email = user.Email
		req.CreateTime = user.CreatedAt.Unix()

		if rsp, err := openImService.ServiceGroupApp.UserRegister(req); err != nil {
			global.GVA_LOG.Error("调用IM失败：UserCreateAction."+err.Error(), zap.Error(err))
			return err
		} else if rsp.Code != 0 {
			global.GVA_LOG.Error("调用IM失败：UserCreateAction."+err.Error(), zap.Error(err))
			return errors.New(fmt.Sprintf("Code:%d ErrMsg:%s", rsp.Code, rsp.ErrMsg))
		} else {
			user.ImToken = rsp.UserToken.Token
		}
	} else if global.GVA_CONFIG.System.ImType == "yunxin-im" {
		var req yunXinImReq.RegisterReq
		req.Accid = utils.UuidTo32String(user.Uuid)
		req.Token = utils.UuidTo32String(user.Uuid)
		req.Name = user.NickName
		req.Icon = user.HeaderImg
		if rsp, err := yunXinImService.ServiceGroupApp.UserCreateAction(req); err != nil {
			global.GVA_LOG.Error("调用IM失败：UserCreateAction."+err.Error(), zap.Error(err))
			return err
		} else if rsp.Code != 414 && rsp.Code != 200 {
			global.GVA_LOG.Error("调用IM失败：UserCreateAction."+err.Error(), zap.Error(err))
			return errors.New(fmt.Sprintf("Code:%d ErrMsg:%s", rsp.Code, rsp.Desc))
		}
	}
	return nil
}
