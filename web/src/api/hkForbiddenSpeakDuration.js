import service from '@/utils/request'

// @Tags HkForbiddenSpeakDuration
// @Summary 创建HkForbiddenSpeakDuration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForbiddenSpeakDuration true "创建HkForbiddenSpeakDuration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForbiddenSpeakDuration/createHkForbiddenSpeakDuration [post]
export const createHkForbiddenSpeakDuration = (data) => {
  return service({
    url: '/hkForbiddenSpeakDuration/createHkForbiddenSpeakDuration',
    method: 'post',
    data
  })
}

// @Tags HkForbiddenSpeakDuration
// @Summary 删除HkForbiddenSpeakDuration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForbiddenSpeakDuration true "删除HkForbiddenSpeakDuration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForbiddenSpeakDuration/deleteHkForbiddenSpeakDuration [delete]
export const deleteHkForbiddenSpeakDuration = (data) => {
  return service({
    url: '/hkForbiddenSpeakDuration/deleteHkForbiddenSpeakDuration',
    method: 'delete',
    data
  })
}

// @Tags HkForbiddenSpeakDuration
// @Summary 删除HkForbiddenSpeakDuration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkForbiddenSpeakDuration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkForbiddenSpeakDuration/deleteHkForbiddenSpeakDuration [delete]
export const deleteHkForbiddenSpeakDurationByIds = (data) => {
  return service({
    url: '/hkForbiddenSpeakDuration/deleteHkForbiddenSpeakDurationByIds',
    method: 'delete',
    data
  })
}

// @Tags HkForbiddenSpeakDuration
// @Summary 更新HkForbiddenSpeakDuration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkForbiddenSpeakDuration true "更新HkForbiddenSpeakDuration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkForbiddenSpeakDuration/updateHkForbiddenSpeakDuration [put]
export const updateHkForbiddenSpeakDuration = (data) => {
  return service({
    url: '/hkForbiddenSpeakDuration/updateHkForbiddenSpeakDuration',
    method: 'put',
    data
  })
}

// @Tags HkForbiddenSpeakDuration
// @Summary 用id查询HkForbiddenSpeakDuration
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkForbiddenSpeakDuration true "用id查询HkForbiddenSpeakDuration"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkForbiddenSpeakDuration/findHkForbiddenSpeakDuration [get]
export const findHkForbiddenSpeakDuration = (params) => {
  return service({
    url: '/hkForbiddenSpeakDuration/findHkForbiddenSpeakDuration',
    method: 'get',
    params
  })
}

// @Tags HkForbiddenSpeakDuration
// @Summary 分页获取HkForbiddenSpeakDuration列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkForbiddenSpeakDuration列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkForbiddenSpeakDuration/getHkForbiddenSpeakDurationList [get]
export const getHkForbiddenSpeakDurationList = (params) => {
  return service({
    url: '/hkForbiddenSpeakDuration/getHkForbiddenSpeakDurationList',
    method: 'get',
    params
  })
}
