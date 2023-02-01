import service from '@/utils/request'

// @Tags HkUser
// @Summary 创建HkUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUser true "创建HkUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUser/createHkUser [post]
export const createHkUser = (data) => {
  return service({
    url: '/hkUser/createHkUser',
    method: 'post',
    data
  })
}

// @Tags HkUser
// @Summary 删除HkUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUser true "删除HkUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUser/deleteHkUser [delete]
export const deleteHkUser = (data) => {
  return service({
    url: '/hkUser/deleteHkUser',
    method: 'delete',
    data
  })
}

// @Tags HkUser
// @Summary 删除HkUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkUser/deleteHkUser [delete]
export const deleteHkUserByIds = (data) => {
  return service({
    url: '/hkUser/deleteHkUserByIds',
    method: 'delete',
    data
  })
}

// @Tags HkUser
// @Summary 更新HkUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkUser true "更新HkUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkUser/updateHkUser [put]
export const updateHkUser = (data) => {
  return service({
    url: '/hkUser/updateHkUser',
    method: 'put',
    data
  })
}

// @Tags HkUser
// @Summary 用id查询HkUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkUser true "用id查询HkUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkUser/findHkUser [get]
export const findHkUser = (params) => {
  return service({
    url: '/hkUser/findHkUser',
    method: 'get',
    params
  })
}

// @Tags HkUser
// @Summary 分页获取HkUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkUser/getHkUserList [get]
export const getHkUserList = (params) => {
  return service({
    url: '/hkUser/getHkUserList',
    method: 'get',
    params
  })
}
