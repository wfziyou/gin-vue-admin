import service from '@/utils/request'

// @Tags HkReport
// @Summary 创建HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkReport true "创建HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkReport/createHkReport [post]
export const createHkReport = (data) => {
  return service({
    url: '/hkReport/createHkReport',
    method: 'post',
    data
  })
}

// @Tags HkReport
// @Summary 删除HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkReport true "删除HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkReport/deleteHkReport [delete]
export const deleteHkReport = (data) => {
  return service({
    url: '/hkReport/deleteHkReport',
    method: 'delete',
    data
  })
}

// @Tags HkReport
// @Summary 删除HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkReport/deleteHkReport [delete]
export const deleteHkReportByIds = (data) => {
  return service({
    url: '/hkReport/deleteHkReportByIds',
    method: 'delete',
    data
  })
}

// @Tags HkReport
// @Summary 更新HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkReport true "更新HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkReport/updateHkReport [put]
export const updateHkReport = (data) => {
  return service({
    url: '/hkReport/updateHkReport',
    method: 'put',
    data
  })
}

// @Tags HkReport
// @Summary 用id查询HkReport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkReport true "用id查询HkReport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkReport/findHkReport [get]
export const findHkReport = (params) => {
  return service({
    url: '/hkReport/findHkReport',
    method: 'get',
    params
  })
}

// @Tags HkReport
// @Summary 分页获取HkReport列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkReport列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkReport/getHkReportList [get]
export const getHkReportList = (params) => {
  return service({
    url: '/hkReport/getHkReportList',
    method: 'get',
    params
  })
}
