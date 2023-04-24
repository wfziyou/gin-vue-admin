import service from '@/utils/request'

// @Tags HkFeedback
// @Summary 创建HkFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkFeedback true "创建HkFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkFeedback/createHkFeedback [post]
export const createHkFeedback = (data) => {
  return service({
    url: '/hkFeedback/createHkFeedback',
    method: 'post',
    data
  })
}

// @Tags HkFeedback
// @Summary 删除HkFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkFeedback true "删除HkFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkFeedback/deleteHkFeedback [delete]
export const deleteHkFeedback = (data) => {
  return service({
    url: '/hkFeedback/deleteHkFeedback',
    method: 'delete',
    data
  })
}

// @Tags HkFeedback
// @Summary 删除HkFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkFeedback/deleteHkFeedback [delete]
export const deleteHkFeedbackByIds = (data) => {
  return service({
    url: '/hkFeedback/deleteHkFeedbackByIds',
    method: 'delete',
    data
  })
}

// @Tags HkFeedback
// @Summary 更新HkFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkFeedback true "更新HkFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkFeedback/updateHkFeedback [put]
export const updateHkFeedback = (data) => {
  return service({
    url: '/hkFeedback/updateHkFeedback',
    method: 'put',
    data
  })
}

// @Tags HkFeedback
// @Summary 用id查询HkFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkFeedback true "用id查询HkFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkFeedback/findHkFeedback [get]
export const findHkFeedback = (params) => {
  return service({
    url: '/hkFeedback/findHkFeedback',
    method: 'get',
    params
  })
}

// @Tags HkFeedback
// @Summary 分页获取HkFeedback列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkFeedback列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkFeedback/getHkFeedbackList [get]
export const getHkFeedbackList = (params) => {
  return service({
    url: '/hkFeedback/getHkFeedbackList',
    method: 'get',
    params
  })
}
