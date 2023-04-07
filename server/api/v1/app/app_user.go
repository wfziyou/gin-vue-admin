package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	//var l communityReq.BindTelephoneReq

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
		time.ParseDuration("20")
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
// @Tags      用户
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
