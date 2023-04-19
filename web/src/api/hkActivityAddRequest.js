import service from '@/utils/request'

// @Tags HkActivityAddRequest
// @Summary 创建HkActivityAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivityAddRequest true "创建HkActivityAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityAddRequest/createHkActivityAddRequest [post]
export const createHkActivityAddRequest = (data) => {
  return service({
    url: '/hkActivityAddRequest/createHkActivityAddRequest',
    method: 'post',
    data
  })
}

// @Tags HkActivityAddRequest
// @Summary 删除HkActivityAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivityAddRequest true "删除HkActivityAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivityAddRequest/deleteHkActivityAddRequest [delete]
export const deleteHkActivityAddRequest = (data) => {
  return service({
    url: '/hkActivityAddRequest/deleteHkActivityAddRequest',
    method: 'delete',
    data
  })
}

// @Tags HkActivityAddRequest
// @Summary 删除HkActivityAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkActivityAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkActivityAddRequest/deleteHkActivityAddRequest [delete]
export const deleteHkActivityAddRequestByIds = (data) => {
  return service({
    url: '/hkActivityAddRequest/deleteHkActivityAddRequestByIds',
    method: 'delete',
    data
  })
}

// @Tags HkActivityAddRequest
// @Summary 更新HkActivityAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkActivityAddRequest true "更新HkActivityAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkActivityAddRequest/updateHkActivityAddRequest [put]
export const updateHkActivityAddRequest = (data) => {
  return service({
    url: '/hkActivityAddRequest/updateHkActivityAddRequest',
    method: 'put',
    data
  })
}

// @Tags HkActivityAddRequest
// @Summary 用id查询HkActivityAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkActivityAddRequest true "用id查询HkActivityAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkActivityAddRequest/findHkActivityAddRequest [get]
export const findHkActivityAddRequest = (params) => {
  return service({
    url: '/hkActivityAddRequest/findHkActivityAddRequest',
    method: 'get',
    params
  })
}

// @Tags HkActivityAddRequest
// @Summary 分页获取HkActivityAddRequest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkActivityAddRequest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkActivityAddRequest/getHkActivityAddRequestList [get]
export const getHkActivityAddRequestList = (params) => {
  return service({
    url: '/hkActivityAddRequest/getHkActivityAddRequestList',
    method: 'get',
    params
  })
}
