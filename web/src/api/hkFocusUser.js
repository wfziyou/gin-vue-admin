import service from '@/utils/request'

// @Tags HkFocusUser
// @Summary 创建HkFocusUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkFocusUser true "创建HkFocusUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkFocusUser/createHkFocusUser [post]
export const createHkFocusUser = (data) => {
  return service({
    url: '/hkFocusUser/createHkFocusUser',
    method: 'post',
    data
  })
}

// @Tags HkFocusUser
// @Summary 删除HkFocusUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkFocusUser true "删除HkFocusUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkFocusUser/deleteHkFocusUser [delete]
export const deleteHkFocusUser = (data) => {
  return service({
    url: '/hkFocusUser/deleteHkFocusUser',
    method: 'delete',
    data
  })
}

// @Tags HkFocusUser
// @Summary 删除HkFocusUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkFocusUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkFocusUser/deleteHkFocusUser [delete]
export const deleteHkFocusUserByIds = (data) => {
  return service({
    url: '/hkFocusUser/deleteHkFocusUserByIds',
    method: 'delete',
    data
  })
}

// @Tags HkFocusUser
// @Summary 更新HkFocusUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkFocusUser true "更新HkFocusUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkFocusUser/updateHkFocusUser [put]
export const updateHkFocusUser = (data) => {
  return service({
    url: '/hkFocusUser/updateHkFocusUser',
    method: 'put',
    data
  })
}

// @Tags HkFocusUser
// @Summary 用id查询HkFocusUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkFocusUser true "用id查询HkFocusUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkFocusUser/findHkFocusUser [get]
export const findHkFocusUser = (params) => {
  return service({
    url: '/hkFocusUser/findHkFocusUser',
    method: 'get',
    params
  })
}

// @Tags HkFocusUser
// @Summary 分页获取HkFocusUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkFocusUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkFocusUser/getHkFocusUserList [get]
export const getHkFocusUserList = (params) => {
  return service({
    url: '/hkFocusUser/getHkFocusUserList',
    method: 'get',
    params
  })
}
