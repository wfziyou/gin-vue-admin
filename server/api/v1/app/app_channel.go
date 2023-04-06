package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HkChannelApi struct {
}

var hkChannelService = service.ServiceGroupApp.AppServiceGroup.Community.HkChannelService

// GetChannelList 获取频道列表
// @Tags HkChannel
// @Summary 获取频道列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200  {object}  response.Response{data=[]community.HkChannel,msg=string}  "返回[]community.HkChannel"
// @Router /hkChannel/getChannelList [get]
func (hkChannelApi *HkChannelApi) GetChannelList(c *gin.Context) {
	if list, _, err := hkChannelService.GetChannelInfoList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.Response{
			Data: list,
		}, "获取成功", c)
	}
}

// GetUserChannelList 获取用户频道
// @Tags HkChannel
// @Summary 获取用户频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]community.HkChannel,msg=string} "返回[]community.HkChannel"
// @Router /hkChannel/getUserChannelList [get]
func (hkChannelApi *HkChannelApi) GetUserChannelList(c *gin.Context) {
	var pageInfo communityReq.HkChannelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	channelIds, err := appUserService.GetUserChannel(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, _, err := hkChannelService.GetChannelInfoListById(channelIds); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.Response{
			Data: list,
		}, "获取成功", c)
	}
}

// SetUserChannel 设置用户频道
// @Tags HkChannel
// @Summary 设置用户频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "设置用户频道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /hkChannel/setUserChannel [post]
func (hkChannelApi *HkChannelApi) SetUserChannel(c *gin.Context) {
	var req communityReq.ParamSetUserChannel
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := appUserService.UpdateUserChannel(userId, req.ChannelIds); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// GetNewsList 获取资讯列表
// @Tags HkChannel
// @Summary 获取资讯列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "获取资讯列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /hkChannel/getNewsList [post]
func (hkChannelApi *HkChannelApi) GetNewsList(c *gin.Context) {

}

// GetNewsTopList 获取资讯置顶列表
// @Tags HkChannel
// @Summary 获取资讯置顶列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "获取资讯置顶列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /hkChannel/getNewsTopList [post]
func (hkChannelApi *HkChannelApi) GetNewsTopList(c *gin.Context) {

}

// GetNearbyHotTopicList 获取附近热门话题列表
// @Tags HkChannel
// @Summary 获取附近热门话题列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "获取附近热门话题列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /hkChannel/getNearbyHotTopicList [post]
func (hkChannelApi *HkChannelApi) GetNearbyHotTopicList(c *gin.Context) {

}

// GetNearbyDynamicList 获取附近动态列表
// @Tags HkChannel
// @Summary 获取附近动态列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "获取附近动态列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /hkChannel/getNearbyDynamicList [post]
func (hkChannelApi *HkChannelApi) GetNearbyDynamicList(c *gin.Context) {

}
