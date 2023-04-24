import service from '@/utils/request'

// @Tags HkFeedbackType
// @Summary 创建HkFeedbackType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkFeedbackType true "创建HkFeedbackType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkFeedbackType/createHkFeedbackType [post]
export const createHkFeedbackType = (data) => {
  return service({
    url: '/hkFeedbackType/createHkFeedbackType',
    method: 'post',
    data
  })
}

// @Tags HkFeedbackType
// @Summary 删除HkFeedbackType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkFeedbackType true "删除HkFeedbackType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkFeedbackType/deleteHkFeedbackType [delete]
export const deleteHkFeedbackType = (data) => {
  return service({
    url: '/hkFeedbackType/deleteHkFeedbackType',
    method: 'delete',
    data
  })
}

// @Tags HkFeedbackType
// @Summary 删除HkFeedbackType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkFeedbackType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkFeedbackType/deleteHkFeedbackType [delete]
export const deleteHkFeedbackTypeByIds = (data) => {
  return service({
    url: '/hkFeedbackType/deleteHkFeedbackTypeByIds',
    method: 'delete',
    data
  })
}

// @Tags HkFeedbackType
// @Summary 更新HkFeedbackType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkFeedbackType true "更新HkFeedbackType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkFeedbackType/updateHkFeedbackType [put]
export const updateHkFeedbackType = (data) => {
  return service({
    url: '/hkFeedbackType/updateHkFeedbackType',
    method: 'put',
    data
  })
}

// @Tags HkFeedbackType
// @Summary 用id查询HkFeedbackType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkFeedbackType true "用id查询HkFeedbackType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkFeedbackType/findHkFeedbackType [get]
export const findHkFeedbackType = (params) => {
  return service({
    url: '/hkFeedbackType/findHkFeedbackType',
    method: 'get',
    params
  })
}

// @Tags HkFeedbackType
// @Summary 分页获取HkFeedbackType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkFeedbackType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkFeedbackType/getHkFeedbackTypeList [get]
export const getHkFeedbackTypeList = (params) => {
  return service({
    url: '/hkFeedbackType/getHkFeedbackTypeList',
    method: 'get',
    params
  })
}
