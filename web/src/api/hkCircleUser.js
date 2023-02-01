import service from '@/utils/request'

// @Tags HkCircleUser
// @Summary 创建HkCircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleUser true "创建HkCircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleUser/createHkCircleUser [post]
export const createHkCircleUser = (data) => {
  return service({
    url: '/hkCircleUser/createHkCircleUser',
    method: 'post',
    data
  })
}

// @Tags HkCircleUser
// @Summary 删除HkCircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleUser true "删除HkCircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleUser/deleteHkCircleUser [delete]
export const deleteHkCircleUser = (data) => {
  return service({
    url: '/hkCircleUser/deleteHkCircleUser',
    method: 'delete',
    data
  })
}

// @Tags HkCircleUser
// @Summary 删除HkCircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleUser/deleteHkCircleUser [delete]
export const deleteHkCircleUserByIds = (data) => {
  return service({
    url: '/hkCircleUser/deleteHkCircleUserByIds',
    method: 'delete',
    data
  })
}

// @Tags HkCircleUser
// @Summary 更新HkCircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleUser true "更新HkCircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleUser/updateHkCircleUser [put]
export const updateHkCircleUser = (data) => {
  return service({
    url: '/hkCircleUser/updateHkCircleUser',
    method: 'put',
    data
  })
}

// @Tags HkCircleUser
// @Summary 用id查询HkCircleUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkCircleUser true "用id查询HkCircleUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleUser/findHkCircleUser [get]
export const findHkCircleUser = (params) => {
  return service({
    url: '/hkCircleUser/findHkCircleUser',
    method: 'get',
    params
  })
}

// @Tags HkCircleUser
// @Summary 分页获取HkCircleUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkCircleUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleUser/getHkCircleUserList [get]
export const getHkCircleUserList = (params) => {
  return service({
    url: '/hkCircleUser/getHkCircleUserList',
    method: 'get',
    params
  })
}
