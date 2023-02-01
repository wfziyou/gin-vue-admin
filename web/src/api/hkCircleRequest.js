import service from '@/utils/request'

// @Tags HkCircleRequest
// @Summary 创建HkCircleRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleRequest true "创建HkCircleRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleRequest/createHkCircleRequest [post]
export const createHkCircleRequest = (data) => {
  return service({
    url: '/hkCircleRequest/createHkCircleRequest',
    method: 'post',
    data
  })
}

// @Tags HkCircleRequest
// @Summary 删除HkCircleRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleRequest true "删除HkCircleRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleRequest/deleteHkCircleRequest [delete]
export const deleteHkCircleRequest = (data) => {
  return service({
    url: '/hkCircleRequest/deleteHkCircleRequest',
    method: 'delete',
    data
  })
}

// @Tags HkCircleRequest
// @Summary 删除HkCircleRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HkCircleRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hkCircleRequest/deleteHkCircleRequest [delete]
export const deleteHkCircleRequestByIds = (data) => {
  return service({
    url: '/hkCircleRequest/deleteHkCircleRequestByIds',
    method: 'delete',
    data
  })
}

// @Tags HkCircleRequest
// @Summary 更新HkCircleRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HkCircleRequest true "更新HkCircleRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hkCircleRequest/updateHkCircleRequest [put]
export const updateHkCircleRequest = (data) => {
  return service({
    url: '/hkCircleRequest/updateHkCircleRequest',
    method: 'put',
    data
  })
}

// @Tags HkCircleRequest
// @Summary 用id查询HkCircleRequest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HkCircleRequest true "用id查询HkCircleRequest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hkCircleRequest/findHkCircleRequest [get]
export const findHkCircleRequest = (params) => {
  return service({
    url: '/hkCircleRequest/findHkCircleRequest',
    method: 'get',
    params
  })
}

// @Tags HkCircleRequest
// @Summary 分页获取HkCircleRequest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取HkCircleRequest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hkCircleRequest/getHkCircleRequestList [get]
export const getHkCircleRequestList = (params) => {
  return service({
    url: '/hkCircleRequest/getHkCircleRequestList',
    method: 'get',
    params
  })
}
