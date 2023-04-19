package app

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	authReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/auth/request"
	authRes "github.com/flipped-aurora/gin-vue-admin/server/model/app/auth/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/community"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	imReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/im/model/request"
	imService "github.com/flipped-aurora/gin-vue-admin/server/plugin/im/service"
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

}

// LoginPwd 用户登录(账号密码)
// @Tags 鉴权认证
// @Summary  用户登录(账号密码)
// @Produce   application/json
// @Param    data  body      authReq.LoginPwd  true  "用户登录(账号密码)"
// @Success  200   {object}  response.Response{data=authRes.LoginResponse,msg=string}  "返回LoginResponse"
// @Router   /app/auth/loginPwd [post]
func (authApi *AuthApi) LoginPwd(c *gin.Context) {
	var l authReq.LoginPwd
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
		TokenNext(c, *user)
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
		TokenNext(c, *user)
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
	//var l authReq.LoginTelephone

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
		err = ImRegiser(userObj)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		TokenNext(c, *user)
		return
	} else {
		if user.Status != 0 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		TokenNext(c, *user)
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
	var r authReq.Register
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
	userInfo, err := appUserService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		err = ImRegiser(userInfo)
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

// TokenNext 登录以后签发jwt
func TokenNext(c *gin.Context, user community.User) {
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

	var req imReq.UserGetUinfosActionReq
	req.Accids = []string{utils.UuidTo32String(user.Uuid)}

	if rsp, err := imService.ServiceGroupApp.UserGetUinfosAction(req); err != nil {
		global.GVA_LOG.Debug("调用IM失败：UserGetUinfosAction."+err.Error(), zap.Error(err))
	} else if rsp.Code == 414 {
		err = ImRegiser(user)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	} else {
		global.GVA_LOG.Debug("调用IM：UserGetUinfosAction code:" + strconv.Itoa(rsp.Code))
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
		response.OkWithDetailed(authRes.LoginResponse{
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
		response.OkWithDetailed(authRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// ImRegiser 调用IM注册
func ImRegiser(user community.User) error {
	var req imReq.RegisterReq
	req.Accid = utils.UuidTo32String(user.Uuid)
	req.Token = utils.UuidTo32String(user.Uuid)
	req.Name = user.NickName
	req.Icon = user.HeaderImg
	if rsp, err := imService.ServiceGroupApp.UserCreateAction(req); err != nil {
		global.GVA_LOG.Error("调用IM失败：UserCreateAction."+err.Error(), zap.Error(err))
		return err
	} else if rsp.Code != 414 && rsp.Code != 200 {
		global.GVA_LOG.Error("调用IM失败：UserCreateAction."+err.Error(), zap.Error(err))
		return err
	}
	return nil
}
