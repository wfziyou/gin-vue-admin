import service from '@/utils/request'

// @Tags HkForbiddenSpeakReason
// @Summary 创建HkForbiddenSpeakReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForbiddenSpeakReason true "创建HkForbiddenSpeakReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForbiddenSpeakReason/createHkForbiddenSpeakReason [post]
export const createHkForbiddenSpeakReason = (data) => {
  return service({
    url: '/hkForbiddenSpeakReason/createHkForbiddenSpeakReason',
    method: 'post',
    data
  })
}

// @Tags HkForbiddenSpeakReason
// @Summary 删除HkForbiddenSpeakReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForbiddenSpeakReason true "删除HkForbiddenSpeakReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForbiddenSpeakReason/deleteHkForbiddenSpeakReason [delete]
export const deleteHkForbiddenSpeakReason = (data) => {
  return service({
    url: '/hkForbiddenSpeakReason/deleteHkForbiddenSpeakReason',
    method: 'delete',
    data
  })
}

// @Tags HkForbiddenSpeakReason
// @Summary 删除HkForbiddenSpeakReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForbiddenSpeakReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForbiddenSpeakReason/deleteHkForbiddenSpeakReason [delete]
export const deleteHkForbiddenSpeakReasonByIds = (data) => {
  return service({
    url: '/hkForbiddenSpeakReason/deleteHkForbiddenSpeakReasonByIds',
    method: 'delete',
    data
  })
}

// @Tags HkForbiddenSpeakReason
// @Summary 更新HkForbiddenSpeakReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForbiddenSpeakReason true "更新HkForbiddenSpeakReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForbiddenSpeakReason/updateHkForbiddenSpeakReason [put]
export const updateHkForbiddenSpeakReason = (data) => {
  return service({
    url: '/hkForbiddenSpeakReason/updateHkForbiddenSpeakReason',
    method: 'put',
    data
  })
}

// @Tags HkForbiddenSpeakReason
// @Summary 用id查询HkForbiddenSpeakReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkForbiddenSpeakReason true "用id查询HkForbiddenSpeakReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForbiddenSpeakReason/findHkForbiddenSpeakReason [get]
export const findHkForbiddenSpeakReason = (params) => {
  return service({
    url: '/hkForbiddenSpeakReason/findHkForbiddenSpeakReason',
    method: 'get',
    params
  })
}

// @Tags HkForbiddenSpeakReason
// @Summary 分页获取HkForbiddenSpeakReason列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkForbiddenSpeakReason列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForbiddenSpeakReason/getHkForbiddenSpeakReasonList [get]
export const getHkForbiddenSpeakReasonList = (params) => {
  return service({
    url: '/hkForbiddenSpeakReason/getHkForbiddenSpeakReasonList',
    method: 'get',
    params
  })
}
