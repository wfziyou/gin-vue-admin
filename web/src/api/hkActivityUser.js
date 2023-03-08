import service from '@/utils/request'

// @Tags HkActivityUser
// @Summary 创建HkActivityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivityUser true "创建HkActivityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityUser/createHkActivityUser [post]
export const createHkActivityUser = (data) => {
  return service({
    url: '/hkActivityUser/createHkActivityUser',
    method: 'post',
    data
  })
}

// @Tags HkActivityUser
// @Summary 删除HkActivityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivityUser true "删除HkActivityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivityUser/deleteHkActivityUser [delete]
export const deleteHkActivityUser = (data) => {
  return service({
    url: '/hkActivityUser/deleteHkActivityUser',
    method: 'delete',
    data
  })
}

// @Tags HkActivityUser
// @Summary 删除HkActivityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkActivityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivityUser/deleteHkActivityUser [delete]
export const deleteHkActivityUserByIds = (data) => {
  return service({
    url: '/hkActivityUser/deleteHkActivityUserByIds',
    method: 'delete',
    data
  })
}

// @Tags HkActivityUser
// @Summary 更新HkActivityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivityUser true "更新HkActivityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkActivityUser/updateHkActivityUser [put]
export const updateHkActivityUser = (data) => {
  return service({
    url: '/hkActivityUser/updateHkActivityUser',
    method: 'put',
    data
  })
}

// @Tags HkActivityUser
// @Summary 用id查询HkActivityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkActivityUser true "用id查询HkActivityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkActivityUser/findHkActivityUser [get]
export const findHkActivityUser = (params) => {
  return service({
    url: '/hkActivityUser/findHkActivityUser',
    method: 'get',
    params
  })
}

// @Tags HkActivityUser
// @Summary 分页获取HkActivityUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkActivityUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityUser/getHkActivityUserList [get]
export const getHkActivityUserList = (params) => {
  return service({
    url: '/hkActivityUser/getHkActivityUserList',
    method: 'get',
    params
  })
}
