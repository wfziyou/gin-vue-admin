package app

import "github.com/gin-gonic/gin"

type FocusApi struct {
}

// GetFrequentBrowsingUserList 获取经常浏览用户列表
// @Tags focus
// @Summary 获取经常浏览用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "获取经常浏览用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /focus/getFrequentBrowsingUserList [post]
func (focusApi *FocusApi) GetFrequentBrowsingUserList(c *gin.Context) {

}

// GetFocusUserDynamicList 获取关注用户动态列表
// @Tags focus
// @Summary 获取关注用户动态列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "获取关注用户动态列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /focus/getFocusUserDynamicList [post]
func (focusApi *FocusApi) GetFocusUserDynamicList(c *gin.Context) {

}

// FocusUser 关注用户
// @Tags focus
// @Summary 关注用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "关注用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /focus/focusUser [post]
func (focusApi *FocusApi) FocusUser(c *gin.Context) {

}

// FocusUserCancel 取消用户关注
// @Tags focus
// @Summary 取消用户关注
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "取消用户关注"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /focus/focusUserCancel [post]
func (focusApi *FocusApi) FocusUserCancel(c *gin.Context) {

}

// GetFansList 获取粉丝列表
// @Tags focus
// @Summary 获取粉丝列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "获取粉丝列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /focus/getFansList [post]
func (focusApi *FocusApi) GetFansList(c *gin.Context) {

}
