import service from '@/utils/request'

// @Tags HkForbiddenSpeak
// @Summary 创建HkForbiddenSpeak
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForbiddenSpeak true "创建HkForbiddenSpeak"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForbiddenSpeak/createHkForbiddenSpeak [post]
export const createHkForbiddenSpeak = (data) => {
  return service({
    url: '/hkForbiddenSpeak/createHkForbiddenSpeak',
    method: 'post',
    data
  })
}

// @Tags HkForbiddenSpeak
// @Summary 删除HkForbiddenSpeak
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForbiddenSpeak true "删除HkForbiddenSpeak"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForbiddenSpeak/deleteHkForbiddenSpeak [delete]
export const deleteHkForbiddenSpeak = (data) => {
  return service({
    url: '/hkForbiddenSpeak/deleteHkForbiddenSpeak',
    method: 'delete',
    data
  })
}

// @Tags HkForbiddenSpeak
// @Summary 删除HkForbiddenSpeak
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForbiddenSpeak"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForbiddenSpeak/deleteHkForbiddenSpeak [delete]
export const deleteHkForbiddenSpeakByIds = (data) => {
  return service({
    url: '/hkForbiddenSpeak/deleteHkForbiddenSpeakByIds',
    method: 'delete',
    data
  })
}

// @Tags HkForbiddenSpeak
// @Summary 更新HkForbiddenSpeak
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForbiddenSpeak true "更新HkForbiddenSpeak"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForbiddenSpeak/updateHkForbiddenSpeak [put]
export const updateHkForbiddenSpeak = (data) => {
  return service({
    url: '/hkForbiddenSpeak/updateHkForbiddenSpeak',
    method: 'put',
    data
  })
}

// @Tags HkForbiddenSpeak
// @Summary 用id查询HkForbiddenSpeak
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkForbiddenSpeak true "用id查询HkForbiddenSpeak"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForbiddenSpeak/findHkForbiddenSpeak [get]
export const findHkForbiddenSpeak = (params) => {
  return service({
    url: '/hkForbiddenSpeak/findHkForbiddenSpeak',
    method: 'get',
    params
  })
}

// @Tags HkForbiddenSpeak
// @Summary 分页获取HkForbiddenSpeak列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkForbiddenSpeak列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForbiddenSpeak/getHkForbiddenSpeakList [get]
export const getHkForbiddenSpeakList = (params) => {
  return service({
    url: '/hkForbiddenSpeak/getHkForbiddenSpeakList',
    method: 'get',
    params
  })
}
