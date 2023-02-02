package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	communityRes "github.com/flipped-aurora/gin-vue-admin/server/model/community/response"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
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

// GetCaptcha 获取验证码
// @Tags      Base
// @Summary   获取验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data query appReq.CaptchaReq true "获取验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router    /app/user/getCaptcha [post]
func (userApi *UserApi) GetCaptcha(c *gin.Context) {
	//// 判断验证码是否开启
	//openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	//openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	//key := c.ClientIP()
	//v, ok := global.BlackCache.Get(key)
	//if !ok {
	//	global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	//}
	//
	//var oc bool
	//if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
	//	oc = true
	//}
	//// 字符,公式,验证码配置
	//// 生成默认数字的driver
	//driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	//// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	//cp := base64Captcha.NewCaptcha(driver, store)
	//id, b64s, err := cp.Generate()
	//if err != nil {
	//	global.GVA_LOG.Error("验证码获取失败!", zap.Error(err))
	//	response.FailWithMessage("验证码获取失败", c)
	//	return
	//}
	//response.OkWithDetailed(systemRes.SysCaptchaResponse{
	//	CaptchaId:     id,
	//	PicPath:       b64s,
	//	CaptchaLength: global.GVA_CONFIG.Captcha.KeyLong,
	//	OpenCaptcha:   oc,
	//}, "验证码获取成功", c)
}

// Register 注册
// @Tags     App_User
// @Summary  注册
// @Produce   application/json
// @Param    data  body     appReq.Register true  "注册"
// @Success  200   {object}  response.Response{data=communityRes.SysUserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /app/user/register [post]
func (userApi *UserApi) Register(c *gin.Context) {
	var r appReq.Register
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

	user := &community.HkUser{Account: r.Account, NickName: r.NickName, Password: r.Password}
	userReturn, err := hkUserService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(communityRes.SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(communityRes.SysUserResponse{User: userReturn}, "注册成功", c)
}

// LoginPwd 用户登录(账号密码)
// @Tags     App_User
// @Summary  用户登录(账号密码)
// @Produce   application/json
// @Param    data  body      appReq.LoginPwd  true  "用户登录(账号密码)"
// @Success  200   {object}  response.Response{data=communityRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /app/user/loginPwd [post]
func (userApi *UserApi) LoginPwd(c *gin.Context) {
	var l appReq.LoginPwd
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

	if !oc || true {
		u := &community.HkUser{Account: l.Account, Password: l.Password}
		user, err := hkUserService.Login(u)
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
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}

// LoginTelephone 用户登录(手机)
// @Tags     App_User
// @Summary  用户登录(手机)
// @Produce   application/json
// @Param    data  body      appReq.LoginTelephone   true  "用户登录(手机)"
// @Success  200   {object}  response.Response{data=communityRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /app/user/loginTelephone [post]
func (userApi *UserApi) LoginTelephone(c *gin.Context) {
	//var l appReq.LoginTelephone

}

// LoginThird 用户登录(第三方授权)
// @Tags     App_User
// @Summary  用户登录(第三方授权)
// @Produce   application/json
// @Param    data  body      appReq.LoginThird    true  "用户登录(第三方授权)"
// @Success  200   {object}  response.Response{data=communityRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /app/user/loginThird [post]
func (userApi *UserApi) LoginThird(c *gin.Context) {
	//var l appReq.LoginThird

}

// TokenNext 登录以后签发jwt
func (userApi *UserApi) TokenNext(c *gin.Context, user community.HkUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.Uuid,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Account,
		AuthorityId: user.RoleId,
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
// @Param     data  body      appReq.ResetPasswordReq    true  "重置密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router    /app/user/resetPassword [post]
func (userApi *UserApi) ResetPassword(c *gin.Context) {
	var req appReq.ResetPasswordReq
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
	u := &community.HkUser{GVA_MODEL: global.GVA_MODEL{ID: uid}, Password: req.Password}
	_, err = hkUserService.ChangePassword(u, req.Password)
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
// @Param    data  body      appReq.BindTelephoneReq   true  "绑定手机"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router   /app/user/bindTelephone [post]
func (userApi *UserApi) BindTelephone(c *gin.Context) {
	//var l appReq.BindTelephoneReq

}

// BindEmail 绑定邮箱
// @Tags     App_User
// @Summary  绑定邮箱
// @Produce   application/json
// @Param    data  body      appReq.BindEmailReq  true  "绑定邮箱"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router   /app/user/bindEmail [post]
func (userApi *UserApi) BindEmail(c *gin.Context) {
	//var l appReq.BindEmailReq

}

// GetUserBaseInfo 用id查询UserBaseInfo
// @Tags App_User
// @Summary 用id查询UserBaseInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.IdSearch true "用id查询UserBaseInfo"
// @Success  200   {object}  response.Response{data=common.UserBaseInfo,msg=string}  "返回UserBaseInfo"
// @Router /app/user/getUserBaseInfo [get]
func (userApi *UserApi) GetUserBaseInfo(c *gin.Context) {
	var idSearch request.IdSearch

	err := c.ShouldBindQuery(&idSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if user, err := hkUserService.GetHkUser(idSearch.ID); err != nil {
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
// @Param     data  body      appReq.SetSelfBaseInfoReq    true  "设置用户基础信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router    /app/user/setSelfBaseInfo [put]
func (userApi *UserApi) SetSelfBaseInfo(c *gin.Context) {
	var user appReq.SetSelfBaseInfoReq
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = utils.GetUserID(c)
	err = hkUserService.UpdateHkUser(community.HkUser{
		GVA_MODEL: global.GVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Birthday:  user.Birthday,
		Sex:       user.Sex,
	})

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
// @Param data query appReq.UserSearch true "分页获取User列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/user/getUserList [get]
func (userApi *UserApi) GetUserList(c *gin.Context) {
	var pageInfo appReq.UserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hkUserService.AppGetHkUserInfoList(pageInfo); err != nil {
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
