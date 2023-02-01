import service from '@/utils/request'

// @Tags HkUserBrowsingHistory
// @Summary 创建HkUserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserBrowsingHistory true "创建HkUserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserBrowsingHistory/createHkUserBrowsingHistory [post]
export const createHkUserBrowsingHistory = (data) => {
  return service({
    url: '/hkUserBrowsingHistory/createHkUserBrowsingHistory',
    method: 'post',
    data
  })
}

// @Tags HkUserBrowsingHistory
// @Summary 删除HkUserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserBrowsingHistory true "删除HkUserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserBrowsingHistory/deleteHkUserBrowsingHistory [delete]
export const deleteHkUserBrowsingHistory = (data) => {
  return service({
    url: '/hkUserBrowsingHistory/deleteHkUserBrowsingHistory',
    method: 'delete',
    data
  })
}

// @Tags HkUserBrowsingHistory
// @Summary 删除HkUserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUserBrowsingHistory/deleteHkUserBrowsingHistory [delete]
export const deleteHkUserBrowsingHistoryByIds = (data) => {
  return service({
    url: '/hkUserBrowsingHistory/deleteHkUserBrowsingHistoryByIds',
    method: 'delete',
    data
  })
}

// @Tags HkUserBrowsingHistory
// @Summary 更新HkUserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUserBrowsingHistory true "更新HkUserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUserBrowsingHistory/updateHkUserBrowsingHistory [put]
export const updateHkUserBrowsingHistory = (data) => {
  return service({
    url: '/hkUserBrowsingHistory/updateHkUserBrowsingHistory',
    method: 'put',
    data
  })
}

// @Tags HkUserBrowsingHistory
// @Summary 用id查询HkUserBrowsingHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkUserBrowsingHistory true "用id查询HkUserBrowsingHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUserBrowsingHistory/findHkUserBrowsingHistory [get]
export const findHkUserBrowsingHistory = (params) => {
  return service({
    url: '/hkUserBrowsingHistory/findHkUserBrowsingHistory',
    method: 'get',
    params
  })
}

// @Tags HkUserBrowsingHistory
// @Summary 分页获取HkUserBrowsingHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkUserBrowsingHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUserBrowsingHistory/getHkUserBrowsingHistoryList [get]
export const getHkUserBrowsingHistoryList = (params) => {
  return service({
    url: '/hkUserBrowsingHistory/getHkUserBrowsingHistoryList',
    method: 'get',
    params
  })
}
