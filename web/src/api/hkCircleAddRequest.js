import service from '@/utils/request'

// @Tags HkCircleAddRequest
// @Summary 创建HkCircleAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleAddRequest true "创建HkCircleAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleAddRequest/createHkCircleAddRequest [post]
export const createHkCircleAddRequest = (data) => {
  return service({
    url: '/hkCircleAddRequest/createHkCircleAddRequest',
    method: 'post',
    data
  })
}

// @Tags HkCircleAddRequest
// @Summary 删除HkCircleAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleAddRequest true "删除HkCircleAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleAddRequest/deleteHkCircleAddRequest [delete]
export const deleteHkCircleAddRequest = (data) => {
  return service({
    url: '/hkCircleAddRequest/deleteHkCircleAddRequest',
    method: 'delete',
    data
  })
}

// @Tags HkCircleAddRequest
// @Summary 删除HkCircleAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleAddRequest/deleteHkCircleAddRequest [delete]
export const deleteHkCircleAddRequestByIds = (data) => {
  return service({
    url: '/hkCircleAddRequest/deleteHkCircleAddRequestByIds',
    method: 'delete',
    data
  })
}

// @Tags HkCircleAddRequest
// @Summary 更新HkCircleAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleAddRequest true "更新HkCircleAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleAddRequest/updateHkCircleAddRequest [put]
export const updateHkCircleAddRequest = (data) => {
  return service({
    url: '/hkCircleAddRequest/updateHkCircleAddRequest',
    method: 'put',
    data
  })
}

// @Tags HkCircleAddRequest
// @Summary 用id查询HkCircleAddRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkCircleAddRequest true "用id查询HkCircleAddRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleAddRequest/findHkCircleAddRequest [get]
export const findHkCircleAddRequest = (params) => {
  return service({
    url: '/hkCircleAddRequest/findHkCircleAddRequest',
    method: 'get',
    params
  })
}

// @Tags HkCircleAddRequest
// @Summary 分页获取HkCircleAddRequest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkCircleAddRequest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleAddRequest/getHkCircleAddRequestList [get]
export const getHkCircleAddRequestList = (params) => {
  return service({
    url: '/hkCircleAddRequest/getHkCircleAddRequestList',
    method: 'get',
    params
  })
}
