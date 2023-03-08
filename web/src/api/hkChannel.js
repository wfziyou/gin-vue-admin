import service from '@/utils/request'

// @Tags HkChannel
// @Summary 创建HkChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkChannel true "创建HkChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkChannel/createHkChannel [post]
export const createHkChannel = (data) => {
  return service({
    url: '/hkChannel/createHkChannel',
    method: 'post',
    data
  })
}

// @Tags HkChannel
// @Summary 删除HkChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkChannel true "删除HkChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkChannel/deleteHkChannel [delete]
export const deleteHkChannel = (data) => {
  return service({
    url: '/hkChannel/deleteHkChannel',
    method: 'delete',
    data
  })
}

// @Tags HkChannel
// @Summary 删除HkChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkChannel/deleteHkChannel [delete]
export const deleteHkChannelByIds = (data) => {
  return service({
    url: '/hkChannel/deleteHkChannelByIds',
    method: 'delete',
    data
  })
}

// @Tags HkChannel
// @Summary 更新HkChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkChannel true "更新HkChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkChannel/updateHkChannel [put]
export const updateHkChannel = (data) => {
  return service({
    url: '/hkChannel/updateHkChannel',
    method: 'put',
    data
  })
}

// @Tags HkChannel
// @Summary 用id查询HkChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkChannel true "用id查询HkChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkChannel/findHkChannel [get]
export const findHkChannel = (params) => {
  return service({
    url: '/hkChannel/findHkChannel',
    method: 'get',
    params
  })
}

// @Tags HkChannel
// @Summary 分页获取HkChannel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkChannel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkChannel/getHkChannelList [get]
export const getHkChannelList = (params) => {
  return service({
    url: '/hkChannel/getHkChannelList',
    method: 'get',
    params
  })
}
