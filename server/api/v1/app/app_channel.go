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
// @Tags 频道
// @Summary 获取频道列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object}  response.Response{data=[]community.ChannelInfo,msg=string} "返回[]community.HkChannel"
// @Router /app/channel/getChannelList [get]
func (hkChannelApi *HkChannelApi) GetChannelList(c *gin.Context) {
	list, _, err := hkChannelService.GetChannelInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	response.OkWithDetailed(list, "获取成功", c)
}

// GetUserChannelList 获取用户频道
// @Tags 频道
// @Summary 获取用户频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object}  response.Response{data=[]community.ChannelInfo,msg=string} "返回[]community.ChannelInfo"
// @Router /app/channel/getUserChannelList [get]
func (hkChannelApi *HkChannelApi) GetUserChannelList(c *gin.Context) {
	userId := utils.GetUserID(c)
	channelIds, err := appUserService.GetUserChannel(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, _, err := hkChannelService.GetChannelInfoListById(channelIds); err != nil {
		tmp := utils.SplitToUint64List(channelIds, ",")
		for index, id := range tmp {
			for _, obj := range list {
				if obj.ID == id {
					obj.Sort = index
					break
				}
			}
		}
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

// SetUserChannel 设置用户频道
// @Tags 频道
// @Summary 设置用户频道
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body communityReq.ParamSetUserChannel true "设置用户频道"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功"}"
// @Router /app/channel/setUserChannel [post]
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
