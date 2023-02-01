import service from '@/utils/request'

// @Tags HkReportReason
// @Summary 创建HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkReportReason true "创建HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkReportReason/createHkReportReason [post]
export const createHkReportReason = (data) => {
  return service({
    url: '/hkReportReason/createHkReportReason',
    method: 'post',
    data
  })
}

// @Tags HkReportReason
// @Summary 删除HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkReportReason true "删除HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkReportReason/deleteHkReportReason [delete]
export const deleteHkReportReason = (data) => {
  return service({
    url: '/hkReportReason/deleteHkReportReason',
    method: 'delete',
    data
  })
}

// @Tags HkReportReason
// @Summary 删除HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkReportReason/deleteHkReportReason [delete]
export const deleteHkReportReasonByIds = (data) => {
  return service({
    url: '/hkReportReason/deleteHkReportReasonByIds',
    method: 'delete',
    data
  })
}

// @Tags HkReportReason
// @Summary 更新HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkReportReason true "更新HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkReportReason/updateHkReportReason [put]
export const updateHkReportReason = (data) => {
  return service({
    url: '/hkReportReason/updateHkReportReason',
    method: 'put',
    data
  })
}

// @Tags HkReportReason
// @Summary 用id查询HkReportReason
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkReportReason true "用id查询HkReportReason"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkReportReason/findHkReportReason [get]
export const findHkReportReason = (params) => {
  return service({
    url: '/hkReportReason/findHkReportReason',
    method: 'get',
    params
  })
}

// @Tags HkReportReason
// @Summary 分页获取HkReportReason列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkReportReason列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkReportReason/getHkReportReasonList [get]
export const getHkReportReasonList = (params) => {
  return service({
    url: '/hkReportReason/getHkReportReasonList',
    method: 'get',
    params
  })
}
