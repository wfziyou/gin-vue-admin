import service from '@/utils/request'

// @Tags HkBugReport
// @Summary 创建HkBugReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkBugReport true "创建HkBugReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkBugReport/createHkBugReport [post]
export const createHkBugReport = (data) => {
  return service({
    url: '/hkBugReport/createHkBugReport',
    method: 'post',
    data
  })
}

// @Tags HkBugReport
// @Summary 删除HkBugReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkBugReport true "删除HkBugReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkBugReport/deleteHkBugReport [delete]
export const deleteHkBugReport = (data) => {
  return service({
    url: '/hkBugReport/deleteHkBugReport',
    method: 'delete',
    data
  })
}

// @Tags HkBugReport
// @Summary 删除HkBugReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkBugReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkBugReport/deleteHkBugReport [delete]
export const deleteHkBugReportByIds = (data) => {
  return service({
    url: '/hkBugReport/deleteHkBugReportByIds',
    method: 'delete',
    data
  })
}

// @Tags HkBugReport
// @Summary 更新HkBugReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkBugReport true "更新HkBugReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkBugReport/updateHkBugReport [put]
export const updateHkBugReport = (data) => {
  return service({
    url: '/hkBugReport/updateHkBugReport',
    method: 'put',
    data
  })
}

// @Tags HkBugReport
// @Summary 用id查询HkBugReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkBugReport true "用id查询HkBugReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkBugReport/findHkBugReport [get]
export const findHkBugReport = (params) => {
  return service({
    url: '/hkBugReport/findHkBugReport',
    method: 'get',
    params
  })
}

// @Tags HkBugReport
// @Summary 分页获取HkBugReport列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkBugReport列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkBugReport/getHkBugReportList [get]
export const getHkBugReportList = (params) => {
  return service({
    url: '/hkBugReport/getHkBugReportList',
    method: 'get',
    params
  })
}
